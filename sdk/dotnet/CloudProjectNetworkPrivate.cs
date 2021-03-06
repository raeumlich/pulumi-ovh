// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    /// <summary>
    /// Creates a private network in a public cloud project.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var net = new Ovh.CloudProjectNetworkPrivate("net", new Ovh.CloudProjectNetworkPrivateArgs
    ///         {
    ///             Regions = 
    ///             {
    ///                 "GRA1",
    ///                 "BHS1",
    ///             },
    ///             ServiceName = "XXXXXX",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [OvhResourceType("ovh:index/cloudProjectNetworkPrivate:CloudProjectNetworkPrivate")]
    public partial class CloudProjectNetworkPrivate : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the network.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// an array of valid OVH public cloud region ID in which the network
        /// will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        /// </summary>
        [Output("regions")]
        public Output<ImmutableArray<string>> Regions { get; private set; } = null!;

        /// <summary>
        /// A map representing information about the region.
        /// * `regions_attributes/region` - The id of the region.
        /// * `regions_attributes/status` - The status of the network in the region.
        /// * `regions_attributes/openstackid` - The private network id in the region.
        /// </summary>
        [Output("regionsAttributes")]
        public Output<ImmutableArray<Outputs.CloudProjectNetworkPrivateRegionsAttribute>> RegionsAttributes { get; private set; } = null!;

        /// <summary>
        /// (Deprecated) A map representing the status of the network per region.
        /// * `regions_status/region` - (Deprecated) The id of the region.
        /// * `regions_status/status` - (Deprecated) The status of the network in the region.
        /// </summary>
        [Output("regionsStatuses")]
        public Output<ImmutableArray<Outputs.CloudProjectNetworkPrivateRegionsStatus>> RegionsStatuses { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// the status of the network. should be normally set to 'ACTIVE'.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// the type of the network. Either 'private' or 'public'.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// a vlan id to associate with the network.
        /// Changing this value recreates the resource. Defaults to 0.
        /// </summary>
        [Output("vlanId")]
        public Output<int?> VlanId { get; private set; } = null!;


        /// <summary>
        /// Create a CloudProjectNetworkPrivate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudProjectNetworkPrivate(string name, CloudProjectNetworkPrivateArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/cloudProjectNetworkPrivate:CloudProjectNetworkPrivate", name, args ?? new CloudProjectNetworkPrivateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudProjectNetworkPrivate(string name, Input<string> id, CloudProjectNetworkPrivateState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/cloudProjectNetworkPrivate:CloudProjectNetworkPrivate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudProjectNetworkPrivate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudProjectNetworkPrivate Get(string name, Input<string> id, CloudProjectNetworkPrivateState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudProjectNetworkPrivate(name, id, state, options);
        }
    }

    public sealed class CloudProjectNetworkPrivateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// an array of valid OVH public cloud region ID in which the network
        /// will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// a vlan id to associate with the network.
        /// Changing this value recreates the resource. Defaults to 0.
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        public CloudProjectNetworkPrivateArgs()
        {
        }
    }

    public sealed class CloudProjectNetworkPrivateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// an array of valid OVH public cloud region ID in which the network
        /// will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        [Input("regionsAttributes")]
        private InputList<Inputs.CloudProjectNetworkPrivateRegionsAttributeGetArgs>? _regionsAttributes;

        /// <summary>
        /// A map representing information about the region.
        /// * `regions_attributes/region` - The id of the region.
        /// * `regions_attributes/status` - The status of the network in the region.
        /// * `regions_attributes/openstackid` - The private network id in the region.
        /// </summary>
        public InputList<Inputs.CloudProjectNetworkPrivateRegionsAttributeGetArgs> RegionsAttributes
        {
            get => _regionsAttributes ?? (_regionsAttributes = new InputList<Inputs.CloudProjectNetworkPrivateRegionsAttributeGetArgs>());
            set => _regionsAttributes = value;
        }

        [Input("regionsStatuses")]
        private InputList<Inputs.CloudProjectNetworkPrivateRegionsStatusGetArgs>? _regionsStatuses;

        /// <summary>
        /// (Deprecated) A map representing the status of the network per region.
        /// * `regions_status/region` - (Deprecated) The id of the region.
        /// * `regions_status/status` - (Deprecated) The status of the network in the region.
        /// </summary>
        [Obsolete(@"use the regions_attributes field instead")]
        public InputList<Inputs.CloudProjectNetworkPrivateRegionsStatusGetArgs> RegionsStatuses
        {
            get => _regionsStatuses ?? (_regionsStatuses = new InputList<Inputs.CloudProjectNetworkPrivateRegionsStatusGetArgs>());
            set => _regionsStatuses = value;
        }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// the status of the network. should be normally set to 'ACTIVE'.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// the type of the network. Either 'private' or 'public'.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// a vlan id to associate with the network.
        /// Changing this value recreates the resource. Defaults to 0.
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        public CloudProjectNetworkPrivateState()
        {
        }
    }
}
