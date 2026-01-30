package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// TerraformPluginSDKExternalNameConfigs contains all external name configurations
// belonging to Terraform Plugin SDKv2 resources to be reconciled
// under the no-fork architecture for this provider.
var TerraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"opensearch_anomaly_detection":         config.IdentifierFromProvider,
	"opensearch_audit_config":              config.IdentifierFromProvider,
	"opensearch_channel_configuration":     config.IdentifierFromProvider,
	"opensearch_cluster_settings":          config.IdentifierFromProvider,
	"opensearch_component_template":        config.IdentifierFromProvider,
	"opensearch_composable_index_template": config.IdentifierFromProvider,
	"opensearch_dashboard_object":          config.IdentifierFromProvider,
	"opensearch_dashboard_tenant":          config.IdentifierFromProvider,
	"opensearch_data_stream":               config.IdentifierFromProvider,
	"opensearch_index":                     config.IdentifierFromProvider,
	"opensearch_index_template":            config.IdentifierFromProvider,
	"opensearch_ingest_pipeline":           config.IdentifierFromProvider,
	"opensearch_ism_policy":                config.IdentifierFromProvider,
	"opensearch_ism_policy_mapping":        config.IdentifierFromProvider,
	"opensearch_monitor":                   config.IdentifierFromProvider,
	"opensearch_role":                      config.IdentifierFromProvider,
	"opensearch_roles_mapping":             config.IdentifierFromProvider,
	"opensearch_script":                    config.IdentifierFromProvider,
	"opensearch_sm_policy":                 config.IdentifierFromProvider,
	"opensearch_snapshot_repository":       config.IdentifierFromProvider,
	"opensearch_user":                      config.IdentifierFromProvider,
}

// cliReconciledExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the CLI-based
// architecture for this provider.
var cliReconciledExternalNameConfigs = map[string]config.ExternalName{}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := TerraformPluginSDKExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		} else if e, ok := cliReconciledExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}
