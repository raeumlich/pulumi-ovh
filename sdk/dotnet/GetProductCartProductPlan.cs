// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetProductCartProductPlan
    {
        /// <summary>
        /// Use this data source to retrieve information of order cart product plan.
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
        ///         var plan = mycart.Apply(mycart =&gt; Output.Create(Ovh.GetProductCartProductPlan.InvokeAsync(new Ovh.GetProductCartProductPlanArgs
        ///         {
        ///             CartId = mycart.Id,
        ///             PriceCapacity = "renew",
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
        public static Task<GetProductCartProductPlanResult> InvokeAsync(GetProductCartProductPlanArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProductCartProductPlanResult>("ovh:index/getProductCartProductPlan:getProductCartProductPlan", args ?? new GetProductCartProductPlanArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information of order cart product plan.
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
        ///         var plan = mycart.Apply(mycart =&gt; Output.Create(Ovh.GetProductCartProductPlan.InvokeAsync(new Ovh.GetProductCartProductPlanArgs
        ///         {
        ///             CartId = mycart.Id,
        ///             PriceCapacity = "renew",
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
        public static Output<GetProductCartProductPlanResult> Invoke(GetProductCartProductPlanInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProductCartProductPlanResult>("ovh:index/getProductCartProductPlan:getProductCartProductPlan", args ?? new GetProductCartProductPlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProductCartProductPlanArgs : Pulumi.InvokeArgs
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
        /// Capacity of the pricing (type of pricing)
        /// </summary>
        [Input("priceCapacity", required: true)]
        public string PriceCapacity { get; set; } = null!;

        /// <summary>
        /// Product
        /// </summary>
        [Input("product", required: true)]
        public string Product { get; set; } = null!;

        public GetProductCartProductPlanArgs()
        {
        }
    }

    public sealed class GetProductCartProductPlanInvokeArgs : Pulumi.InvokeArgs
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
        /// Capacity of the pricing (type of pricing)
        /// </summary>
        [Input("priceCapacity", required: true)]
        public Input<string> PriceCapacity { get; set; } = null!;

        /// <summary>
        /// Product
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        public GetProductCartProductPlanInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProductCartProductPlanResult
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
        public readonly string PriceCapacity;
        /// <summary>
        /// Prices of the product offer
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProductCartProductPlanPriceResult> Prices;
        public readonly string Product;
        /// <summary>
        /// Name of the product
        /// </summary>
        public readonly string ProductName;
        /// <summary>
        /// Product type
        /// </summary>
        public readonly string ProductType;
        /// <summary>
        /// Selected Price according to capacity
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProductCartProductPlanSelectedPriceResult> SelectedPrices;

        [OutputConstructor]
        private GetProductCartProductPlanResult(
            string cartId,

            string? catalogName,

            string id,

            string planCode,

            string priceCapacity,

            ImmutableArray<Outputs.GetProductCartProductPlanPriceResult> prices,

            string product,

            string productName,

            string productType,

            ImmutableArray<Outputs.GetProductCartProductPlanSelectedPriceResult> selectedPrices)
        {
            CartId = cartId;
            CatalogName = catalogName;
            Id = id;
            PlanCode = planCode;
            PriceCapacity = priceCapacity;
            Prices = prices;
            Product = product;
            ProductName = productName;
            ProductType = productType;
            SelectedPrices = selectedPrices;
        }
    }
}