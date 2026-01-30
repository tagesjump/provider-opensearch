package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	opensearch "github.com/opensearch-project/terraform-provider-opensearch/provider"

	"github.com/crossplane/upjet/v2/pkg/registry/reference"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	rolesClustered "github.com/tagesjump/provider-opensearch/config/cluster/roles"
	namespacedClustered "github.com/tagesjump/provider-opensearch/config/namespaced/roles"
)

const (
	resourcePrefix = "opensearch"
	modulePath     = "github.com/tagesjump/provider-opensearch"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// workaround for the TF Google v2.41.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(generationProvider bool) (*ujconfig.Provider, error) {
	var providerInstance *schema.Provider
	var err error
	if generationProvider {
		providerInstance, err = getProviderSchema(providerSchema)
	} else {
		providerInstance = opensearch.Provider()
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("opensearch.upbound.io"),
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformProvider(providerInstance),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(TerraformPluginSDKExternalNameConfigs)),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		rolesClustered.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// GetProvider returns provider configuration
func GetProviderNamespaced(generationProvider bool) (*ujconfig.Provider, error) {
	var providerInstance *schema.Provider
	var err error
	if generationProvider {
		providerInstance, err = getProviderSchema(providerSchema)
	} else {
		providerInstance = opensearch.Provider()
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("opensearch.m.upbound.io"),
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformProvider(providerInstance),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(TerraformPluginSDKExternalNameConfigs)),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		namespacedClustered.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

func resourceList(t map[string]ujconfig.ExternalName) []string {
	l := make([]string, len(t))
	i := 0
	for n := range t {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}
