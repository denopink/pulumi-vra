// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
// This is an example of how to lookup fabric vSphere storage policies.
//
// **Fabric vSphere storage policy by Id:**
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
//			_, err := vra.GetFabricStoragePolicyVsphere(ctx, &GetFabricStoragePolicyVsphereArgs{
//				Id: pulumi.StringRef(_var.Fabric_storage_policy_vsphere_id),
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
// **Fabric vSphere storage policy by filter query:**
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/schmidtw/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.GetFabricStoragePolicyVsphere(ctx, &GetFabricStoragePolicyVsphereArgs{
//				Filter: pulumi.StringRef(fmt.Sprintf("name eq '%v'", _var.Name)),
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
// A fabric vSphere storage policy supports the following arguments:
func GetFabricStoragePolicyVsphere(ctx *pulumi.Context, args *GetFabricStoragePolicyVsphereArgs, opts ...pulumi.InvokeOption) (*GetFabricStoragePolicyVsphereResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFabricStoragePolicyVsphereResult
	err := ctx.Invoke("vra:index/getFabricStoragePolicyVsphere:getFabricStoragePolicyVsphere", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFabricStoragePolicyVsphere.
type GetFabricStoragePolicyVsphereArgs struct {
	// Search criteria to narrow down the fabric vSphere storage policy. Only one of 'filter' or 'id' must be specified.
	Filter *string `pulumi:"filter"`
	// The id of the fabric vSphere storage policy. Only one of 'filter' or 'id' must be specified.
	Id *string `pulumi:"id"`
}

// A collection of values returned by getFabricStoragePolicyVsphere.
type GetFabricStoragePolicyVsphereResult struct {
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds []string `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// External entity Id on the provider side.
	ExternalId string `pulumi:"externalId"`
	// The id of the region for which this entity is defined.
	ExternalRegionId string  `pulumi:"externalRegionId"`
	Filter           *string `pulumi:"filter"`
	Id               string  `pulumi:"id"`
	// HATEOAS of the entity
	Links []GetFabricStoragePolicyVsphereLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.  Only one of 'filter', 'id', 'name' or 'region_id' must be specified.
	Name string `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId string `pulumi:"orgId"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetFabricStoragePolicyVsphereOutput(ctx *pulumi.Context, args GetFabricStoragePolicyVsphereOutputArgs, opts ...pulumi.InvokeOption) GetFabricStoragePolicyVsphereResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFabricStoragePolicyVsphereResult, error) {
			args := v.(GetFabricStoragePolicyVsphereArgs)
			r, err := GetFabricStoragePolicyVsphere(ctx, &args, opts...)
			var s GetFabricStoragePolicyVsphereResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFabricStoragePolicyVsphereResultOutput)
}

// A collection of arguments for invoking getFabricStoragePolicyVsphere.
type GetFabricStoragePolicyVsphereOutputArgs struct {
	// Search criteria to narrow down the fabric vSphere storage policy. Only one of 'filter' or 'id' must be specified.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The id of the fabric vSphere storage policy. Only one of 'filter' or 'id' must be specified.
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (GetFabricStoragePolicyVsphereOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFabricStoragePolicyVsphereArgs)(nil)).Elem()
}

// A collection of values returned by getFabricStoragePolicyVsphere.
type GetFabricStoragePolicyVsphereResultOutput struct{ *pulumi.OutputState }

func (GetFabricStoragePolicyVsphereResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFabricStoragePolicyVsphereResult)(nil)).Elem()
}

func (o GetFabricStoragePolicyVsphereResultOutput) ToGetFabricStoragePolicyVsphereResultOutput() GetFabricStoragePolicyVsphereResultOutput {
	return o
}

func (o GetFabricStoragePolicyVsphereResultOutput) ToGetFabricStoragePolicyVsphereResultOutputWithContext(ctx context.Context) GetFabricStoragePolicyVsphereResultOutput {
	return o
}

// Set of ids of the cloud accounts this entity belongs to.
func (o GetFabricStoragePolicyVsphereResultOutput) CloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFabricStoragePolicyVsphereResult) []string { return v.CloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o GetFabricStoragePolicyVsphereResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricStoragePolicyVsphereResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// External entity Id on the provider side.
func (o GetFabricStoragePolicyVsphereResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricStoragePolicyVsphereResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// The id of the region for which this entity is defined.
func (o GetFabricStoragePolicyVsphereResultOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricStoragePolicyVsphereResult) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

func (o GetFabricStoragePolicyVsphereResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFabricStoragePolicyVsphereResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o GetFabricStoragePolicyVsphereResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricStoragePolicyVsphereResult) string { return v.Id }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o GetFabricStoragePolicyVsphereResultOutput) Links() GetFabricStoragePolicyVsphereLinkArrayOutput {
	return o.ApplyT(func(v GetFabricStoragePolicyVsphereResult) []GetFabricStoragePolicyVsphereLink { return v.Links }).(GetFabricStoragePolicyVsphereLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.  Only one of 'filter', 'id', 'name' or 'region_id' must be specified.
func (o GetFabricStoragePolicyVsphereResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricStoragePolicyVsphereResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o GetFabricStoragePolicyVsphereResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricStoragePolicyVsphereResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o GetFabricStoragePolicyVsphereResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricStoragePolicyVsphereResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFabricStoragePolicyVsphereResultOutput{})
}