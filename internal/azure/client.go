package azure

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/web/mgmt/web"
)

func (a *Authorizer) NewAppServiceClient() web.AppsClient {
	client := web.NewAppsClient(a.Config.SubscriptionID)
	client.Authorizer = a.authorizer
	return client
}

func (a *Authorizer) NewAppServicePlanClient() web.AppServicePlansClient {
	client := web.NewAppServicePlansClient(a.Config.SubscriptionID)
	client.Authorizer = a.authorizer
	return client
}
