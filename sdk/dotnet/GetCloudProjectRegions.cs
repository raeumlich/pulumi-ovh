// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetCloudProjectRegions
    {
        /// <summary>
        /// Use this data source to get the regions of a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var regions = Output.Create(Ovh.GetCloudProjectRegions.InvokeAsync(new Ovh.GetCloudProjectRegionsArgs
        ///         {
        ///             HasServicesUps = 
        ///             {
        ///                 "network",
        ///             },
        ///             ServiceName = "XXXXXX",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudProjectRegionsResult> InvokeAsync(GetCloudProjectRegionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectRegionsResult>("ovh:index/getCloudProjectRegions:getCloudProjectRegions", args ?? new GetCloudProjectRegionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the regions of a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var regions = Output.Create(Ovh.GetCloudProjectRegions.InvokeAsync(new Ovh.GetCloudProjectRegionsArgs
        ///         {
        ///             HasServicesUps = 
        ///             {
        ///                 "network",
        ///             },
        ///             ServiceName = "XXXXXX",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudProjectRegionsResult> Invoke(GetCloudProjectRegionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProjectRegionsResult>("ovh:index/getCloudProjectRegions:getCloudProjectRegions", args ?? new GetCloudProjectRegionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProjectRegionsArgs : Pulumi.InvokeArgs
    {
        [Input("hasServicesUps")]
        private List<string>? _hasServicesUps;

        /// <summary>
        /// List of services which has to be UP in regions.
        /// Example: "image", "instance", "network", "storage", "volume", "workflow", ...
        /// If left blank, returns all regions associated with the service_name.
        /// </summary>
        public List<string> HasServicesUps
        {
            get => _hasServicesUps ?? (_hasServicesUps = new List<string>());
            set => _hasServicesUps = value;
        }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectRegionsArgs()
        {
        }
    }

    public sealed class GetCloudProjectRegionsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("hasServicesUps")]
        private InputList<string>? _hasServicesUps;

        /// <summary>
        /// List of services which has to be UP in regions.
        /// Example: "image", "instance", "network", "storage", "volume", "workflow", ...
        /// If left blank, returns all regions associated with the service_name.
        /// </summary>
        public InputList<string> HasServicesUps
        {
            get => _hasServicesUps ?? (_hasServicesUps = new InputList<string>());
            set => _hasServicesUps = value;
        }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCloudProjectRegionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudProjectRegionsResult
    {
        public readonly ImmutableArray<string> HasServicesUps;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of regions associated with the project, filtered by services UP.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetCloudProjectRegionsResult(
            ImmutableArray<string> hasServicesUps,

            string id,

            ImmutableArray<string> names,

            string serviceName)
        {
            HasServicesUps = hasServicesUps;
            Id = id;
            Names = names;
            ServiceName = serviceName;
        }
    }
}
