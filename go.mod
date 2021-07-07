module github.com/magodo/waypoint-plugin-azure-app-service

go 1.14

require (
	github.com/Azure/azure-sdk-for-go v55.5.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.19
	github.com/davecgh/go-spew v1.1.1
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/hashicorp/go-azure-helpers v0.16.5
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/waypoint v0.4.1
	github.com/hashicorp/waypoint-plugin-sdk v0.0.0-20210609145036-5c5b44751ee6
	google.golang.org/protobuf v1.26.0
)

// replace github.com/hashicorp/waypoint-plugin-sdk => ../../waypoint-plugin-sdk
