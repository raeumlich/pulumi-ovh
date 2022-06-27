// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a vrack network for your IP Loadbalancing service.
type IPLoadBalancingVRackNetwork struct {
	pulumi.CustomResourceState

	// Human readable name for your vrack network
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
	FarmIds pulumi.IntArrayOutput `pulumi:"farmIds"`
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp pulumi.StringOutput `pulumi:"natIp"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// IP block of the private network in the vRack
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan pulumi.IntOutput `pulumi:"vlan"`
	// (Required) Internal Load Balancer identifier of the vRack private network
	VrackNetworkId pulumi.IntOutput `pulumi:"vrackNetworkId"`
}

// NewIPLoadBalancingVRackNetwork registers a new resource with the given unique name, arguments, and options.
func NewIPLoadBalancingVRackNetwork(ctx *pulumi.Context,
	name string, args *IPLoadBalancingVRackNetworkArgs, opts ...pulumi.ResourceOption) (*IPLoadBalancingVRackNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NatIp == nil {
		return nil, errors.New("invalid value for required argument 'NatIp'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	var resource IPLoadBalancingVRackNetwork
	err := ctx.RegisterResource("ovh:index/iPLoadBalancingVRackNetwork:IPLoadBalancingVRackNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPLoadBalancingVRackNetwork gets an existing IPLoadBalancingVRackNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPLoadBalancingVRackNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPLoadBalancingVRackNetworkState, opts ...pulumi.ResourceOption) (*IPLoadBalancingVRackNetwork, error) {
	var resource IPLoadBalancingVRackNetwork
	err := ctx.ReadResource("ovh:index/iPLoadBalancingVRackNetwork:IPLoadBalancingVRackNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPLoadBalancingVRackNetwork resources.
type iploadBalancingVRackNetworkState struct {
	// Human readable name for your vrack network
	DisplayName *string `pulumi:"displayName"`
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
	FarmIds []int `pulumi:"farmIds"`
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp *string `pulumi:"natIp"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// IP block of the private network in the vRack
	Subnet *string `pulumi:"subnet"`
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan *int `pulumi:"vlan"`
	// (Required) Internal Load Balancer identifier of the vRack private network
	VrackNetworkId *int `pulumi:"vrackNetworkId"`
}

type IPLoadBalancingVRackNetworkState struct {
	// Human readable name for your vrack network
	DisplayName pulumi.StringPtrInput
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
	FarmIds pulumi.IntArrayInput
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// IP block of the private network in the vRack
	Subnet pulumi.StringPtrInput
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan pulumi.IntPtrInput
	// (Required) Internal Load Balancer identifier of the vRack private network
	VrackNetworkId pulumi.IntPtrInput
}

func (IPLoadBalancingVRackNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadBalancingVRackNetworkState)(nil)).Elem()
}

