// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 * ### S
 *
 * This is an example of how to lookup a region enumeration data source.
 *
 * DeprecationMessage: 'region_enumeration' is deprecated. Use 'region_enumeration_vsphere' instead.
 *
 * **Region enumeration data source for vSphere**
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const this = vra.getRegionEnumerationVsphere({
 *     acceptSelfSignedCert: false,
 *     dcid: _var.vra_data_collector_id,
 *     hostname: _var.hostname,
 *     password: _var.password,
 *     username: _var.username,
 * });
 * ```
 *
 * The region enumeration data source supports the following arguments:
 */
export function getRegionEnumeration(args: GetRegionEnumerationArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionEnumerationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:index/getRegionEnumeration:getRegionEnumeration", {
        "acceptSelfSignedCert": args.acceptSelfSignedCert,
        "dcid": args.dcid,
        "hostname": args.hostname,
        "password": args.password,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegionEnumeration.
 */
export interface GetRegionEnumerationArgs {
    /**
     * Accept self signed certificate when connecting to vSphere. Example: false
     */
    acceptSelfSignedCert?: boolean;
    /**
     * ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
     */
    dcid?: string;
    /**
     * Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com
     */
    hostname: string;
    /**
     * Password for the user used to authenticate with the cloud Account
     */
    password: string;
    /**
     * Username to authenticate with the cloud account
     */
    username: string;
}

/**
 * A collection of values returned by getRegionEnumeration.
 */
export interface GetRegionEnumerationResult {
    readonly acceptSelfSignedCert?: boolean;
    readonly dcid?: string;
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly password: string;
    /**
     * A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2
     */
    readonly regions: string[];
    readonly username: string;
}

export function getRegionEnumerationOutput(args: GetRegionEnumerationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionEnumerationResult> {
    return pulumi.output(args).apply(a => getRegionEnumeration(a, opts))
}

/**
 * A collection of arguments for invoking getRegionEnumeration.
 */
export interface GetRegionEnumerationOutputArgs {
    /**
     * Accept self signed certificate when connecting to vSphere. Example: false
     */
    acceptSelfSignedCert?: pulumi.Input<boolean>;
    /**
     * ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
     */
    dcid?: pulumi.Input<string>;
    /**
     * Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com
     */
    hostname: pulumi.Input<string>;
    /**
     * Password for the user used to authenticate with the cloud Account
     */
    password: pulumi.Input<string>;
    /**
     * Username to authenticate with the cloud account
     */
    username: pulumi.Input<string>;
}