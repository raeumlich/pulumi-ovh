// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information about a container registry associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const registry = pulumi.output(ovh.getCloudProjectContainerregistry({
 *     registryId: "yyyy",
 *     serviceName: "XXXXXX",
 * }));
 * ```
 */
export function getCloudProjectContainerregistry(args: GetCloudProjectContainerregistryArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudProjectContainerregistryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("ovh:index/getCloudProjectContainerregistry:getCloudProjectContainerregistry", {
        "registryId": args.registryId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudProjectContainerregistry.
 */
export interface GetCloudProjectContainerregistryArgs {
    /**
     * Registry ID
     */
    registryId: string;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getCloudProjectContainerregistry.
 */
export interface GetCloudProjectContainerregistryResult {
    /**
     * Registry creation date
     */
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Registry name
     */
    readonly name: string;
    /**
     * Project ID of your registry
     */
    readonly projectId: string;
    /**
     * Region of the registry
     */
    readonly region: string;
    readonly registryId: string;
    readonly serviceName: string;
    /**
     * Current size of the registry (bytes)
     */
    readonly size: number;
    /**
     * Registry status
     */
    readonly status: string;
    /**
     * Registry last update date
     */
    readonly updatedAt: string;
    /**
     * Access url of the registry
     */
    readonly url: string;
    /**
     * Version of your registry
     */
    readonly version: string;
}

export function getCloudProjectContainerregistryOutput(args: GetCloudProjectContainerregistryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudProjectContainerregistryResult> {
    return pulumi.output(args).apply(a => getCloudProjectContainerregistry(a, opts))
}

/**
 * A collection of arguments for invoking getCloudProjectContainerregistry.
 */
export interface GetCloudProjectContainerregistryOutputArgs {
    /**
     * Registry ID
     */
    registryId: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
