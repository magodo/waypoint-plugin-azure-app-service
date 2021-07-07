package web

import (
	"fmt"
	"regexp"
	"strings"
)

type AppServiceSlotId struct {
	SubscriptionId string
	ResourceGroup  string
	SiteName       string
	SlotName       string
}

func NewAppServiceSlotID(subscriptionId, resourceGroup, siteName, slotName string) AppServiceSlotId {
	return AppServiceSlotId{
		SubscriptionId: subscriptionId,
		ResourceGroup:  resourceGroup,
		SiteName:       siteName,
		SlotName:       slotName,
	}
}

func (id AppServiceSlotId) String() string {
	segments := []string{
		fmt.Sprintf("Slot Name %q", id.SlotName),
		fmt.Sprintf("Site Name %q", id.SiteName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "App Service Slot", segmentsStr)
}

func (id AppServiceSlotId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.SiteName, id.SlotName)
}

func AppServiceSlotID(input string) (*AppServiceSlotId, error) {
	p := regexp.MustCompile(`/subscriptions/([^/]+)/resourceGroups/([^/]+)/providers/Microsoft.Web/sites/([^/]+)/slots/([^/]+)`)
	m := p.FindAllStringSubmatch(input, 1)
	if len(m) != 1 {
		return nil, fmt.Errorf("invalid ID format of App Service Slot: %s", input)
	}
	if len(m[0]) != 5 {
		return nil, fmt.Errorf("invalid ID format of App Service Slot: %s", input)
	}

	return &AppServiceSlotId{
		SubscriptionId: m[0][1],
		ResourceGroup:  m[0][2],
		SiteName:       m[0][3],
		SlotName:       m[0][4],
	}, nil
}
