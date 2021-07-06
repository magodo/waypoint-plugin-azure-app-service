package platform

import (
	"context"
	"fmt"

	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"github.com/hashicorp/waypoint/builtin/docker"
	"github.com/magodo/waypoint-plugin-azure-app-service/azure"
	"github.com/magodo/waypoint-plugin-azure-app-service/azure/appservice"
)

type DeployConfig struct {
	AppServiceID string `hcl:"app_service_id,optional"`
}

type Platform struct {
	config DeployConfig
}

// Implement Configurable
func (p *Platform) Config() (interface{}, error) {
	return &p.config, nil
}

// Implement ConfigurableNotify
func (p *Platform) ConfigSet(config interface{}) error {
	c, ok := config.(*DeployConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *DeployConfig as parameter")
	}

	// validate the config
	if c.AppServiceID == "" {
		return fmt.Errorf("`app_service_id` not set")
	}
	if _, err := appservice.AppServiceID(c.AppServiceID); err != nil {
		return fmt.Errorf("invalid value for `app_service_id`: %+v", err)
	}

	return nil
}

// Implement Builder
func (p *Platform) DeployFunc() interface{} {
	return p.deploy
}

func (p *Platform) deploy(
	ctx context.Context,
	img *docker.Image,
	ui terminal.UI,
) (*Deployment, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Deploy application")

	// validated in the ConfigSet()
	id, _ := appservice.AppServiceID(p.config.AppServiceID)

	authorizer, err := azure.NewAuthorizer(ctx)
	if err != nil {
		return nil, err
	}
	client := authorizer.NewAppServiceClient()

	resp, err := client.Get(ctx, id.ResourceGroup, id.SiteName)
	if err != nil {
		return nil, err
	}

	u.Update(*resp.ID)

	return &Deployment{}, nil
}
