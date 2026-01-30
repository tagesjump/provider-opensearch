package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	clusterv1beta1 "github.com/tagesjump/provider-opensearch/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/tagesjump/provider-opensearch/apis/namespaced/v1beta1"
)

const (
	// Required
	// OpenSearch URL
	url = "url"

	// Optional
	// The access key for use with AWS OpenSearch Service domains
	awsAccessKey = "aws_access_key"
	// Amazon Resource Name of an IAM Role to assume prior to making AWS API calls.
	awsAssumeRoleArn = "aws_assume_role_arn"
	// External ID configured in the IAM policy of the IAM Role to assume prior to making AWS API calls.
	awsAssumeRoleExternalId = "aws_assume_role_external_id"
	// The AWS profile for use with AWS OpenSearch Service domains
	awsProfile = "aws_profile"
	// The AWS region for use in signing of AWS OpenSearch requests. Must be specified in order to use AWS URL signing with AWS OpenSearch endpoint exposed on a custom DNS domain.
	awsRegion = "aws_region"
	// The secret key for use with AWS OpenSearch Service domains
	awsSecretKey = "aws_secret_key"
	// AWS service name used in the credential scope of signed requests to opensearch.
	awsSignatureService = "aws_signature_service"
	// The session token for use with AWS OpenSearch Service domains
	awsToken = "aws_token"
	// A Custom CA certificate
	cacertFile = "cacert_file"
	// A X509 certificate to connect to opensearch
	clientCertPath = "client_cert_path"
	// A X509 key to connect to opensearch
	clientKeyPath = "client_key_path"
	// Set the client healthcheck option for the opensearch client. Healthchecking is designed for direct access to the cluster.
	healthcheck = "healthcheck"
	// If provided, sets the 'Host' header of requests and the 'ServerName' for certificate validation to this value. See the documentation on connecting to opensearch via an SSH tunnel.
	hostOverride = "host_override"
	// Disable SSL verification of API calls
	insecure = "insecure"
	// opensearch Version
	opensearchVersion = "opensearch_version"
	// Password to use to connect to opensearch using basic auth
	password = "password"
	// Proxy URL to use for requests to opensearch.
	proxy = "proxy"
	// Enable signing of AWS opensearch requests. The url must refer to AWS ES domain (*.<region>.es.amazonaws.com), or aws_region must be specified explicitly.
	signAwsRequests = "sign_aws_requests"
	// Set the node sniffing option for the opensearch client. Client won't work with sniffing if nodes are not routable.
	sniff = "sniff"
	// A bearer token or ApiKey for an Authorization header, e.g. Active Directory API key.
	token = "token"
	// The type of token, usually ApiKey or Bearer
	tokenName = "token_name"
	// Username to use to connect to opensearch using basic auth
	username = "username"
	// Version ping timeout in seconds
	versionPingTimeout = "version_ping_timeout"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal opensearch credentials as JSON"
	errNoRequiredFieldURL   = "Missing required field - url"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot resolve provider config")
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// set provider configuration
		ps.Configuration = map[string]any{}

		if value, ok := creds[url]; ok {
			ps.Configuration[url] = value
		} else {
			return ps, errors.New(errNoRequiredFieldURL)
		}

		for _, setting := range []string{
			awsAccessKey, awsAssumeRoleArn, awsAssumeRoleExternalId, awsProfile, awsRegion,
			awsSecretKey, awsSignatureService, awsToken, cacertFile, clientCertPath, clientKeyPath,
			healthcheck, hostOverride, insecure, opensearchVersion, password, proxy, signAwsRequests,
			sniff, token, tokenName, username, versionPingTimeout,
		} {
			if value, ok := creds[setting]; ok {
				ps.Configuration[setting] = value
			}
		}

		return ps, errors.Wrap(configureNoForkOpenSearchClient(ctx, &ps, *tfProvider), "failed to configure the Terraform OpenSearch provider meta")
	}
}

func configureNoForkOpenSearchClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	// Please be aware that this implementation relies on the schema.Provider
	// parameter `p` being a non-pointer. This is because normally
	// the Terraform plugin SDK normally configures the provider
	// only once and using a pointer argument here will cause
	// race conditions between resources referring to different
	// ProviderConfigs.
	diag := p.Configure(ctx, &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}

func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.New(" is not an Object")
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	pcu := &namespacedv1beta1.ProviderConfigUsage{}
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New("unknown provider config type")
	}
	t := resource.NewProviderConfigUsageTracker(crClient, pcu)
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}
