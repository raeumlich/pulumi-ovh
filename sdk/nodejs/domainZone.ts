// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates a domain zone.
 *
 * ## Important
 *
 * This resource is in beta state. Use with caution.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mycart = ovh.getOrderCart({
 *     ovhSubsidiary: "fr",
 * });
 * const zoneProductCartProductPlan = mycart.then(mycart => ovh.getProductCartProductPlan({
 *     cartId: mycart.id,
 *     priceCapacity: "renew",
 *     product: "dns",
 *     planCode: "zone",
 * }));
 * const zoneDomainZone = new ovh.DomainZone("zoneDomainZone", {
 *     ovhSubsidiary: mycart.then(mycart => mycart.ovhSubsidiary),
 *     paymentMean: "fidelity",
 *     plan: {
 *         duration: zoneProductCartProductPlan.then(zoneProductCartProductPlan => zoneProductCartProductPlan.selectedPrices?[0]?.duration),
 *         planCode: zoneProductCartProductPlan.then(zoneProductCartProductPlan => zoneProductCartProductPlan.planCode),
 *         pricingMode: zoneProductCartProductPlan.then(zoneProductCartProductPlan => zoneProductCartProductPlan.selectedPrices?[0]?.pricingMode),
 *         configurations: [
 *             {
 *                 label: "zone",
 *                 value: "myzone.mydomain.com",
 *             },
 *             {
 *                 label: "template",
 *                 value: "minimized",
 *             },
 *         ],
 *     },
 * });
 * ```
 */
export class DomainZone extends pulumi.CustomResource {
    /**
     * Get an existing DomainZone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainZoneState, opts?: pulumi.CustomResourceOptions): DomainZone {
        return new DomainZone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/domainZone:DomainZone';

    /**
     * Returns true if the given object is an instance of DomainZone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainZone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainZone.__pulumiType;
    }

    /**
     * Is DNSSEC supported by this zone
     */
    public /*out*/ readonly dnssecSupported!: pulumi.Output<boolean>;
    /**
     * hasDnsAnycast flag of the DNS zone
     */
    public /*out*/ readonly hasDnsAnycast!: pulumi.Output<boolean>;
    /**
     * Last update date of the DNS zone
     */
    public /*out*/ readonly lastUpdate!: pulumi.Output<string>;
    /**
     * Zone name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Name servers that host the DNS zone
     */
    public /*out*/ readonly nameServers!: pulumi.Output<string[]>;
    /**
     * Details about an Order
     */
    public /*out*/ readonly orders!: pulumi.Output<outputs.DomainZoneOrder[]>;
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
    public readonly plan!: pulumi.Output<outputs.DomainZonePlan>;
    /**
     * Product Plan to order
     */
    public readonly planOptions!: pulumi.Output<outputs.DomainZonePlanOption[] | undefined>;

    /**
     * Create a DomainZone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainZoneArgs | DomainZoneState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainZoneState | undefined;
            resourceInputs["dnssecSupported"] = state ? state.dnssecSupported : undefined;
            resourceInputs["hasDnsAnycast"] = state ? state.hasDnsAnycast : undefined;
            resourceInputs["lastUpdate"] = state ? state.lastUpdate : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nameServers"] = state ? state.nameServers : undefined;
            resourceInputs["orders"] = state ? state.orders : undefined;
            resourceInputs["ovhSubsidiary"] = state ? state.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = state ? state.paymentMean : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["planOptions"] = state ? state.planOptions : undefined;
        } else {
            const args = argsOrState as DomainZoneArgs | undefined;
            if ((!args || args.ovhSubsidiary === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ovhSubsidiary'");
            }
            if ((!args || args.paymentMean === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentMean'");
            }
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            resourceInputs["ovhSubsidiary"] = args ? args.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = args ? args.paymentMean : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["planOptions"] = args ? args.planOptions : undefined;
            resourceInputs["dnssecSupported"] = undefined /*out*/;
            resourceInputs["hasDnsAnycast"] = undefined /*out*/;
            resourceInputs["lastUpdate"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nameServers"] = undefined /*out*/;
            resourceInputs["orders"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainZone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainZone resources.
 */
export interface DomainZoneState {
    /**
     * Is DNSSEC supported by this zone
     */
    dnssecSupported?: pulumi.Input<boolean>;
    /**
     * hasDnsAnycast flag of the DNS zone
     */
    hasDnsAnycast?: pulumi.Input<boolean>;
    /**
     * Last update date of the DNS zone
     */
    lastUpdate?: pulumi.Input<string>;
    /**
     * Zone name
     */
    name?: pulumi.Input<string>;
    /**
     * Name servers that host the DNS zone
     */
    nameServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Details about an Order
     */
    orders?: pulumi.Input<pulumi.Input<inputs.DomainZoneOrder>[]>;
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
    plan?: pulumi.Input<inputs.DomainZonePlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.DomainZonePlanOption>[]>;
}

/**
 * The set of arguments for constructing a DomainZone resource.
 */
export interface DomainZoneArgs {
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
    plan: pulumi.Input<inputs.DomainZonePlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.DomainZonePlanOption>[]>;
}