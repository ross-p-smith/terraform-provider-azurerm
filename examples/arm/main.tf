provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = "ross"
  location = "westeurope"
}

resource "azurerm_log_analytics_workspace" "oms" {
  name                = "laross123"
  location            = "westeurope"
  resource_group_name = azurerm_resource_group.example.name
  sku                 = "pergb2018"
  retention_in_days   = 30
}

resource "azurerm_resource_group_template_deployment" "azure_monitor_arm_template" {
  name                = "azure_monitor_arm_template"
  resource_group_name = azurerm_resource_group.example.name
  deployment_mode     = "Incremental"
  template_content    = file("/workspaces/terraform-provider-azurerm/examples/arm/azure_monitor_armtemplate.json")
  depends_on          = [azurerm_log_analytics_workspace.oms]
}