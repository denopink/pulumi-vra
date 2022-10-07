// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudaccount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware vRA cloudaccount.Nsxv data source.
//
// ## Example Usage
// ### S
//
// **NSX-V cloud account data source by its id:**
//
// This is an example of how to read the cloud account data source using its id.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/cloudaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/cloudaccount"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudaccount.LookupNsxv(ctx, &cloudaccount.LookupNsxvArgs{
//				Id: pulumi.StringRef(_var.Vra_cloud_account_nsxv_id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// **NSX-V cloud account data source by its name:**
//
// This is an example of how to read the cloud account data source using its name.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/cloudaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/cloudaccount"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudaccount.LookupNsxv(ctx, &cloudaccount.LookupNsxvArgs{
//				Name: pulumi.StringRef(_var.Vra_cloud_account_nsxv_name),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxv(ctx *pulumi.Context, args *LookupNsxvArgs, opts ...pulumi.InvokeOption) (*LookupNsxvResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNsxvResult
	err := ctx.Invoke("vra:cloudaccount/getNsxv:getNsxv", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxv.
type LookupNsxvArgs struct {
	// The id of this NSX-V cloud account.
	Id *string `pulumi:"id"`
	// The name of this NSX-V cloud account.
	Name *string `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetNsxvTag `pulumi:"tags"`
}

// A collection of values returned by getNsxv.
type LookupNsxvResult struct {
	// Cloud accounts associated with this cloud account.
	AssociatedCloudAccountIds []string `pulumi:"associatedCloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// Identifier of a data collector vm deployed in the on premise infrastructure.
	DcId string `pulumi:"dcId"`
	// A human-friendly description.
	Description string `pulumi:"description"`
	// Host name for the NSX-V cloud account.
	Hostname string `pulumi:"hostname"`
	Id       string `pulumi:"id"`
	// HATEOAS of the entity.
	Links []GetNsxvLink `pulumi:"links"`
	Name  string        `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId string `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetNsxvTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
	// Username to authenticate with the cloud account.
	Username string `pulumi:"username"`
}

func LookupNsxvOutput(ctx *pulumi.Context, args LookupNsxvOutputArgs, opts ...pulumi.InvokeOption) LookupNsxvResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxvResult, error) {
			args := v.(LookupNsxvArgs)
			r, err := LookupNsxv(ctx, &args, opts...)
			var s LookupNsxvResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxvResultOutput)
}

// A collection of arguments for invoking getNsxv.
type LookupNsxvOutputArgs struct {
	// The id of this NSX-V cloud account.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of this NSX-V cloud account.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags GetNsxvTagArrayInput `pulumi:"tags"`
}

func (LookupNsxvOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxvArgs)(nil)).Elem()
}

// A collection of values returned by getNsxv.
type LookupNsxvResultOutput struct{ *pulumi.OutputState }

func (LookupNsxvResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxvResult)(nil)).Elem()
}

func (o LookupNsxvResultOutput) ToLookupNsxvResultOutput() LookupNsxvResultOutput {
	return o
}

func (o LookupNsxvResultOutput) ToLookupNsxvResultOutputWithContext(ctx context.Context) LookupNsxvResultOutput {
	return o
}

// Cloud accounts associated with this cloud account.
func (o LookupNsxvResultOutput) AssociatedCloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxvResult) []string { return v.AssociatedCloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o LookupNsxvResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Identifier of a data collector vm deployed in the on premise infrastructure.
func (o LookupNsxvResultOutput) DcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvResult) string { return v.DcId }).(pulumi.StringOutput)
}

// A human-friendly description.
func (o LookupNsxvResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvResult) string { return v.Description }).(pulumi.StringOutput)
}

// Host name for the NSX-V cloud account.
func (o LookupNsxvResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvResult) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o LookupNsxvResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvResult) string { return v.Id }).(pulumi.StringOutput)
}

// HATEOAS of the entity.
func (o LookupNsxvResultOutput) Links() GetNsxvLinkArrayOutput {
	return o.ApplyT(func(v LookupNsxvResult) []GetNsxvLink { return v.Links }).(GetNsxvLinkArrayOutput)
}

func (o LookupNsxvResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o LookupNsxvResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o LookupNsxvResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvResult) string { return v.Owner }).(pulumi.StringOutput)
}

// A set of tag keys and optional values that were set on this resource.
// example:[ { "key" : "vmware", "value": "provider" } ]
func (o LookupNsxvResultOutput) Tags() GetNsxvTagArrayOutput {
	return o.ApplyT(func(v LookupNsxvResult) []GetNsxvTag { return v.Tags }).(GetNsxvTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o LookupNsxvResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Username to authenticate with the cloud account.
func (o LookupNsxvResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxvResultOutput{})
}