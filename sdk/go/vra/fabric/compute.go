// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Updates a VMware vRealize Automation fabricCompute resource.
//
// ## Example Usage
// ### S
//
// You cannot create a fabric compute resource, however you can import it using the command specified in the import section below.
//
// Once a resource is imported, you can update it as shown below:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/fabric"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/fabric"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fabric.NewCompute(ctx, "this", &fabric.ComputeArgs{
//				Tags: fabric.ComputeTagArray{
//					&fabric.ComputeTagArgs{
//						Key:   pulumi.String("foo"),
//						Value: pulumi.String("bar"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # To import the fabric compute resource, use the ID as in the following example
//
// ```sh
//
//	$ pulumi import vra:fabric/compute:Compute this 88fdea8b-92ed-4aa9-b6ee-4670412961b0
//
// ```
type Compute struct {
	pulumi.CustomResourceState

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A list of key value pair of custom properties for the fabric compute resource.
	CustomProperties pulumi.MapOutput `pulumi:"customProperties"`
	// A human-friendly description.
	Description pulumi.StringOutput `pulumi:"description"`
	// The id of the external entity on the provider side.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// The external region id of the fabric compute.
	ExternalRegionId pulumi.StringOutput `pulumi:"externalRegionId"`
	// The external zone id of the fabric compute.
	ExternalZoneId pulumi.StringOutput `pulumi:"externalZoneId"`
	// Lifecycle status of the compute instance.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// HATEOAS of the entity.
	Links ComputeLinkArrayOutput `pulumi:"links"`
	// A human-friendly name used as an identifier for the fabric compute resource instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Power state of fabric compute instance.
	PowerState pulumi.StringOutput `pulumi:"powerState"`
	// A set of tag keys and optional values that were set on this resource:
	Tags ComputeTagArrayOutput `pulumi:"tags"`
	// Type of the fabric compute instance.
	Type pulumi.StringOutput `pulumi:"type"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewCompute registers a new resource with the given unique name, arguments, and options.
func NewCompute(ctx *pulumi.Context,
	name string, args *ComputeArgs, opts ...pulumi.ResourceOption) (*Compute, error) {
	if args == nil {
		args = &ComputeArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Compute
	err := ctx.RegisterResource("vra:fabric/compute:Compute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCompute gets an existing Compute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCompute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeState, opts ...pulumi.ResourceOption) (*Compute, error) {
	var resource Compute
	err := ctx.ReadResource("vra:fabric/compute:Compute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Compute resources.
type computeState struct {
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// A list of key value pair of custom properties for the fabric compute resource.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// A human-friendly description.
	Description *string `pulumi:"description"`
	// The id of the external entity on the provider side.
	ExternalId *string `pulumi:"externalId"`
	// The external region id of the fabric compute.
	ExternalRegionId *string `pulumi:"externalRegionId"`
	// The external zone id of the fabric compute.
	ExternalZoneId *string `pulumi:"externalZoneId"`
	// Lifecycle status of the compute instance.
	LifecycleState *string `pulumi:"lifecycleState"`
	// HATEOAS of the entity.
	Links []ComputeLink `pulumi:"links"`
	// A human-friendly name used as an identifier for the fabric compute resource instance.
	Name *string `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner *string `pulumi:"owner"`
	// Power state of fabric compute instance.
	PowerState *string `pulumi:"powerState"`
	// A set of tag keys and optional values that were set on this resource:
	Tags []ComputeTag `pulumi:"tags"`
	// Type of the fabric compute instance.
	Type *string `pulumi:"type"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ComputeState struct {
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// A list of key value pair of custom properties for the fabric compute resource.
	CustomProperties pulumi.MapInput
	// A human-friendly description.
	Description pulumi.StringPtrInput
	// The id of the external entity on the provider side.
	ExternalId pulumi.StringPtrInput
	// The external region id of the fabric compute.
	ExternalRegionId pulumi.StringPtrInput
	// The external zone id of the fabric compute.
	ExternalZoneId pulumi.StringPtrInput
	// Lifecycle status of the compute instance.
	LifecycleState pulumi.StringPtrInput
	// HATEOAS of the entity.
	Links ComputeLinkArrayInput
	// A human-friendly name used as an identifier for the fabric compute resource instance.
	Name pulumi.StringPtrInput
	// The id of the organization this entity belongs to.
	OrgId pulumi.StringPtrInput
	// Email of the user that owns the entity.
	Owner pulumi.StringPtrInput
	// Power state of fabric compute instance.
	PowerState pulumi.StringPtrInput
	// A set of tag keys and optional values that were set on this resource:
	Tags ComputeTagArrayInput
	// Type of the fabric compute instance.
	Type pulumi.StringPtrInput
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (ComputeState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeState)(nil)).Elem()
}

type computeArgs struct {
	// A set of tag keys and optional values that were set on this resource:
	Tags []ComputeTag `pulumi:"tags"`
}

// The set of arguments for constructing a Compute resource.
type ComputeArgs struct {
	// A set of tag keys and optional values that were set on this resource:
	Tags ComputeTagArrayInput
}

func (ComputeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeArgs)(nil)).Elem()
}

type ComputeInput interface {
	pulumi.Input

	ToComputeOutput() ComputeOutput
	ToComputeOutputWithContext(ctx context.Context) ComputeOutput
}

func (*Compute) ElementType() reflect.Type {
	return reflect.TypeOf((**Compute)(nil)).Elem()
}

func (i *Compute) ToComputeOutput() ComputeOutput {
	return i.ToComputeOutputWithContext(context.Background())
}

func (i *Compute) ToComputeOutputWithContext(ctx context.Context) ComputeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeOutput)
}

// ComputeArrayInput is an input type that accepts ComputeArray and ComputeArrayOutput values.
// You can construct a concrete instance of `ComputeArrayInput` via:
//
//	ComputeArray{ ComputeArgs{...} }
type ComputeArrayInput interface {
	pulumi.Input

	ToComputeArrayOutput() ComputeArrayOutput
	ToComputeArrayOutputWithContext(context.Context) ComputeArrayOutput
}

type ComputeArray []ComputeInput

func (ComputeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Compute)(nil)).Elem()
}

func (i ComputeArray) ToComputeArrayOutput() ComputeArrayOutput {
	return i.ToComputeArrayOutputWithContext(context.Background())
}

func (i ComputeArray) ToComputeArrayOutputWithContext(ctx context.Context) ComputeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeArrayOutput)
}

// ComputeMapInput is an input type that accepts ComputeMap and ComputeMapOutput values.
// You can construct a concrete instance of `ComputeMapInput` via:
//
//	ComputeMap{ "key": ComputeArgs{...} }
type ComputeMapInput interface {
	pulumi.Input

	ToComputeMapOutput() ComputeMapOutput
	ToComputeMapOutputWithContext(context.Context) ComputeMapOutput
}

type ComputeMap map[string]ComputeInput

func (ComputeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Compute)(nil)).Elem()
}

func (i ComputeMap) ToComputeMapOutput() ComputeMapOutput {
	return i.ToComputeMapOutputWithContext(context.Background())
}

func (i ComputeMap) ToComputeMapOutputWithContext(ctx context.Context) ComputeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeMapOutput)
}

type ComputeOutput struct{ *pulumi.OutputState }

func (ComputeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Compute)(nil)).Elem()
}

func (o ComputeOutput) ToComputeOutput() ComputeOutput {
	return o
}

func (o ComputeOutput) ToComputeOutputWithContext(ctx context.Context) ComputeOutput {
	return o
}

// Date when the entity was created. The date is in ISO 8601 and UTC.
func (o ComputeOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A list of key value pair of custom properties for the fabric compute resource.
func (o ComputeOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v *Compute) pulumi.MapOutput { return v.CustomProperties }).(pulumi.MapOutput)
}

// A human-friendly description.
func (o ComputeOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The id of the external entity on the provider side.
func (o ComputeOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.ExternalId }).(pulumi.StringOutput)
}

// The external region id of the fabric compute.
func (o ComputeOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.ExternalRegionId }).(pulumi.StringOutput)
}

// The external zone id of the fabric compute.
func (o ComputeOutput) ExternalZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.ExternalZoneId }).(pulumi.StringOutput)
}

// Lifecycle status of the compute instance.
func (o ComputeOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

// HATEOAS of the entity.
func (o ComputeOutput) Links() ComputeLinkArrayOutput {
	return o.ApplyT(func(v *Compute) ComputeLinkArrayOutput { return v.Links }).(ComputeLinkArrayOutput)
}

// A human-friendly name used as an identifier for the fabric compute resource instance.
func (o ComputeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o ComputeOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o ComputeOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Power state of fabric compute instance.
func (o ComputeOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.PowerState }).(pulumi.StringOutput)
}

// A set of tag keys and optional values that were set on this resource:
func (o ComputeOutput) Tags() ComputeTagArrayOutput {
	return o.ApplyT(func(v *Compute) ComputeTagArrayOutput { return v.Tags }).(ComputeTagArrayOutput)
}

// Type of the fabric compute instance.
func (o ComputeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o ComputeOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ComputeArrayOutput struct{ *pulumi.OutputState }

func (ComputeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Compute)(nil)).Elem()
}

func (o ComputeArrayOutput) ToComputeArrayOutput() ComputeArrayOutput {
	return o
}

func (o ComputeArrayOutput) ToComputeArrayOutputWithContext(ctx context.Context) ComputeArrayOutput {
	return o
}

func (o ComputeArrayOutput) Index(i pulumi.IntInput) ComputeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Compute {
		return vs[0].([]*Compute)[vs[1].(int)]
	}).(ComputeOutput)
}

type ComputeMapOutput struct{ *pulumi.OutputState }

func (ComputeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Compute)(nil)).Elem()
}

func (o ComputeMapOutput) ToComputeMapOutput() ComputeMapOutput {
	return o
}

func (o ComputeMapOutput) ToComputeMapOutputWithContext(ctx context.Context) ComputeMapOutput {
	return o
}

func (o ComputeMapOutput) MapIndex(k pulumi.StringInput) ComputeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Compute {
		return vs[0].(map[string]*Compute)[vs[1].(string)]
	}).(ComputeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInput)(nil)).Elem(), &Compute{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeArrayInput)(nil)).Elem(), ComputeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeMapInput)(nil)).Elem(), ComputeMap{})
	pulumi.RegisterOutputType(ComputeOutput{})
	pulumi.RegisterOutputType(ComputeArrayOutput{})
	pulumi.RegisterOutputType(ComputeMapOutput{})
}
