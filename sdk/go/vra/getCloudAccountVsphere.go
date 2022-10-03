// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware vRA CloudAccountVsphere data source.
//
// ## Example Usage
// ### S
//
// **vSphere cloud account data source by its id:**
//
// This is an example of how to read the cloud account data source using its id.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/schmidtw/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.LookupCloudAccountVsphere(ctx, &GetCloudAccountVsphereArgs{
//				Id: pulumi.StringRef(_var.Vra_cloud_account_vsphere_id),
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
// **vSphere cloud account data source by its name:**
//
// This is an example of how to read the cloud account data source using its name.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/schmidtw/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.LookupCloudAccountVsphere(ctx, &GetCloudAccountVsphereArgs{
//				Name: pulumi.StringRef(_var.Vra_cloud_account_vsphere_name),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCloudAccountVsphere(ctx *pulumi.Context, args *LookupCloudAccountVsphereArgs, opts ...pulumi.InvokeOption) (*LookupCloudAccountVsphereResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCloudAccountVsphereResult
	err := ctx.Invoke("vra:index/getCloudAccountVsphere:getCloudAccountVsphere", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudAccountVsphere.
type LookupCloudAccountVsphereArgs struct {
	// The id of this vSphere cloud account.
	Id *string `pulumi:"id"`
	// The name of this vSphere cloud account.
	Name *string `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetCloudAccountVsphereTag `pulumi:"tags"`
}

// A collection of values returned by getCloudAccountVsphere.
type LookupCloudAccountVsphereResult struct {
	// Cloud accounts associated with this cloud account.
	AssociatedCloudAccountIds []string `pulumi:"associatedCloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	Dcid      string `pulumi:"dcid"`
	// A human-friendly description.
	Description string `pulumi:"description"`
	// The IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
	Hostname string                       `pulumi:"hostname"`
	Id       string                       `pulumi:"id"`
	Links    []GetCloudAccountVsphereLink `pulumi:"links"`
	Name     string                       `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId string `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// A set of region IDs that are enabled for this account.
	Regions []string `pulumi:"regions"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetCloudAccountVsphereTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
	// The vSphere username to authenticate the vsphere account.
	Username string `pulumi:"username"`
}

func LookupCloudAccountVsphereOutput(ctx *pulumi.Context, args LookupCloudAccountVsphereOutputArgs, opts ...pulumi.InvokeOption) LookupCloudAccountVsphereResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudAccountVsphereResult, error) {
			args := v.(LookupCloudAccountVsphereArgs)
			r, err := LookupCloudAccountVsphere(ctx, &args, opts...)
			var s LookupCloudAccountVsphereResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudAccountVsphereResultOutput)
}

// A collection of arguments for invoking getCloudAccountVsphere.
type LookupCloudAccountVsphereOutputArgs struct {
	// The id of this vSphere cloud account.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of this vSphere cloud account.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags GetCloudAccountVsphereTagArrayInput `pulumi:"tags"`
}

func (LookupCloudAccountVsphereOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudAccountVsphereArgs)(nil)).Elem()
}

// A collection of values returned by getCloudAccountVsphere.
type LookupCloudAccountVsphereResultOutput struct{ *pulumi.OutputState }

func (LookupCloudAccountVsphereResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudAccountVsphereResult)(nil)).Elem()
}

func (o LookupCloudAccountVsphereResultOutput) ToLookupCloudAccountVsphereResultOutput() LookupCloudAccountVsphereResultOutput {
	return o
}

func (o LookupCloudAccountVsphereResultOutput) ToLookupCloudAccountVsphereResultOutputWithContext(ctx context.Context) LookupCloudAccountVsphereResultOutput {
	return o
}

// Cloud accounts associated with this cloud account.
func (o LookupCloudAccountVsphereResultOutput) AssociatedCloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) []string { return v.AssociatedCloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o LookupCloudAccountVsphereResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupCloudAccountVsphereResultOutput) Dcid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) string { return v.Dcid }).(pulumi.StringOutput)
}

// A human-friendly description.
func (o LookupCloudAccountVsphereResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) string { return v.Description }).(pulumi.StringOutput)
}

// The IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
func (o LookupCloudAccountVsphereResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o LookupCloudAccountVsphereResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudAccountVsphereResultOutput) Links() GetCloudAccountVsphereLinkArrayOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) []GetCloudAccountVsphereLink { return v.Links }).(GetCloudAccountVsphereLinkArrayOutput)
}

func (o LookupCloudAccountVsphereResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o LookupCloudAccountVsphereResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o LookupCloudAccountVsphereResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) string { return v.Owner }).(pulumi.StringOutput)
}

// A set of region IDs that are enabled for this account.
func (o LookupCloudAccountVsphereResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

// A set of tag keys and optional values that were set on this resource.
// example:[ { "key" : "vmware", "value": "provider" } ]
func (o LookupCloudAccountVsphereResultOutput) Tags() GetCloudAccountVsphereTagArrayOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) []GetCloudAccountVsphereTag { return v.Tags }).(GetCloudAccountVsphereTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o LookupCloudAccountVsphereResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The vSphere username to authenticate the vsphere account.
func (o LookupCloudAccountVsphereResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountVsphereResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudAccountVsphereResultOutput{})
}