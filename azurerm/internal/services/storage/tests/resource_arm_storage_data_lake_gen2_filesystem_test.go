package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/storage/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMStorageDataLakeGen2FileSystem_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_storage_data_lake_gen2_filesystem", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMStorageDataLakeGen2FileSystemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMStorageDataLakeGen2FileSystem_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMStorageDataLakeGen2FileSystemExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMStorageDataLakeGen2FileSystem_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_storage_data_lake_gen2_filesystem", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMStorageDataLakeGen2FileSystemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMStorageDataLakeGen2FileSystem_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMStorageDataLakeGen2FileSystemExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccAzureRMStorageDataLakeGen2FileSystem_requiresImport),
		},
	})
}

func TestAccAzureRMStorageDataLakeGen2FileSystem_withDefaultACL(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_storage_data_lake_gen2_filesystem", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMStorageDataLakeGen2FileSystemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMStorageDataLakeGen2FileSystem_withDefaultACL(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMStorageDataLakeGen2FileSystemExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccAzureRMStorageDataLakeGen2FileSystem_requiresImport),
		},
	})
}

func TestAccAzureRMStorageDataLakeGen2FileSystem_UpdateDefaultACL(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_storage_data_lake_gen2_filesystem", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMStorageDataLakeGen2FileSystemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMStorageDataLakeGen2FileSystem_withDefaultACL(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMStorageDataLakeGen2FileSystemExists(data.ResourceName),
				),
			},
			data.ImportStep(),
			{
				Config: testAccAzureRMStorageDataLakeGen2FileSystem_withExecuteACLForSPN(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMStorageDataLakeGen2FileSystemExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMStorageDataLakeGen2FileSystem_properties(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_storage_data_lake_gen2_filesystem", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMStorageDataLakeGen2FileSystemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMStorageDataLakeGen2FileSystem_properties(data, "aGVsbG8="),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMStorageDataLakeGen2FileSystemExists(data.ResourceName),
				),
			},
			data.ImportStep(),
			{
				Config: testAccAzureRMStorageDataLakeGen2FileSystem_properties(data, "ZXll"),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMStorageDataLakeGen2FileSystemExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMStorageDataLakeGen2FileSystem_handlesStorageAccountDeletion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_storage_data_lake_gen2_filesystem", "test")
	config := testAccAzureRMStorageDataLakeGen2FileSystem_basic(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMStorageDataLakeGen2FileSystemDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMStorageDataLakeGen2FileSystemExists(data.ResourceName),
					testAzureRMStorageDataLakeGen2StorageAccountDelete(data.ResourceName),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMStorageDataLakeGen2FileSystemExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func testCheckAzureRMStorageDataLakeGen2FileSystemExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).Storage.FileSystemsClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		fileSystemName := rs.Primary.Attributes["name"]
		storageID, err := parse.StorageAccountID(rs.Primary.Attributes["storage_account_id"])
		if err != nil {
			return err
		}

		resp, err := client.GetProperties(ctx, storageID.Name, fileSystemName)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: File System %q (Account %q) does not exist", fileSystemName, storageID.Name)
			}

			return fmt.Errorf("Bad: Get on FileSystemsClient: %+v", err)
		}

		return nil
	}
}

func testAzureRMStorageDataLakeGen2StorageAccountDelete(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).Storage.AccountsClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		storageID, err := parse.StorageAccountID(rs.Primary.Attributes["storage_account_id"])
		if err != nil {
			return err
		}

		if _, err := client.Delete(ctx, storageID.ResourceGroup, storageID.Name); err != nil {
			return fmt.Errorf("Unable to delete azurerm_storage_account: %+v", storageID.Name)
		}

		return nil
	}
}

