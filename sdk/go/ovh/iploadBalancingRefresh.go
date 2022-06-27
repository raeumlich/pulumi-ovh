// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Applies changes from other ovh_iploadbalancing_* resourcesto the production configuration of loadbalancers.
type IPLoadBalancingRefresh struct {
	pulumi.CustomResourceState

	// List of values traccked to trigger refresh, used also to form implicit dependencies
	Keepers pulumi.StringArrayOutput `pulumi:"keepers"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewIPLoadBalancingRefresh registers a new resource with the given unique name, arguments, and options.
func NewIPLoadBalancingRefresh(ctx *pulumi.Context,
	name string, args *IPLoadBalancingRefreshArgs, opts ...pulumi.ResourceOption) (*IPLoadBalancingRefresh, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Keepers == nil {
		return nil, errors.New("invalid value for required argument 'Keepers'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource IPLoadBalancingRefresh
	err := ctx.RegisterResource("ovh:index/iPLoadBalancingRefresh:IPLoadBalancingRefresh", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPLoadBalancingRefresh gets an existing IPLoadBalancingRefresh resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPLoadBalancingRefresh(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPLoadBalancingRefreshState, opts ...pulumi.ResourceOption) (*IPLoadBalancingRefresh, error) {
	var resource IPLoadBalancingRefresh
	err := ctx.ReadResource("ovh:index/iPLoadBalancingRefresh:IPLoadBalancingRefresh", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPLoadBalancingRefresh resources.
type iploadBalancingRefreshState struct {
	// List of values traccked to trigger refresh, used also to form implicit dependencies
	Keepers []string `pulumi:"keepers"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
}

type IPLoadBalancingRefreshState struct {
	// List of values traccked to trigger refresh, used also to form implicit dependencies
	Keepers pulumi.StringArrayInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
}

func (IPLoadBalancingRefreshState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadBalancingRefreshState)(nil)).Elem()
}

