// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Blueprint
{
    /// <summary>
    /// Creates a VMware vRealize Automation cloud template (blueprint) version resource.
    /// 
    /// ## Example Usage
    /// ### S
    /// 
    /// The following example shows how to create a cloud template (blueprint) version resource.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = Pulumiverse.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Vra.Blueprint.BlueprintVersion("this", new()
    ///     {
    ///         BlueprintId = @var.Vra_blueprint_id,
    ///         ChangeLog = "First version",
    ///         Description = "Released from vRA terraform provider",
    ///         Release = true,
    ///         Version = random_integer.Suffix.Result / random_integer.Suffix.Result,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// To import the cloud template (blueprint) version, use the ID as in the following example
    /// 
    /// ```sh
    ///  $ pulumi import vra:blueprint/blueprintVersion:BlueprintVersion this 05956583-6488-4e7d-84c9-92a7b7219a15`
    /// ```
    /// </summary>
    [VraResourceType("vra:blueprint/blueprintVersion:BlueprintVersion")]
    public partial class BlueprintVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of cloud template (blueprint).
        /// </summary>
        [Output("blueprintDescription")]
        public Output<string> BlueprintDescription { get; private set; } = null!;

        /// <summary>
        /// ID of the cloud template  (blueprint).
        /// </summary>
        [Output("blueprintId")]
        public Output<string> BlueprintId { get; private set; } = null!;

        /// <summary>
        /// Cloud template  (blueprint) version log.
        /// </summary>
        [Output("changeLog")]
        public Output<string?> ChangeLog { get; private set; } = null!;

        /// <summary>
        /// Blueprint YAML content.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was created. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// User who created the entity.
        /// </summary>
        [Output("createdBy")]
        public Output<string> CreatedBy { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description for the cloud template  (blueprint) version.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of cloud template (blueprint) version.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of organization that entity belongs to.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// ID of project that entity belongs to.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Name of project that entity belongs to.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// Flag to indicate whether to release the version.
        /// </summary>
        [Output("release")]
        public Output<bool?> Release { get; private set; } = null!;

        /// <summary>
        /// Status of the cloud template (blueprint). Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// User who last updated the entity.
        /// </summary>
        [Output("updatedBy")]
        public Output<string> UpdatedBy { get; private set; } = null!;

        /// <summary>
        /// Flag to indicate if the current content of the cloud template (blueprint) is valid.
        /// </summary>
        [Output("valid")]
        public Output<string> Valid { get; private set; } = null!;

        /// <summary>
        /// Cloud template  (blueprint) version.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a BlueprintVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BlueprintVersion(string name, BlueprintVersionArgs args, CustomResourceOptions? options = null)
            : base("vra:blueprint/blueprintVersion:BlueprintVersion", name, args ?? new BlueprintVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BlueprintVersion(string name, Input<string> id, BlueprintVersionState? state = null, CustomResourceOptions? options = null)
            : base("vra:blueprint/blueprintVersion:BlueprintVersion", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BlueprintVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BlueprintVersion Get(string name, Input<string> id, BlueprintVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new BlueprintVersion(name, id, state, options);
        }
    }

    public sealed class BlueprintVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the cloud template  (blueprint).
        /// </summary>
        [Input("blueprintId", required: true)]
        public Input<string> BlueprintId { get; set; } = null!;

        /// <summary>
        /// Cloud template  (blueprint) version log.
        /// </summary>
        [Input("changeLog")]
        public Input<string>? ChangeLog { get; set; }

        /// <summary>
        /// Human-friendly description for the cloud template  (blueprint) version.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Flag to indicate whether to release the version.
        /// </summary>
        [Input("release")]
        public Input<bool>? Release { get; set; }

        /// <summary>
        /// Cloud template  (blueprint) version.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public BlueprintVersionArgs()
        {
        }
        public static new BlueprintVersionArgs Empty => new BlueprintVersionArgs();
    }

    public sealed class BlueprintVersionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of cloud template (blueprint).
        /// </summary>
        [Input("blueprintDescription")]
        public Input<string>? BlueprintDescription { get; set; }

        /// <summary>
        /// ID of the cloud template  (blueprint).
        /// </summary>
        [Input("blueprintId")]
        public Input<string>? BlueprintId { get; set; }

        /// <summary>
        /// Cloud template  (blueprint) version log.
        /// </summary>
        [Input("changeLog")]
        public Input<string>? ChangeLog { get; set; }

        /// <summary>
        /// Blueprint YAML content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Date when the entity was created. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// User who created the entity.
        /// </summary>
        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        /// <summary>
        /// Human-friendly description for the cloud template  (blueprint) version.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of cloud template (blueprint) version.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of organization that entity belongs to.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of project that entity belongs to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Name of project that entity belongs to.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Flag to indicate whether to release the version.
        /// </summary>
        [Input("release")]
        public Input<bool>? Release { get; set; }

        /// <summary>
        /// Status of the cloud template (blueprint). Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// User who last updated the entity.
        /// </summary>
        [Input("updatedBy")]
        public Input<string>? UpdatedBy { get; set; }

        /// <summary>
        /// Flag to indicate if the current content of the cloud template (blueprint) is valid.
        /// </summary>
        [Input("valid")]
        public Input<string>? Valid { get; set; }

        /// <summary>
        /// Cloud template  (blueprint) version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public BlueprintVersionState()
        {
        }
        public static new BlueprintVersionState Empty => new BlueprintVersionState();
    }
}