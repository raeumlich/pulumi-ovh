// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Outputs
{

    [OutputType]
    public sealed class GetOrderCartProductOptionsResultResult
    {
        /// <summary>
        /// Define if options of this family are exclusive with each other
        /// </summary>
        public readonly bool Exclusive;
        /// <summary>
        /// Option family
        /// </summary>
        public readonly string Family;
        /// <summary>
        /// Define if an option of this family is mandatory
        /// </summary>
        public readonly bool Mandatory;
        /// <summary>
        /// Product offer identifier
        /// </summary>
        public readonly string PlanCode;
        /// <summary>
        /// Prices of the product offer
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrderCartProductOptionsResultPriceResult> Prices;
        /// <summary>
        /// Name of the product
        /// </summary>
        public readonly string ProductName;
        /// <summary>
        /// Product type
        /// </summary>
        public readonly string ProductType;

        [OutputConstructor]
        private GetOrderCartProductOptionsResultResult(
            bool exclusive,

            string family,

            bool mandatory,

            string planCode,

            ImmutableArray<Outputs.GetOrderCartProductOptionsResultPriceResult> prices,

            string productName,

            string productType)
        {
            Exclusive = exclusive;
            Family = family;
            Mandatory = mandatory;
            PlanCode = planCode;
            Prices = prices;
            ProductName = productName;
            ProductType = productType;
        }
    }
}
