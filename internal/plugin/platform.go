package plugin

import (
	"context"
	"fmt"
	azureweb "github.com/Azure/azure-sdk-for-go/profiles/latest/web/mgmt/web"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"github.com/hashicorp/waypoint/builtin/docker"
	"github.com/magodo/waypoint-plugin-azure-app-service/azure"
	"github.com/magodo/waypoint-plugin-azure-app-service/azure/web"
	"net/http"
)

type DeployConfig struct {
	ResourceGroupName string `hcl:"resource_group_name"`
	AppServiceName    string `hcl:"app_service_name"`
}

type Platform struct {
	config DeployConfig
}

func (p *Platform) Config() (interface{}, error) {
	return &p.config, nil
}

func (p *Platform) ConfigSet(config interface{}) error {
	c, ok := config.(*DeployConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *DeployConfig as parameter")
	}
	_ = c
	return nil
}

func (p *Platform) DeployFunc() interface{} {
	return p.deploy
}

func (p *Platform) deploy(
	ctx context.Context,
	ui terminal.UI,
	log hclog.Logger,
	img *docker.Image,
) (*DeploymentOutput, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Deploy application")

	authorizer, err := azure.NewAuthorizer(ctx)
	if err != nil {
		return nil, err
	}

	id := web.NewAppServiceID(authorizer.Config.SubscriptionID, p.config.ResourceGroupName, p.config.AppServiceName)

	// Check the App Service Plan to see whether its tier (only when the tier is in one of "Standard", "Premium", "Isolated")
	// allows to create separate deployment slots.
	// See: https://docs.microsoft.com/en-us/azure/app-service/deploy-staging-slots.
	ok, err := web.AppServiceSupportsSlot(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("checking whether %s supports slot: %+v", id, err)
	}

	if !ok {
		return p.deployToAppService(ctx, u, log, id, img)
	}
	return p.deployToAppServiceSlot(ctx, u, log, id, img)
}

func (p *Platform) buildSiteConfigPatch(img *docker.Image) azureweb.SiteConfig {
	linuxFxVersion := fmt.Sprintf("DOCKER|%s:%s", img.Image, img.Tag)
	return azureweb.SiteConfig{
		LinuxFxVersion: &linuxFxVersion,
	}
}

func (p *Platform) deployToAppService(ctx context.Context, u terminal.Status, log hclog.Logger, appServiceId web.AppServiceId, img *docker.Image) (*DeploymentOutput, error) {
	authorizer, err := azure.NewAuthorizer(ctx)
	if err != nil {
		return nil, err
	}
	client := authorizer.NewAppServiceClient()
	siteConfig := p.buildSiteConfigPatch(img)

	u.Update(fmt.Sprintf("Update %s", appServiceId))
	siteEnvelope := azureweb.SitePatchResource{
		SitePatchResourceProperties: &azureweb.SitePatchResourceProperties{
			SiteConfig: &siteConfig,
		},
	}
	if _, err := client.Update(ctx, appServiceId.ResourceGroup, appServiceId.SiteName, siteEnvelope); err != nil {
		return nil, err
	}

	u.Update(fmt.Sprintf("Update the configuration for %s", appServiceId))
	if _, err := client.UpdateConfiguration(ctx, appServiceId.ResourceGroup, appServiceId.SiteName, azureweb.SiteConfigResource{
		SiteConfig: &siteConfig,
	}); err != nil {
		return nil, err
	}

	return &DeploymentOutput{
		AppServiceId: appServiceId.ID(),
	}, nil
}