type iploadBalancingRefreshArgs struct {
	// List of values traccked to trigger refresh, used also to form implicit dependencies
	Keepers []string `pulumi:"keepers"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a IPLoadBalancingRefresh resource.
type IPLoadBalancingRefreshArgs struct {
	// List of values traccked to trigger refresh, used also to form implicit dependencies
	Keepers pulumi.StringArrayInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
}

func (IPLoadBalancingRefreshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadBalancingRefreshArgs)(nil)).Elem()
}

type IPLoadBalancingRefreshInput interface {
	pulumi.Input

	ToIPLoadBalancingRefreshOutput() IPLoadBalancingRefreshOutput
	ToIPLoadBalancingRefreshOutputWithContext(ctx context.Context) IPLoadBalancingRefreshOutput
}

func (*IPLoadBalancingRefresh) ElementType() reflect.Type {
	return reflect.TypeOf((**IPLoadBalancingRefresh)(nil)).Elem()
}

func (i *IPLoadBalancingRefresh) ToIPLoadBalancingRefreshOutput() IPLoadBalancingRefreshOutput {
	return i.ToIPLoadBalancingRefreshOutputWithContext(context.Background())
}

func (i *IPLoadBalancingRefresh) ToIPLoadBalancingRefreshOutputWithContext(ctx context.Context) IPLoadBalancingRefreshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingRefreshOutput)
}

// IPLoadBalancingRefreshArrayInput is an input type that accepts IPLoadBalancingRefreshArray and IPLoadBalancingRefreshArrayOutput values.
// You can construct a concrete instance of `IPLoadBalancingRefreshArrayInput` via:
//
//          IPLoadBalancingRefreshArray{ IPLoadBalancingRefreshArgs{...} }
type IPLoadBalancingRefreshArrayInput interface {
	pulumi.Input

	ToIPLoadBalancingRefreshArrayOutput() IPLoadBalancingRefreshArrayOutput
	ToIPLoadBalancingRefreshArrayOutputWithContext(context.Context) IPLoadBalancingRefreshArrayOutput
}

type IPLoadBalancingRefreshArray []IPLoadBalancingRefreshInput

func (IPLoadBalancingRefreshArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IPLoadBalancingRefresh)(nil)).Elem()
}

func (i IPLoadBalancingRefreshArray) ToIPLoadBalancingRefreshArrayOutput() IPLoadBalancingRefreshArrayOutput {
	return i.ToIPLoadBalancingRefreshArrayOutputWithContext(context.Background())
}

func (i IPLoadBalancingRefreshArray) ToIPLoadBalancingRefreshArrayOutputWithContext(ctx context.Context) IPLoadBalancingRefreshArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingRefreshArrayOutput)
}

// IPLoadBalancingRefreshMapInput is an input type that accepts IPLoadBalancingRefreshMap and IPLoadBalancingRefreshMapOutput values.
// You can construct a concrete instance of `IPLoadBalancingRefreshMapInput` via:
//
//          IPLoadBalancingRefreshMap{ "key": IPLoadBalancingRefreshArgs{...} }
type IPLoadBalancingRefreshMapInput interface {
	pulumi.Input

	ToIPLoadBalancingRefreshMapOutput() IPLoadBalancingRefreshMapOutput
	ToIPLoadBalancingRefreshMapOutputWithContext(context.Context) IPLoadBalancingRefreshMapOutput
}

type IPLoadBalancingRefreshMap map[string]IPLoadBalancingRefreshInput

func (IPLoadBalancingRefreshMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IPLoadBalancingRefresh)(nil)).Elem()
}

func (i IPLoadBalancingRefreshMap) ToIPLoadBalancingRefreshMapOutput() IPLoadBalancingRefreshMapOutput {
	return i.ToIPLoadBalancingRefreshMapOutputWithContext(context.Background())
}

func (i IPLoadBalancingRefreshMap) ToIPLoadBalancingRefreshMapOutputWithContext(ctx context.Context) IPLoadBalancingRefreshMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingRefreshMapOutput)
}

type IPLoadBalancingRefreshOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingRefreshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPLoadBalancingRefresh)(nil)).Elem()
}

func (o IPLoadBalancingRefreshOutput) ToIPLoadBalancingRefreshOutput() IPLoadBalancingRefreshOutput {
	return o
}

func (o IPLoadBalancingRefreshOutput) ToIPLoadBalancingRefreshOutputWithContext(ctx context.Context) IPLoadBalancingRefreshOutput {
	return o
}

// List of values traccked to trigger refresh, used also to form implicit dependencies
func (o IPLoadBalancingRefreshOutput) Keepers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IPLoadBalancingRefresh) pulumi.StringArrayOutput { return v.Keepers }).(pulumi.StringArrayOutput)
}

// The internal name of your IP load balancing
func (o IPLoadBalancingRefreshOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IPLoadBalancingRefresh) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type IPLoadBalancingRefreshArrayOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingRefreshArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IPLoadBalancingRefresh)(nil)).Elem()
}

func (o IPLoadBalancingRefreshArrayOutput) ToIPLoadBalancingRefreshArrayOutput() IPLoadBalancingRefreshArrayOutput {
	return o
}

func (o IPLoadBalancingRefreshArrayOutput) ToIPLoadBalancingRefreshArrayOutputWithContext(ctx context.Context) IPLoadBalancingRefreshArrayOutput {
	return o
}

func (o IPLoadBalancingRefreshArrayOutput) Index(i pulumi.IntInput) IPLoadBalancingRefreshOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IPLoadBalancingRefresh {
		return vs[0].([]*IPLoadBalancingRefresh)[vs[1].(int)]
	}).(IPLoadBalancingRefreshOutput)
}

type IPLoadBalancingRefreshMapOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingRefreshMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IPLoadBalancingRefresh)(nil)).Elem()
}

func (o IPLoadBalancingRefreshMapOutput) ToIPLoadBalancingRefreshMapOutput() IPLoadBalancingRefreshMapOutput {
	return o
}

func (o IPLoadBalancingRefreshMapOutput) ToIPLoadBalancingRefreshMapOutputWithContext(ctx context.Context) IPLoadBalancingRefreshMapOutput {
	return o
}

func (o IPLoadBalancingRefreshMapOutput) MapIndex(k pulumi.StringInput) IPLoadBalancingRefreshOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IPLoadBalancingRefresh {
		return vs[0].(map[string]*IPLoadBalancingRefresh)[vs[1].(string)]
	}).(IPLoadBalancingRefreshOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingRefreshInput)(nil)).Elem(), &IPLoadBalancingRefresh{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingRefreshArrayInput)(nil)).Elem(), IPLoadBalancingRefreshArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingRefreshMapInput)(nil)).Elem(), IPLoadBalancingRefreshMap{})
	pulumi.RegisterOutputType(IPLoadBalancingRefreshOutput{})
	pulumi.RegisterOutputType(IPLoadBalancingRefreshArrayOutput{})
	pulumi.RegisterOutputType(IPLoadBalancingRefreshMapOutput{})
}