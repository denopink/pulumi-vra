// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This resource provides a way to create a deployment in vRealize Automation(vRA) by either using a blueprint, or catalog item, or an inline blueprint.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Deployment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import vra:deployment/deployment:Deployment this 05956583-6488-4e7d-84c9-92a7b7219a15`
 * ```
 */
export class Deployment extends pulumi.CustomResource {
    /**
     * Get an existing Deployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeploymentState, opts?: pulumi.CustomResourceOptions): Deployment {
        return new Deployment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vra:deployment/deployment:Deployment';

    /**
     * Returns true if the given object is an instance of Deployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Deployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Deployment.__pulumiType;
    }

    /**
     * vRA Cloud template content. Conflicts with `blueprintId` and `catalogItemId`.
     */
    public readonly blueprintContent!: pulumi.Output<string | undefined>;
    /**
     * Identifier of the requested blueprint in the form ‘UUID:version’.
     */
    public readonly blueprintId!: pulumi.Output<string>;
    /**
     * The version of the vRA cloud template to request the deployment. Used only when `blueprintId` is provided.
     */
    public readonly blueprintVersion!: pulumi.Output<string>;
    /**
     * The id of the vRA catalog item to request the deployment. Conflicts with `blueprintId` and `blueprintContent`.
     */
    public readonly catalogItemId!: pulumi.Output<string>;
    /**
     * The version of the vRA catalog item to request the deployment. Used only when `catalogItemId` is provided.
     */
    public readonly catalogItemVersion!: pulumi.Output<string>;
    /**
     * Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The user the entity was created by.
     */
    public /*out*/ readonly createdBy!: pulumi.Output<string>;
    /**
     * A human-friendly description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * @deprecated Deprecated. True by default even if not provided.
     */
    public readonly expandLastRequest!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to indicate whether to expand project information.
     */
    public readonly expandProject!: pulumi.Output<boolean | undefined>;
    /**
     * @deprecated Deprecated. True by default even if not provided.
     */
    public readonly expandResources!: pulumi.Output<boolean | undefined>;
    /**
     * Expense incurred for the deployment.
     */
    public /*out*/ readonly expenses!: pulumi.Output<outputs.deployment.DeploymentExpense[]>;
    /**
     * Inputs provided by the user. For inputs including those with default values, refer to `inputsIncludingDefaults`.
     */
    public readonly inputs!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * All the inputs applied during last create/update operation, including those with default values. For the list of inputs provided by the user in the configuration, refer to `inputs`.
     */
    public /*out*/ readonly inputsIncludingDefaults!: pulumi.Output<{[key: string]: string}>;
    /**
     * Represents deployment requests.
     */
    public /*out*/ readonly lastRequests!: pulumi.Output<outputs.deployment.DeploymentLastRequest[]>;
    /**
     * Time at which the deployment was last updated.
     */
    public /*out*/ readonly lastUpdatedAt!: pulumi.Output<string>;
    /**
     * The user that last updated the deployment.
     */
    public /*out*/ readonly lastUpdatedBy!: pulumi.Output<string>;
    /**
     * Time at which the deployment lease expires.
     */
    public readonly leaseExpireAt!: pulumi.Output<string>;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the organization this deployment belongs to.
     */
    public /*out*/ readonly orgId!: pulumi.Output<string>;
    /**
     * The user this deployment belongs to. At create, the owner is ignored but is used to update during next apply.
     */
    public readonly owner!: pulumi.Output<string>;
    /**
     * The id of the project this entity belongs to.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The project this entity belongs to.
     */
    public /*out*/ readonly projects!: pulumi.Output<outputs.deployment.DeploymentProject[]>;
    /**
     * Expanded resources for the deployment. Content of this property will not be maintained backward compatible.
     */
    public /*out*/ readonly resources!: pulumi.Output<outputs.deployment.DeploymentResource[]>;
    /**
     * Deployment status. Supported values are: `CREATE_SUCCESSFUL`, `CREATE_INPROGRESS`, `CREATE_FAILED`, `UPDATE_SUCCESSFUL`, `UPDATE_INPROGRESS`, `UPDATE_FAILED`, `DELETE_SUCCESSFUL`, `DELETE_INPROGRESS`, `DELETE_FAILED`, `ACTION_SUCCESSFUL`, `ACTION_INPROGRESS`, `ACTION_FAILED`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Deployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentArgs | DeploymentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeploymentState | undefined;
            resourceInputs["blueprintContent"] = state ? state.blueprintContent : undefined;
            resourceInputs["blueprintId"] = state ? state.blueprintId : undefined;
            resourceInputs["blueprintVersion"] = state ? state.blueprintVersion : undefined;
            resourceInputs["catalogItemId"] = state ? state.catalogItemId : undefined;
            resourceInputs["catalogItemVersion"] = state ? state.catalogItemVersion : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["createdBy"] = state ? state.createdBy : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expandLastRequest"] = state ? state.expandLastRequest : undefined;
            resourceInputs["expandProject"] = state ? state.expandProject : undefined;
            resourceInputs["expandResources"] = state ? state.expandResources : undefined;
            resourceInputs["expenses"] = state ? state.expenses : undefined;
            resourceInputs["inputs"] = state ? state.inputs : undefined;
            resourceInputs["inputsIncludingDefaults"] = state ? state.inputsIncludingDefaults : undefined;
            resourceInputs["lastRequests"] = state ? state.lastRequests : undefined;
            resourceInputs["lastUpdatedAt"] = state ? state.lastUpdatedAt : undefined;
            resourceInputs["lastUpdatedBy"] = state ? state.lastUpdatedBy : undefined;
            resourceInputs["leaseExpireAt"] = state ? state.leaseExpireAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["projects"] = state ? state.projects : undefined;
            resourceInputs["resources"] = state ? state.resources : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as DeploymentArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["blueprintContent"] = args ? args.blueprintContent : undefined;
            resourceInputs["blueprintId"] = args ? args.blueprintId : undefined;
            resourceInputs["blueprintVersion"] = args ? args.blueprintVersion : undefined;
            resourceInputs["catalogItemId"] = args ? args.catalogItemId : undefined;
            resourceInputs["catalogItemVersion"] = args ? args.catalogItemVersion : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["expandLastRequest"] = args ? args.expandLastRequest : undefined;
            resourceInputs["expandProject"] = args ? args.expandProject : undefined;
            resourceInputs["expandResources"] = args ? args.expandResources : undefined;
            resourceInputs["inputs"] = args ? args.inputs : undefined;
            resourceInputs["leaseExpireAt"] = args ? args.leaseExpireAt : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["expenses"] = undefined /*out*/;
            resourceInputs["inputsIncludingDefaults"] = undefined /*out*/;
            resourceInputs["lastRequests"] = undefined /*out*/;
            resourceInputs["lastUpdatedAt"] = undefined /*out*/;
            resourceInputs["lastUpdatedBy"] = undefined /*out*/;
            resourceInputs["orgId"] = undefined /*out*/;
            resourceInputs["projects"] = undefined /*out*/;
            resourceInputs["resources"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Deployment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Deployment resources.
 */
export interface DeploymentState {
    /**
     * vRA Cloud template content. Conflicts with `blueprintId` and `catalogItemId`.
     */
    blueprintContent?: pulumi.Input<string>;
    /**
     * Identifier of the requested blueprint in the form ‘UUID:version’.
     */
    blueprintId?: pulumi.Input<string>;
    /**
     * The version of the vRA cloud template to request the deployment. Used only when `blueprintId` is provided.
     */
    blueprintVersion?: pulumi.Input<string>;
    /**
     * The id of the vRA catalog item to request the deployment. Conflicts with `blueprintId` and `blueprintContent`.
     */
    catalogItemId?: pulumi.Input<string>;
    /**
     * The version of the vRA catalog item to request the deployment. Used only when `catalogItemId` is provided.
     */
    catalogItemVersion?: pulumi.Input<string>;
    /**
     * Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The user the entity was created by.
     */
    createdBy?: pulumi.Input<string>;
    /**
     * A human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated. True by default even if not provided.
     */
    expandLastRequest?: pulumi.Input<boolean>;
    /**
     * Flag to indicate whether to expand project information.
     */
    expandProject?: pulumi.Input<boolean>;
    /**
     * @deprecated Deprecated. True by default even if not provided.
     */
    expandResources?: pulumi.Input<boolean>;
    /**
     * Expense incurred for the deployment.
     */
    expenses?: pulumi.Input<pulumi.Input<inputs.deployment.DeploymentExpense>[]>;
    /**
     * Inputs provided by the user. For inputs including those with default values, refer to `inputsIncludingDefaults`.
     */
    inputs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * All the inputs applied during last create/update operation, including those with default values. For the list of inputs provided by the user in the configuration, refer to `inputs`.
     */
    inputsIncludingDefaults?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Represents deployment requests.
     */
    lastRequests?: pulumi.Input<pulumi.Input<inputs.deployment.DeploymentLastRequest>[]>;
    /**
     * Time at which the deployment was last updated.
     */
    lastUpdatedAt?: pulumi.Input<string>;
    /**
     * The user that last updated the deployment.
     */
    lastUpdatedBy?: pulumi.Input<string>;
    /**
     * Time at which the deployment lease expires.
     */
    leaseExpireAt?: pulumi.Input<string>;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the organization this deployment belongs to.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The user this deployment belongs to. At create, the owner is ignored but is used to update during next apply.
     */
    owner?: pulumi.Input<string>;
    /**
     * The id of the project this entity belongs to.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The project this entity belongs to.
     */
    projects?: pulumi.Input<pulumi.Input<inputs.deployment.DeploymentProject>[]>;
    /**
     * Expanded resources for the deployment. Content of this property will not be maintained backward compatible.
     */
    resources?: pulumi.Input<pulumi.Input<inputs.deployment.DeploymentResource>[]>;
    /**
     * Deployment status. Supported values are: `CREATE_SUCCESSFUL`, `CREATE_INPROGRESS`, `CREATE_FAILED`, `UPDATE_SUCCESSFUL`, `UPDATE_INPROGRESS`, `UPDATE_FAILED`, `DELETE_SUCCESSFUL`, `DELETE_INPROGRESS`, `DELETE_FAILED`, `ACTION_SUCCESSFUL`, `ACTION_INPROGRESS`, `ACTION_FAILED`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Deployment resource.
 */
export interface DeploymentArgs {
    /**
     * vRA Cloud template content. Conflicts with `blueprintId` and `catalogItemId`.
     */
    blueprintContent?: pulumi.Input<string>;
    /**
     * Identifier of the requested blueprint in the form ‘UUID:version’.
     */
    blueprintId?: pulumi.Input<string>;
    /**
     * The version of the vRA cloud template to request the deployment. Used only when `blueprintId` is provided.
     */
    blueprintVersion?: pulumi.Input<string>;
    /**
     * The id of the vRA catalog item to request the deployment. Conflicts with `blueprintId` and `blueprintContent`.
     */
    catalogItemId?: pulumi.Input<string>;
    /**
     * The version of the vRA catalog item to request the deployment. Used only when `catalogItemId` is provided.
     */
    catalogItemVersion?: pulumi.Input<string>;
    /**
     * A human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated. True by default even if not provided.
     */
    expandLastRequest?: pulumi.Input<boolean>;
    /**
     * Flag to indicate whether to expand project information.
     */
    expandProject?: pulumi.Input<boolean>;
    /**
     * @deprecated Deprecated. True by default even if not provided.
     */
    expandResources?: pulumi.Input<boolean>;
    /**
     * Inputs provided by the user. For inputs including those with default values, refer to `inputsIncludingDefaults`.
     */
    inputs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Time at which the deployment lease expires.
     */
    leaseExpireAt?: pulumi.Input<string>;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * The user this deployment belongs to. At create, the owner is ignored but is used to update during next apply.
     */
    owner?: pulumi.Input<string>;
    /**
     * The id of the project this entity belongs to.
     */
    projectId: pulumi.Input<string>;
}