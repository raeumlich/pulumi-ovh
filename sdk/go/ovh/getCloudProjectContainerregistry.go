// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a container registry associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-ovh/sdk/go/ovh"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ovh.GetCloudProjectContainerregistry(ctx, &GetCloudProjectContainerregistryArgs{
// 			RegistryId:  "yyyy",
// 			ServiceName: "XXXXXX",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCloudProjectContainerregistry(ctx *pulumi.Context, args *GetCloudProjectContainerregistryArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectContainerregistryResult, error) {
	var rv GetCloudProjectContainerregistryResult
	err := ctx.Invoke("ovh:index/getCloudProjectContainerregistry:getCloudProjectContainerregistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectContainerregistry.
type GetCloudProjectContainerregistryArgs struct {
	// Registry ID
	RegistryId string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectContainerregistry.
type GetCloudProjectContainerregistryResult struct {
	// Registry creation date
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Registry name
	Name string `pulumi:"name"`
	// Project ID of your registry
	ProjectId string `pulumi:"projectId"`
	// Region of the registry
	Region      string `pulumi:"region"`
	RegistryId  string `pulumi:"registryId"`
	ServiceName string `pulumi:"serviceName"`
	// Current size of the registry (bytes)
	Size int `pulumi:"size"`
	// Registry status
	Status string `pulumi:"status"`
	// Registry last update date
	UpdatedAt string `pulumi:"updatedAt"`
	// Access url of the registry
	Url string `pulumi:"url"`
	// Version of your registry
	Version string `pulumi:"version"`
}

func GetCloudProjectContainerregistryOutput(ctx *pulumi.Context, args GetCloudProjectContainerregistryOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectContainerregistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectContainerregistryResult, error) {
			args := v.(GetCloudProjectContainerregistryArgs)
			r, err := GetCloudProjectContainerregistry(ctx, &args, opts...)
			var s GetCloudProjectContainerregistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectContainerregistryResultOutput)
}

// A collection of arguments for invoking getCloudProjectContainerregistry.
type GetCloudProjectContainerregistryOutputArgs struct {
	// Registry ID
	RegistryId pulumi.StringInput `pulumi:"registryId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetCloudProjectContainerregistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectContainerregistryArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectContainerregistry.
type GetCloudProjectContainerregistryResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectContainerregistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectContainerregistryResult)(nil)).Elem()
}

func (o GetCloudProjectContainerregistryResultOutput) ToGetCloudProjectContainerregistryResultOutput() GetCloudProjectContainerregistryResultOutput {
	return o
}

func (o GetCloudProjectContainerregistryResultOutput) ToGetCloudProjectContainerregistryResultOutputWithContext(ctx context.Context) GetCloudProjectContainerregistryResultOutput {
	return o
}

// Registry creation date
func (o GetCloudProjectContainerregistryResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCloudProjectContainerregistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Registry name
func (o GetCloudProjectContainerregistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project ID of your registry
func (o GetCloudProjectContainerregistryResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Region of the registry
func (o GetCloudProjectContainerregistryResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetCloudProjectContainerregistryResultOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.RegistryId }).(pulumi.StringOutput)
}

func (o GetCloudProjectContainerregistryResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Current size of the registry (bytes)
func (o GetCloudProjectContainerregistryResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) int { return v.Size }).(pulumi.IntOutput)
}

// Registry status
func (o GetCloudProjectContainerregistryResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.Status }).(pulumi.StringOutput)
}

// Registry last update date
func (o GetCloudProjectContainerregistryResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Access url of the registry
func (o GetCloudProjectContainerregistryResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.Url }).(pulumi.StringOutput)
}

// Version of your registry
func (o GetCloudProjectContainerregistryResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectContainerregistryResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectContainerregistryResultOutput{})
}
