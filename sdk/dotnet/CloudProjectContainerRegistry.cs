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
    /// Creates a container registry associated with a public cloud project.
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
    ///         var regcap = Output.Create(Ovh.GetCloudProjectCapabilitiesContainerregistryFilter.InvokeAsync(new Ovh.GetCloudProjectCapabilitiesContainerregistryFilterArgs
    ///         {
    ///             ServiceName = "XXXXXX",
    ///             PlanName = "SMALL",
    ///             Region = "GRA",
    ///         }));
    ///         var reg = new Ovh.CloudProjectContainerRegistry("reg", new Ovh.CloudProjectContainerRegistryArgs
    ///         {
    ///             ServiceName = regcap.Apply(regcap =&gt; regcap.ServiceName),
    ///             PlanId = regcap.Apply(regcap =&gt; regcap.Id),
    ///             Region = regcap.Apply(regcap =&gt; regcap.Region),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [OvhResourceType("ovh:index/cloudProjectContainerRegistry:CloudProjectContainerRegistry")]
    public partial class CloudProjectContainerRegistry : Pulumi.CustomResource
    {
        /// <summary>
        /// Plan creation date
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Registry name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Plan ID of the registry
        /// </summary>
        [Output("planId")]
        public Output<string> PlanId { get; private set; } = null!;

        /// <summary>
        /// Plan of the registry
        /// </summary>
        [Output("plans")]
        public Output<ImmutableArray<Outputs.CloudProjectContainerRegistryPlan>> Plans { get; private set; } = null!;

        /// <summary>
        /// Project ID of your registry
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Region of the registry
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Current size of the registry (bytes)
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// Registry status
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Registry last update date
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Access url of the registry
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Version of your registry
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a CloudProjectContainerRegistry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudProjectContainerRegistry(string name, CloudProjectContainerRegistryArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/cloudProjectContainerRegistry:CloudProjectContainerRegistry", name, args ?? new CloudProjectContainerRegistryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudProjectContainerRegistry(string name, Input<string> id, CloudProjectContainerRegistryState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/cloudProjectContainerRegistry:CloudProjectContainerRegistry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CloudProjectContainerRegistry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudProjectContainerRegistry Get(string name, Input<string> id, CloudProjectContainerRegistryState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudProjectContainerRegistry(name, id, state, options);
        }
    }

    public sealed class CloudProjectContainerRegistryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Registry name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Plan ID of the registry
        /// </summary>
        [Input("planId")]
        public Input<string>? PlanId { get; set; }

        /// <summary>
        /// Region of the registry
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public CloudProjectContainerRegistryArgs()
        {
        }
    }

    public sealed class CloudProjectContainerRegistryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Plan creation date
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Registry name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Plan ID of the registry
        /// </summary>
        [Input("planId")]
        public Input<string>? PlanId { get; set; }

        [Input("plans")]
        private InputList<Inputs.CloudProjectContainerRegistryPlanGetArgs>? _plans;

        /// <summary>
        /// Plan of the registry
        /// </summary>
        public InputList<Inputs.CloudProjectContainerRegistryPlanGetArgs> Plans
        {
            get => _plans ?? (_plans = new InputList<Inputs.CloudProjectContainerRegistryPlanGetArgs>());
            set => _plans = value;
        }

        /// <summary>
        /// Project ID of your registry
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Region of the registry
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Current size of the registry (bytes)
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Registry status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Registry last update date
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// Access url of the registry
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Version of your registry
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public CloudProjectContainerRegistryState()
        {
        }
    }
}