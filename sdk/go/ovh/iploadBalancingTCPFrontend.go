// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a backend server group (frontend) to be used by loadbalancing frontend(s)
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
// 		lb, err := ovh.LookupIPLoadBalancing(ctx, &GetIPLoadBalancingArgs{
// 			ServiceName: pulumi.StringRef("ip-1.2.3.4"),
// 			State:       pulumi.StringRef("ok"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		farm80, err := ovh.NewIPLoadBalancingTCPFarm(ctx, "farm80", &ovh.IPLoadBalancingTCPFarmArgs{
// 			DisplayName: pulumi.String("ingress-8080-gra"),
// 			Port:        pulumi.Int(80),
// 			ServiceName: pulumi.String(lb.ServiceName),
// 			Zone:        pulumi.String("all"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ovh.NewIPLoadBalancingTCPFrontend(ctx, "testfrontend", &ovh.IPLoadBalancingTCPFrontendArgs{
// 			DefaultFarmId: farm80.ID(),
// 			DisplayName:   pulumi.String("ingress-8080-gra"),
// 			Port:          pulumi.String("80,443"),
// 			ServiceName:   pulumi.String(lb.ServiceName),
// 			Zone:          pulumi.String("all"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IPLoadBalancingTCPFrontend struct {
	pulumi.CustomResourceState

	// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
	AllowedSources pulumi.StringArrayOutput `pulumi:"allowedSources"`
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos pulumi.StringArrayOutput `pulumi:"dedicatedIpfos"`
	// Default TCP Farm of your frontend
	DefaultFarmId pulumi.IntOutput `pulumi:"defaultFarmId"`
	// Default ssl served to your customer
	DefaultSslId pulumi.IntOutput `pulumi:"defaultSslId"`
	// Disable your frontend. Default: 'false'
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Human readable name for your frontend, this field is for you
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Port(s) attached to your frontend. Supports single port (numerical value),
	// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
	// and/or 'range'. Each port must be in the [1;49151] range
	Port pulumi.StringOutput `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// SSL deciphering. Default: 'false'
	Ssl pulumi.BoolPtrOutput `pulumi:"ssl"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewIPLoadBalancingTCPFrontend registers a new resource with the given unique name, arguments, and options.
