package appservice

import (
	"fmt"
	"regexp"
	"strings"
)

type AppServiceId struct {
	SubscriptionId string
	ResourceGroup  string
	SiteName       string
}

func NewAppServiceID(subscriptionId, resourceGroup, siteName string) AppServiceId {
	return AppServiceId{
		SubscriptionId: subscriptionId,
		ResourceGroup:  resourceGroup,
		SiteName:       siteName,
	}
}

func (id AppServiceId) String() string {
	segments := []string{
		fmt.Sprintf("Site Name %q", id.SiteName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "App Service", segmentsStr)
}

func (id AppServiceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.SiteName)
}

// AppServiceID parses a AppService ID into an AppServiceId struct
func AppServiceID(input string) (*AppServiceId, error) {
	p := regexp.MustCompile(`/subscriptions/([^/]+)/resourceGroups/([^/]+)/providers/Microsoft.Web/sites/([^/]+)`)
	m := p.FindAllStringSubmatch(input, 1)
	if len(m) != 1 {
		return nil, fmt.Errorf("invalid ID format of App Service: %s", input)
	}
	if len(m[0]) != 4 {
		return nil, fmt.Errorf("invalid ID format of App Service: %s", input)
	}

	return &AppServiceId{
		SubscriptionId: m[0][1],
		ResourceGroup:  m[0][2],
		SiteName:       m[0][3],
	}, nil
}