func (p *Platform) deployToAppServiceSlot(ctx context.Context, u terminal.Status, log hclog.Logger, appServiceId web.AppServiceId, img *docker.Image) (*DeploymentOutput, error) {
	authorizer, err := azure.NewAuthorizer(ctx)
	if err != nil {
		return nil, err
	}
	client := authorizer.NewAppServiceClient()

	deploymentId, err := component.Id()
	if err != nil {
		return nil, err
	}
	slotName := fmt.Sprintf("%s-%s", appServiceId.SiteName, deploymentId)
	slotId := web.NewAppServiceSlotID(appServiceId.SubscriptionId, appServiceId.ResourceGroup, appServiceId.SiteName, slotName)

	// Check whether it exists
	resp, err := client.GetSlot(ctx, slotId.ResourceGroup, slotId.SiteName, slotId.SlotName)
	if err != nil {
		// Since the Slot API regard 404 as non error, so if err != nil, it always mean something goes wrong.
		return nil, fmt.Errorf("getting %s: %+v", slotId, err)
	}

	embeddedResp := resp.Response.Response
	if embeddedResp == nil {
		return nil, fmt.Errorf("unexpected nil embedded response")
	}

	// Update the existing slot.
	if embeddedResp.StatusCode == http.StatusOK {
		if err := p.updateSlotImage(ctx, u, img, client, slotId); err != nil {
			return nil, err
		}
		return &DeploymentOutput{
			AppServiceId:     appServiceId.ID(),
			AppServiceSlotId: slotId.ID(),
		}, nil
	}

	// Create a new slot with the same settings from the app service.
	u.Update(fmt.Sprintf("Get %s", appServiceId))
	asresp, err := client.Get(ctx, slotId.ResourceGroup, slotId.SiteName)
	if err != nil {
		return nil, fmt.Errorf("getting %s: %+v", appServiceId, err)
	}

	u.Update(fmt.Sprintf("Create %s", slotId))
	future, err := client.CreateOrUpdateSlot(ctx, slotId.ResourceGroup, slotId.SiteName, asresp, slotId.SlotName)
	if err != nil {
		return nil, fmt.Errorf("creating %s: %+v", slotId, err)
	}

	u.Update(fmt.Sprintf("Watinig for the creation of %s", slotId))
	if err := future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return nil, fmt.Errorf("watinig for the creation of %s: %+v", slotId, err)
	}

	if err := p.updateSlotImage(ctx, u, img, client, slotId); err != nil {
		return nil, err
	}

	return &DeploymentOutput{
		AppServiceId:     appServiceId.ID(),
		AppServiceSlotId: slotId.ID(),
	}, nil
}

func (p *Platform) updateSlotImage(ctx context.Context, u terminal.Status, img *docker.Image, client azureweb.AppsClient, slotId web.AppServiceSlotId) error {
	siteConfig := p.buildSiteConfigPatch(img)
	u.Update(fmt.Sprintf("Update %s", slotId))
	siteEnvelope := azureweb.SitePatchResource{
		SitePatchResourceProperties: &azureweb.SitePatchResourceProperties{
			SiteConfig: &siteConfig,
		},
	}
	if _, err := client.UpdateSlot(ctx, slotId.ResourceGroup, slotId.SiteName, siteEnvelope, slotId.SlotName); err != nil {
		return err
	}

	u.Update(fmt.Sprintf("Update the site config for %s", slotId))
	if _, err := client.UpdateConfigurationSlot(ctx, slotId.ResourceGroup, slotId.SiteName, azureweb.SiteConfigResource{SiteConfig: &siteConfig}, slotId.SlotName); err != nil {
		return err
	}
	return nil
}

func (p *Platform) DestroyFunc() interface{} {
	return p.destroy
}

func (p *Platform) destroy(
	ctx context.Context,
	ui terminal.UI,
	deployment *DeploymentOutput,
) error {
	slotIdLit := deployment.AppServiceSlotId
	if slotIdLit == "" {
		return nil
	}

	st := ui.Status()
	defer st.Close()

	id, err := web.AppServiceSlotID(slotIdLit)
	if err != nil {
		st.Step(terminal.ErrorStyle, fmt.Sprintf("Unable to delete the App Service Slot: %s", slotIdLit))
		return err
	}

	authorizer, err := azure.NewAuthorizer(ctx)
	if err != nil {
		return err
	}
	client := authorizer.NewAppServiceClient()

	deleteMetrics := true
	deleteEmptyServerFarm := false

	if _, err := client.DeleteSlot(ctx, id.ResourceGroup, id.SiteName, id.SlotName, &deleteMetrics, &deleteEmptyServerFarm); err != nil {
		st.Step(terminal.ErrorStyle, fmt.Sprintf("Unable to delete: %s", id))
		return err
	}

	st.Step(terminal.StatusOK, fmt.Sprintf("Deleted %s", id))
	return nil
}
