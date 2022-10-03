// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 * ### S
 * This is an example of how to create a storage profile aws resource.
 *
 * **Storage profile aws data source by its id:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const this = vra.getStorageProfileAws({
 *     id: vra_storage_profile_aws["this"].id,
 * });
 * ```
 *
 * **Vra storage profile data source filter by external region id:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const thisStorageProfileAws = pulumi.output(vra.getStorageProfileAws({
 *     filter: "externalRegionId eq 'foobar'",
 * }));
 * ```
 *
 * A storage profile aws data source supports the following arguments:
 */
export function getStorageProfileAws(args?: GetStorageProfileAwsArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageProfileAwsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:index/getStorageProfileAws:getStorageProfileAws", {
        "filter": args.filter,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorageProfileAws.
 */
export interface GetStorageProfileAwsArgs {
    /**
     * Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
     */
    filter?: string;
    /**
     * The id of the image profile instance.
     */
    id?: string;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    tags?: inputs.GetStorageProfileAwsTag[];
}

/**
 * A collection of values returned by getStorageProfileAws.
 */
export interface GetStorageProfileAwsResult {
    readonly cloudAccountId: string;
    /**
     * Date when the entity was created. The date is in ISO 6801 and UTC.
     */
    readonly createdAt: string;
    /**
     * Indicates if this storage profile is a default profile.
     */
    readonly defaultItem: boolean;
    readonly description: string;
    readonly deviceType: string;
    /**
     * The id of the region as seen in the cloud provider for which this profile is defined.
     */
    readonly externalRegionId: string;
    readonly filter?: string;
    readonly id: string;
    readonly iops: string;
    /**
     * HATEOAS of the entity
     */
    readonly links: outputs.GetStorageProfileAwsLink[];
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    readonly name: string;
    readonly orgId: string;
    /**
     * Email of the user that owns the entity.
     */
    readonly owner: string;
    readonly supportsEncryption: boolean;
    readonly tags: outputs.GetStorageProfileAwsTag[];
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    readonly updatedAt: string;
    readonly volumeType: string;
}

export function getStorageProfileAwsOutput(args?: GetStorageProfileAwsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStorageProfileAwsResult> {
    return pulumi.output(args).apply(a => getStorageProfileAws(a, opts))
}

/**
 * A collection of arguments for invoking getStorageProfileAws.
 */
export interface GetStorageProfileAwsOutputArgs {
    /**
     * Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
     */
    filter?: pulumi.Input<string>;
    /**
     * The id of the image profile instance.
     */
    id?: pulumi.Input<string>;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.GetStorageProfileAwsTagArgs>[]>;
}