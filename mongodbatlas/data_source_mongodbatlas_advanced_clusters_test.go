package mongodbatlas

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	matlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestAccDataSourceMongoDBAtlasAdvancedClusters_basic(t *testing.T) {
	var (
		cluster        matlas.AdvancedCluster
		resourceName   = "mongodbatlas_advanced_cluster.test"
		dataSourceName = "data.mongodbatlas_advanced_clusters.test"
		projectID      = os.Getenv("MONGODB_ATLAS_PROJECT_ID")
		name           = acctest.RandomWithPrefix("test-acc")
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckMongoDBAtlasAdvancedClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMongoDBAtlasAdvancedClustersConfig(projectID, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMongoDBAtlasAdvancedClusterExists(resourceName, &cluster),
					testAccCheckMongoDBAtlasAdvancedClusterAttributes(&cluster, name),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "replication_specs.#"),
					resource.TestCheckResourceAttrSet(resourceName, "replication_specs.0.region_configs.#"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.replication_specs.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.name"),
				),
			},
		},
	})
}

func TestAccDataSourceMongoDBAtlasAdvancedClusters_multicloud(t *testing.T) {
	var (
		cluster        matlas.AdvancedCluster
		resourceName   = "mongodbatlas_advanced_cluster.test"
		dataSourceName = "data.mongodbatlas_advanced_clusters.test"
		projectID      = os.Getenv("MONGODB_ATLAS_PROJECT_ID")
		name           = acctest.RandomWithPrefix("test-acc")
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckMongoDBAtlasAdvancedClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMongoDBAtlasAdvancedClustersMultiCloudConfig(projectID, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMongoDBAtlasAdvancedClusterExists(resourceName, &cluster),
					testAccCheckMongoDBAtlasAdvancedClusterAttributes(&cluster, name),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "replication_specs.#"),
					resource.TestCheckResourceAttrSet(resourceName, "replication_specs.0.region_configs.#"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.replication_specs.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.name"),
				),
			},
		},
	})
}

func testAccDataSourceMongoDBAtlasAdvancedClustersConfig(projectID, name string) string {
	return fmt.Sprintf(`
%s

data "mongodbatlas_advanced_clusters" "test" {
  project_id = mongodbatlas_advanced_cluster.test.project_id
}
	`, testAccMongoDBAtlasAdvancedClusterConfigTenant(projectID, name))
}

func testAccDataSourceMongoDBAtlasAdvancedClustersMultiCloudConfig(projectID, name string) string {
	return fmt.Sprintf(`
%s

data "mongodbatlas_advanced_clusters" "test" {
  project_id = mongodbatlas_advanced_cluster.test.project_id
}
	`, testAccMongoDBAtlasAdvancedClusterConfigMultiCloud(projectID, name))
}
