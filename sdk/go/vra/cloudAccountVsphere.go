// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VMware vRealize Automation vSphere cloud account resource.
//
// ## Example Usage
// ### S
//
// The following example shows how to create a vSphere cloud account resource.
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
//			_, err := vra.NewCloudAccountVsphere(ctx, "this", &vra.CloudAccountVsphereArgs{
//				Description: pulumi.String("foobar"),
//				Username:    pulumi.Any(_var.Username),
//				Password:    pulumi.Any(_var.Password),
//				Hostname:    pulumi.Any(_var.Hostname),
//				Dcid:        pulumi.Any(_var.Vra_data_collector_id),
//				Regions:     pulumi.Any(_var.Regions),
//				AssociatedCloudAccountIds: pulumi.StringArray{
//					pulumi.Any(_var.Vra_cloud_account_nsxt_id),
//				},
//				AcceptSelfSignedCert: pulumi.Bool(true),
//				Tags: CloudAccountVsphereTagArray{
//					&CloudAccountVsphereTagArgs{
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
// # To import the vSphere cloud account, use the ID as in the following example
//
// ```sh
//
//	$ pulumi import vra:index/cloudAccountVsphere:CloudAccountVsphere new_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15`
//
// ```
type CloudAccountVsphere struct {
	pulumi.CustomResourceState

	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert pulumi.BoolPtrOutput `pulumi:"acceptSelfSignedCert"`
	// Cloud accounts associated with the cloud account.
	AssociatedCloudAccountIds pulumi.StringArrayOutput `pulumi:"associatedCloudAccountIds"`
	// Date when  entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Identifier of a data collector vm deployed in the on premise infrastructure.
	Dcid pulumi.StringPtrOutput `pulumi:"dcid"`
	// Human-friendly description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// HATEOAS of entity.
	Links CloudAccountVsphereLinkArrayOutput `pulumi:"links"`
	// Name of the vSphere cloud account.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Email of entity owner.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Password used to authenticate to the cloud account.
	Password pulumi.StringOutput `pulumi:"password"`
	// A set of region names that are enabled for the cloud account.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// A set of tag keys and optional values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags CloudAccountVsphereTagArrayOutput `pulumi:"tags"`
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// vSphere username used to authenticate to the cloud account.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewCloudAccountVsphere registers a new resource with the given unique name, arguments, and options.
func NewCloudAccountVsphere(ctx *pulumi.Context,
	name string, args *CloudAccountVsphereArgs, opts ...pulumi.ResourceOption) (*CloudAccountVsphere, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Regions == nil {
		return nil, errors.New("invalid value for required argument 'Regions'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CloudAccountVsphere
	err := ctx.RegisterResource("vra:index/cloudAccountVsphere:CloudAccountVsphere", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudAccountVsphere gets an existing CloudAccountVsphere resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudAccountVsphere(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudAccountVsphereState, opts ...pulumi.ResourceOption) (*CloudAccountVsphere, error) {
	var resource CloudAccountVsphere
	err := ctx.ReadResource("vra:index/cloudAccountVsphere:CloudAccountVsphere", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudAccountVsphere resources.
type cloudAccountVsphereState struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert *bool `pulumi:"acceptSelfSignedCert"`
	// Cloud accounts associated with the cloud account.
	AssociatedCloudAccountIds []string `pulumi:"associatedCloudAccountIds"`
	// Date when  entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Identifier of a data collector vm deployed in the on premise infrastructure.
	Dcid *string `pulumi:"dcid"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
	Hostname *string `pulumi:"hostname"`
	// HATEOAS of entity.
	Links []CloudAccountVsphereLink `pulumi:"links"`
	// Name of the vSphere cloud account.
	Name *string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// Email of entity owner.
	Owner *string `pulumi:"owner"`
	// Password used to authenticate to the cloud account.
	Password *string `pulumi:"password"`
	// A set of region names that are enabled for the cloud account.
	Regions []string `pulumi:"regions"`
	// A set of tag keys and optional values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []CloudAccountVsphereTag `pulumi:"tags"`
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
	// vSphere username used to authenticate to the cloud account.
	Username *string `pulumi:"username"`
}

type CloudAccountVsphereState struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert pulumi.BoolPtrInput
	// Cloud accounts associated with the cloud account.
	AssociatedCloudAccountIds pulumi.StringArrayInput
	// Date when  entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// Identifier of a data collector vm deployed in the on premise infrastructure.
	Dcid pulumi.StringPtrInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
	Hostname pulumi.StringPtrInput
	// HATEOAS of entity.
	Links CloudAccountVsphereLinkArrayInput
	// Name of the vSphere cloud account.
	Name pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrgId pulumi.StringPtrInput
	// Email of entity owner.
	Owner pulumi.StringPtrInput
	// Password used to authenticate to the cloud account.
	Password pulumi.StringPtrInput
	// A set of region names that are enabled for the cloud account.
	Regions pulumi.StringArrayInput
	// A set of tag keys and optional values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags CloudAccountVsphereTagArrayInput
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
	// vSphere username used to authenticate to the cloud account.
	Username pulumi.StringPtrInput
}

func (CloudAccountVsphereState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudAccountVsphereState)(nil)).Elem()
}

type cloudAccountVsphereArgs struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert *bool `pulumi:"acceptSelfSignedCert"`
	// Cloud accounts associated with the cloud account.
	AssociatedCloudAccountIds []string `pulumi:"associatedCloudAccountIds"`
	// Identifier of a data collector vm deployed in the on premise infrastructure.
	Dcid *string `pulumi:"dcid"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
	Hostname string `pulumi:"hostname"`
	// Name of the vSphere cloud account.
	Name *string `pulumi:"name"`
	// Password used to authenticate to the cloud account.
	Password string `pulumi:"password"`
	// A set of region names that are enabled for the cloud account.
	Regions []string `pulumi:"regions"`
	// A set of tag keys and optional values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []CloudAccountVsphereTag `pulumi:"tags"`
	// vSphere username used to authenticate to the cloud account.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a CloudAccountVsphere resource.
type CloudAccountVsphereArgs struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert pulumi.BoolPtrInput
	// Cloud accounts associated with the cloud account.
	AssociatedCloudAccountIds pulumi.StringArrayInput
	// Identifier of a data collector vm deployed in the on premise infrastructure.
	Dcid pulumi.StringPtrInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
	Hostname pulumi.StringInput
	// Name of the vSphere cloud account.
	Name pulumi.StringPtrInput
	// Password used to authenticate to the cloud account.
	Password pulumi.StringInput
	// A set of region names that are enabled for the cloud account.
	Regions pulumi.StringArrayInput
	// A set of tag keys and optional values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags CloudAccountVsphereTagArrayInput
	// vSphere username used to authenticate to the cloud account.
	Username pulumi.StringInput
}

func (CloudAccountVsphereArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudAccountVsphereArgs)(nil)).Elem()
}

type CloudAccountVsphereInput interface {
	pulumi.Input

	ToCloudAccountVsphereOutput() CloudAccountVsphereOutput
	ToCloudAccountVsphereOutputWithContext(ctx context.Context) CloudAccountVsphereOutput
}

func (*CloudAccountVsphere) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudAccountVsphere)(nil)).Elem()
}

func (i *CloudAccountVsphere) ToCloudAccountVsphereOutput() CloudAccountVsphereOutput {
	return i.ToCloudAccountVsphereOutputWithContext(context.Background())
}

func (i *CloudAccountVsphere) ToCloudAccountVsphereOutputWithContext(ctx context.Context) CloudAccountVsphereOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountVsphereOutput)
}

