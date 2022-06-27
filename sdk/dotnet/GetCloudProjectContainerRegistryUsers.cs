// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetCloudProjectContainerRegistryUsers
    {
        /// <summary>
        /// Use this data source to get the list of users of a container registry associated with a public cloud project.
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
        ///         var registry = Output.Create(Ovh.GetCloudProjectContainerregistry.InvokeAsync(new Ovh.GetCloudProjectContainerregistryArgs
        ///         {
        ///             ServiceName = "XXXXXX",
        ///             RegistryId = "yyyy",
        ///         }));
        ///         var users = Output.Create(Ovh.GetCloudProjectContainerRegistryUsers.InvokeAsync(new Ovh.GetCloudProjectContainerRegistryUsersArgs
        ///         {
        ///             ServiceName = ovh_cloud_project_containerregistry.Registry.Service_name,
        ///             RegistryId = ovh_cloud_project_containerregistry.Registry.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudProjectContainerRegistryUsersResult> InvokeAsync(GetCloudProjectContainerRegistryUsersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectContainerRegistryUsersResult>("ovh:index/getCloudProjectContainerRegistryUsers:getCloudProjectContainerRegistryUsers", args ?? new GetCloudProjectContainerRegistryUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the list of users of a container registry associated with a public cloud project.
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
        ///         var registry = Output.Create(Ovh.GetCloudProjectContainerregistry.InvokeAsync(new Ovh.GetCloudProjectContainerregistryArgs
        ///         {
        ///             ServiceName = "XXXXXX",
        ///             RegistryId = "yyyy",
        ///         }));
        ///         var users = Output.Create(Ovh.GetCloudProjectContainerRegistryUsers.InvokeAsync(new Ovh.GetCloudProjectContainerRegistryUsersArgs
        ///         {
        ///             ServiceName = ovh_cloud_project_containerregistry.Registry.Service_name,
        ///             RegistryId = ovh_cloud_project_containerregistry.Registry.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudProjectContainerRegistryUsersResult> Invoke(GetCloudProjectContainerRegistryUsersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProjectContainerRegistryUsersResult>("ovh:index/getCloudProjectContainerRegistryUsers:getCloudProjectContainerRegistryUsers", args ?? new GetCloudProjectContainerRegistryUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProjectContainerRegistryUsersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Registry ID
        /// </summary>
        [Input("registryId", required: true)]
        public string RegistryId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectContainerRegistryUsersArgs()
        {
        }
    }

    public sealed class GetCloudProjectContainerRegistryUsersInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Registry ID
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCloudProjectContainerRegistryUsersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudProjectContainerRegistryUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string RegistryId;
        /// <summary>
        /// The list of users of the container registry associated with the project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudProjectContainerRegistryUsersResultResult> Results;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetCloudProjectContainerRegistryUsersResult(
            string id,

            string registryId,

            ImmutableArray<Outputs.GetCloudProjectContainerRegistryUsersResultResult> results,

            string serviceName)
        {
            Id = id;
            RegistryId = registryId;
            Results = results;
            ServiceName = serviceName;
        }
    }
}