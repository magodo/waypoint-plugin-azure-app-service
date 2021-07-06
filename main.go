package main

import (
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
	"github.com/magodo/waypoint-plugin-azure-app-service/platform"
)

func main() {
	sdk.Main(sdk.WithComponents(
		&platform.Platform{},
	))
}
