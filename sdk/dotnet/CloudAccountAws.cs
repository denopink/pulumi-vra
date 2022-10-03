// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace schmidtw.Vra
{
    /// <summary>
    /// Creates a VMware vRealize Automation AWS cloud account resource.
    /// 
    /// ## Example Usage
    /// ### S
    /// 
    /// The following example shows how to create an AWS cloud account resource.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = schmidtw.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Vra.CloudAccountAws("this", new()
    ///     {
    ///         Description = "terraform test cloud account aws",
    ///         AccessKey = @var.Access_key,
    ///         SecretKey = @var.Secret_key,
    ///         Regions = new[]
    ///         {
    ///             "us-east-1",
    ///             "us-west-1",
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Vra.Inputs.CloudAccountAwsTagArgs
    ///             {
    ///                 Key = "foo",
    ///                 Value = "bar",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// To import the AWS cloud account, use the ID as in the following example
    /// 
    /// ```sh
    ///  $ pulumi import vra:index/cloudAccountAws:CloudAccountAws new_aws 05956583-6488-4e7d-84c9-92a7b7219a15`
    /// ```
    /// </summary>
    [VraResourceType("vra:index/cloudAccountAws:CloudAccountAws")]
    public partial class CloudAccountAws : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access key ID for AWS.
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// Date when entity was created. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// HATEOAS of entity.
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.CloudAccountAwsLink>> Links { get; private set; } = null!;

        /// <summary>
        /// Name of AWS cloud account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of organization that entity belongs to.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Email of entity owner.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// Set of region names enabled for the cloud account.
        /// </summary>
        [Output("regions")]
        public Output<ImmutableArray<string>> Regions { get; private set; } = null!;

        /// <summary>
        /// AWS Secret Access Key
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;

        /// <summary>
        /// Set of tag keys and values to apply to the cloud account.  
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.CloudAccountAwsTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a CloudAccountAws resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudAccountAws(string name, CloudAccountAwsArgs args, CustomResourceOptions? options = null)
            : base("vra:index/cloudAccountAws:CloudAccountAws", name, args ?? new CloudAccountAwsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudAccountAws(string name, Input<string> id, CloudAccountAwsState? state = null, CustomResourceOptions? options = null)
            : base("vra:index/cloudAccountAws:CloudAccountAws", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/schmidtw/pulumi-vra/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudAccountAws resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudAccountAws Get(string name, Input<string> id, CloudAccountAwsState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudAccountAws(name, id, state, options);
        }
    }

    public sealed class CloudAccountAwsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access key ID for AWS.
        /// </summary>
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of AWS cloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regions", required: true)]
        private InputList<string>? _regions;

        /// <summary>
        /// Set of region names enabled for the cloud account.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// AWS Secret Access Key
        /// </summary>
        [Input("secretKey", required: true)]
        public Input<string> SecretKey { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.CloudAccountAwsTagArgs>? _tags;

        /// <summary>
        /// Set of tag keys and values to apply to the cloud account.  
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public InputList<Inputs.CloudAccountAwsTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CloudAccountAwsTagArgs>());
            set => _tags = value;
        }

        public CloudAccountAwsArgs()
        {
        }
        public static new CloudAccountAwsArgs Empty => new CloudAccountAwsArgs();
    }

    public sealed class CloudAccountAwsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access key ID for AWS.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// Date when entity was created. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("links")]
        private InputList<Inputs.CloudAccountAwsLinkGetArgs>? _links;

        /// <summary>
        /// HATEOAS of entity.
        /// </summary>
        public InputList<Inputs.CloudAccountAwsLinkGetArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.CloudAccountAwsLinkGetArgs>());
            set => _links = value;
        }

        /// <summary>
        /// Name of AWS cloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of organization that entity belongs to.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Email of entity owner.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// Set of region names enabled for the cloud account.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// AWS Secret Access Key
        /// </summary>
        [Input("secretKey")]
        public Input<string>? SecretKey { get; set; }

        [Input("tags")]
        private InputList<Inputs.CloudAccountAwsTagGetArgs>? _tags;

        /// <summary>
        /// Set of tag keys and values to apply to the cloud account.  
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public InputList<Inputs.CloudAccountAwsTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CloudAccountAwsTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public CloudAccountAwsState()
        {
        }
        public static new CloudAccountAwsState Empty => new CloudAccountAwsState();
    }
}