package roles

import (
	"fmt"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/tagesjump/provider-opensearch/config/namespaced/opensearch"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-opensearch/apis/namespaced/roles/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-opensearch/config/namespaced/roles"
)

// Configure adds configurations for the roles group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opensearch_roles_mapping", func(r *config.Resource) {
		r.References["role_name"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", opensearch.ApisPackagePath, "Role"),
		}
		r.References["users"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", opensearch.ApisPackagePath, "User"),
		}
	})

}
