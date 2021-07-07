package main

import (
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
	"github.com/magodo/waypoint-plugin-azure-app-service/internal/plugin"
)

func main() {
	sdk.Main(sdk.WithComponents(
		&plugin.Platform{},
		//&plugin.Releaser{},
	))
}