type iploadBalancingVRackNetworkArgs struct {
	// Human readable name for your vrack network
	DisplayName *string `pulumi:"displayName"`
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
	FarmIds []int `pulumi:"farmIds"`
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp string `pulumi:"natIp"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// IP block of the private network in the vRack
	Subnet string `pulumi:"subnet"`
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan *int `pulumi:"vlan"`
}

// The set of arguments for constructing a IPLoadBalancingVRackNetwork resource.
type IPLoadBalancingVRackNetworkArgs struct {
	// Human readable name for your vrack network
	DisplayName pulumi.StringPtrInput
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
	FarmIds pulumi.IntArrayInput
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp pulumi.StringInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// IP block of the private network in the vRack
	Subnet pulumi.StringInput
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan pulumi.IntPtrInput
}

func (IPLoadBalancingVRackNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadBalancingVRackNetworkArgs)(nil)).Elem()
}

type IPLoadBalancingVRackNetworkInput interface {
	pulumi.Input

	ToIPLoadBalancingVRackNetworkOutput() IPLoadBalancingVRackNetworkOutput
	ToIPLoadBalancingVRackNetworkOutputWithContext(ctx context.Context) IPLoadBalancingVRackNetworkOutput
}

func (*IPLoadBalancingVRackNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**IPLoadBalancingVRackNetwork)(nil)).Elem()
}

func (i *IPLoadBalancingVRackNetwork) ToIPLoadBalancingVRackNetworkOutput() IPLoadBalancingVRackNetworkOutput {
	return i.ToIPLoadBalancingVRackNetworkOutputWithContext(context.Background())
}

func (i *IPLoadBalancingVRackNetwork) ToIPLoadBalancingVRackNetworkOutputWithContext(ctx context.Context) IPLoadBalancingVRackNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingVRackNetworkOutput)
}

// IPLoadBalancingVRackNetworkArrayInput is an input type that accepts IPLoadBalancingVRackNetworkArray and IPLoadBalancingVRackNetworkArrayOutput values.
// You can construct a concrete instance of `IPLoadBalancingVRackNetworkArrayInput` via:
//
//          IPLoadBalancingVRackNetworkArray{ IPLoadBalancingVRackNetworkArgs{...} }
type IPLoadBalancingVRackNetworkArrayInput interface {
	pulumi.Input

	ToIPLoadBalancingVRackNetworkArrayOutput() IPLoadBalancingVRackNetworkArrayOutput
	ToIPLoadBalancingVRackNetworkArrayOutputWithContext(context.Context) IPLoadBalancingVRackNetworkArrayOutput
}

type IPLoadBalancingVRackNetworkArray []IPLoadBalancingVRackNetworkInput

func (IPLoadBalancingVRackNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IPLoadBalancingVRackNetwork)(nil)).Elem()
}

func (i IPLoadBalancingVRackNetworkArray) ToIPLoadBalancingVRackNetworkArrayOutput() IPLoadBalancingVRackNetworkArrayOutput {
	return i.ToIPLoadBalancingVRackNetworkArrayOutputWithContext(context.Background())
}

func (i IPLoadBalancingVRackNetworkArray) ToIPLoadBalancingVRackNetworkArrayOutputWithContext(ctx context.Context) IPLoadBalancingVRackNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingVRackNetworkArrayOutput)
}

// IPLoadBalancingVRackNetworkMapInput is an input type that accepts IPLoadBalancingVRackNetworkMap and IPLoadBalancingVRackNetworkMapOutput values.
// You can construct a concrete instance of `IPLoadBalancingVRackNetworkMapInput` via:
//
//          IPLoadBalancingVRackNetworkMap{ "key": IPLoadBalancingVRackNetworkArgs{...} }
type IPLoadBalancingVRackNetworkMapInput interface {
	pulumi.Input

	ToIPLoadBalancingVRackNetworkMapOutput() IPLoadBalancingVRackNetworkMapOutput
	ToIPLoadBalancingVRackNetworkMapOutputWithContext(context.Context) IPLoadBalancingVRackNetworkMapOutput
}

type IPLoadBalancingVRackNetworkMap map[string]IPLoadBalancingVRackNetworkInput

func (IPLoadBalancingVRackNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IPLoadBalancingVRackNetwork)(nil)).Elem()
}

func (i IPLoadBalancingVRackNetworkMap) ToIPLoadBalancingVRackNetworkMapOutput() IPLoadBalancingVRackNetworkMapOutput {
	return i.ToIPLoadBalancingVRackNetworkMapOutputWithContext(context.Background())
}

func (i IPLoadBalancingVRackNetworkMap) ToIPLoadBalancingVRackNetworkMapOutputWithContext(ctx context.Context) IPLoadBalancingVRackNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPLoadBalancingVRackNetworkMapOutput)
}

type IPLoadBalancingVRackNetworkOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingVRackNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPLoadBalancingVRackNetwork)(nil)).Elem()
}

func (o IPLoadBalancingVRackNetworkOutput) ToIPLoadBalancingVRackNetworkOutput() IPLoadBalancingVRackNetworkOutput {
	return o
}

func (o IPLoadBalancingVRackNetworkOutput) ToIPLoadBalancingVRackNetworkOutputWithContext(ctx context.Context) IPLoadBalancingVRackNetworkOutput {
	return o
}

// Human readable name for your vrack network
func (o IPLoadBalancingVRackNetworkOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPLoadBalancingVRackNetwork) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrackNetworkId` attribute
func (o IPLoadBalancingVRackNetworkOutput) FarmIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *IPLoadBalancingVRackNetwork) pulumi.IntArrayOutput { return v.FarmIds }).(pulumi.IntArrayOutput)
}

// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
func (o IPLoadBalancingVRackNetworkOutput) NatIp() pulumi.StringOutput {
	return o.ApplyT(func(v *IPLoadBalancingVRackNetwork) pulumi.StringOutput { return v.NatIp }).(pulumi.StringOutput)
}

// The internal name of your IP load balancing
func (o IPLoadBalancingVRackNetworkOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IPLoadBalancingVRackNetwork) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// IP block of the private network in the vRack
func (o IPLoadBalancingVRackNetworkOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *IPLoadBalancingVRackNetwork) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
func (o IPLoadBalancingVRackNetworkOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v *IPLoadBalancingVRackNetwork) pulumi.IntOutput { return v.Vlan }).(pulumi.IntOutput)
}

// (Required) Internal Load Balancer identifier of the vRack private network
func (o IPLoadBalancingVRackNetworkOutput) VrackNetworkId() pulumi.IntOutput {
	return o.ApplyT(func(v *IPLoadBalancingVRackNetwork) pulumi.IntOutput { return v.VrackNetworkId }).(pulumi.IntOutput)
}

type IPLoadBalancingVRackNetworkArrayOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingVRackNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IPLoadBalancingVRackNetwork)(nil)).Elem()
}

func (o IPLoadBalancingVRackNetworkArrayOutput) ToIPLoadBalancingVRackNetworkArrayOutput() IPLoadBalancingVRackNetworkArrayOutput {
	return o
}

func (o IPLoadBalancingVRackNetworkArrayOutput) ToIPLoadBalancingVRackNetworkArrayOutputWithContext(ctx context.Context) IPLoadBalancingVRackNetworkArrayOutput {
	return o
}

func (o IPLoadBalancingVRackNetworkArrayOutput) Index(i pulumi.IntInput) IPLoadBalancingVRackNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IPLoadBalancingVRackNetwork {
		return vs[0].([]*IPLoadBalancingVRackNetwork)[vs[1].(int)]
	}).(IPLoadBalancingVRackNetworkOutput)
}

type IPLoadBalancingVRackNetworkMapOutput struct{ *pulumi.OutputState }

func (IPLoadBalancingVRackNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IPLoadBalancingVRackNetwork)(nil)).Elem()
}

func (o IPLoadBalancingVRackNetworkMapOutput) ToIPLoadBalancingVRackNetworkMapOutput() IPLoadBalancingVRackNetworkMapOutput {
	return o
}

func (o IPLoadBalancingVRackNetworkMapOutput) ToIPLoadBalancingVRackNetworkMapOutputWithContext(ctx context.Context) IPLoadBalancingVRackNetworkMapOutput {
	return o
}

func (o IPLoadBalancingVRackNetworkMapOutput) MapIndex(k pulumi.StringInput) IPLoadBalancingVRackNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IPLoadBalancingVRackNetwork {
		return vs[0].(map[string]*IPLoadBalancingVRackNetwork)[vs[1].(string)]
	}).(IPLoadBalancingVRackNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingVRackNetworkInput)(nil)).Elem(), &IPLoadBalancingVRackNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingVRackNetworkArrayInput)(nil)).Elem(), IPLoadBalancingVRackNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPLoadBalancingVRackNetworkMapInput)(nil)).Elem(), IPLoadBalancingVRackNetworkMap{})
	pulumi.RegisterOutputType(IPLoadBalancingVRackNetworkOutput{})
	pulumi.RegisterOutputType(IPLoadBalancingVRackNetworkArrayOutput{})
	pulumi.RegisterOutputType(IPLoadBalancingVRackNetworkMapOutput{})
}