func NewIPLoadBalancingTCPFrontend(ctx *pulumi.Context,
	name string, args *IPLoadBalancingTCPFrontendArgs, opts ...pulumi.ResourceOption) (*IPLoadBalancingTCPFrontend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource IPLoadBalancingTCPFrontend
	err := ctx.RegisterResource("ovh:index/iPLoadBalancingTCPFrontend:IPLoadBalancingTCPFrontend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPLoadBalancingTCPFrontend gets an existing IPLoadBalancingTCPFrontend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPLoadBalancingTCPFrontend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPLoadBalancingTCPFrontendState, opts ...pulumi.ResourceOption) (*IPLoadBalancingTCPFrontend, error) {
	var resource IPLoadBalancingTCPFrontend
	err := ctx.ReadResource("ovh:index/iPLoadBalancingTCPFrontend:IPLoadBalancingTCPFrontend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPLoadBalancingTCPFrontend resources.
type iploadBalancingTCPFrontendState struct {
	// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
	AllowedSources []string `pulumi:"allowedSources"`
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos []string `pulumi:"dedicatedIpfos"`
	// Default TCP Farm of your frontend
	DefaultFarmId *int `pulumi:"defaultFarmId"`
	// Default ssl served to your customer
	DefaultSslId *int `pulumi:"defaultSslId"`
	// Disable your frontend. Default: 'false'
	Disabled *bool `pulumi:"disabled"`
	// Human readable name for your frontend, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Port(s) attached to your frontend. Supports single port (numerical value),
	// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
	// and/or 'range'. Each port must be in the [1;49151] range
	Port *string `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// SSL deciphering. Default: 'false'
	Ssl *bool `pulumi:"ssl"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone *string `pulumi:"zone"`
}

type IPLoadBalancingTCPFrontendState struct {
	// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
	AllowedSources pulumi.StringArrayInput
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos pulumi.StringArrayInput
	// Default TCP Farm of your frontend
	DefaultFarmId pulumi.IntPtrInput
	// Default ssl served to your customer
	DefaultSslId pulumi.IntPtrInput
	// Disable your frontend. Default: 'false'
	Disabled pulumi.BoolPtrInput
	// Human readable name for your frontend, this field is for you
	DisplayName pulumi.StringPtrInput
	// Port(s) attached to your frontend. Supports single port (numerical value),
	// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
	// and/or 'range'. Each port must be in the [1;49151] range
	Port pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// SSL deciphering. Default: 'false'
	Ssl pulumi.BoolPtrInput
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringPtrInput
}

func (IPLoadBalancingTCPFrontendState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadBalancingTCPFrontendState)(nil)).Elem()
}

type iploadBalancingTCPFrontendArgs struct {
	// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
	AllowedSources []string `pulumi:"allowedSources"`
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos []string `pulumi:"dedicatedIpfos"`
	// Default TCP Farm of your frontend
	DefaultFarmId *int `pulumi:"defaultFarmId"`
	// Default ssl served to your customer
	DefaultSslId *int `pulumi:"defaultSslId"`
	// Disable your frontend. Default: 'false'
	Disabled *bool `pulumi:"disabled"`
	// Human readable name for your frontend, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Port(s) attached to your frontend. Supports single port (numerical value),
	// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
	// and/or 'range'. Each port must be in the [1;49151] range
	Port string `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// SSL deciphering. Default: 'false'
	Ssl *bool `pulumi:"ssl"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a IPLoadBalancingTCPFrontend resource.
type IPLoadBalancingTCPFrontendArgs struct {
	// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
	AllowedSources pulumi.StringArrayInput
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos pulumi.StringArrayInput
	// Default TCP Farm of your frontend
	DefaultFarmId pulumi.IntPtrInput
	// Default ssl served to your customer
	DefaultSslId pulumi.IntPtrInput
	// Disable your frontend. Default: 'false'
	Disabled pulumi.BoolPtrInput
	// Human readable name for your frontend, this field is for you
	DisplayName pulumi.StringPtrInput
	// Port(s) attached to your frontend. Supports single port (numerical value),
	// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
	// and/or 'range'. Each port must be in the [1;49151] range
	Port pulumi.StringInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// SSL deciphering. Default: 'false'
	Ssl pulumi.BoolPtrInput
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringInput
}

func (IPLoadBalancingTCPFrontendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadBalancingTCPFrontendArgs)(nil)).Elem()
}

type IPLoadBalancingTCPFrontendInput interface {
	pulumi.Input

	ToIPLoadBalancingTCPFrontendOutput() IPLoadBalancingTCPFrontendOutput
	ToIPLoadBalancingTCPFrontendOutputWithContext(ctx context.Context) IPLoadBalancingTCPFrontendOutput
}

func (*IPLoadBalancingTCPFrontend) ElementType() reflect.Type {
	return reflect.TypeOf((**IPLoadBalancingTCPFrontend)(nil)).Elem()
}

func (i *IPLoadBalancingTCPFrontend) ToIPLoadBalancingTCPFrontendOutput() IPLoadBalancingTCPFrontendOutput {
	return i.ToIPLoadBalancingTCPFrontendOutputWithContext(context.Background())
}

func (i *IPLoadBalancingTCPFrontend) ToIPLoadBalancingTCPFrontendOutputWithContext(ctx context.Context) IPLoadBalancingTCPFrontendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingTCPFrontendOutput)
}

