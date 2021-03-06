// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the details of a failover ip address of a service in a public cloud project.
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
// 		_, err := ovh.LookupCloudProjectFailoverIPAttach(ctx, &GetCloudProjectFailoverIPAttachArgs{
// 			Ip:          pulumi.StringRef("XXXXXX"),
// 			ServiceName: "XXXXXX",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCloudProjectFailoverIPAttach(ctx *pulumi.Context, args *LookupCloudProjectFailoverIPAttachArgs, opts ...pulumi.InvokeOption) (*LookupCloudProjectFailoverIPAttachResult, error) {
	var rv LookupCloudProjectFailoverIPAttachResult
	err := ctx.Invoke("ovh:index/getCloudProjectFailoverIPAttach:getCloudProjectFailoverIPAttach", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectFailoverIPAttach.
type LookupCloudProjectFailoverIPAttachArgs struct {
	// The IP block
	Block         *string `pulumi:"block"`
	ContinentCode *string `pulumi:"continentCode"`
	GeoLoc        *string `pulumi:"geoLoc"`
	// The failover ip address to query
	Ip       *string `pulumi:"ip"`
	RoutedTo *string `pulumi:"routedTo"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectFailoverIPAttach.
type LookupCloudProjectFailoverIPAttachResult struct {
	// The IP block
	Block         string `pulumi:"block"`
	ContinentCode string `pulumi:"continentCode"`
	GeoLoc        string `pulumi:"geoLoc"`
	// The Ip id
	Id string `pulumi:"id"`
	// The Ip Address
	Ip string `pulumi:"ip"`
	// Current operation progress in percent
	Progress    int    `pulumi:"progress"`
	RoutedTo    string `pulumi:"routedTo"`
	ServiceName string `pulumi:"serviceName"`
	// Ip status, can be `ok` or `operationPending`
	Status  string `pulumi:"status"`
	SubType string `pulumi:"subType"`
}

func LookupCloudProjectFailoverIPAttachOutput(ctx *pulumi.Context, args LookupCloudProjectFailoverIPAttachOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProjectFailoverIPAttachResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProjectFailoverIPAttachResult, error) {
			args := v.(LookupCloudProjectFailoverIPAttachArgs)
			r, err := LookupCloudProjectFailoverIPAttach(ctx, &args, opts...)
			var s LookupCloudProjectFailoverIPAttachResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudProjectFailoverIPAttachResultOutput)
}

// A collection of arguments for invoking getCloudProjectFailoverIPAttach.
type LookupCloudProjectFailoverIPAttachOutputArgs struct {
	// The IP block
	Block         pulumi.StringPtrInput `pulumi:"block"`
	ContinentCode pulumi.StringPtrInput `pulumi:"continentCode"`
	GeoLoc        pulumi.StringPtrInput `pulumi:"geoLoc"`
	// The failover ip address to query
	Ip       pulumi.StringPtrInput `pulumi:"ip"`
	RoutedTo pulumi.StringPtrInput `pulumi:"routedTo"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCloudProjectFailoverIPAttachOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectFailoverIPAttachArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectFailoverIPAttach.
type LookupCloudProjectFailoverIPAttachResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProjectFailoverIPAttachResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectFailoverIPAttachResult)(nil)).Elem()
}

func (o LookupCloudProjectFailoverIPAttachResultOutput) ToLookupCloudProjectFailoverIPAttachResultOutput() LookupCloudProjectFailoverIPAttachResultOutput {
	return o
}

func (o LookupCloudProjectFailoverIPAttachResultOutput) ToLookupCloudProjectFailoverIPAttachResultOutputWithContext(ctx context.Context) LookupCloudProjectFailoverIPAttachResultOutput {
	return o
}

// The IP block
func (o LookupCloudProjectFailoverIPAttachResultOutput) Block() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectFailoverIPAttachResult) string { return v.Block }).(pulumi.StringOutput)
}

func (o LookupCloudProjectFailoverIPAttachResultOutput) ContinentCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectFailoverIPAttachResult) string { return v.ContinentCode }).(pulumi.StringOutput)
}

func (o LookupCloudProjectFailoverIPAttachResultOutput) GeoLoc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectFailoverIPAttachResult) string { return v.GeoLoc }).(pulumi.StringOutput)
}

// The Ip id
func (o LookupCloudProjectFailoverIPAttachResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectFailoverIPAttachResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Ip Address
func (o LookupCloudProjectFailoverIPAttachResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectFailoverIPAttachResult) string { return v.Ip }).(pulumi.StringOutput)
}

// Current operation progress in percent
func (o LookupCloudProjectFailoverIPAttachResultOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudProjectFailoverIPAttachResult) int { return v.Progress }).(pulumi.IntOutput)
}

func (o LookupCloudProjectFailoverIPAttachResultOutput) RoutedTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectFailoverIPAttachResult) string { return v.RoutedTo }).(pulumi.StringOutput)
}

func (o LookupCloudProjectFailoverIPAttachResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectFailoverIPAttachResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Ip status, can be `ok` or `operationPending`
func (o LookupCloudProjectFailoverIPAttachResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectFailoverIPAttachResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupCloudProjectFailoverIPAttachResultOutput) SubType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectFailoverIPAttachResult) string { return v.SubType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProjectFailoverIPAttachResultOutput{})
}
