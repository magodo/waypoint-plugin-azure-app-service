package web

import (
	"fmt"
	"regexp"
	"strings"
)

type AppServicePlanId struct {
	SubscriptionId string
	ResourceGroup  string
	ServerfarmName string
}

func NewAppServicePlanID(subscriptionId, resourceGroup, serverfarmName string) AppServicePlanId {
	return AppServicePlanId{
		SubscriptionId: subscriptionId,
		ResourceGroup:  resourceGroup,
		ServerfarmName: serverfarmName,
	}
}

func (id AppServicePlanId) String() string {
	segments := []string{
		fmt.Sprintf("Serverfarm Name %q", id.ServerfarmName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "App Service Plan", segmentsStr)
}

func (id AppServicePlanId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/serverfarms/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.ServerfarmName)
}

func AppServicePlanID(input string) (*AppServicePlanId, error) {
	p := regexp.MustCompile(`/subscriptions/([^/]+)/resourceGroups/([^/]+)/providers/Microsoft.Web/serverfarms/([^/]+)`)
	m := p.FindAllStringSubmatch(input, 1)
	if len(m) != 1 {
		return nil, fmt.Errorf("invalid ID format of App Service Plan: %s", input)
	}
	if len(m[0]) != 4 {
		return nil, fmt.Errorf("invalid ID format of App Service Plan: %s", input)
	}

	return &AppServicePlanId{
		SubscriptionId: m[0][1],
		ResourceGroup:  m[0][2],
		ServerfarmName: m[0][3],
	}, nil
}