// IPLoadBalancingTCPFrontendArrayInput is an input type that accepts IPLoadBalancingTCPFrontendArray and IPLoadBalancingTCPFrontendArrayOutput values.
// You can construct a concrete instance of `IPLoadBalancingTCPFrontendArrayInput` via:
//
//          IPLoadBalancingTCPFrontendArray{ IPLoadBalancingTCPFrontendArgs{...} }
type IPLoadBalancingTCPFrontendArrayInput interface {
	pulumi.Input

	ToIPLoadBalancingTCPFrontendArrayOutput() IPLoadBalancingTCPFrontendArrayOutput
	ToIPLoadBalancingTCPFrontendArrayOutputWithContext(context.Context) IPLoadBalancingTCPFrontendArrayOutput
}

type IPLoadBalancingTCPFrontendArray []IPLoadBalancingTCPFrontendInput

func (IPLoadBalancingTCPFrontendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IPLoadBalancingTCPFrontend)(nil)).Elem()
}

func (i IPLoadBalancingTCPFrontendArray) ToIPLoadBalancingTCPFrontendArrayOutput() IPLoadBalancingTCPFrontendArrayOutput {
	return i.ToIPLoadBalancingTCPFrontendArrayOutputWithContext(context.Background())
}

func (i IPLoadBalancingTCPFrontendArray) ToIPLoadBalancingTCPFrontendArrayOutputWithContext(ctx context.Context) IPLoadBalancingTCPFrontendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingTCPFrontendArrayOutput)
}

// IPLoadBalancingTCPFrontendMapInput is an input type that accepts IPLoadBalancingTCPFrontendMap and IPLoadBalancingTCPFrontendMapOutput values.
// You can construct a concrete instance of `IPLoadBalancingTCPFrontendMapInput` via:
//
//          IPLoadBalancingTCPFrontendMap{ "key": IPLoadBalancingTCPFrontendArgs{...} }
type IPLoadBalancingTCPFrontendMapInput interface {
	pulumi.Input

	ToIPLoadBalancingTCPFrontendMapOutput() IPLoadBalancingTCPFrontendMapOutput
	ToIPLoadBalancingTCPFrontendMapOutputWithContext(context.Context) IPLoadBalancingTCPFrontendMapOutput
}

type IPLoadBalancingTCPFrontendMap map[string]IPLoadBalancingTCPFrontendInput

func (IPLoadBalancingTCPFrontendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IPLoadBalancingTCPFrontend)(nil)).Elem()
}

func (i IPLoadBalancingTCPFrontendMap) ToIPLoadBalancingTCPFrontendMapOutput() IPLoadBalancingTCPFrontendMapOutput {
	return i.ToIPLoadBalancingTCPFrontendMapOutputWithContext(context.Background())
}

func (i IPLoadBalancingTCPFrontendMap) ToIPLoadBalancingTCPFrontendMapOutputWithContext(ctx context.Context) IPLoadBalancingTCPFrontendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingTCPFrontendMapOutput)
}

type IPLoadBalancingTCPFrontendOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingTCPFrontendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPLoadBalancingTCPFrontend)(nil)).Elem()
}

func (o IPLoadBalancingTCPFrontendOutput) ToIPLoadBalancingTCPFrontendOutput() IPLoadBalancingTCPFrontendOutput {
	return o
}

func (o IPLoadBalancingTCPFrontendOutput) ToIPLoadBalancingTCPFrontendOutputWithContext(ctx context.Context) IPLoadBalancingTCPFrontendOutput {
	return o
}

// Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
func (o IPLoadBalancingTCPFrontendOutput) AllowedSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IPLoadBalancingTCPFrontend) pulumi.StringArrayOutput { return v.AllowedSources }).(pulumi.StringArrayOutput)
}

// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
func (o IPLoadBalancingTCPFrontendOutput) DedicatedIpfos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IPLoadBalancingTCPFrontend) pulumi.StringArrayOutput { return v.DedicatedIpfos }).(pulumi.StringArrayOutput)
}

// Default TCP Farm of your frontend
func (o IPLoadBalancingTCPFrontendOutput) DefaultFarmId() pulumi.IntOutput {
	return o.ApplyT(func(v *IPLoadBalancingTCPFrontend) pulumi.IntOutput { return v.DefaultFarmId }).(pulumi.IntOutput)
}

