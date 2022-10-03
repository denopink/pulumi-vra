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
    public static class GetCloudAccountGcp
    {
        /// <summary>
        /// Provides a VMware vRA vra.CloudAccountGcp data source.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// **GCP cloud account data source by its id:**
        /// 
        /// This is an example of how to create an GCP cloud account resource and read it as a data source using its id.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCloudAccountGcp.Invoke(new()
        ///     {
        ///         Id = @var.Vra_cloud_account_gcp_id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// **GCP cloud account data source by its name:**
        /// 
        /// This is an example of how to read the cloud account data source using its name.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCloudAccountGcp.Invoke(new()
        ///     {
        ///         Name = @var.Vra_cloud_account_gcp_name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudAccountGcpResult> InvokeAsync(GetCloudAccountGcpArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudAccountGcpResult>("vra:index/getCloudAccountGcp:getCloudAccountGcp", args ?? new GetCloudAccountGcpArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware vRA vra.CloudAccountGcp data source.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// **GCP cloud account data source by its id:**
        /// 
        /// This is an example of how to create an GCP cloud account resource and read it as a data source using its id.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCloudAccountGcp.Invoke(new()
        ///     {
        ///         Id = @var.Vra_cloud_account_gcp_id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// **GCP cloud account data source by its name:**
        /// 
        /// This is an example of how to read the cloud account data source using its name.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCloudAccountGcp.Invoke(new()
        ///     {
        ///         Name = @var.Vra_cloud_account_gcp_name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudAccountGcpResult> Invoke(GetCloudAccountGcpInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudAccountGcpResult>("vra:index/getCloudAccountGcp:getCloudAccountGcp", args ?? new GetCloudAccountGcpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudAccountGcpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of this GCP cloud account.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of this GCP cloud account.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("tags")]
        private List<Inputs.GetCloudAccountGcpTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public List<Inputs.GetCloudAccountGcpTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetCloudAccountGcpTagArgs>());
            set => _tags = value;
        }

        public GetCloudAccountGcpArgs()
        {
        }
        public static new GetCloudAccountGcpArgs Empty => new GetCloudAccountGcpArgs();
    }

    public sealed class GetCloudAccountGcpInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of this GCP cloud account.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of this GCP cloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetCloudAccountGcpTagInputArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public InputList<Inputs.GetCloudAccountGcpTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetCloudAccountGcpTagInputArgs>());
            set => _tags = value;
        }

        public GetCloudAccountGcpInvokeArgs()
        {
        }
        public static new GetCloudAccountGcpInvokeArgs Empty => new GetCloudAccountGcpInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudAccountGcpResult
    {
        /// <summary>
        /// GCP Client email.
        /// </summary>
        public readonly string ClientEmail;
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        public readonly string Description;
        public readonly string Id;
        /// <summary>
        /// HATEOAS of the entity.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudAccountGcpLinkResult> Links;
        public readonly string Name;
        /// <summary>
        /// The id of the organization this entity belongs to.
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// GCP Private key ID.
        /// </summary>
        public readonly string PrivateKeyId;
        /// <summary>
        /// GCP Project ID.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// A set of region names that are enabled for this account.
        /// </summary>
        public readonly ImmutableArray<string> Regions;
        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudAccountGcpTagResult> Tags;
        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetCloudAccountGcpResult(
            string clientEmail,

            string createdAt,

            string description,

            string id,

            ImmutableArray<Outputs.GetCloudAccountGcpLinkResult> links,

            string name,

            string orgId,

            string owner,

            string privateKeyId,

            string projectId,

            ImmutableArray<string> regions,

            ImmutableArray<Outputs.GetCloudAccountGcpTagResult> tags,

            string updatedAt)
        {
            ClientEmail = clientEmail;
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Links = links;
            Name = name;
            OrgId = orgId;
            Owner = owner;
            PrivateKeyId = privateKeyId;
            ProjectId = projectId;
            Regions = regions;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}