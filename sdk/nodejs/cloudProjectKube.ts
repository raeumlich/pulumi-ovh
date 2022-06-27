// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates a OVH Managed Kubernetes Service cluster in a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mykube = new ovh.CloudProjectKube("mykube", {
 *     region: "GRA7",
 *     serviceName: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 * });
 * ```
 *
 * ## Import
 *
 * OVHcloud Managed Kubernetes Service clusters can be imported using the `serviceName` and the `id` of the cluster, separated by "/" E.g.,
 *
 * ```sh
 *  $ pulumi import ovh:index/cloudProjectKube:CloudProjectKube my_kube_cluster a6678gggjh76hggjh7f59/a123bc45-a1b2-34c5-678d-678ghg7676ebc
 * ```
 */
export class CloudProjectKube extends pulumi.CustomResource {
    /**
     * Get an existing CloudProjectKube resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudProjectKubeState, opts?: pulumi.CustomResourceOptions): CloudProjectKube {
        return new CloudProjectKube(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/cloudProjectKube:CloudProjectKube';

    /**
     * Returns true if the given object is an instance of CloudProjectKube.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudProjectKube {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudProjectKube.__pulumiType;
    }

    /**
     * True if control-plane is up to date.
     */
    public /*out*/ readonly controlPlaneIsUpToDate!: pulumi.Output<boolean>;
    /**
     * True if all nodes and control-plane are up to date.
     */
    public /*out*/ readonly isUpToDate!: pulumi.Output<boolean>;
    /**
     * The kubeconfig file. Use this file to connect to your kubernetes cluster.
     */
    public /*out*/ readonly kubeconfig!: pulumi.Output<string>;
    /**
     * The name of the kubernetes cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Kubernetes versions available for upgrade.
     */
    public /*out*/ readonly nextUpgradeVersions!: pulumi.Output<string[]>;
    /**
     * Cluster nodes URL.
     */
    public /*out*/ readonly nodesUrl!: pulumi.Output<string>;
    /**
     * OpenStack private network (or vrack) ID to use.
     * Changing this value recreates the resource. Defaults - not use private network.
     */
    public readonly privateNetworkId!: pulumi.Output<string | undefined>;
    /**
     * a valid OVH public cloud region ID in which the kubernetes
     * cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
     * Changing this value recreates the resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Cluster status. Should be normally set to 'READY'.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
     */
    public /*out*/ readonly updatePolicy!: pulumi.Output<string>;
    /**
     * Management URL of your cluster.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * kubernetes version to use.
     * Changing this value recreates the resource. Defaults to latest available.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a CloudProjectKube resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudProjectKubeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudProjectKubeArgs | CloudProjectKubeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudProjectKubeState | undefined;
            resourceInputs["controlPlaneIsUpToDate"] = state ? state.controlPlaneIsUpToDate : undefined;
            resourceInputs["isUpToDate"] = state ? state.isUpToDate : undefined;
            resourceInputs["kubeconfig"] = state ? state.kubeconfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nextUpgradeVersions"] = state ? state.nextUpgradeVersions : undefined;
            resourceInputs["nodesUrl"] = state ? state.nodesUrl : undefined;
            resourceInputs["privateNetworkId"] = state ? state.privateNetworkId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updatePolicy"] = state ? state.updatePolicy : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as CloudProjectKubeArgs | undefined;
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privateNetworkId"] = args ? args.privateNetworkId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["controlPlaneIsUpToDate"] = undefined /*out*/;
            resourceInputs["isUpToDate"] = undefined /*out*/;
            resourceInputs["kubeconfig"] = undefined /*out*/;
            resourceInputs["nextUpgradeVersions"] = undefined /*out*/;
            resourceInputs["nodesUrl"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatePolicy"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CloudProjectKube.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudProjectKube resources.
 */
export interface CloudProjectKubeState {
    /**
     * True if control-plane is up to date.
     */
    controlPlaneIsUpToDate?: pulumi.Input<boolean>;
    /**
     * True if all nodes and control-plane are up to date.
     */
    isUpToDate?: pulumi.Input<boolean>;
    /**
     * The kubeconfig file. Use this file to connect to your kubernetes cluster.
     */
    kubeconfig?: pulumi.Input<string>;
    /**
     * The name of the kubernetes cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * Kubernetes versions available for upgrade.
     */
    nextUpgradeVersions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Cluster nodes URL.
     */
    nodesUrl?: pulumi.Input<string>;
    /**
     * OpenStack private network (or vrack) ID to use.
     * Changing this value recreates the resource. Defaults - not use private network.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * a valid OVH public cloud region ID in which the kubernetes
     * cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
     * Changing this value recreates the resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Cluster status. Should be normally set to 'READY'.
     */
    status?: pulumi.Input<string>;
    /**
     * Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
     */
    updatePolicy?: pulumi.Input<string>;
    /**
     * Management URL of your cluster.
     */
    url?: pulumi.Input<string>;
    /**
     * kubernetes version to use.
     * Changing this value recreates the resource. Defaults to latest available.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudProjectKube resource.
 */
export interface CloudProjectKubeArgs {
    /**
     * The name of the kubernetes cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * OpenStack private network (or vrack) ID to use.
     * Changing this value recreates the resource. Defaults - not use private network.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * a valid OVH public cloud region ID in which the kubernetes
     * cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
     * Changing this value recreates the resource.
     */
    region: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
    /**
     * kubernetes version to use.
     * Changing this value recreates the resource. Defaults to latest available.
     */
    version?: pulumi.Input<string>;
}