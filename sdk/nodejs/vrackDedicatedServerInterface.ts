// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Attach a Dedicated Server Network Interface to a VRack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const vdsi = new ovh.VRackDedicatedServerInterface("vdsi", {
 *     interfaceId: "67890",
 *     serviceName: "12345",
 * });
 * ```
 */
export class VRackDedicatedServerInterface extends pulumi.CustomResource {
    /**
     * Get an existing VRackDedicatedServerInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VRackDedicatedServerInterfaceState, opts?: pulumi.CustomResourceOptions): VRackDedicatedServerInterface {
        return new VRackDedicatedServerInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/vRackDedicatedServerInterface:VRackDedicatedServerInterface';

    /**
     * Returns true if the given object is an instance of VRackDedicatedServerInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VRackDedicatedServerInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VRackDedicatedServerInterface.__pulumiType;
    }

    /**
     * The id of dedicated server network interface.
     */
    public readonly interfaceId!: pulumi.Output<string>;
    /**
     * The id of the vrack. If omitted,
     * the `OVH_VRACK_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a VRackDedicatedServerInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VRackDedicatedServerInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VRackDedicatedServerInterfaceArgs | VRackDedicatedServerInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VRackDedicatedServerInterfaceState | undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as VRackDedicatedServerInterfaceArgs | undefined;
            if ((!args || args.interfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaceId'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VRackDedicatedServerInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VRackDedicatedServerInterface resources.
 */
export interface VRackDedicatedServerInterfaceState {
    /**
     * The id of dedicated server network interface.
     */
    interfaceId?: pulumi.Input<string>;
    /**
     * The id of the vrack. If omitted,
     * the `OVH_VRACK_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VRackDedicatedServerInterface resource.
 */
export interface VRackDedicatedServerInterfaceArgs {
    /**
     * The id of dedicated server network interface.
     */
    interfaceId: pulumi.Input<string>;
    /**
     * The id of the vrack. If omitted,
     * the `OVH_VRACK_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
