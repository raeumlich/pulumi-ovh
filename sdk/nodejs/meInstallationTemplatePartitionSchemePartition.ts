// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this resource to create a partition in the partition scheme of a custom installation template available for dedicated servers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mytemplate = new ovh.MeInstallationTemplate("mytemplate", {
 *     baseTemplateName: "centos7_64",
 *     templateName: "mytemplate",
 *     defaultLanguage: "fr",
 * });
 * const scheme = new ovh.MeInstallationTemplatePartitionScheme("scheme", {
 *     templateName: mytemplate.templateName,
 *     priority: 1,
 * });
 * const root = new ovh.MeInstallationTemplatePartitionSchemePartition("root", {
 *     templateName: scheme.templateName,
 *     schemeName: scheme.name,
 *     mountpoint: "/",
 *     filesystem: "ext4",
 *     size: 400,
 *     order: 1,
 *     type: "primary",
 * });
 * ```
 *
 * ## Import
 *
 * Use the fake id format to import the resource `template_name/scheme_name/mountpoint` (example"mytemplate/myscheme//").
 */
export class MeInstallationTemplatePartitionSchemePartition extends pulumi.CustomResource {
    /**
     * Get an existing MeInstallationTemplatePartitionSchemePartition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MeInstallationTemplatePartitionSchemePartitionState, opts?: pulumi.CustomResourceOptions): MeInstallationTemplatePartitionSchemePartition {
        return new MeInstallationTemplatePartitionSchemePartition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/meInstallationTemplatePartitionSchemePartition:MeInstallationTemplatePartitionSchemePartition';

    /**
     * Returns true if the given object is an instance of MeInstallationTemplatePartitionSchemePartition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MeInstallationTemplatePartitionSchemePartition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MeInstallationTemplatePartitionSchemePartition.__pulumiType;
    }

    /**
     * Partition filesystem
     */
    public readonly filesystem!: pulumi.Output<string>;
    /**
     * partition mount point
     */
    public readonly mountpoint!: pulumi.Output<string>;
    /**
     * step or order. specifies the creation order of the partition on the disk
     */
    public readonly order!: pulumi.Output<number>;
    /**
     * raid partition type
     */
    public readonly raid!: pulumi.Output<string>;
    /**
     * name of this partitioning scheme
     */
    public readonly schemeName!: pulumi.Output<string>;
    /**
     * size of partition in MB, 0 => rest of the space
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * Template name
     */
    public readonly templateName!: pulumi.Output<string>;
    /**
     * partition type
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The volume name needed for proxmox distribution
     */
    public readonly volumeName!: pulumi.Output<string>;

    /**
     * Create a MeInstallationTemplatePartitionSchemePartition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MeInstallationTemplatePartitionSchemePartitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MeInstallationTemplatePartitionSchemePartitionArgs | MeInstallationTemplatePartitionSchemePartitionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MeInstallationTemplatePartitionSchemePartitionState | undefined;
            resourceInputs["filesystem"] = state ? state.filesystem : undefined;
            resourceInputs["mountpoint"] = state ? state.mountpoint : undefined;
            resourceInputs["order"] = state ? state.order : undefined;
            resourceInputs["raid"] = state ? state.raid : undefined;
            resourceInputs["schemeName"] = state ? state.schemeName : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["volumeName"] = state ? state.volumeName : undefined;
        } else {
            const args = argsOrState as MeInstallationTemplatePartitionSchemePartitionArgs | undefined;
            if ((!args || args.filesystem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filesystem'");
            }
            if ((!args || args.mountpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mountpoint'");
            }
            if ((!args || args.order === undefined) && !opts.urn) {
                throw new Error("Missing required property 'order'");
            }
            if ((!args || args.schemeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schemeName'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["filesystem"] = args ? args.filesystem : undefined;
            resourceInputs["mountpoint"] = args ? args.mountpoint : undefined;
            resourceInputs["order"] = args ? args.order : undefined;
            resourceInputs["raid"] = args ? args.raid : undefined;
            resourceInputs["schemeName"] = args ? args.schemeName : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["volumeName"] = args ? args.volumeName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MeInstallationTemplatePartitionSchemePartition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MeInstallationTemplatePartitionSchemePartition resources.
 */
export interface MeInstallationTemplatePartitionSchemePartitionState {
    /**
     * Partition filesystem
     */
    filesystem?: pulumi.Input<string>;
    /**
     * partition mount point
     */
    mountpoint?: pulumi.Input<string>;
    /**
     * step or order. specifies the creation order of the partition on the disk
     */
    order?: pulumi.Input<number>;
    /**
     * raid partition type
     */
    raid?: pulumi.Input<string>;
    /**
     * name of this partitioning scheme
     */
    schemeName?: pulumi.Input<string>;
    /**
     * size of partition in MB, 0 => rest of the space
     */
    size?: pulumi.Input<number>;
    /**
     * Template name
     */
    templateName?: pulumi.Input<string>;
    /**
     * partition type
     */
    type?: pulumi.Input<string>;
    /**
     * The volume name needed for proxmox distribution
     */
    volumeName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MeInstallationTemplatePartitionSchemePartition resource.
 */
export interface MeInstallationTemplatePartitionSchemePartitionArgs {
    /**
     * Partition filesystem
     */
    filesystem: pulumi.Input<string>;
    /**
     * partition mount point
     */
    mountpoint: pulumi.Input<string>;
    /**
     * step or order. specifies the creation order of the partition on the disk
     */
    order: pulumi.Input<number>;
    /**
     * raid partition type
     */
    raid?: pulumi.Input<string>;
    /**
     * name of this partitioning scheme
     */
    schemeName: pulumi.Input<string>;
    /**
     * size of partition in MB, 0 => rest of the space
     */
    size: pulumi.Input<number>;
    /**
     * Template name
     */
    templateName: pulumi.Input<string>;
    /**
     * partition type
     */
    type: pulumi.Input<string>;
    /**
     * The volume name needed for proxmox distribution
     */
    volumeName?: pulumi.Input<string>;
}
