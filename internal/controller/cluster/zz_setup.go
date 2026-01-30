package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	detection "github.com/tagesjump/provider-opensearch/internal/controller/cluster/anomaly/detection"
	config "github.com/tagesjump/provider-opensearch/internal/controller/cluster/audit/config"
	configuration "github.com/tagesjump/provider-opensearch/internal/controller/cluster/channel/configuration"
	settings "github.com/tagesjump/provider-opensearch/internal/controller/cluster/cluster/settings"
	template "github.com/tagesjump/provider-opensearch/internal/controller/cluster/component/template"
	indextemplate "github.com/tagesjump/provider-opensearch/internal/controller/cluster/composable/indextemplate"
	object "github.com/tagesjump/provider-opensearch/internal/controller/cluster/dashboard/object"
	tenant "github.com/tagesjump/provider-opensearch/internal/controller/cluster/dashboard/tenant"
	stream "github.com/tagesjump/provider-opensearch/internal/controller/cluster/data/stream"
	templateindex "github.com/tagesjump/provider-opensearch/internal/controller/cluster/index/template"
	pipeline "github.com/tagesjump/provider-opensearch/internal/controller/cluster/ingest/pipeline"
	policy "github.com/tagesjump/provider-opensearch/internal/controller/cluster/ism/policy"
	policymapping "github.com/tagesjump/provider-opensearch/internal/controller/cluster/ism/policymapping"
	index "github.com/tagesjump/provider-opensearch/internal/controller/cluster/opensearch/index"
	monitor "github.com/tagesjump/provider-opensearch/internal/controller/cluster/opensearch/monitor"
	role "github.com/tagesjump/provider-opensearch/internal/controller/cluster/opensearch/role"
	script "github.com/tagesjump/provider-opensearch/internal/controller/cluster/opensearch/script"
	user "github.com/tagesjump/provider-opensearch/internal/controller/cluster/opensearch/user"
	providerconfig "github.com/tagesjump/provider-opensearch/internal/controller/cluster/providerconfig"
	mapping "github.com/tagesjump/provider-opensearch/internal/controller/cluster/roles/mapping"
	policysm "github.com/tagesjump/provider-opensearch/internal/controller/cluster/sm/policy"
	repository "github.com/tagesjump/provider-opensearch/internal/controller/cluster/snapshot/repository"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		detection.Setup,
		config.Setup,
		configuration.Setup,
		settings.Setup,
		template.Setup,
		indextemplate.Setup,
		object.Setup,
		tenant.Setup,
		stream.Setup,
		templateindex.Setup,
		pipeline.Setup,
		policy.Setup,
		policymapping.Setup,
		index.Setup,
		monitor.Setup,
		role.Setup,
		script.Setup,
		user.Setup,
		providerconfig.Setup,
		mapping.Setup,
		policysm.Setup,
		repository.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		detection.SetupGated,
		config.SetupGated,
		configuration.SetupGated,
		settings.SetupGated,
		template.SetupGated,
		indextemplate.SetupGated,
		object.SetupGated,
		tenant.SetupGated,
		stream.SetupGated,
		templateindex.SetupGated,
		pipeline.SetupGated,
		policy.SetupGated,
		policymapping.SetupGated,
		index.SetupGated,
		monitor.SetupGated,
		role.SetupGated,
		script.SetupGated,
		user.SetupGated,
		providerconfig.SetupGated,
		mapping.SetupGated,
		policysm.SetupGated,
		repository.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
