// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information of order cart product options.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mycart = ovh.getOrderCart({
 *     ovhSubsidiary: "fr",
 *     description: "my cart",
 * });
 * const options = mycart.then(mycart => ovh.getOrderCartProductOptions({
 *     cartId: mycart.id,
 *     product: "cloud",
 *     planCode: "project",
 * }));
 * ```
 */
export function getOrderCartProductOptions(args: GetOrderCartProductOptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetOrderCartProductOptionsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("ovh:index/getOrderCartProductOptions:getOrderCartProductOptions", {
        "cartId": args.cartId,
        "catalogName": args.catalogName,
        "planCode": args.planCode,
        "product": args.product,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrderCartProductOptions.
 */
export interface GetOrderCartProductOptionsArgs {
    /**
     * Cart identifier
     */
    cartId: string;
    /**
     * Catalog name
     */
    catalogName?: string;
    /**
     * Product offer identifier
     */
    planCode: string;
    /**
     * Product
     */
    product: string;
}

/**
 * A collection of values returned by getOrderCartProductOptions.
 */
export interface GetOrderCartProductOptionsResult {
    readonly cartId: string;
    readonly catalogName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Product offer identifier
     */
    readonly planCode: string;
    readonly product: string;
    /**
     * products results
     */
    readonly results: outputs.GetOrderCartProductOptionsResult[];
}

export function getOrderCartProductOptionsOutput(args: GetOrderCartProductOptionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrderCartProductOptionsResult> {
    return pulumi.output(args).apply(a => getOrderCartProductOptions(a, opts))
}

/**
 * A collection of arguments for invoking getOrderCartProductOptions.
 */
export interface GetOrderCartProductOptionsOutputArgs {
    /**
     * Cart identifier
     */
    cartId: pulumi.Input<string>;
    /**
     * Catalog name
     */
    catalogName?: pulumi.Input<string>;
    /**
     * Product offer identifier
     */
    planCode: pulumi.Input<string>;
    /**
     * Product
     */
    product: pulumi.Input<string>;
}
