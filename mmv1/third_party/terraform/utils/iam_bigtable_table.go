package google

import (
	"github.com/hashicorp/terraform-provider-google/google/services/bigtable"
	"github.com/hashicorp/terraform-provider-google/google/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var IamBigtableTableSchema = bigtable.IamBigtableTableSchema

func NewBigtableTableUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	return bigtable.NewBigtableTableUpdater(d, config)
}

func BigtableTableIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	return bigtable.BigtableTableIdParseFunc(d, config)
}