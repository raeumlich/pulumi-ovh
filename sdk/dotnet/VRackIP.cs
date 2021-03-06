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
    /// Attach a Ip block to a VRack.
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
    ///         var mycart = Output.Create(Ovh.GetOrderCart.InvokeAsync(new Ovh.GetOrderCartArgs
    ///         {
    ///             OvhSubsidiary = "fr",
    ///             Description = "my cart",
    ///         }));
    ///         var vrackProductCartProductPlan = mycart.Apply(mycart =&gt; Output.Create(Ovh.GetProductCartProductPlan.InvokeAsync(new Ovh.GetProductCartProductPlanArgs
    ///         {
    ///             CartId = mycart.Id,
    ///             PriceCapacity = "renew",
    ///             Product = "vrack",
    ///             PlanCode = "vrack",
    ///         })));
    ///         var vrackVRack = new Ovh.VRack("vrackVRack", new Ovh.VRackArgs
    ///         {
    ///             Description = mycart.Apply(mycart =&gt; mycart.Description),
    ///             OvhSubsidiary = mycart.Apply(mycart =&gt; mycart.OvhSubsidiary),
    ///             PaymentMean = "fidelity",
    ///             Plan = new Ovh.Inputs.VRackPlanArgs
    ///             {
    ///                 Duration = vrackProductCartProductPlan.Apply(vrackProductCartProductPlan =&gt; vrackProductCartProductPlan.SelectedPrices?[0]?.Duration),
    ///                 PlanCode = vrackProductCartProductPlan.Apply(vrackProductCartProductPlan =&gt; vrackProductCartProductPlan.PlanCode),
    ///                 PricingMode = vrackProductCartProductPlan.Apply(vrackProductCartProductPlan =&gt; vrackProductCartProductPlan.SelectedPrices?[0]?.PricingMode),
    ///             },
    ///         });
    ///         var ipblockProductCartProductPlan = mycart.Apply(mycart =&gt; Output.Create(Ovh.GetProductCartProductPlan.InvokeAsync(new Ovh.GetProductCartProductPlanArgs
    ///         {
    ///             CartId = mycart.Id,
    ///             PriceCapacity = "renew",
    ///             Product = "ip",
    ///             PlanCode = "ip-v4-s30-ripe",
    ///         })));
    ///         var ipblockIPService = new Ovh.IPService("ipblockIPService", new Ovh.IPServiceArgs
    ///         {
    ///             OvhSubsidiary = mycart.Apply(mycart =&gt; mycart.OvhSubsidiary),
    ///             PaymentMean = "ovh-account",
    ///             Description = mycart.Apply(mycart =&gt; mycart.Description),
    ///             Plan = new Ovh.Inputs.IPServicePlanArgs
    ///             {
    ///                 Duration = ipblockProductCartProductPlan.Apply(ipblockProductCartProductPlan =&gt; ipblockProductCartProductPlan.SelectedPrices?[0]?.Duration),
    ///                 PlanCode = ipblockProductCartProductPlan.Apply(ipblockProductCartProductPlan =&gt; ipblockProductCartProductPlan.PlanCode),
    ///                 PricingMode = ipblockProductCartProductPlan.Apply(ipblockProductCartProductPlan =&gt; ipblockProductCartProductPlan.SelectedPrices?[0]?.PricingMode),
    ///                 Configurations = 
    ///                 {
    ///                     new Ovh.Inputs.IPServicePlanConfigurationArgs
    ///                     {
    ///                         Label = "country",
    ///                         Value = "FR",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var vrackblock = new Ovh.VRackIP("vrackblock", new Ovh.VRackIPArgs
    ///         {
    ///             ServiceName = vrackVRack.ServiceName,
    ///             Block = ipblockIPService.Ip,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [OvhResourceType("ovh:index/vRackIP:VRackIP")]
    public partial class VRackIP : Pulumi.CustomResource
    {
        /// <summary>
        /// Your IP block.
        /// </summary>
        [Output("block")]
        public Output<string> Block { get; private set; } = null!;

        /// <summary>
        /// Your gateway
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// Your IP block
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Where you want your block announced on the network
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a VRackIP resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VRackIP(string name, VRackIPArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/vRackIP:VRackIP", name, args ?? new VRackIPArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VRackIP(string name, Input<string> id, VRackIPState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/vRackIP:VRackIP", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VRackIP resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VRackIP Get(string name, Input<string> id, VRackIPState? state = null, CustomResourceOptions? options = null)
        {
            return new VRackIP(name, id, state, options);
        }
    }

    public sealed class VRackIPArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Your IP block.
        /// </summary>
        [Input("block", required: true)]
        public Input<string> Block { get; set; } = null!;

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public VRackIPArgs()
        {
        }
    }

    public sealed class VRackIPState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Your IP block.
        /// </summary>
        [Input("block")]
        public Input<string>? Block { get; set; }

        /// <summary>
        /// Your gateway
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Your IP block
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Where you want your block announced on the network
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VRackIPState()
        {
        }
    }
}
