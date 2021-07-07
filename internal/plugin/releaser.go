package plugin

import (
	"context"
	"fmt"
	azureweb "github.com/Azure/azure-sdk-for-go/profiles/latest/web/mgmt/web"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/docs"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"github.com/magodo/waypoint-plugin-azure-app-service/azure"
	"github.com/magodo/waypoint-plugin-azure-app-service/azure/web"
)

type ReleaseConfig struct {
	// Whether to create a new virtual network after Release
	NewVnet bool `hcl:"new_vnet,optional"`
}

type Releaser struct {
	config ReleaseConfig
}

var (
	_ component.ReleaseManager = (*Releaser)(nil)
	_ component.Configurable   = (*Releaser)(nil)
	_ component.Documented     = (*Releaser)(nil)
)

func (rm *Releaser) Config() (interface{}, error) {
	return &rm.config, nil
}

func (rm *Releaser) ConfigSet(config interface{}) error {
	_, ok := config.(*ReleaseConfig)
	if !ok {
		return fmt.Errorf("Expected *ReleaseConfig as parameter")
	}

	return nil
}

func (rm *Releaser) ReleaseFunc() interface{} {
	return rm.Release
}

// Release releases a new deployment to the Azure App Service.
//
// In case the App Service allows to create slot, it will swap the staged slot (created in deployment step) to production.
// Otherwise, it will do nothing, since in deployment step, the image has already been "released".
func (rm *Releaser) Release(ctx context.Context, ui terminal.UI, deployment *DeploymentOutput) (*ReleaseOutput, error) {
	u := ui.Status()
	defer u.Close()
	ui.Output("Release application")

	appServiceId, err := web.AppServiceID(deployment.AppServiceId)
	if err != nil {
		return nil, err
	}

	authorizer, err := azure.NewAuthorizer(ctx)
	if err != nil {
		return nil, err
	}
	client := authorizer.NewAppServiceClient()

	// Get the URL of the App Service.
	resp, err := client.Get(ctx, appServiceId.ResourceGroup, appServiceId.SiteName)
	if err != nil {
		return nil, err
	}
	url, err := web.AppServiceDefaultHost(resp)
	if err != nil {
		return nil, err
	}

	// If the AppServiceSlotId is not set in the deployment, it means this APp Service Plan doesn't support slot creation.
	// In this case, the deployment step has already done the Release, and there is nothing needed to be done here.
	if deployment.AppServiceSlotId == "" {
		u.Step(terminal.StatusOK, "Created Release")
		return &ReleaseOutput{
			Url: url,
		}, nil
	}

	// Otherwise, swap the slot and the production slot.
	slotId, err := web.AppServiceSlotID(deployment.AppServiceSlotId)
	if err != nil {
		return nil, err
	}
	preserveVnet := !rm.config.NewVnet

	slotEntity := azureweb.CsmSlotEntity{
		TargetSlot:   &slotId.SlotName,
		PreserveVnet: &preserveVnet,
	}

	future, err := client.SwapSlotWithProduction(ctx, appServiceId.ResourceGroup, appServiceId.SiteName, slotEntity)
	if err != nil {
		return nil, fmt.Errorf("swapping slot on %s: %+v", appServiceId, err)
	}
	if err := future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return nil, fmt.Errorf("waiting for slot swap on %s: %+v", appServiceId, err)
	}

	u.Step(terminal.StatusOK, "Created Release")
	return &ReleaseOutput{
		Url: url,
	}, nil
}

func (r *Releaser) Documentation() (*docs.Documentation, error) {
	doc, err := docs.New(docs.FromConfig(&ReleaseConfig{}), docs.FromFunc(r.ReleaseFunc()))
	if err != nil {
		return nil, err
	}

	return doc, nil
}