// Default ssl served to your customer
func (o IPLoadBalancingTCPFrontendOutput) DefaultSslId() pulumi.IntOutput {
	return o.ApplyT(func(v *IPLoadBalancingTCPFrontend) pulumi.IntOutput { return v.DefaultSslId }).(pulumi.IntOutput)
}

// Disable your frontend. Default: 'false'
func (o IPLoadBalancingTCPFrontendOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IPLoadBalancingTCPFrontend) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Human readable name for your frontend, this field is for you
func (o IPLoadBalancingTCPFrontendOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPLoadBalancingTCPFrontend) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Port(s) attached to your frontend. Supports single port (numerical value),
// range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
// and/or 'range'. Each port must be in the [1;49151] range
func (o IPLoadBalancingTCPFrontendOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *IPLoadBalancingTCPFrontend) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// The internal name of your IP load balancing
func (o IPLoadBalancingTCPFrontendOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IPLoadBalancingTCPFrontend) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// SSL deciphering. Default: 'false'
func (o IPLoadBalancingTCPFrontendOutput) Ssl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IPLoadBalancingTCPFrontend) pulumi.BoolPtrOutput { return v.Ssl }).(pulumi.BoolPtrOutput)
}

// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
func (o IPLoadBalancingTCPFrontendOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *IPLoadBalancingTCPFrontend) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type IPLoadBalancingTCPFrontendArrayOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingTCPFrontendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IPLoadBalancingTCPFrontend)(nil)).Elem()
}

func (o IPLoadBalancingTCPFrontendArrayOutput) ToIPLoadBalancingTCPFrontendArrayOutput() IPLoadBalancingTCPFrontendArrayOutput {
	return o
}

func (o IPLoadBalancingTCPFrontendArrayOutput) ToIPLoadBalancingTCPFrontendArrayOutputWithContext(ctx context.Context) IPLoadBalancingTCPFrontendArrayOutput {
	return o
}

func (o IPLoadBalancingTCPFrontendArrayOutput) Index(i pulumi.IntInput) IPLoadBalancingTCPFrontendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IPLoadBalancingTCPFrontend {
		return vs[0].([]*IPLoadBalancingTCPFrontend)[vs[1].(int)]
	}).(IPLoadBalancingTCPFrontendOutput)
}

type IPLoadBalancingTCPFrontendMapOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingTCPFrontendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IPLoadBalancingTCPFrontend)(nil)).Elem()
}

func (o IPLoadBalancingTCPFrontendMapOutput) ToIPLoadBalancingTCPFrontendMapOutput() IPLoadBalancingTCPFrontendMapOutput {
	return o
}

func (o IPLoadBalancingTCPFrontendMapOutput) ToIPLoadBalancingTCPFrontendMapOutputWithContext(ctx context.Context) IPLoadBalancingTCPFrontendMapOutput {
	return o
}

func (o IPLoadBalancingTCPFrontendMapOutput) MapIndex(k pulumi.StringInput) IPLoadBalancingTCPFrontendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IPLoadBalancingTCPFrontend {
		return vs[0].(map[string]*IPLoadBalancingTCPFrontend)[vs[1].(string)]
	}).(IPLoadBalancingTCPFrontendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingTCPFrontendInput)(nil)).Elem(), &IPLoadBalancingTCPFrontend{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingTCPFrontendArrayInput)(nil)).Elem(), IPLoadBalancingTCPFrontendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingTCPFrontendMapInput)(nil)).Elem(), IPLoadBalancingTCPFrontendMap{})
	pulumi.RegisterOutputType(IPLoadBalancingTCPFrontendOutput{})
	pulumi.RegisterOutputType(IPLoadBalancingTCPFrontendArrayOutput{})
	pulumi.RegisterOutputType(IPLoadBalancingTCPFrontendMapOutput{})
}
