package web

import (
	"context"
	"fmt"
	"github.com/magodo/waypoint-plugin-azure-app-service/azure"
)

func AppServiceSupportsSlot(ctx context.Context, id *AppServiceId) (bool, error) {
	authorizer, err := azure.NewAuthorizer(ctx)
	if err != nil {
		return false, err
	}
	client := authorizer.NewAppServiceClient()

	resp, err := client.Get(ctx, id.ResourceGroup, id.SiteName)
	if err != nil {
		return false, err
	}

	if resp.SiteProperties== nil {
		return false, fmt.Errorf(`unexpected nil "siteProperties" of the App Service`)
	}
	if resp.SiteProperties.ServerFarmID == nil {
		return false, fmt.Errorf(`unexpected nil "siteProperties.serverFarmID" of the App Service`)
	}

	planId, err := AppServicePlanID(*resp.SiteProperties.ServerFarmID)
	if err != nil {
		return false, err
	}

	planClient := authorizer.NewAppServicePlanClient()
	planResp, err := planClient.Get(ctx, planId.ResourceGroup, planId.ServerfarmName)
	if err != nil {
		return false, err
	}

	if planResp.Sku == nil {
		return false, fmt.Errorf(`unexpected nil "sku" of the App Service Plan`)
	}
	if planResp.Sku.Tier == nil {
		return false, fmt.Errorf(`unexpected nil "sku.tier" of the App Service Plan`)
	}

	tier := *planResp.Sku.Tier
	return tier == "Standard" || tier == "Premium" || tier == "Isolated", nil
}
