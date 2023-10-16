// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
// This is an example of how to lookup Network domain objects.
//
// **Network Domain by filter query:**
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/network"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/network"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := network.GetDomain(ctx, &network.GetDomainArgs{
//				Filter: fmt.Sprintf("name eq '%v'", _var.Name),
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
// A network domain object supports the following arguments:
func GetDomain(ctx *pulumi.Context, args *GetDomainArgs, opts ...pulumi.InvokeOption) (*GetDomainResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDomainResult
	err := ctx.Invoke("vra:network/getDomain:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomain.
type GetDomainArgs struct {
	// Search criteria to narrow down the network domain objects.
	Filter string `pulumi:"filter"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetDomainTag `pulumi:"tags"`
}

// A collection of values returned by getDomain.
type GetDomainResult struct {
	Cidr string `pulumi:"cidr"`
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds []string `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt        string                 `pulumi:"createdAt"`
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// A human-friendly description of the fabric vSphere storage account.
	Description string `pulumi:"description"`
	// External entity Id on the provider side.
	ExternalId string `pulumi:"externalId"`
	// The id of the region for which this entity is defined.
	ExternalRegionId string `pulumi:"externalRegionId"`
	Filter           string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// HATEOAS of the entity
	Links []GetDomainLink `pulumi:"links"`
	// Name of the network domain.
	Name           string `pulumi:"name"`
	OrganizationId string `pulumi:"organizationId"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetDomainTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetDomainOutput(ctx *pulumi.Context, args GetDomainOutputArgs, opts ...pulumi.InvokeOption) GetDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainResult, error) {
			args := v.(GetDomainArgs)
			r, err := GetDomain(ctx, &args, opts...)
			var s GetDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainResultOutput)
}

// A collection of arguments for invoking getDomain.
type GetDomainOutputArgs struct {
	// Search criteria to narrow down the network domain objects.
	Filter pulumi.StringInput `pulumi:"filter"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags GetDomainTagArrayInput `pulumi:"tags"`
}

func (GetDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainArgs)(nil)).Elem()
}

// A collection of values returned by getDomain.
type GetDomainResultOutput struct{ *pulumi.OutputState }

func (GetDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainResult)(nil)).Elem()
}

func (o GetDomainResultOutput) ToGetDomainResultOutput() GetDomainResultOutput {
	return o
}

func (o GetDomainResultOutput) ToGetDomainResultOutputWithContext(ctx context.Context) GetDomainResultOutput {
	return o
}

func (o GetDomainResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.Cidr }).(pulumi.StringOutput)
}

// Set of ids of the cloud accounts this entity belongs to.
func (o GetDomainResultOutput) CloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDomainResult) []string { return v.CloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o GetDomainResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o GetDomainResultOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v GetDomainResult) map[string]interface{} { return v.CustomProperties }).(pulumi.MapOutput)
}

// A human-friendly description of the fabric vSphere storage account.
func (o GetDomainResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.Description }).(pulumi.StringOutput)
}

// External entity Id on the provider side.
func (o GetDomainResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// The id of the region for which this entity is defined.
func (o GetDomainResultOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

func (o GetDomainResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.Filter }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o GetDomainResultOutput) Links() GetDomainLinkArrayOutput {
	return o.ApplyT(func(v GetDomainResult) []GetDomainLink { return v.Links }).(GetDomainLinkArrayOutput)
}

// Name of the network domain.
func (o GetDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDomainResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o GetDomainResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the resource.
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o GetDomainResultOutput) Tags() GetDomainTagArrayOutput {
	return o.ApplyT(func(v GetDomainResult) []GetDomainTag { return v.Tags }).(GetDomainTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o GetDomainResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainResultOutput{})
}
