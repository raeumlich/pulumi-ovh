// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the container registry capabilities of a public cloud project.
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
// 		_, err := ovh.LookupCloudProjectCapabilitiesContainerRegistry(ctx, &GetCloudProjectCapabilitiesContainerRegistryArgs{
// 			ServiceName: "XXXXXX",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCloudProjectCapabilitiesContainerRegistry(ctx *pulumi.Context, args *LookupCloudProjectCapabilitiesContainerRegistryArgs, opts ...pulumi.InvokeOption) (*LookupCloudProjectCapabilitiesContainerRegistryResult, error) {
	var rv LookupCloudProjectCapabilitiesContainerRegistryResult
	err := ctx.Invoke("ovh:index/getCloudProjectCapabilitiesContainerRegistry:getCloudProjectCapabilitiesContainerRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectCapabilitiesContainerRegistry.
type LookupCloudProjectCapabilitiesContainerRegistryArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectCapabilitiesContainerRegistry.
type LookupCloudProjectCapabilitiesContainerRegistryResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of container registry capability for a single region
	Results     []GetCloudProjectCapabilitiesContainerRegistryResult `pulumi:"results"`
	ServiceName string                                               `pulumi:"serviceName"`
}

func LookupCloudProjectCapabilitiesContainerRegistryOutput(ctx *pulumi.Context, args LookupCloudProjectCapabilitiesContainerRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProjectCapabilitiesContainerRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProjectCapabilitiesContainerRegistryResult, error) {
			args := v.(LookupCloudProjectCapabilitiesContainerRegistryArgs)
			r, err := LookupCloudProjectCapabilitiesContainerRegistry(ctx, &args, opts...)
			var s LookupCloudProjectCapabilitiesContainerRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudProjectCapabilitiesContainerRegistryResultOutput)
}

// A collection of arguments for invoking getCloudProjectCapabilitiesContainerRegistry.
type LookupCloudProjectCapabilitiesContainerRegistryOutputArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCloudProjectCapabilitiesContainerRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectCapabilitiesContainerRegistryArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectCapabilitiesContainerRegistry.
type LookupCloudProjectCapabilitiesContainerRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProjectCapabilitiesContainerRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectCapabilitiesContainerRegistryResult)(nil)).Elem()
}

func (o LookupCloudProjectCapabilitiesContainerRegistryResultOutput) ToLookupCloudProjectCapabilitiesContainerRegistryResultOutput() LookupCloudProjectCapabilitiesContainerRegistryResultOutput {
	return o
}

func (o LookupCloudProjectCapabilitiesContainerRegistryResultOutput) ToLookupCloudProjectCapabilitiesContainerRegistryResultOutputWithContext(ctx context.Context) LookupCloudProjectCapabilitiesContainerRegistryResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudProjectCapabilitiesContainerRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectCapabilitiesContainerRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of container registry capability for a single region
func (o LookupCloudProjectCapabilitiesContainerRegistryResultOutput) Results() GetCloudProjectCapabilitiesContainerRegistryResultArrayOutput {
	return o.ApplyT(func(v LookupCloudProjectCapabilitiesContainerRegistryResult) []GetCloudProjectCapabilitiesContainerRegistryResult {
		return v.Results
	}).(GetCloudProjectCapabilitiesContainerRegistryResultArrayOutput)
}

func (o LookupCloudProjectCapabilitiesContainerRegistryResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectCapabilitiesContainerRegistryResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProjectCapabilitiesContainerRegistryResultOutput{})
}
