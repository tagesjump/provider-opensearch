package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	detection "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/anomaly/detection"
	config "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/audit/config"
	configuration "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/channel/configuration"
	settings "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/cluster/settings"
	template "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/component/template"
	indextemplate "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/composable/indextemplate"
	object "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/dashboard/object"
	tenant "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/dashboard/tenant"
	stream "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/data/stream"
	templateindex "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/index/template"
	pipeline "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/ingest/pipeline"
	policy "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/ism/policy"
	policymapping "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/ism/policymapping"
	index "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/opensearch/index"
	monitor "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/opensearch/monitor"
	role "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/opensearch/role"
	script "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/opensearch/script"
	user "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/opensearch/user"
	providerconfig "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/providerconfig"
	mapping "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/roles/mapping"
	policysm "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/sm/policy"
	repository "github.com/tagesjump/provider-opensearch/internal/controller/namespaced/snapshot/repository"
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
