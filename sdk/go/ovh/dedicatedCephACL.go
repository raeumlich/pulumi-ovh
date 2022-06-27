// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Add a new access ACL for the given network/mask.
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
// 		my_ceph, err := ovh.GetDedicatedCeph(ctx, &GetDedicatedCephArgs{
// 			ServiceName: "94d423da-0e55-45f2-9812-836460a19939",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ovh.NewDedicatedCephACL(ctx, "my-acl", &ovh.DedicatedCephACLArgs{
// 			ServiceName: pulumi.String(my_ceph.Id),
// 			Network:     pulumi.String("1.2.3.4"),
// 			Netmask:     pulumi.String("255.255.255.255"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DedicatedCephACL struct {
	pulumi.CustomResourceState

	// IP family. `IPv4` or `IPv6`
	Family pulumi.StringOutput `pulumi:"family"`
	// The network mask to apply
	Netmask pulumi.StringOutput `pulumi:"netmask"`
	// The network IP to authorize
	Network pulumi.StringOutput `pulumi:"network"`
	// The internal name of your dedicated CEPH
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewDedicatedCephACL registers a new resource with the given unique name, arguments, and options.
func NewDedicatedCephACL(ctx *pulumi.Context,
	name string, args *DedicatedCephACLArgs, opts ...pulumi.ResourceOption) (*DedicatedCephACL, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Netmask == nil {
		return nil, errors.New("invalid value for required argument 'Netmask'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource DedicatedCephACL
	err := ctx.RegisterResource("ovh:index/dedicatedCephACL:DedicatedCephACL", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedCephACL gets an existing DedicatedCephACL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedCephACL(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedCephACLState, opts ...pulumi.ResourceOption) (*DedicatedCephACL, error) {
	var resource DedicatedCephACL
	err := ctx.ReadResource("ovh:index/dedicatedCephACL:DedicatedCephACL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedCephACL resources.
type dedicatedCephACLState struct {
	// IP family. `IPv4` or `IPv6`
	Family *string `pulumi:"family"`
	// The network mask to apply
	Netmask *string `pulumi:"netmask"`
	// The network IP to authorize
	Network *string `pulumi:"network"`
	// The internal name of your dedicated CEPH
	ServiceName *string `pulumi:"serviceName"`
}

type DedicatedCephACLState struct {
	// IP family. `IPv4` or `IPv6`
	Family pulumi.StringPtrInput
	// The network mask to apply
	Netmask pulumi.StringPtrInput
	// The network IP to authorize
	Network pulumi.StringPtrInput
	// The internal name of your dedicated CEPH
	ServiceName pulumi.StringPtrInput
}

func (DedicatedCephACLState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCephACLState)(nil)).Elem()
}

type dedicatedCephACLArgs struct {
	// The network mask to apply
	Netmask string `pulumi:"netmask"`
	// The network IP to authorize
	Network string `pulumi:"network"`
	// The internal name of your dedicated CEPH
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a DedicatedCephACL resource.
type DedicatedCephACLArgs struct {
	// The network mask to apply
	Netmask pulumi.StringInput
	// The network IP to authorize
	Network pulumi.StringInput
	// The internal name of your dedicated CEPH
	ServiceName pulumi.StringInput
}

func (DedicatedCephACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCephACLArgs)(nil)).Elem()
}

type DedicatedCephACLInput interface {
	pulumi.Input

	ToDedicatedCephACLOutput() DedicatedCephACLOutput
	ToDedicatedCephACLOutputWithContext(ctx context.Context) DedicatedCephACLOutput
}

func (*DedicatedCephACL) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCephACL)(nil)).Elem()
}

func (i *DedicatedCephACL) ToDedicatedCephACLOutput() DedicatedCephACLOutput {
	return i.ToDedicatedCephACLOutputWithContext(context.Background())
}

func (i *DedicatedCephACL) ToDedicatedCephACLOutputWithContext(ctx context.Context) DedicatedCephACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCephACLOutput)
}

// DedicatedCephACLArrayInput is an input type that accepts DedicatedCephACLArray and DedicatedCephACLArrayOutput values.
// You can construct a concrete instance of `DedicatedCephACLArrayInput` via:
//
//          DedicatedCephACLArray{ DedicatedCephACLArgs{...} }
type DedicatedCephACLArrayInput interface {
	pulumi.Input

	ToDedicatedCephACLArrayOutput() DedicatedCephACLArrayOutput
	ToDedicatedCephACLArrayOutputWithContext(context.Context) DedicatedCephACLArrayOutput
}

