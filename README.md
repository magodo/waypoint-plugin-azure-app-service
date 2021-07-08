# waypoint-plugin-azure-app-service

Plugin for waypoint that adds support for the Azure App Service.

## Install

Check out the [waypoint doc](https://www.waypointproject.io/docs/plugins) about how to install and use an external plugin.

## Authentication

This plugin leverages the some basic auth package as the [Terraform AzureRM provider](https://github.com/terraform-providers/terraform-provider-azurerm). This means the way to authenticate against Azure is the same.

For the sake of DRY, please check out the Terraform AzureRM provider [authentication guide](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs#authenticating-to-azure).

## Prerequisite

Since there are a lot of properties can be configured for Azure App Service and its dependencies, to make it simply do one thing (i.e. CD), this plugin decides to leave the provisioning of the App Service and its dependencies to the users.

A simple provisioning can be done via below Terraform configuration:

```hcl
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = "example-rg"
  location = "west europe"
}

resource "azurerm_app_service_plan" "example" {
  name                = "example-asp"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  kind = "Linux"
  reserved = true
  
  sku {
    tier = "Standard"
    size = "S1"
  }
}

resource "azurerm_app_service" "test" {
  name                = "example-as"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id
}
```

## Usage

### Azure App Service Deployment

```hcl
deploy {
 use "azure-app-service" {
   resource_group_name = "example-rg"
   app_service_name = "example-as"
 }
}
```

For App Service that allows to create a slot (which means the sku of the App Service Plan is one of "Standard", "Premium", "Isolated"), the plguin will create a new slot for each deployment.

Otherwise, it will deploy the app to the App Service's default production slot.

### Azure App Service Release

```hcl
release {
   use "azure-app-service" {}
}
```

For App Service that allows to create a slot (which means the sku of the App Service Plan is one of "Standard", "Premium", "Isolated"), the plguin will swap the slot that is created in the deployment step with the production slot.

Otherwise, it will do nothing.
