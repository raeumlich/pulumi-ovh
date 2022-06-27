// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates a private network in a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const net = new ovh.CloudProjectNetworkPrivate("net", {
 *     regions: [
 *         "GRA1",
 *         "BHS1",
 *     ],
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export class CloudProjectNetworkPrivate extends pulumi.CustomResource {
    /**
     * Get an existing CloudProjectNetworkPrivate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudProjectNetworkPrivateState, opts?: pulumi.CustomResourceOptions): CloudProjectNetworkPrivate {
        return new CloudProjectNetworkPrivate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/cloudProjectNetworkPrivate:CloudProjectNetworkPrivate';

    /**
     * Returns true if the given object is an instance of CloudProjectNetworkPrivate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudProjectNetworkPrivate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudProjectNetworkPrivate.__pulumiType;
    }

    /**
     * The name of the network.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * an array of valid OVH public cloud region ID in which the network
     * will be available. Ex.: "GRA1". Defaults to all public cloud regions.
     */
    public readonly regions!: pulumi.Output<string[]>;
    /**
     * A map representing information about the region.
     * * `regions_attributes/region` - The id of the region.
     * * `regions_attributes/status` - The status of the network in the region.
     * * `regions_attributes/openstackid` - The private network id in the region.
     */
    public /*out*/ readonly regionsAttributes!: pulumi.Output<outputs.CloudProjectNetworkPrivateRegionsAttribute[]>;
    /**
     * (Deprecated) A map representing the status of the network per region.
     * * `regions_status/region` - (Deprecated) The id of the region.
     * * `regions_status/status` - (Deprecated) The status of the network in the region.
     *
     * @deprecated use the regions_attributes field instead
     */
    public /*out*/ readonly regionsStatuses!: pulumi.Output<outputs.CloudProjectNetworkPrivateRegionsStatus[]>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * the status of the network. should be normally set to 'ACTIVE'.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * the type of the network. Either 'private' or 'public'.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * a vlan id to associate with the network.
     * Changing this value recreates the resource. Defaults to 0.
     */
    public readonly vlanId!: pulumi.Output<number | undefined>;

    /**
     * Create a CloudProjectNetworkPrivate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudProjectNetworkPrivateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudProjectNetworkPrivateArgs | CloudProjectNetworkPrivateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudProjectNetworkPrivateState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["regions"] = state ? state.regions : undefined;
            resourceInputs["regionsAttributes"] = state ? state.regionsAttributes : undefined;
            resourceInputs["regionsStatuses"] = state ? state.regionsStatuses : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vlanId"] = state ? state.vlanId : undefined;
        } else {
            const args = argsOrState as CloudProjectNetworkPrivateArgs | undefined;
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regions"] = args ? args.regions : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["vlanId"] = args ? args.vlanId : undefined;
            resourceInputs["regionsAttributes"] = undefined /*out*/;
            resourceInputs["regionsStatuses"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CloudProjectNetworkPrivate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudProjectNetworkPrivate resources.
 */
export interface CloudProjectNetworkPrivateState {
    /**
     * The name of the network.
     */
    name?: pulumi.Input<string>;
    /**
     * an array of valid OVH public cloud region ID in which the network
     * will be available. Ex.: "GRA1". Defaults to all public cloud regions.
     */
    regions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map representing information about the region.
     * * `regions_attributes/region` - The id of the region.
     * * `regions_attributes/status` - The status of the network in the region.
     * * `regions_attributes/openstackid` - The private network id in the region.
     */
    regionsAttributes?: pulumi.Input<pulumi.Input<inputs.CloudProjectNetworkPrivateRegionsAttribute>[]>;
    /**
     * (Deprecated) A map representing the status of the network per region.
     * * `regions_status/region` - (Deprecated) The id of the region.
     * * `regions_status/status` - (Deprecated) The status of the network in the region.
     *
     * @deprecated use the regions_attributes field instead
     */
    regionsStatuses?: pulumi.Input<pulumi.Input<inputs.CloudProjectNetworkPrivateRegionsStatus>[]>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * the status of the network. should be normally set to 'ACTIVE'.
     */
    status?: pulumi.Input<string>;
    /**
     * the type of the network. Either 'private' or 'public'.
     */
    type?: pulumi.Input<string>;
    /**
     * a vlan id to associate with the network.
     * Changing this value recreates the resource. Defaults to 0.
     */
    vlanId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CloudProjectNetworkPrivate resource.
 */
export interface CloudProjectNetworkPrivateArgs {
    /**
     * The name of the network.
     */
    name?: pulumi.Input<string>;
    /**
     * an array of valid OVH public cloud region ID in which the network
     * will be available. Ex.: "GRA1". Defaults to all public cloud regions.
     */
    regions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
    /**
     * a vlan id to associate with the network.
     * Changing this value recreates the resource. Defaults to 0.
     */
    vlanId?: pulumi.Input<number>;
}