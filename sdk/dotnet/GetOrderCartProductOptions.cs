// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOrderCartProductOptions
    {
        /// <summary>
        /// Use this data source to retrieve information of order cart product options.
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
        ///         var mycart = Output.Create(Ovh.GetOrderCart.InvokeAsync(new Ovh.GetOrderCartArgs
        ///         {
        ///             OvhSubsidiary = "fr",
        ///             Description = "my cart",
        ///         }));
        ///         var options = mycart.Apply(mycart =&gt; Output.Create(Ovh.GetOrderCartProductOptions.InvokeAsync(new Ovh.GetOrderCartProductOptionsArgs
        ///         {
        ///             CartId = mycart.Id,
        ///             Product = "cloud",
        ///             PlanCode = "project",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrderCartProductOptionsResult> InvokeAsync(GetOrderCartProductOptionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrderCartProductOptionsResult>("ovh:index/getOrderCartProductOptions:getOrderCartProductOptions", args ?? new GetOrderCartProductOptionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information of order cart product options.
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
        ///         var mycart = Output.Create(Ovh.GetOrderCart.InvokeAsync(new Ovh.GetOrderCartArgs
        ///         {
        ///             OvhSubsidiary = "fr",
        ///             Description = "my cart",
        ///         }));
        ///         var options = mycart.Apply(mycart =&gt; Output.Create(Ovh.GetOrderCartProductOptions.InvokeAsync(new Ovh.GetOrderCartProductOptionsArgs
        ///         {
        ///             CartId = mycart.Id,
        ///             Product = "cloud",
        ///             PlanCode = "project",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOrderCartProductOptionsResult> Invoke(GetOrderCartProductOptionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrderCartProductOptionsResult>("ovh:index/getOrderCartProductOptions:getOrderCartProductOptions", args ?? new GetOrderCartProductOptionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrderCartProductOptionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cart identifier
        /// </summary>
        [Input("cartId", required: true)]
        public string CartId { get; set; } = null!;

        /// <summary>
        /// Catalog name
        /// </summary>
        [Input("catalogName")]
        public string? CatalogName { get; set; }

        /// <summary>
        /// Product offer identifier
        /// </summary>
        [Input("planCode", required: true)]
        public string PlanCode { get; set; } = null!;

        /// <summary>
        /// Product
        /// </summary>
        [Input("product", required: true)]
        public string Product { get; set; } = null!;

        public GetOrderCartProductOptionsArgs()
        {
        }
    }

    public sealed class GetOrderCartProductOptionsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cart identifier
        /// </summary>
        [Input("cartId", required: true)]
        public Input<string> CartId { get; set; } = null!;

        /// <summary>
        /// Catalog name
        /// </summary>
        [Input("catalogName")]
        public Input<string>? CatalogName { get; set; }

        /// <summary>
        /// Product offer identifier
        /// </summary>
        [Input("planCode", required: true)]
        public Input<string> PlanCode { get; set; } = null!;

        /// <summary>
        /// Product
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        public GetOrderCartProductOptionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOrderCartProductOptionsResult
    {
        public readonly string CartId;
        public readonly string? CatalogName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Product offer identifier
        /// </summary>
        public readonly string PlanCode;
        public readonly string Product;
        /// <summary>
        /// products results
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrderCartProductOptionsResultResult> Results;

        [OutputConstructor]
        private GetOrderCartProductOptionsResult(
            string cartId,

            string? catalogName,

            string id,

            string planCode,

            string product,

            ImmutableArray<Outputs.GetOrderCartProductOptionsResultResult> results)
        {
            CartId = cartId;
            CatalogName = catalogName;
            Id = id;
            PlanCode = planCode;
            Product = product;
            Results = results;
        }
    }
}
