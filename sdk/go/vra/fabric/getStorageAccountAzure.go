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
// This is an example of how to lookup fabric Azure storage account.
//
// **Fabric Azure storage account by Id:**
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
//			_, err := fabric.GetStorageAccountAzure(ctx, &fabric.GetStorageAccountAzureArgs{
//				Id: pulumi.StringRef(_var.Fabric_storage_account_azure_id),
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
// **Fabric Azure storage by filter query:**
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
//			_, err := fabric.GetStorageAccountAzure(ctx, &fabric.GetStorageAccountAzureArgs{
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
// A fabric Azure storage account supports the following arguments:
func GetStorageAccountAzure(ctx *pulumi.Context, args *GetStorageAccountAzureArgs, opts ...pulumi.InvokeOption) (*GetStorageAccountAzureResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetStorageAccountAzureResult
	err := ctx.Invoke("vra:fabric/getStorageAccountAzure:getStorageAccountAzure", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorageAccountAzure.
type GetStorageAccountAzureArgs struct {
	// Search criteria to narrow down the fabric Azure storage accounts. Only one of 'filter' or 'id' must be specified.
	Filter *string `pulumi:"filter"`
	// The id of the fabric Azure storage account. Only one of 'filter' or 'id' must be specified.
	Id *string `pulumi:"id"`
}

// A collection of values returned by getStorageAccountAzure.
type GetStorageAccountAzureResult struct {
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds []string `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// A human-friendly description of the fabric Azure storage account.
	Description string `pulumi:"description"`
	// External entity Id on the provider side.
	ExternalId string `pulumi:"externalId"`
	// The id of the region for which this entity is defined.
	ExternalRegionId string  `pulumi:"externalRegionId"`
	Filter           *string `pulumi:"filter"`
	Id               string  `pulumi:"id"`
	// HATEOAS of the entity
	Links []GetStorageAccountAzureLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId string `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed. example: Standard_LRS / Premium_LRS
	Type string `pulumi:"type"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetStorageAccountAzureOutput(ctx *pulumi.Context, args GetStorageAccountAzureOutputArgs, opts ...pulumi.InvokeOption) GetStorageAccountAzureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStorageAccountAzureResult, error) {
			args := v.(GetStorageAccountAzureArgs)
			r, err := GetStorageAccountAzure(ctx, &args, opts...)
			var s GetStorageAccountAzureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStorageAccountAzureResultOutput)
}

// A collection of arguments for invoking getStorageAccountAzure.
type GetStorageAccountAzureOutputArgs struct {
	// Search criteria to narrow down the fabric Azure storage accounts. Only one of 'filter' or 'id' must be specified.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The id of the fabric Azure storage account. Only one of 'filter' or 'id' must be specified.
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (GetStorageAccountAzureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStorageAccountAzureArgs)(nil)).Elem()
}

// A collection of values returned by getStorageAccountAzure.
type GetStorageAccountAzureResultOutput struct{ *pulumi.OutputState }

func (GetStorageAccountAzureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStorageAccountAzureResult)(nil)).Elem()
}

func (o GetStorageAccountAzureResultOutput) ToGetStorageAccountAzureResultOutput() GetStorageAccountAzureResultOutput {
	return o
}

func (o GetStorageAccountAzureResultOutput) ToGetStorageAccountAzureResultOutputWithContext(ctx context.Context) GetStorageAccountAzureResultOutput {
	return o
}

// Set of ids of the cloud accounts this entity belongs to.
func (o GetStorageAccountAzureResultOutput) CloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) []string { return v.CloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o GetStorageAccountAzureResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A human-friendly description of the fabric Azure storage account.
func (o GetStorageAccountAzureResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) string { return v.Description }).(pulumi.StringOutput)
}

// External entity Id on the provider side.
func (o GetStorageAccountAzureResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// The id of the region for which this entity is defined.
func (o GetStorageAccountAzureResultOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

func (o GetStorageAccountAzureResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o GetStorageAccountAzureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) string { return v.Id }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o GetStorageAccountAzureResultOutput) Links() GetStorageAccountAzureLinkArrayOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) []GetStorageAccountAzureLink { return v.Links }).(GetStorageAccountAzureLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o GetStorageAccountAzureResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o GetStorageAccountAzureResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o GetStorageAccountAzureResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed. example: Standard_LRS / Premium_LRS
func (o GetStorageAccountAzureResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) string { return v.Type }).(pulumi.StringOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o GetStorageAccountAzureResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageAccountAzureResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStorageAccountAzureResultOutput{})
}