type DedicatedCephACLArray []DedicatedCephACLInput

func (DedicatedCephACLArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedCephACL)(nil)).Elem()
}

func (i DedicatedCephACLArray) ToDedicatedCephACLArrayOutput() DedicatedCephACLArrayOutput {
	return i.ToDedicatedCephACLArrayOutputWithContext(context.Background())
}

func (i DedicatedCephACLArray) ToDedicatedCephACLArrayOutputWithContext(ctx context.Context) DedicatedCephACLArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCephACLArrayOutput)
}

// DedicatedCephACLMapInput is an input type that accepts DedicatedCephACLMap and DedicatedCephACLMapOutput values.
// You can construct a concrete instance of `DedicatedCephACLMapInput` via:
//
//          DedicatedCephACLMap{ "key": DedicatedCephACLArgs{...} }
type DedicatedCephACLMapInput interface {
	pulumi.Input

	ToDedicatedCephACLMapOutput() DedicatedCephACLMapOutput
	ToDedicatedCephACLMapOutputWithContext(context.Context) DedicatedCephACLMapOutput
}

type DedicatedCephACLMap map[string]DedicatedCephACLInput

func (DedicatedCephACLMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedCephACL)(nil)).Elem()
}

func (i DedicatedCephACLMap) ToDedicatedCephACLMapOutput() DedicatedCephACLMapOutput {
	return i.ToDedicatedCephACLMapOutputWithContext(context.Background())
}

func (i DedicatedCephACLMap) ToDedicatedCephACLMapOutputWithContext(ctx context.Context) DedicatedCephACLMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCephACLMapOutput)
}

type DedicatedCephACLOutput struct{ *pulumi.OutputState }

func (DedicatedCephACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCephACL)(nil)).Elem()
}

func (o DedicatedCephACLOutput) ToDedicatedCephACLOutput() DedicatedCephACLOutput {
	return o
}

func (o DedicatedCephACLOutput) ToDedicatedCephACLOutputWithContext(ctx context.Context) DedicatedCephACLOutput {
	return o
}

// IP family. `IPv4` or `IPv6`
func (o DedicatedCephACLOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCephACL) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

// The network mask to apply
func (o DedicatedCephACLOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCephACL) pulumi.StringOutput { return v.Netmask }).(pulumi.StringOutput)
}

// The network IP to authorize
func (o DedicatedCephACLOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCephACL) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// The internal name of your dedicated CEPH
func (o DedicatedCephACLOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCephACL) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type DedicatedCephACLArrayOutput struct{ *pulumi.OutputState }

func (DedicatedCephACLArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedCephACL)(nil)).Elem()
}

func (o DedicatedCephACLArrayOutput) ToDedicatedCephACLArrayOutput() DedicatedCephACLArrayOutput {
	return o
}

func (o DedicatedCephACLArrayOutput) ToDedicatedCephACLArrayOutputWithContext(ctx context.Context) DedicatedCephACLArrayOutput {
	return o
}

func (o DedicatedCephACLArrayOutput) Index(i pulumi.IntInput) DedicatedCephACLOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DedicatedCephACL {
		return vs[0].([]*DedicatedCephACL)[vs[1].(int)]
	}).(DedicatedCephACLOutput)
}

type DedicatedCephACLMapOutput struct{ *pulumi.OutputState }

func (DedicatedCephACLMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedCephACL)(nil)).Elem()
}

func (o DedicatedCephACLMapOutput) ToDedicatedCephACLMapOutput() DedicatedCephACLMapOutput {
	return o
}

func (o DedicatedCephACLMapOutput) ToDedicatedCephACLMapOutputWithContext(ctx context.Context) DedicatedCephACLMapOutput {
	return o
}

func (o DedicatedCephACLMapOutput) MapIndex(k pulumi.StringInput) DedicatedCephACLOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DedicatedCephACL {
		return vs[0].(map[string]*DedicatedCephACL)[vs[1].(string)]
	}).(DedicatedCephACLOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCephACLInput)(nil)).Elem(), &DedicatedCephACL{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCephACLArrayInput)(nil)).Elem(), DedicatedCephACLArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCephACLMapInput)(nil)).Elem(), DedicatedCephACLMap{})
	pulumi.RegisterOutputType(DedicatedCephACLOutput{})
	pulumi.RegisterOutputType(DedicatedCephACLArrayOutput{})
	pulumi.RegisterOutputType(DedicatedCephACLMapOutput{})
}