func testCheckAzureRMStorageDataLakeGen2FileSystemDestroy(s *terraform.State) error {
	client := acceptance.AzureProvider.Meta().(*clients.Client).Storage.FileSystemsClient
	ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_storage_data_lake_gen2_filesystem" {
			continue
		}

		fileSystemName := rs.Primary.Attributes["name"]
		storageID, err := parse.StorageAccountID(rs.Primary.Attributes["storage_account_id"])
		if err != nil {
			return err
		}

		props, err := client.GetProperties(ctx, storageID.Name, fileSystemName)
		if err != nil {
			return nil
		}

		return fmt.Errorf("File System still exists: %+v", props)
	}

	return nil
}

func testAccAzureRMStorageDataLakeGen2FileSystem_basic(data acceptance.TestData) string {
	template := testAccAzureRMStorageDataLakeGen2FileSystem_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_storage_data_lake_gen2_filesystem" "test" {
  name               = "acctest-%d"
  storage_account_id = azurerm_storage_account.test.id
}
`, template, data.RandomInteger)
}

func testAccAzureRMStorageDataLakeGen2FileSystem_requiresImport(data acceptance.TestData) string {
	template := testAccAzureRMStorageDataLakeGen2FileSystem_basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_storage_data_lake_gen2_filesystem" "import" {
  name               = azurerm_storage_data_lake_gen2_filesystem.test.name
  storage_account_id = azurerm_storage_data_lake_gen2_filesystem.test.storage_account_id
}
`, template)
}

func testAccAzureRMStorageDataLakeGen2FileSystem_withDefaultACL(data acceptance.TestData) string {
	template := testAccAzureRMStorageDataLakeGen2FileSystem_template(data)
	return fmt.Sprintf(`
%s

data "azurerm_client_config" "current" {
}

resource "azurerm_role_assignment" "storageAccountRoleAssignment" {
  scope                = azurerm_storage_account.test.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = data.azurerm_client_config.current.object_id
}

resource "azurerm_storage_data_lake_gen2_filesystem" "test" {
  name               = "acctest-%[2]d"
  storage_account_id = azurerm_storage_account.test.id
  ace {
    type        = "user"
    permissions = "rwx"
  }
  ace {
    type        = "group"
    permissions = "r-x"
  }
  ace {
    type        = "other"
    permissions = "---"
  }
  depends_on = [
    azurerm_role_assignment.storageAccountRoleAssignment
  ]
}
`, template, data.RandomInteger)
}

func testAccAzureRMStorageDataLakeGen2FileSystem_withExecuteACLForSPN(data acceptance.TestData) string {
	template := testAccAzureRMStorageDataLakeGen2FileSystem_template(data)
	return fmt.Sprintf(`
%s

data "azurerm_client_config" "current" {
}

resource "azurerm_role_assignment" "storageAccountRoleAssignment" {
  scope                = azurerm_storage_account.test.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = data.azurerm_client_config.current.object_id
}

resource "azuread_application" "test" {
  name = "acctestspa%[2]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azurerm_storage_data_lake_gen2_filesystem" "test" {
  name               = "acctest-%[2]d"
  storage_account_id = azurerm_storage_account.test.id
  ace {
    type        = "user"
    permissions = "rwx"
  }
  ace {
    type        = "user"
    id          = azuread_service_principal.test.object_id
    permissions = "--x"
  }
  ace {
    type        = "group"
    permissions = "r-x"
  }
  ace {
    type        = "mask"
    permissions = "r-x"
  }
  ace {
    type        = "other"
    permissions = "---"
  }
  depends_on = [
	azurerm_role_assignment.storageAccountRoleAssignment,
	azuread_service_principal.test
  ]
}
`, template, data.RandomInteger)
}

func testAccAzureRMStorageDataLakeGen2FileSystem_properties(data acceptance.TestData, value string) string {
	template := testAccAzureRMStorageDataLakeGen2FileSystem_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_storage_data_lake_gen2_filesystem" "test" {
  name               = "acctest-%d"
  storage_account_id = azurerm_storage_account.test.id

  properties = {
    key = "%s"
  }
}
`, template, data.RandomInteger, value)
}

func testAccAzureRMStorageDataLakeGen2FileSystem_template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestacc%s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_kind             = "BlobStorage"
  account_tier             = "Standard"
  account_replication_type = "LRS"
  is_hns_enabled           = true
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString)
}
