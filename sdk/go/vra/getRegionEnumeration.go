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
//
// This is an example of how to lookup a region enumeration data source.
//
// DeprecationMessage: 'region_enumeration' is deprecated. Use 'region_enumeration_vsphere' instead.
//
// **Region enumeration data source for vSphere**
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
//			_, err := vra.GetRegionEnumerationVsphere(ctx, &GetRegionEnumerationVsphereArgs{
//				AcceptSelfSignedCert: pulumi.BoolRef(false),
//				Dcid:                 pulumi.StringRef(_var.Vra_data_collector_id),
//				Hostname:             _var.Hostname,
//				Password:             _var.Password,
//				Username:             _var.Username,
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
// The region enumeration data source supports the following arguments:
func GetRegionEnumeration(ctx *pulumi.Context, args *GetRegionEnumerationArgs, opts ...pulumi.InvokeOption) (*GetRegionEnumerationResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRegionEnumerationResult
	err := ctx.Invoke("vra:index/getRegionEnumeration:getRegionEnumeration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegionEnumeration.
type GetRegionEnumerationArgs struct {
	// Accept self signed certificate when connecting to vSphere. Example: false
	AcceptSelfSignedCert *bool `pulumi:"acceptSelfSignedCert"`
	// ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
	Dcid *string `pulumi:"dcid"`
	// Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com
	Hostname string `pulumi:"hostname"`
	// Password for the user used to authenticate with the cloud Account
	Password string `pulumi:"password"`
	// Username to authenticate with the cloud account
	Username string `pulumi:"username"`
}

// A collection of values returned by getRegionEnumeration.
type GetRegionEnumerationResult struct {
	AcceptSelfSignedCert *bool   `pulumi:"acceptSelfSignedCert"`
	Dcid                 *string `pulumi:"dcid"`
	Hostname             string  `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Password string `pulumi:"password"`
	// A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2
	Regions  []string `pulumi:"regions"`
	Username string   `pulumi:"username"`
}

func GetRegionEnumerationOutput(ctx *pulumi.Context, args GetRegionEnumerationOutputArgs, opts ...pulumi.InvokeOption) GetRegionEnumerationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegionEnumerationResult, error) {
			args := v.(GetRegionEnumerationArgs)
			r, err := GetRegionEnumeration(ctx, &args, opts...)
			var s GetRegionEnumerationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegionEnumerationResultOutput)
}

// A collection of arguments for invoking getRegionEnumeration.
type GetRegionEnumerationOutputArgs struct {
	// Accept self signed certificate when connecting to vSphere. Example: false
	AcceptSelfSignedCert pulumi.BoolPtrInput `pulumi:"acceptSelfSignedCert"`
	// ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
	Dcid pulumi.StringPtrInput `pulumi:"dcid"`
	// Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com
	Hostname pulumi.StringInput `pulumi:"hostname"`
	// Password for the user used to authenticate with the cloud Account
	Password pulumi.StringInput `pulumi:"password"`
	// Username to authenticate with the cloud account
	Username pulumi.StringInput `pulumi:"username"`
}

func (GetRegionEnumerationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionEnumerationArgs)(nil)).Elem()
}

// A collection of values returned by getRegionEnumeration.
type GetRegionEnumerationResultOutput struct{ *pulumi.OutputState }

func (GetRegionEnumerationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionEnumerationResult)(nil)).Elem()
}

func (o GetRegionEnumerationResultOutput) ToGetRegionEnumerationResultOutput() GetRegionEnumerationResultOutput {
	return o
}

func (o GetRegionEnumerationResultOutput) ToGetRegionEnumerationResultOutputWithContext(ctx context.Context) GetRegionEnumerationResultOutput {
	return o
}

func (o GetRegionEnumerationResultOutput) AcceptSelfSignedCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetRegionEnumerationResult) *bool { return v.AcceptSelfSignedCert }).(pulumi.BoolPtrOutput)
}

func (o GetRegionEnumerationResultOutput) Dcid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegionEnumerationResult) *string { return v.Dcid }).(pulumi.StringPtrOutput)
}

func (o GetRegionEnumerationResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRegionEnumerationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRegionEnumerationResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationResult) string { return v.Password }).(pulumi.StringOutput)
}

// A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2
func (o GetRegionEnumerationResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRegionEnumerationResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

func (o GetRegionEnumerationResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegionEnumerationResultOutput{})
}