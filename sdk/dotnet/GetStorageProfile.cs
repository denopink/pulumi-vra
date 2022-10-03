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
    public static class GetStorageProfile
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// This is an example of how to create a storage profile data source.
        /// 
        /// **Storage profile data source by its id:**
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetStorageProfile.Invoke(new()
        ///     {
        ///         Id = vra_storage_profile.This.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// **Vra storage profile data source filter by external region id:**
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetStorageProfile.Invoke(new()
        ///     {
        ///         Filter = "externalRegionId eq 'foobar'",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A storage profile data source supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetStorageProfileResult> InvokeAsync(GetStorageProfileArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStorageProfileResult>("vra:index/getStorageProfile:getStorageProfile", args ?? new GetStorageProfileArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// This is an example of how to create a storage profile data source.
        /// 
        /// **Storage profile data source by its id:**
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetStorageProfile.Invoke(new()
        ///     {
        ///         Id = vra_storage_profile.This.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// **Vra storage profile data source filter by external region id:**
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetStorageProfile.Invoke(new()
        ///     {
        ///         Filter = "externalRegionId eq 'foobar'",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A storage profile data source supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetStorageProfileResult> Invoke(GetStorageProfileInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStorageProfileResult>("vra:index/getStorageProfile:getStorageProfile", args ?? new GetStorageProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStorageProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        [Input("diskProperties")]
        private Dictionary<string, object>? _diskProperties;

        /// <summary>
        /// Map of storage properties that are to be applied on disk while provisioning.
        /// </summary>
        public Dictionary<string, object> DiskProperties
        {
            get => _diskProperties ?? (_diskProperties = new Dictionary<string, object>());
            set => _diskProperties = value;
        }

        /// <summary>
        /// Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '&lt;regionId&gt;' and cloudAccountId eq '&lt;cloudAccountId&gt;'.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// The id of the image profile instance.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        [Input("tags")]
        private List<Inputs.GetStorageProfileTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public List<Inputs.GetStorageProfileTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetStorageProfileTagArgs>());
            set => _tags = value;
        }

        public GetStorageProfileArgs()
        {
        }
        public static new GetStorageProfileArgs Empty => new GetStorageProfileArgs();
    }

    public sealed class GetStorageProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("diskProperties")]
        private InputMap<object>? _diskProperties;

        /// <summary>
        /// Map of storage properties that are to be applied on disk while provisioning.
        /// </summary>
        public InputMap<object> DiskProperties
        {
            get => _diskProperties ?? (_diskProperties = new InputMap<object>());
            set => _diskProperties = value;
        }

        /// <summary>
        /// Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '&lt;regionId&gt;' and cloudAccountId eq '&lt;cloudAccountId&gt;'.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The id of the image profile instance.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetStorageProfileTagInputArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public InputList<Inputs.GetStorageProfileTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetStorageProfileTagInputArgs>());
            set => _tags = value;
        }

        public GetStorageProfileInvokeArgs()
        {
        }
        public static new GetStorageProfileInvokeArgs Empty => new GetStorageProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetStorageProfileResult
    {
        /// <summary>
        /// Id of the cloud account this storage profile belongs to.
        /// </summary>
        public readonly string CloudAccountId;
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Indicates if this storage profile is a default profile.
        /// </summary>
        public readonly bool DefaultItem;
        public readonly string? Description;
        public readonly ImmutableDictionary<string, object> DiskProperties;
        /// <summary>
        /// The id of the region as seen in the cloud provider for which this profile is defined.
        /// </summary>
        public readonly string ExternalRegionId;
        public readonly string? Filter;
        public readonly string Id;
        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStorageProfileLinkResult> Links;
        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
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
        /// Indicates whether this storage profile supports encryption or not.
        /// </summary>
        public readonly bool SupportsEncryption;
        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStorageProfileTagResult> Tags;
        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetStorageProfileResult(
            string cloudAccountId,

            string createdAt,

            bool defaultItem,

            string? description,

            ImmutableDictionary<string, object> diskProperties,

            string externalRegionId,

            string? filter,

            string id,

            ImmutableArray<Outputs.GetStorageProfileLinkResult> links,

            string name,

            string orgId,

            string owner,

            bool supportsEncryption,

            ImmutableArray<Outputs.GetStorageProfileTagResult> tags,

            string updatedAt)
        {
            CloudAccountId = cloudAccountId;
            CreatedAt = createdAt;
            DefaultItem = defaultItem;
            Description = description;
            DiskProperties = diskProperties;
            ExternalRegionId = externalRegionId;
            Filter = filter;
            Id = id;
            Links = links;
            Name = name;
            OrgId = orgId;
            Owner = owner;
            SupportsEncryption = supportsEncryption;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}