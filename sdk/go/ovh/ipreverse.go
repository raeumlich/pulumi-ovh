// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a OVH IP reverse.
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
// 		_, err := ovh.NewIPReverse(ctx, "test", &ovh.IPReverseArgs{
// 			Ip:        pulumi.String("192.0.2.0/24"),
// 			IpReverse: pulumi.String("192.0.2.1"),
// 			Reverse:   pulumi.String("example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IPReverse struct {
	pulumi.CustomResourceState

	// The IP block to which the IP belongs
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The IP to set the reverse of
	IpReverse pulumi.StringOutput `pulumi:"ipReverse"`
	// The value of the reverse
	Reverse pulumi.StringOutput `pulumi:"reverse"`
}

// NewIPReverse registers a new resource with the given unique name, arguments, and options.
func NewIPReverse(ctx *pulumi.Context,
	name string, args *IPReverseArgs, opts ...pulumi.ResourceOption) (*IPReverse, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.IpReverse == nil {
		return nil, errors.New("invalid value for required argument 'IpReverse'")
	}
	if args.Reverse == nil {
		return nil, errors.New("invalid value for required argument 'Reverse'")
	}
	var resource IPReverse
	err := ctx.RegisterResource("ovh:index/iPReverse:IPReverse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPReverse gets an existing IPReverse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPReverse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPReverseState, opts ...pulumi.ResourceOption) (*IPReverse, error) {
	var resource IPReverse
	err := ctx.ReadResource("ovh:index/iPReverse:IPReverse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPReverse resources.
type ipreverseState struct {
	// The IP block to which the IP belongs
	Ip *string `pulumi:"ip"`
	// The IP to set the reverse of
	IpReverse *string `pulumi:"ipReverse"`
	// The value of the reverse
	Reverse *string `pulumi:"reverse"`
}

type IPReverseState struct {
	// The IP block to which the IP belongs
	Ip pulumi.StringPtrInput
	// The IP to set the reverse of
	IpReverse pulumi.StringPtrInput
	// The value of the reverse
	Reverse pulumi.StringPtrInput
}

func (IPReverseState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipreverseState)(nil)).Elem()
}

type ipreverseArgs struct {
	// The IP block to which the IP belongs
	Ip string `pulumi:"ip"`
	// The IP to set the reverse of
	IpReverse string `pulumi:"ipReverse"`
	// The value of the reverse
	Reverse string `pulumi:"reverse"`
}

// The set of arguments for constructing a IPReverse resource.
type IPReverseArgs struct {
	// The IP block to which the IP belongs
	Ip pulumi.StringInput
	// The IP to set the reverse of
	IpReverse pulumi.StringInput
	// The value of the reverse
	Reverse pulumi.StringInput
}

func (IPReverseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipreverseArgs)(nil)).Elem()
}

type IPReverseInput interface {
	pulumi.Input

	ToIPReverseOutput() IPReverseOutput
	ToIPReverseOutputWithContext(ctx context.Context) IPReverseOutput
}

func (*IPReverse) ElementType() reflect.Type {
	return reflect.TypeOf((**IPReverse)(nil)).Elem()
}

func (i *IPReverse) ToIPReverseOutput() IPReverseOutput {
	return i.ToIPReverseOutputWithContext(context.Background())
}

func (i *IPReverse) ToIPReverseOutputWithContext(ctx context.Context) IPReverseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPReverseOutput)
}

// IPReverseArrayInput is an input type that accepts IPReverseArray and IPReverseArrayOutput values.
// You can construct a concrete instance of `IPReverseArrayInput` via:
//
//          IPReverseArray{ IPReverseArgs{...} }
type IPReverseArrayInput interface {
	pulumi.Input

	ToIPReverseArrayOutput() IPReverseArrayOutput
	ToIPReverseArrayOutputWithContext(context.Context) IPReverseArrayOutput
}

type IPReverseArray []IPReverseInput

func (IPReverseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IPReverse)(nil)).Elem()
}

func (i IPReverseArray) ToIPReverseArrayOutput() IPReverseArrayOutput {
	return i.ToIPReverseArrayOutputWithContext(context.Background())
}

func (i IPReverseArray) ToIPReverseArrayOutputWithContext(ctx context.Context) IPReverseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPReverseArrayOutput)
}

// IPReverseMapInput is an input type that accepts IPReverseMap and IPReverseMapOutput values.
// You can construct a concrete instance of `IPReverseMapInput` via:
//
//          IPReverseMap{ "key": IPReverseArgs{...} }
type IPReverseMapInput interface {
	pulumi.Input

	ToIPReverseMapOutput() IPReverseMapOutput
	ToIPReverseMapOutputWithContext(context.Context) IPReverseMapOutput
}

type IPReverseMap map[string]IPReverseInput

func (IPReverseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IPReverse)(nil)).Elem()
}

func (i IPReverseMap) ToIPReverseMapOutput() IPReverseMapOutput {
	return i.ToIPReverseMapOutputWithContext(context.Background())
}

func (i IPReverseMap) ToIPReverseMapOutputWithContext(ctx context.Context) IPReverseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPReverseMapOutput)
}

type IPReverseOutput struct{ *pulumi.OutputState }

func (IPReverseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPReverse)(nil)).Elem()
}

func (o IPReverseOutput) ToIPReverseOutput() IPReverseOutput {
	return o
}

func (o IPReverseOutput) ToIPReverseOutputWithContext(ctx context.Context) IPReverseOutput {
	return o
}

// The IP block to which the IP belongs
func (o IPReverseOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *IPReverse) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The IP to set the reverse of
func (o IPReverseOutput) IpReverse() pulumi.StringOutput {
	return o.ApplyT(func(v *IPReverse) pulumi.StringOutput { return v.IpReverse }).(pulumi.StringOutput)
}

// The value of the reverse
func (o IPReverseOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *IPReverse) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

type IPReverseArrayOutput struct{ *pulumi.OutputState }

func (IPReverseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IPReverse)(nil)).Elem()
}

func (o IPReverseArrayOutput) ToIPReverseArrayOutput() IPReverseArrayOutput {
	return o
}

func (o IPReverseArrayOutput) ToIPReverseArrayOutputWithContext(ctx context.Context) IPReverseArrayOutput {
	return o
}

func (o IPReverseArrayOutput) Index(i pulumi.IntInput) IPReverseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IPReverse {
		return vs[0].([]*IPReverse)[vs[1].(int)]
	}).(IPReverseOutput)
}

type IPReverseMapOutput struct{ *pulumi.OutputState }

func (IPReverseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IPReverse)(nil)).Elem()
}

func (o IPReverseMapOutput) ToIPReverseMapOutput() IPReverseMapOutput {
	return o
}

func (o IPReverseMapOutput) ToIPReverseMapOutputWithContext(ctx context.Context) IPReverseMapOutput {
	return o
}

func (o IPReverseMapOutput) MapIndex(k pulumi.StringInput) IPReverseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IPReverse {
		return vs[0].(map[string]*IPReverse)[vs[1].(string)]
	}).(IPReverseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IPReverseInput)(nil)).Elem(), &IPReverse{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPReverseArrayInput)(nil)).Elem(), IPReverseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPReverseMapInput)(nil)).Elem(), IPReverseMap{})
	pulumi.RegisterOutputType(IPReverseOutput{})
	pulumi.RegisterOutputType(IPReverseArrayOutput{})
	pulumi.RegisterOutputType(IPReverseMapOutput{})
}
