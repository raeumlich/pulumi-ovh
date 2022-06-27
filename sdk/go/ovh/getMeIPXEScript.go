// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an IPXE Script.
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
// 		_, err := ovh.GetMeIPXEScript(ctx, &GetMeIPXEScriptArgs{
// 			Name: "myscript",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetMeIPXEScript(ctx *pulumi.Context, args *GetMeIPXEScriptArgs, opts ...pulumi.InvokeOption) (*GetMeIPXEScriptResult, error) {
	var rv GetMeIPXEScriptResult
	err := ctx.Invoke("ovh:index/getMeIPXEScript:getMeIPXEScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMeIPXEScript.
type GetMeIPXEScriptArgs struct {
	// The name of the IPXE Script.
	Name string `pulumi:"name"`
}

// A collection of values returned by getMeIPXEScript.
type GetMeIPXEScriptResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// The content of the script.
	Script string `pulumi:"script"`
}

func GetMeIPXEScriptOutput(ctx *pulumi.Context, args GetMeIPXEScriptOutputArgs, opts ...pulumi.InvokeOption) GetMeIPXEScriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMeIPXEScriptResult, error) {
			args := v.(GetMeIPXEScriptArgs)
			r, err := GetMeIPXEScript(ctx, &args, opts...)
			var s GetMeIPXEScriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMeIPXEScriptResultOutput)
}

// A collection of arguments for invoking getMeIPXEScript.
type GetMeIPXEScriptOutputArgs struct {
	// The name of the IPXE Script.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetMeIPXEScriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMeIPXEScriptArgs)(nil)).Elem()
}

// A collection of values returned by getMeIPXEScript.
type GetMeIPXEScriptResultOutput struct{ *pulumi.OutputState }

func (GetMeIPXEScriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMeIPXEScriptResult)(nil)).Elem()
}

func (o GetMeIPXEScriptResultOutput) ToGetMeIPXEScriptResultOutput() GetMeIPXEScriptResultOutput {
	return o
}

func (o GetMeIPXEScriptResultOutput) ToGetMeIPXEScriptResultOutputWithContext(ctx context.Context) GetMeIPXEScriptResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetMeIPXEScriptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMeIPXEScriptResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetMeIPXEScriptResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetMeIPXEScriptResult) string { return v.Name }).(pulumi.StringOutput)
}

// The content of the script.
func (o GetMeIPXEScriptResultOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v GetMeIPXEScriptResult) string { return v.Script }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMeIPXEScriptResultOutput{})
}
