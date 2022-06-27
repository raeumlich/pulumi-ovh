// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a http backend server group (farm) to be used by loadbalancing frontend(s)
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
// 		_, err = ovh.NewIPLoadBalancingHTTPFarm(ctx, "farmname", &ovh.IPLoadBalancingHTTPFarmArgs{
// 			DisplayName: pulumi.String("ingress-8080-gra"),
// 			ServiceName: pulumi.String(lb.Id),
// 			Zone:        pulumi.String("GRA"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IPLoadBalancingHTTPFarm struct {
	pulumi.CustomResourceState

	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
	Balance pulumi.StringPtrOutput `pulumi:"balance"`
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Port for backends to recieve traffic on.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// define a backend healthcheck probe
	Probe IPLoadBalancingHTTPFarmProbePtrOutput `pulumi:"probe"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
	Stickiness pulumi.StringPtrOutput `pulumi:"stickiness"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.IntPtrOutput `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewIPLoadBalancingHTTPFarm registers a new resource with the given unique name, arguments, and options.
func NewIPLoadBalancingHTTPFarm(ctx *pulumi.Context,
	name string, args *IPLoadBalancingHTTPFarmArgs, opts ...pulumi.ResourceOption) (*IPLoadBalancingHTTPFarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource IPLoadBalancingHTTPFarm
	err := ctx.RegisterResource("ovh:index/iPLoadBalancingHTTPFarm:IPLoadBalancingHTTPFarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPLoadBalancingHTTPFarm gets an existing IPLoadBalancingHTTPFarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPLoadBalancingHTTPFarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPLoadBalancingHTTPFarmState, opts ...pulumi.ResourceOption) (*IPLoadBalancingHTTPFarm, error) {
	var resource IPLoadBalancingHTTPFarm
	err := ctx.ReadResource("ovh:index/iPLoadBalancingHTTPFarm:IPLoadBalancingHTTPFarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPLoadBalancingHTTPFarm resources.
type iploadBalancingHTTPFarmState struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
	Balance *string `pulumi:"balance"`
	// Readable label for loadbalancer farm
	DisplayName *string `pulumi:"displayName"`
	// Port for backends to recieve traffic on.
	Port *int `pulumi:"port"`
	// define a backend healthcheck probe
	Probe *IPLoadBalancingHTTPFarmProbe `pulumi:"probe"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
	Stickiness *string `pulumi:"stickiness"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId *int `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone *string `pulumi:"zone"`
}

type IPLoadBalancingHTTPFarmState struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
	Balance pulumi.StringPtrInput
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrInput
	// Port for backends to recieve traffic on.
	Port pulumi.IntPtrInput
	// define a backend healthcheck probe
	Probe IPLoadBalancingHTTPFarmProbePtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
	Stickiness pulumi.StringPtrInput
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.IntPtrInput
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone pulumi.StringPtrInput
}

func (IPLoadBalancingHTTPFarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadBalancingHTTPFarmState)(nil)).Elem()
}

