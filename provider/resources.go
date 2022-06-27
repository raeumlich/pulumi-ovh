// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ovh

import (
	"fmt"
	"path/filepath"

	"github.com/ovh/terraform-provider-ovh/ovh"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/raeumlich/pulumi-ovh/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "ovh"
	// modules:
	mainMod = "index" // the ovh module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(ovh.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "ovh",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Pulumi",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing ovh cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "ovh", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/raeumlich/pulumi-ovh",
		// The GitHub Org for the provider - defaults to `terraform-providers`
		GitHubOrg: "ovh",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"ovh_cloud_project": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudProject"),
			},
			"ovh_cloud_project_containerregistry": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudProjectContainerRegistry"),
			},
			"ovh_cloud_project_containerregistry_user": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudProjectContainerRegistryUser"),
			},
			"ovh_cloud_project_failover_ip_attach": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudProjectFailoverIPAttach"),
			},
			"ovh_cloud_project_kube": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudProjectKube"),
			},
			"ovh_cloud_project_kube_nodepool": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudProjectKubeNodepool"),
			},
			"ovh_cloud_project_network_private": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudProjectNetworkPrivate"),
			},
			"ovh_cloud_project_network_private_subnet": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudProjectNetworkPrivateSubnet"),
			},
			"ovh_cloud_project_user": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudProjectUser"),
			},
			"ovh_dbaas_logs_input": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DBaaSLogsInput"),
			},
			"ovh_dbaas_logs_output_graylog_stream": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DBaaSLogsOutputGraylogStream"),
			},
			"ovh_dedicated_ceph_acl": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DedicatedCephACL"),
			},
			"ovh_dedicated_server_install_task": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DedicatedServerInstallTask"),
			},
			"ovh_dedicated_server_reboot_task": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DedicatedServerRebootTask"),
			},
			"ovh_dedicated_server_update": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DedicatedServerUpdate"),
			},
			"ovh_domain_zone": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DomainZone"),
			},
			"ovh_domain_zone_record": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DomainZoneRecord"),
			},
			"ovh_domain_zone_redirection": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DomainZoneRedirection"),
			},
			"ovh_ip_reverse": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPReverse"),
			},
			"ovh_ip_service": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPService"),
			},
			"ovh_iploadbalancing": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancing"),
			},
			"ovh_iploadbalancing_http_farm": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingHTTPFarm"),
			},
			"ovh_iploadbalancing_http_farm_server": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingHTTPFarmServer"),
			},
			"ovh_iploadbalancing_http_frontend": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingHTTPFrontend"),
			},
			"ovh_iploadbalancing_http_route": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingHTTPRoute"),
			},
			"ovh_iploadbalancing_http_route_rule": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingHTTPRouteRule"),
			},
			"ovh_iploadbalancing_refresh": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingRefresh"),
			},
			"ovh_iploadbalancing_tcp_farm": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingTCPFarm"),
			},
			"ovh_iploadbalancing_tcp_farm_server": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingTCPFarmServer"),
			},
			"ovh_iploadbalancing_tcp_frontend": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingTCPFrontend"),
			},
			"ovh_iploadbalancing_tcp_route": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingTCPRoute"),
			},
			"ovh_iploadbalancing_tcp_route_rule": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingTCPRouteRule"),
			},
			"ovh_iploadbalancing_vrack_network": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPLoadBalancingVRackNetwork"),
			},
			"ovh_me_identity_user": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "MeIdentityUser"),
			},
			"ovh_me_installation_template": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "MeInstallationTemplate"),
			},
			"ovh_me_installation_template_partition_scheme": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "MeInstallationTemplatePartitionScheme"),
			},
			"ovh_me_installation_template_partition_scheme_hardware_raid": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "MeInstallationTemplatePartitionSchemeHardwareRAID"),
			},
			"ovh_me_installation_template_partition_scheme_partition": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "MeInstallationTemplatePartitionSchemePartition"),
			},
			"ovh_me_ipxe_script": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IPXEScript"),
			},
			"ovh_me_ssh_key": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "MeSSHKey"),
			},
			"ovh_vrack": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "VRack"),
			},
			"ovh_vrack_cloudproject": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "VRackCloudProject"),
			},
			"ovh_vrack_dedicated_server": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "VRackDedicatedServer"),
			},
			"ovh_vrack_dedicated_server_interface": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "VRackDedicatedServerInterface"),
			},
			"ovh_vrack_ip": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "VRackIP"),
			},
			"ovh_vrack_iploadbalancing": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "VRackIPLoadBalancing"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"ovh_cloud_project_capabilities_containerregistry": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudProjectCapabilitiesContainerRegistry"),
			},
			"ovh_cloud_project_capabilities_containerregistry_filter": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudProjectCapabilitiesContainerregistryFilter"),
			},
			"ovh_cloud_project_containerregistries": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudProjectContainerRegistries"),
			},
			"ovh_cloud_project_containerregistry": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudProjectContainerregistry"),
			},
			"ovh_cloud_project_containerregistry_users": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudProjectContainerRegistryUsers"),
			},
			"ovh_cloud_project_failover_ip_attach": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudProjectFailoverIPAttach"),
			},
			"ovh_cloud_project_kube": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudProjectKube"),
			},
			"ovh_cloud_project_region": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudProjectRegion"),
			},
			"ovh_cloud_project_regions": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudProjectRegions"),
			},
			"ovh_dbaas_logs_input_engine": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDBaaSLogsInputEngine"),
			},
			"ovh_dbaas_logs_output_graylog_stream": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDBaaSLogsOutputGraylogStream"),
			},
			"ovh_dedicated_ceph": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDedicatedCeph"),
			},
			"ovh_dedicated_installation_templates": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDedicatedInstallationTemplates"),
			},
			"ovh_dedicated_server": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDedicatedServer"),
			},
			"ovh_dedicated_server_boots": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDedicatedServerBoots"),
			},
			"ovh_dedicated_servers": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDedicatedServers"),
			},
			"ovh_domain_zone": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDomainZone"),
			},
			"ovh_ip_service": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIPService"),
			},
			"ovh_iploadbalancing": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIPLoadBalancing"),
			},
			"ovh_iploadbalancing_vrack_network": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIPLoadBalancingVRackNetwork"),
			},
			"ovh_iploadbalancing_vrack_networks": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIPLoadBalancingVRackNetworks"),
			},
			"ovh_me": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMe"),
			},
			"ovh_me_identity_user": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMeIdentityUser"),
			},
			"ovh_me_identity_users": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMeIdentityUsers"),
			},
			"ovh_me_installation_template": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMeInstallationTemplate"),
			},
			"ovh_me_installation_templates": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMeInstallationTemplates"),
			},
			"ovh_me_ipxe_script": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMeIPXEScript"),
			},
			"ovh_me_ipxe_scripts": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMeIPXEScripts"),
			},
			"ovh_me_paymentmean_bankaccount": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMePaymentMeanBankAccount"),
			},
			"ovh_me_paymentmean_creditcard": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMePaymentMeanCreditCard"),
			},
			"ovh_me_ssh_key": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMeSSHKey"),
			},
			"ovh_me_ssh_keys": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMeSSHKeys"),
			},
			"ovh_order_cart": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrderCart"),
			},
			"ovh_order_cart_product": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrderCartProduct"),
			},
			"ovh_order_cart_product_options": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrderCartProductOptions"),
			},
			"ovh_order_cart_product_options_plan": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrderCartProductOptionsPlan"),
			},
			"ovh_order_cart_product_plan": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getProductCartProductPlan"),
			},
			"ovh_vps": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVPS"),
			},
			"ovh_vracks": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVRacks"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
