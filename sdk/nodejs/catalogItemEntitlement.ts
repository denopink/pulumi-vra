// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource provides a way to create a catalog item entitlement in VMware vRealize Automation.
 *
 * ## Example Usage
 * ### S
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@schmidtw/vra";
 *
 * const _this = new vra.CatalogItemEntitlement("this", {
 *     catalogItemId: _var.catalog_item_id,
 *     projectId: _var.project_id,
 * });
 * ```
 *
 * ## Import
 *
 * Catalog item entitlement can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import vra:index/catalogItemEntitlement:CatalogItemEntitlement this 05956583-6488-4e7d-84c9-92a7b7219a15`
 * ```
 */
export class CatalogItemEntitlement extends pulumi.CustomResource {
    /**
     * Get an existing CatalogItemEntitlement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CatalogItemEntitlementState, opts?: pulumi.CustomResourceOptions): CatalogItemEntitlement {
        return new CatalogItemEntitlement(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vra:index/catalogItemEntitlement:CatalogItemEntitlement';

    /**
     * Returns true if the given object is an instance of CatalogItemEntitlement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CatalogItemEntitlement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CatalogItemEntitlement.__pulumiType;
    }

    /**
     * The id of the catalog item to create the entitlement.
     */
    public readonly catalogItemId!: pulumi.Output<string>;
    /**
     * Represents a catalog item that is linked to a project via an entitlement.
     */
    public /*out*/ readonly definitions!: pulumi.Output<outputs.CatalogItemEntitlementDefinition[]>;
    /**
     * The id of the project this entity belongs to.
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a CatalogItemEntitlement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CatalogItemEntitlementArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CatalogItemEntitlementArgs | CatalogItemEntitlementState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CatalogItemEntitlementState | undefined;
            resourceInputs["catalogItemId"] = state ? state.catalogItemId : undefined;
            resourceInputs["definitions"] = state ? state.definitions : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as CatalogItemEntitlementArgs | undefined;
            if ((!args || args.catalogItemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalogItemId'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["catalogItemId"] = args ? args.catalogItemId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["definitions"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CatalogItemEntitlement.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CatalogItemEntitlement resources.
 */
export interface CatalogItemEntitlementState {
    /**
     * The id of the catalog item to create the entitlement.
     */
    catalogItemId?: pulumi.Input<string>;
    /**
     * Represents a catalog item that is linked to a project via an entitlement.
     */
    definitions?: pulumi.Input<pulumi.Input<inputs.CatalogItemEntitlementDefinition>[]>;
    /**
     * The id of the project this entity belongs to.
     */
    projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CatalogItemEntitlement resource.
 */
export interface CatalogItemEntitlementArgs {
    /**
     * The id of the catalog item to create the entitlement.
     */
    catalogItemId: pulumi.Input<string>;
    /**
     * The id of the project this entity belongs to.
     */
    projectId: pulumi.Input<string>;
}