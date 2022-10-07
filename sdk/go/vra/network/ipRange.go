// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VMware vRealize Automation networkIpRange resource.
//
// ## Example Usage
//
// ## Import
//
// # To import the vRA Network IP range, use the ID as in the following example
//
// ```sh
//
//	$ pulumi import vra:network/ipRange:IpRange new_ip_range 05956583-6488-4e7d-84c9-92a7b7219a15`
//
// ```
type IpRange struct {
	pulumi.CustomResourceState

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// State object representing a network on a external cloud provider.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// End IP address of the range.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// External entity Id on the provider side.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// Fabric network Id.
	FabricNetworkId pulumi.StringPtrOutput `pulumi:"fabricNetworkId"`
	// IP address version: IPv4 or IPv6. Default: IPv4.
	IpVersion pulumi.StringOutput `pulumi:"ipVersion"`
	// HATEOAS of the entity
	Links IpRangeLinkArrayOutput `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Start IP address of the range.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags IpRangeTagArrayOutput `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewIpRange registers a new resource with the given unique name, arguments, and options.
func NewIpRange(ctx *pulumi.Context,
	name string, args *IpRangeArgs, opts ...pulumi.ResourceOption) (*IpRange, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.IpVersion == nil {
		return nil, errors.New("invalid value for required argument 'IpVersion'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IpRange
	err := ctx.RegisterResource("vra:network/ipRange:IpRange", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpRange gets an existing IpRange resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpRange(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpRangeState, opts ...pulumi.ResourceOption) (*IpRange, error) {
	var resource IpRange
	err := ctx.ReadResource("vra:network/ipRange:IpRange", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpRange resources.
type ipRangeState struct {
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// State object representing a network on a external cloud provider.
	Description *string `pulumi:"description"`
	// End IP address of the range.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// External entity Id on the provider side.
	ExternalId *string `pulumi:"externalId"`
	// Fabric network Id.
	FabricNetworkId *string `pulumi:"fabricNetworkId"`
	// IP address version: IPv4 or IPv6. Default: IPv4.
	IpVersion *string `pulumi:"ipVersion"`
	// HATEOAS of the entity
	Links []IpRangeLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// Start IP address of the range.
	StartIpAddress *string `pulumi:"startIpAddress"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []IpRangeTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type IpRangeState struct {
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt pulumi.StringPtrInput
	// State object representing a network on a external cloud provider.
	Description pulumi.StringPtrInput
	// End IP address of the range.
	EndIpAddress pulumi.StringPtrInput
	// External entity Id on the provider side.
	ExternalId pulumi.StringPtrInput
	// Fabric network Id.
	FabricNetworkId pulumi.StringPtrInput
	// IP address version: IPv4 or IPv6. Default: IPv4.
	IpVersion pulumi.StringPtrInput
	// HATEOAS of the entity
	Links IpRangeLinkArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrgId pulumi.StringPtrInput
	// Start IP address of the range.
	StartIpAddress pulumi.StringPtrInput
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags IpRangeTagArrayInput
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (IpRangeState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipRangeState)(nil)).Elem()
}

type ipRangeArgs struct {
	// State object representing a network on a external cloud provider.
	Description *string `pulumi:"description"`
	// End IP address of the range.
	EndIpAddress string `pulumi:"endIpAddress"`
	// Fabric network Id.
	FabricNetworkId *string `pulumi:"fabricNetworkId"`
	// IP address version: IPv4 or IPv6. Default: IPv4.
	IpVersion string `pulumi:"ipVersion"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// Start IP address of the range.
	StartIpAddress string `pulumi:"startIpAddress"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []IpRangeTag `pulumi:"tags"`
}

// The set of arguments for constructing a IpRange resource.
type IpRangeArgs struct {
	// State object representing a network on a external cloud provider.
	Description pulumi.StringPtrInput
	// End IP address of the range.
	EndIpAddress pulumi.StringInput
	// Fabric network Id.
	FabricNetworkId pulumi.StringPtrInput
	// IP address version: IPv4 or IPv6. Default: IPv4.
	IpVersion pulumi.StringInput
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// Start IP address of the range.
	StartIpAddress pulumi.StringInput
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags IpRangeTagArrayInput
}

func (IpRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipRangeArgs)(nil)).Elem()
}

type IpRangeInput interface {
	pulumi.Input

	ToIpRangeOutput() IpRangeOutput
	ToIpRangeOutputWithContext(ctx context.Context) IpRangeOutput
}

func (*IpRange) ElementType() reflect.Type {
	return reflect.TypeOf((**IpRange)(nil)).Elem()
}

func (i *IpRange) ToIpRangeOutput() IpRangeOutput {
	return i.ToIpRangeOutputWithContext(context.Background())
}

func (i *IpRange) ToIpRangeOutputWithContext(ctx context.Context) IpRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRangeOutput)
}

// IpRangeArrayInput is an input type that accepts IpRangeArray and IpRangeArrayOutput values.
// You can construct a concrete instance of `IpRangeArrayInput` via:
//
//	IpRangeArray{ IpRangeArgs{...} }
type IpRangeArrayInput interface {
	pulumi.Input

	ToIpRangeArrayOutput() IpRangeArrayOutput
	ToIpRangeArrayOutputWithContext(context.Context) IpRangeArrayOutput
}

type IpRangeArray []IpRangeInput

func (IpRangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpRange)(nil)).Elem()
}

func (i IpRangeArray) ToIpRangeArrayOutput() IpRangeArrayOutput {
	return i.ToIpRangeArrayOutputWithContext(context.Background())
}

func (i IpRangeArray) ToIpRangeArrayOutputWithContext(ctx context.Context) IpRangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRangeArrayOutput)
}

// IpRangeMapInput is an input type that accepts IpRangeMap and IpRangeMapOutput values.
// You can construct a concrete instance of `IpRangeMapInput` via:
//
//	IpRangeMap{ "key": IpRangeArgs{...} }
type IpRangeMapInput interface {
	pulumi.Input

	ToIpRangeMapOutput() IpRangeMapOutput
	ToIpRangeMapOutputWithContext(context.Context) IpRangeMapOutput
}

type IpRangeMap map[string]IpRangeInput

func (IpRangeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpRange)(nil)).Elem()
}

func (i IpRangeMap) ToIpRangeMapOutput() IpRangeMapOutput {
	return i.ToIpRangeMapOutputWithContext(context.Background())
}

func (i IpRangeMap) ToIpRangeMapOutputWithContext(ctx context.Context) IpRangeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRangeMapOutput)
}

type IpRangeOutput struct{ *pulumi.OutputState }

func (IpRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpRange)(nil)).Elem()
}

func (o IpRangeOutput) ToIpRangeOutput() IpRangeOutput {
	return o
}

func (o IpRangeOutput) ToIpRangeOutputWithContext(ctx context.Context) IpRangeOutput {
	return o
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o IpRangeOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// State object representing a network on a external cloud provider.
func (o IpRangeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// End IP address of the range.
func (o IpRangeOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringOutput { return v.EndIpAddress }).(pulumi.StringOutput)
}

// External entity Id on the provider side.
func (o IpRangeOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringOutput { return v.ExternalId }).(pulumi.StringOutput)
}

// Fabric network Id.
func (o IpRangeOutput) FabricNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringPtrOutput { return v.FabricNetworkId }).(pulumi.StringPtrOutput)
}

// IP address version: IPv4 or IPv6. Default: IPv4.
func (o IpRangeOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringOutput { return v.IpVersion }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o IpRangeOutput) Links() IpRangeLinkArrayOutput {
	return o.ApplyT(func(v *IpRange) IpRangeLinkArrayOutput { return v.Links }).(IpRangeLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o IpRangeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o IpRangeOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Start IP address of the range.
func (o IpRangeOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringOutput { return v.StartIpAddress }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the resource.
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o IpRangeOutput) Tags() IpRangeTagArrayOutput {
	return o.ApplyT(func(v *IpRange) IpRangeTagArrayOutput { return v.Tags }).(IpRangeTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o IpRangeOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type IpRangeArrayOutput struct{ *pulumi.OutputState }

func (IpRangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpRange)(nil)).Elem()
}

func (o IpRangeArrayOutput) ToIpRangeArrayOutput() IpRangeArrayOutput {
	return o
}

func (o IpRangeArrayOutput) ToIpRangeArrayOutputWithContext(ctx context.Context) IpRangeArrayOutput {
	return o
}

func (o IpRangeArrayOutput) Index(i pulumi.IntInput) IpRangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpRange {
		return vs[0].([]*IpRange)[vs[1].(int)]
	}).(IpRangeOutput)
}

type IpRangeMapOutput struct{ *pulumi.OutputState }

func (IpRangeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpRange)(nil)).Elem()
}

func (o IpRangeMapOutput) ToIpRangeMapOutput() IpRangeMapOutput {
	return o
}

func (o IpRangeMapOutput) ToIpRangeMapOutputWithContext(ctx context.Context) IpRangeMapOutput {
	return o
}

func (o IpRangeMapOutput) MapIndex(k pulumi.StringInput) IpRangeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpRange {
		return vs[0].(map[string]*IpRange)[vs[1].(string)]
	}).(IpRangeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpRangeInput)(nil)).Elem(), &IpRange{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpRangeArrayInput)(nil)).Elem(), IpRangeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpRangeMapInput)(nil)).Elem(), IpRangeMap{})
	pulumi.RegisterOutputType(IpRangeOutput{})
	pulumi.RegisterOutputType(IpRangeArrayOutput{})
	pulumi.RegisterOutputType(IpRangeMapOutput{})
}
