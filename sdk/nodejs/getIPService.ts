// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about an ip service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myip = pulumi.output(ovh.getIPService({
 *     serviceName: "XXXXXX",
 * }));
 * ```
 */
export function getIPService(args: GetIPServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetIPServiceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("ovh:index/getIPService:getIPService", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getIPService.
 */
export interface GetIPServiceArgs {
    /**
     * The service name
     */
    serviceName: string;
}

/**
 * A collection of values returned by getIPService.
 */
export interface GetIPServiceResult {
    /**
     * can be terminated
     */
    readonly canBeTerminated: boolean;
    /**
     * country
     */
    readonly country: string;
    /**
     * Custom description on your ip
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * ip block
     */
    readonly ip: string;
    /**
     * IP block organisation Id
     */
    readonly organisationId: string;
    /**
     * Routage information
     */
    readonly routedTos: outputs.GetIPServiceRoutedTo[];
    /**
     * Service where ip is routed to
     */
    readonly serviceName: string;
    /**
     * Possible values for ip type (    "cdn", "cloud", "dedicated", "failover", "hostedSsl", "housing", "loadBalancing", "mail", "overthebox", "pcc", "pci", "private", "vpn", "vps", "vrack", "xdsl")
     */
    readonly type: string;
}

export function getIPServiceOutput(args: GetIPServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIPServiceResult> {
    return pulumi.output(args).apply(a => getIPService(a, opts))
}

/**
 * A collection of arguments for invoking getIPService.
 */
export interface GetIPServiceOutputArgs {
    /**
     * The service name
     */
    serviceName: pulumi.Input<string>;
}
