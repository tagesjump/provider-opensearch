// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	detection "github.com/tagesjump/provider-opensearch/internal/controller/anomaly/detection"
	config "github.com/tagesjump/provider-opensearch/internal/controller/audit/config"
	configuration "github.com/tagesjump/provider-opensearch/internal/controller/channel/configuration"
	settings "github.com/tagesjump/provider-opensearch/internal/controller/cluster/settings"
	template "github.com/tagesjump/provider-opensearch/internal/controller/component/template"
	indextemplate "github.com/tagesjump/provider-opensearch/internal/controller/composable/indextemplate"
	object "github.com/tagesjump/provider-opensearch/internal/controller/dashboard/object"
	tenant "github.com/tagesjump/provider-opensearch/internal/controller/dashboard/tenant"
	stream "github.com/tagesjump/provider-opensearch/internal/controller/data/stream"
	templateindex "github.com/tagesjump/provider-opensearch/internal/controller/index/template"
	pipeline "github.com/tagesjump/provider-opensearch/internal/controller/ingest/pipeline"
	policy "github.com/tagesjump/provider-opensearch/internal/controller/ism/policy"
	policymapping "github.com/tagesjump/provider-opensearch/internal/controller/ism/policymapping"
	index "github.com/tagesjump/provider-opensearch/internal/controller/opensearch/index"
	monitor "github.com/tagesjump/provider-opensearch/internal/controller/opensearch/monitor"
	role "github.com/tagesjump/provider-opensearch/internal/controller/opensearch/role"
	script "github.com/tagesjump/provider-opensearch/internal/controller/opensearch/script"
	user "github.com/tagesjump/provider-opensearch/internal/controller/opensearch/user"
	providerconfig "github.com/tagesjump/provider-opensearch/internal/controller/providerconfig"
	mapping "github.com/tagesjump/provider-opensearch/internal/controller/roles/mapping"
	policysm "github.com/tagesjump/provider-opensearch/internal/controller/sm/policy"
	repository "github.com/tagesjump/provider-opensearch/internal/controller/snapshot/repository"
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
