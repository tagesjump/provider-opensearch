package opensearch

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-opensearch/apis/cluster/opensearch/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-opensearch/config/cluster/opensearch"
)

// Configure adds configurations for the opensearch group.
func Configure(p *config.Provider) {}
