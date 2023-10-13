// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
// This is an example of how to lookup fabric networks.
//
// **Fabric network by filter query:**
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/fabric"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/fabric"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fabric.GetNetwork(ctx, &fabric.GetNetworkArgs{
//				Filter: fmt.Sprintf("name eq '%v' and externalRegionId eq '%v'", _var.Name, _var.External_region_id),
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
// A fabric network data source supports the following arguments:
func GetNetwork(ctx *pulumi.Context, args *GetNetworkArgs, opts ...pulumi.InvokeOption) (*GetNetworkResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetNetworkResult
	err := ctx.Invoke("vra:fabric/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type GetNetworkArgs struct {
	// Filter query string that is supported by vRA multi-cloud IaaS API.
	Filter string `pulumi:"filter"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetNetworkTag `pulumi:"tags"`
}

// A collection of values returned by getNetwork.
type GetNetworkResult struct {
	// Network CIDR to be used.
	Cidr string `pulumi:"cidr"`
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds []string `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// Additional properties that may be used to extend the base resource.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// State object representing a network on a external cloud provider.
	Description string `pulumi:"description"`
	// External entity Id on the provider side.
	ExternalId string `pulumi:"externalId"`
	// The id of the region for which this network is defined.
	ExternalRegionId string `pulumi:"externalRegionId"`
	Filter           string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Indicates whether this is the default subnet for the zone.
	IsDefault bool `pulumi:"isDefault"`
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic bool `pulumi:"isPublic"`
	// HATEOAS of the entity
	Links []GetNetworkLink `pulumi:"links"`
	// Name of the fabric network.
	Name string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrganizationId string `pulumi:"organizationId"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetNetworkTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetNetworkOutput(ctx *pulumi.Context, args GetNetworkOutputArgs, opts ...pulumi.InvokeOption) GetNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNetworkResult, error) {
			args := v.(GetNetworkArgs)
			r, err := GetNetwork(ctx, &args, opts...)
			var s GetNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNetworkResultOutput)
}

// A collection of arguments for invoking getNetwork.
type GetNetworkOutputArgs struct {
	// Filter query string that is supported by vRA multi-cloud IaaS API.
	Filter pulumi.StringInput `pulumi:"filter"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags GetNetworkTagArrayInput `pulumi:"tags"`
}

func (GetNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getNetwork.
type GetNetworkResultOutput struct{ *pulumi.OutputState }

func (GetNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkResult)(nil)).Elem()
}

func (o GetNetworkResultOutput) ToGetNetworkResultOutput() GetNetworkResultOutput {
	return o
}

func (o GetNetworkResultOutput) ToGetNetworkResultOutputWithContext(ctx context.Context) GetNetworkResultOutput {
	return o
}

// Network CIDR to be used.
func (o GetNetworkResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.Cidr }).(pulumi.StringOutput)
}

// Set of ids of the cloud accounts this entity belongs to.
func (o GetNetworkResultOutput) CloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNetworkResult) []string { return v.CloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o GetNetworkResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Additional properties that may be used to extend the base resource.
func (o GetNetworkResultOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v GetNetworkResult) map[string]interface{} { return v.CustomProperties }).(pulumi.MapOutput)
}

// State object representing a network on a external cloud provider.
func (o GetNetworkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.Description }).(pulumi.StringOutput)
}

// External entity Id on the provider side.
func (o GetNetworkResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// The id of the region for which this network is defined.
func (o GetNetworkResultOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

func (o GetNetworkResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.Filter }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whether this is the default subnet for the zone.
func (o GetNetworkResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNetworkResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

// Indicates whether the sub-network supports public IP assignment.
func (o GetNetworkResultOutput) IsPublic() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNetworkResult) bool { return v.IsPublic }).(pulumi.BoolOutput)
}

// HATEOAS of the entity
func (o GetNetworkResultOutput) Links() GetNetworkLinkArrayOutput {
	return o.ApplyT(func(v GetNetworkResult) []GetNetworkLink { return v.Links }).(GetNetworkLinkArrayOutput)
}

// Name of the fabric network.
func (o GetNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o GetNetworkResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o GetNetworkResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the resource.
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o GetNetworkResultOutput) Tags() GetNetworkTagArrayOutput {
	return o.ApplyT(func(v GetNetworkResult) []GetNetworkTag { return v.Tags }).(GetNetworkTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o GetNetworkResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNetworkResultOutput{})
}
