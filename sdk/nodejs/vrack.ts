// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Orders a vrack.
 *
 * ## Important
 *
 * This resource is in beta state. Use with caution.
 *
 * __NOTE__ : Currently, the OVH api doesn't support Vrack termination. You have to open a support ticket to ask for vrack termination. Otherwise, you may hit vrack quota issues.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mycart = ovh.getOrderCart({
 *     ovhSubsidiary: "fr",
 *     description: "my vrack order cart",
 * });
 * const vrackProductCartProductPlan = mycart.then(mycart => ovh.getProductCartProductPlan({
 *     cartId: mycart.id,
 *     priceCapacity: "renew",
 *     product: "vrack",
 *     planCode: "vrack",
 * }));
 * const vrackVRack = new ovh.VRack("vrackVRack", {
 *     ovhSubsidiary: mycart.then(mycart => mycart.ovhSubsidiary),
 *     paymentMean: "fidelity",
 *     description: "my vrack",
 *     plan: {
 *         duration: vrackProductCartProductPlan.then(vrackProductCartProductPlan => vrackProductCartProductPlan.selectedPrices?[0]?.duration),
 *         planCode: vrackProductCartProductPlan.then(vrackProductCartProductPlan => vrackProductCartProductPlan.planCode),
 *         pricingMode: vrackProductCartProductPlan.then(vrackProductCartProductPlan => vrackProductCartProductPlan.selectedPrices?[0]?.pricingMode),
 *     },
 * });
 * ```
 */
export class VRack extends pulumi.CustomResource {
    /**
     * Get an existing VRack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VRackState, opts?: pulumi.CustomResourceOptions): VRack {
        return new VRack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/vRack:VRack';

    /**
     * Returns true if the given object is an instance of VRack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VRack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VRack.__pulumiType;
    }

    /**
     * yourvrackdescription
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * yourvrackname
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Details about an Order
     */
    public /*out*/ readonly orders!: pulumi.Output<outputs.VRackOrder[]>;
    /**
     * Ovh Subsidiary
     */
    public readonly ovhSubsidiary!: pulumi.Output<string>;
    /**
     * Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
     */
    public readonly paymentMean!: pulumi.Output<string>;
    /**
     * Product Plan to order
     */
    public readonly plan!: pulumi.Output<outputs.VRackPlan>;
    /**
     * Product Plan to order
     */
    public readonly planOptions!: pulumi.Output<outputs.VRackPlanOption[] | undefined>;
    /**
     * The internal name of your vrack
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a VRack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VRackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VRackArgs | VRackState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VRackState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orders"] = state ? state.orders : undefined;
            resourceInputs["ovhSubsidiary"] = state ? state.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = state ? state.paymentMean : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["planOptions"] = state ? state.planOptions : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as VRackArgs | undefined;
            if ((!args || args.ovhSubsidiary === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ovhSubsidiary'");
            }
            if ((!args || args.paymentMean === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentMean'");
            }
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ovhSubsidiary"] = args ? args.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = args ? args.paymentMean : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["planOptions"] = args ? args.planOptions : undefined;
            resourceInputs["orders"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VRack.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VRack resources.
 */
export interface VRackState {
    /**
     * yourvrackdescription
     */
    description?: pulumi.Input<string>;
    /**
     * yourvrackname
     */
    name?: pulumi.Input<string>;
    /**
     * Details about an Order
     */
    orders?: pulumi.Input<pulumi.Input<inputs.VRackOrder>[]>;
    /**
     * Ovh Subsidiary
     */
    ovhSubsidiary?: pulumi.Input<string>;
    /**
     * Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
     */
    paymentMean?: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan?: pulumi.Input<inputs.VRackPlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.VRackPlanOption>[]>;
    /**
     * The internal name of your vrack
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VRack resource.
 */
export interface VRackArgs {
    /**
     * yourvrackdescription
     */
    description?: pulumi.Input<string>;
    /**
     * yourvrackname
     */
    name?: pulumi.Input<string>;
    /**
     * Ovh Subsidiary
     */
    ovhSubsidiary: pulumi.Input<string>;
    /**
     * Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
     */
    paymentMean: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan: pulumi.Input<inputs.VRackPlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.VRackPlanOption>[]>;
}