// CloudAccountVsphereArrayInput is an input type that accepts CloudAccountVsphereArray and CloudAccountVsphereArrayOutput values.
// You can construct a concrete instance of `CloudAccountVsphereArrayInput` via:
//
//	CloudAccountVsphereArray{ CloudAccountVsphereArgs{...} }
type CloudAccountVsphereArrayInput interface {
	pulumi.Input

	ToCloudAccountVsphereArrayOutput() CloudAccountVsphereArrayOutput
	ToCloudAccountVsphereArrayOutputWithContext(context.Context) CloudAccountVsphereArrayOutput
}

type CloudAccountVsphereArray []CloudAccountVsphereInput

func (CloudAccountVsphereArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudAccountVsphere)(nil)).Elem()
}

func (i CloudAccountVsphereArray) ToCloudAccountVsphereArrayOutput() CloudAccountVsphereArrayOutput {
	return i.ToCloudAccountVsphereArrayOutputWithContext(context.Background())
}

func (i CloudAccountVsphereArray) ToCloudAccountVsphereArrayOutputWithContext(ctx context.Context) CloudAccountVsphereArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountVsphereArrayOutput)
}

// CloudAccountVsphereMapInput is an input type that accepts CloudAccountVsphereMap and CloudAccountVsphereMapOutput values.
// You can construct a concrete instance of `CloudAccountVsphereMapInput` via:
//
//	CloudAccountVsphereMap{ "key": CloudAccountVsphereArgs{...} }
type CloudAccountVsphereMapInput interface {
	pulumi.Input

	ToCloudAccountVsphereMapOutput() CloudAccountVsphereMapOutput
	ToCloudAccountVsphereMapOutputWithContext(context.Context) CloudAccountVsphereMapOutput
}

type CloudAccountVsphereMap map[string]CloudAccountVsphereInput

func (CloudAccountVsphereMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudAccountVsphere)(nil)).Elem()
}

