// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get the list of dedicated servers associated with your OVH Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const servers = pulumi.output(ovh.getDedicatedServers());
 * ```
 */
export function getDedicatedServers(opts?: pulumi.InvokeOptions): Promise<GetDedicatedServersResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("ovh:index/getDedicatedServers:getDedicatedServers", {
    }, opts);
}

/**
 * A collection of values returned by getDedicatedServers.
 */
export interface GetDedicatedServersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of dedicated servers IDs associated with your OVH Account.
     */
    readonly results: string[];
}
