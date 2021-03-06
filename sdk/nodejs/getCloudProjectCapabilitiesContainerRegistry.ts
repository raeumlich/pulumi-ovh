// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get the container registry capabilities of a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const capabilities = pulumi.output(ovh.getCloudProjectCapabilitiesContainerRegistry({
 *     serviceName: "XXXXXX",
 * }));
 * ```
 */
export function getCloudProjectCapabilitiesContainerRegistry(args: GetCloudProjectCapabilitiesContainerRegistryArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudProjectCapabilitiesContainerRegistryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("ovh:index/getCloudProjectCapabilitiesContainerRegistry:getCloudProjectCapabilitiesContainerRegistry", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudProjectCapabilitiesContainerRegistry.
 */
export interface GetCloudProjectCapabilitiesContainerRegistryArgs {
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getCloudProjectCapabilitiesContainerRegistry.
 */
export interface GetCloudProjectCapabilitiesContainerRegistryResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of container registry capability for a single region
     */
    readonly results: outputs.GetCloudProjectCapabilitiesContainerRegistryResult[];
    readonly serviceName: string;
}

export function getCloudProjectCapabilitiesContainerRegistryOutput(args: GetCloudProjectCapabilitiesContainerRegistryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudProjectCapabilitiesContainerRegistryResult> {
    return pulumi.output(args).apply(a => getCloudProjectCapabilitiesContainerRegistry(a, opts))
}

/**
 * A collection of arguments for invoking getCloudProjectCapabilitiesContainerRegistry.
 */
export interface GetCloudProjectCapabilitiesContainerRegistryOutputArgs {
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