type iploadBalancingHTTPFarmArgs struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
	Balance *string `pulumi:"balance"`
	// Readable label for loadbalancer farm
	DisplayName *string `pulumi:"displayName"`
	// Port for backends to recieve traffic on.
	Port *int `pulumi:"port"`
	// define a backend healthcheck probe
	Probe *IPLoadBalancingHTTPFarmProbe `pulumi:"probe"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
	Stickiness *string `pulumi:"stickiness"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId *int `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a IPLoadBalancingHTTPFarm resource.
type IPLoadBalancingHTTPFarmArgs struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
	Balance pulumi.StringPtrInput
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrInput
	// Port for backends to recieve traffic on.
	Port pulumi.IntPtrInput
	// define a backend healthcheck probe
	Probe IPLoadBalancingHTTPFarmProbePtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
	Stickiness pulumi.StringPtrInput
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.IntPtrInput
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone pulumi.StringInput
}

func (IPLoadBalancingHTTPFarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadBalancingHTTPFarmArgs)(nil)).Elem()
}

type IPLoadBalancingHTTPFarmInput interface {
	pulumi.Input

	ToIPLoadBalancingHTTPFarmOutput() IPLoadBalancingHTTPFarmOutput
	ToIPLoadBalancingHTTPFarmOutputWithContext(ctx context.Context) IPLoadBalancingHTTPFarmOutput
}

func (*IPLoadBalancingHTTPFarm) ElementType() reflect.Type {
	return reflect.TypeOf((**IPLoadBalancingHTTPFarm)(nil)).Elem()
}

func (i *IPLoadBalancingHTTPFarm) ToIPLoadBalancingHTTPFarmOutput() IPLoadBalancingHTTPFarmOutput {
	return i.ToIPLoadBalancingHTTPFarmOutputWithContext(context.Background())
}

func (i *IPLoadBalancingHTTPFarm) ToIPLoadBalancingHTTPFarmOutputWithContext(ctx context.Context) IPLoadBalancingHTTPFarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingHTTPFarmOutput)
}

// IPLoadBalancingHTTPFarmArrayInput is an input type that accepts IPLoadBalancingHTTPFarmArray and IPLoadBalancingHTTPFarmArrayOutput values.
// You can construct a concrete instance of `IPLoadBalancingHTTPFarmArrayInput` via:
//
//          IPLoadBalancingHTTPFarmArray{ IPLoadBalancingHTTPFarmArgs{...} }
type IPLoadBalancingHTTPFarmArrayInput interface {
	pulumi.Input

	ToIPLoadBalancingHTTPFarmArrayOutput() IPLoadBalancingHTTPFarmArrayOutput
	ToIPLoadBalancingHTTPFarmArrayOutputWithContext(context.Context) IPLoadBalancingHTTPFarmArrayOutput
}

type IPLoadBalancingHTTPFarmArray []IPLoadBalancingHTTPFarmInput

func (IPLoadBalancingHTTPFarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IPLoadBalancingHTTPFarm)(nil)).Elem()
}

func (i IPLoadBalancingHTTPFarmArray) ToIPLoadBalancingHTTPFarmArrayOutput() IPLoadBalancingHTTPFarmArrayOutput {
	return i.ToIPLoadBalancingHTTPFarmArrayOutputWithContext(context.Background())
}

func (i IPLoadBalancingHTTPFarmArray) ToIPLoadBalancingHTTPFarmArrayOutputWithContext(ctx context.Context) IPLoadBalancingHTTPFarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingHTTPFarmArrayOutput)
}

// IPLoadBalancingHTTPFarmMapInput is an input type that accepts IPLoadBalancingHTTPFarmMap and IPLoadBalancingHTTPFarmMapOutput values.
// You can construct a concrete instance of `IPLoadBalancingHTTPFarmMapInput` via:
//
//          IPLoadBalancingHTTPFarmMap{ "key": IPLoadBalancingHTTPFarmArgs{...} }
type IPLoadBalancingHTTPFarmMapInput interface {
	pulumi.Input

	ToIPLoadBalancingHTTPFarmMapOutput() IPLoadBalancingHTTPFarmMapOutput
	ToIPLoadBalancingHTTPFarmMapOutputWithContext(context.Context) IPLoadBalancingHTTPFarmMapOutput
}

type IPLoadBalancingHTTPFarmMap map[string]IPLoadBalancingHTTPFarmInput

func (IPLoadBalancingHTTPFarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IPLoadBalancingHTTPFarm)(nil)).Elem()
}

func (i IPLoadBalancingHTTPFarmMap) ToIPLoadBalancingHTTPFarmMapOutput() IPLoadBalancingHTTPFarmMapOutput {
	return i.ToIPLoadBalancingHTTPFarmMapOutputWithContext(context.Background())
}

func (i IPLoadBalancingHTTPFarmMap) ToIPLoadBalancingHTTPFarmMapOutputWithContext(ctx context.Context) IPLoadBalancingHTTPFarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingHTTPFarmMapOutput)
}

type IPLoadBalancingHTTPFarmOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingHTTPFarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPLoadBalancingHTTPFarm)(nil)).Elem()
}

func (o IPLoadBalancingHTTPFarmOutput) ToIPLoadBalancingHTTPFarmOutput() IPLoadBalancingHTTPFarmOutput {
	return o
}

func (o IPLoadBalancingHTTPFarmOutput) ToIPLoadBalancingHTTPFarmOutputWithContext(ctx context.Context) IPLoadBalancingHTTPFarmOutput {
	return o
}

// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
func (o IPLoadBalancingHTTPFarmOutput) Balance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPLoadBalancingHTTPFarm) pulumi.StringPtrOutput { return v.Balance }).(pulumi.StringPtrOutput)
}

// Readable label for loadbalancer farm
func (o IPLoadBalancingHTTPFarmOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPLoadBalancingHTTPFarm) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Port for backends to recieve traffic on.
func (o IPLoadBalancingHTTPFarmOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IPLoadBalancingHTTPFarm) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// define a backend healthcheck probe
func (o IPLoadBalancingHTTPFarmOutput) Probe() IPLoadBalancingHTTPFarmProbePtrOutput {
	return o.ApplyT(func(v *IPLoadBalancingHTTPFarm) IPLoadBalancingHTTPFarmProbePtrOutput { return v.Probe }).(IPLoadBalancingHTTPFarmProbePtrOutput)
}

// The internal name of your IP load balancing
func (o IPLoadBalancingHTTPFarmOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IPLoadBalancingHTTPFarm) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
func (o IPLoadBalancingHTTPFarmOutput) Stickiness() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPLoadBalancingHTTPFarm) pulumi.StringPtrOutput { return v.Stickiness }).(pulumi.StringPtrOutput)
}

// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
func (o IPLoadBalancingHTTPFarmOutput) VrackNetworkId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IPLoadBalancingHTTPFarm) pulumi.IntPtrOutput { return v.VrackNetworkId }).(pulumi.IntPtrOutput)
}

// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
func (o IPLoadBalancingHTTPFarmOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *IPLoadBalancingHTTPFarm) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type IPLoadBalancingHTTPFarmArrayOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingHTTPFarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IPLoadBalancingHTTPFarm)(nil)).Elem()
}

func (o IPLoadBalancingHTTPFarmArrayOutput) ToIPLoadBalancingHTTPFarmArrayOutput() IPLoadBalancingHTTPFarmArrayOutput {
	return o
}

func (o IPLoadBalancingHTTPFarmArrayOutput) ToIPLoadBalancingHTTPFarmArrayOutputWithContext(ctx context.Context) IPLoadBalancingHTTPFarmArrayOutput {
	return o
}

func (o IPLoadBalancingHTTPFarmArrayOutput) Index(i pulumi.IntInput) IPLoadBalancingHTTPFarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IPLoadBalancingHTTPFarm {
		return vs[0].([]*IPLoadBalancingHTTPFarm)[vs[1].(int)]
	}).(IPLoadBalancingHTTPFarmOutput)
}

type IPLoadBalancingHTTPFarmMapOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingHTTPFarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IPLoadBalancingHTTPFarm)(nil)).Elem()
}

func (o IPLoadBalancingHTTPFarmMapOutput) ToIPLoadBalancingHTTPFarmMapOutput() IPLoadBalancingHTTPFarmMapOutput {
	return o
}

func (o IPLoadBalancingHTTPFarmMapOutput) ToIPLoadBalancingHTTPFarmMapOutputWithContext(ctx context.Context) IPLoadBalancingHTTPFarmMapOutput {
	return o
}

func (o IPLoadBalancingHTTPFarmMapOutput) MapIndex(k pulumi.StringInput) IPLoadBalancingHTTPFarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IPLoadBalancingHTTPFarm {
		return vs[0].(map[string]*IPLoadBalancingHTTPFarm)[vs[1].(string)]
	}).(IPLoadBalancingHTTPFarmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingHTTPFarmInput)(nil)).Elem(), &IPLoadBalancingHTTPFarm{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingHTTPFarmArrayInput)(nil)).Elem(), IPLoadBalancingHTTPFarmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingHTTPFarmMapInput)(nil)).Elem(), IPLoadBalancingHTTPFarmMap{})
	pulumi.RegisterOutputType(IPLoadBalancingHTTPFarmOutput{})
	pulumi.RegisterOutputType(IPLoadBalancingHTTPFarmArrayOutput{})
	pulumi.RegisterOutputType(IPLoadBalancingHTTPFarmMapOutput{})
}
