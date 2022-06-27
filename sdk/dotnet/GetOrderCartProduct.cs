// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOrderCartProduct
    {
        /// <summary>
        /// Use this data source to retrieve information of order cart product products.
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
        ///         var plans = mycart.Apply(mycart =&gt; Output.Create(Ovh.GetOrderCartProduct.InvokeAsync(new Ovh.GetOrderCartProductArgs
        ///         {
        ///             CartId = mycart.Id,
        ///             Product = "...",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrderCartProductResult> InvokeAsync(GetOrderCartProductArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrderCartProductResult>("ovh:index/getOrderCartProduct:getOrderCartProduct", args ?? new GetOrderCartProductArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information of order cart product products.
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
        ///         var plans = mycart.Apply(mycart =&gt; Output.Create(Ovh.GetOrderCartProduct.InvokeAsync(new Ovh.GetOrderCartProductArgs
        ///         {
        ///             CartId = mycart.Id,
        ///             Product = "...",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOrderCartProductResult> Invoke(GetOrderCartProductInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrderCartProductResult>("ovh:index/getOrderCartProduct:getOrderCartProduct", args ?? new GetOrderCartProductInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrderCartProductArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cart identifier
        /// </summary>
        [Input("cartId", required: true)]
        public string CartId { get; set; } = null!;

        /// <summary>
        /// product
        /// </summary>
        [Input("product", required: true)]
        public string Product { get; set; } = null!;

        public GetOrderCartProductArgs()
        {
        }
    }

    public sealed class GetOrderCartProductInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cart identifier
        /// </summary>
        [Input("cartId", required: true)]
        public Input<string> CartId { get; set; } = null!;

        /// <summary>
        /// product
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        public GetOrderCartProductInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOrderCartProductResult
    {
        public readonly string CartId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Product;
        /// <summary>
        /// products results
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrderCartProductResultResult> Results;

        [OutputConstructor]
        private GetOrderCartProductResult(
            string cartId,

            string id,

            string product,

            ImmutableArray<Outputs.GetOrderCartProductResultResult> results)
        {
            CartId = cartId;
            Id = id;
            Product = product;
            Results = results;
        }
    }
}