func (i CloudAccountVsphereMap) ToCloudAccountVsphereMapOutput() CloudAccountVsphereMapOutput {
	return i.ToCloudAccountVsphereMapOutputWithContext(context.Background())
}

func (i CloudAccountVsphereMap) ToCloudAccountVsphereMapOutputWithContext(ctx context.Context) CloudAccountVsphereMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountVsphereMapOutput)
}

type CloudAccountVsphereOutput struct{ *pulumi.OutputState }

func (CloudAccountVsphereOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudAccountVsphere)(nil)).Elem()
}

func (o CloudAccountVsphereOutput) ToCloudAccountVsphereOutput() CloudAccountVsphereOutput {
	return o
}

func (o CloudAccountVsphereOutput) ToCloudAccountVsphereOutputWithContext(ctx context.Context) CloudAccountVsphereOutput {
	return o
}

// Accept self-signed certificate when connecting to the cloud account.
func (o CloudAccountVsphereOutput) AcceptSelfSignedCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.BoolPtrOutput { return v.AcceptSelfSignedCert }).(pulumi.BoolPtrOutput)
}

// Cloud accounts associated with the cloud account.
func (o CloudAccountVsphereOutput) AssociatedCloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringArrayOutput { return v.AssociatedCloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when  entity was created. Date and time format is ISO 8601 and UTC.
func (o CloudAccountVsphereOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Identifier of a data collector vm deployed in the on premise infrastructure.
func (o CloudAccountVsphereOutput) Dcid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringPtrOutput { return v.Dcid }).(pulumi.StringPtrOutput)
}

// Human-friendly description.
func (o CloudAccountVsphereOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
func (o CloudAccountVsphereOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// HATEOAS of entity.
func (o CloudAccountVsphereOutput) Links() CloudAccountVsphereLinkArrayOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) CloudAccountVsphereLinkArrayOutput { return v.Links }).(CloudAccountVsphereLinkArrayOutput)
}

// Name of the vSphere cloud account.
func (o CloudAccountVsphereOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o CloudAccountVsphereOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Email of entity owner.
func (o CloudAccountVsphereOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Password used to authenticate to the cloud account.
func (o CloudAccountVsphereOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// A set of region names that are enabled for the cloud account.
func (o CloudAccountVsphereOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// A set of tag keys and optional values to apply to the cloud account.\
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o CloudAccountVsphereOutput) Tags() CloudAccountVsphereTagArrayOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) CloudAccountVsphereTagArrayOutput { return v.Tags }).(CloudAccountVsphereTagArrayOutput)
}

// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
func (o CloudAccountVsphereOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// vSphere username used to authenticate to the cloud account.
func (o CloudAccountVsphereOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountVsphere) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type CloudAccountVsphereArrayOutput struct{ *pulumi.OutputState }

func (CloudAccountVsphereArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudAccountVsphere)(nil)).Elem()
}

func (o CloudAccountVsphereArrayOutput) ToCloudAccountVsphereArrayOutput() CloudAccountVsphereArrayOutput {
	return o
}

func (o CloudAccountVsphereArrayOutput) ToCloudAccountVsphereArrayOutputWithContext(ctx context.Context) CloudAccountVsphereArrayOutput {
	return o
}

func (o CloudAccountVsphereArrayOutput) Index(i pulumi.IntInput) CloudAccountVsphereOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudAccountVsphere {
		return vs[0].([]*CloudAccountVsphere)[vs[1].(int)]
	}).(CloudAccountVsphereOutput)
}

type CloudAccountVsphereMapOutput struct{ *pulumi.OutputState }

func (CloudAccountVsphereMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudAccountVsphere)(nil)).Elem()
}

func (o CloudAccountVsphereMapOutput) ToCloudAccountVsphereMapOutput() CloudAccountVsphereMapOutput {
	return o
}

func (o CloudAccountVsphereMapOutput) ToCloudAccountVsphereMapOutputWithContext(ctx context.Context) CloudAccountVsphereMapOutput {
	return o
}

func (o CloudAccountVsphereMapOutput) MapIndex(k pulumi.StringInput) CloudAccountVsphereOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudAccountVsphere {
		return vs[0].(map[string]*CloudAccountVsphere)[vs[1].(string)]
	}).(CloudAccountVsphereOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudAccountVsphereInput)(nil)).Elem(), &CloudAccountVsphere{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudAccountVsphereArrayInput)(nil)).Elem(), CloudAccountVsphereArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudAccountVsphereMapInput)(nil)).Elem(), CloudAccountVsphereMap{})
	pulumi.RegisterOutputType(CloudAccountVsphereOutput{})
	pulumi.RegisterOutputType(CloudAccountVsphereArrayOutput{})
	pulumi.RegisterOutputType(CloudAccountVsphereMapOutput{})
}