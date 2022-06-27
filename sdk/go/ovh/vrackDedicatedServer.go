// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attach a dedicated server to a VRack.
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
// 		_, err := ovh.NewVRackDedicatedServer(ctx, "vds", &ovh.VRackDedicatedServerArgs{
// 			ServerId:    pulumi.String("67890"),
// 			ServiceName: pulumi.String("XXXX"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VRackDedicatedServer struct {
	pulumi.CustomResourceState

	// The id of the dedicated server.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewVRackDedicatedServer registers a new resource with the given unique name, arguments, and options.
func NewVRackDedicatedServer(ctx *pulumi.Context,
	name string, args *VRackDedicatedServerArgs, opts ...pulumi.ResourceOption) (*VRackDedicatedServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource VRackDedicatedServer
	err := ctx.RegisterResource("ovh:index/vRackDedicatedServer:VRackDedicatedServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVRackDedicatedServer gets an existing VRackDedicatedServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVRackDedicatedServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VRackDedicatedServerState, opts ...pulumi.ResourceOption) (*VRackDedicatedServer, error) {
	var resource VRackDedicatedServer
	err := ctx.ReadResource("ovh:index/vRackDedicatedServer:VRackDedicatedServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VRackDedicatedServer resources.
type vrackDedicatedServerState struct {
	// The id of the dedicated server.
	ServerId *string `pulumi:"serverId"`
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

type VRackDedicatedServerState struct {
	// The id of the dedicated server.
	ServerId pulumi.StringPtrInput
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (VRackDedicatedServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackDedicatedServerState)(nil)).Elem()
}

type vrackDedicatedServerArgs struct {
	// The id of the dedicated server.
	ServerId string `pulumi:"serverId"`
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a VRackDedicatedServer resource.
type VRackDedicatedServerArgs struct {
	// The id of the dedicated server.
	ServerId pulumi.StringInput
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (VRackDedicatedServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackDedicatedServerArgs)(nil)).Elem()
}

type VRackDedicatedServerInput interface {
	pulumi.Input

	ToVRackDedicatedServerOutput() VRackDedicatedServerOutput
	ToVRackDedicatedServerOutputWithContext(ctx context.Context) VRackDedicatedServerOutput
}

func (*VRackDedicatedServer) ElementType() reflect.Type {
	return reflect.TypeOf((**VRackDedicatedServer)(nil)).Elem()
}

func (i *VRackDedicatedServer) ToVRackDedicatedServerOutput() VRackDedicatedServerOutput {
	return i.ToVRackDedicatedServerOutputWithContext(context.Background())
}

func (i *VRackDedicatedServer) ToVRackDedicatedServerOutputWithContext(ctx context.Context) VRackDedicatedServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VRackDedicatedServerOutput)
}

// VRackDedicatedServerArrayInput is an input type that accepts VRackDedicatedServerArray and VRackDedicatedServerArrayOutput values.
// You can construct a concrete instance of `VRackDedicatedServerArrayInput` via:
//
//          VRackDedicatedServerArray{ VRackDedicatedServerArgs{...} }
type VRackDedicatedServerArrayInput interface {
	pulumi.Input

	ToVRackDedicatedServerArrayOutput() VRackDedicatedServerArrayOutput
	ToVRackDedicatedServerArrayOutputWithContext(context.Context) VRackDedicatedServerArrayOutput
}

type VRackDedicatedServerArray []VRackDedicatedServerInput

func (VRackDedicatedServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VRackDedicatedServer)(nil)).Elem()
}

func (i VRackDedicatedServerArray) ToVRackDedicatedServerArrayOutput() VRackDedicatedServerArrayOutput {
	return i.ToVRackDedicatedServerArrayOutputWithContext(context.Background())
}

func (i VRackDedicatedServerArray) ToVRackDedicatedServerArrayOutputWithContext(ctx context.Context) VRackDedicatedServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VRackDedicatedServerArrayOutput)
}

// VRackDedicatedServerMapInput is an input type that accepts VRackDedicatedServerMap and VRackDedicatedServerMapOutput values.
// You can construct a concrete instance of `VRackDedicatedServerMapInput` via:
//
//          VRackDedicatedServerMap{ "key": VRackDedicatedServerArgs{...} }
type VRackDedicatedServerMapInput interface {
	pulumi.Input

	ToVRackDedicatedServerMapOutput() VRackDedicatedServerMapOutput
	ToVRackDedicatedServerMapOutputWithContext(context.Context) VRackDedicatedServerMapOutput
}

type VRackDedicatedServerMap map[string]VRackDedicatedServerInput

func (VRackDedicatedServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VRackDedicatedServer)(nil)).Elem()
}

func (i VRackDedicatedServerMap) ToVRackDedicatedServerMapOutput() VRackDedicatedServerMapOutput {
	return i.ToVRackDedicatedServerMapOutputWithContext(context.Background())
}

func (i VRackDedicatedServerMap) ToVRackDedicatedServerMapOutputWithContext(ctx context.Context) VRackDedicatedServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VRackDedicatedServerMapOutput)
}

type VRackDedicatedServerOutput struct{ *pulumi.OutputState }

func (VRackDedicatedServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VRackDedicatedServer)(nil)).Elem()
}

func (o VRackDedicatedServerOutput) ToVRackDedicatedServerOutput() VRackDedicatedServerOutput {
	return o
}

func (o VRackDedicatedServerOutput) ToVRackDedicatedServerOutputWithContext(ctx context.Context) VRackDedicatedServerOutput {
	return o
}

// The id of the dedicated server.
func (o VRackDedicatedServerOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *VRackDedicatedServer) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

// The id of the vrack. If omitted,
// the `OVH_VRACK_SERVICE` environment variable is used.
func (o VRackDedicatedServerOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VRackDedicatedServer) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type VRackDedicatedServerArrayOutput struct{ *pulumi.OutputState }

func (VRackDedicatedServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VRackDedicatedServer)(nil)).Elem()
}

func (o VRackDedicatedServerArrayOutput) ToVRackDedicatedServerArrayOutput() VRackDedicatedServerArrayOutput {
	return o
}

func (o VRackDedicatedServerArrayOutput) ToVRackDedicatedServerArrayOutputWithContext(ctx context.Context) VRackDedicatedServerArrayOutput {
	return o
}

func (o VRackDedicatedServerArrayOutput) Index(i pulumi.IntInput) VRackDedicatedServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VRackDedicatedServer {
		return vs[0].([]*VRackDedicatedServer)[vs[1].(int)]
	}).(VRackDedicatedServerOutput)
}

type VRackDedicatedServerMapOutput struct{ *pulumi.OutputState }

func (VRackDedicatedServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VRackDedicatedServer)(nil)).Elem()
}

func (o VRackDedicatedServerMapOutput) ToVRackDedicatedServerMapOutput() VRackDedicatedServerMapOutput {
	return o
}

func (o VRackDedicatedServerMapOutput) ToVRackDedicatedServerMapOutputWithContext(ctx context.Context) VRackDedicatedServerMapOutput {
	return o
}

func (o VRackDedicatedServerMapOutput) MapIndex(k pulumi.StringInput) VRackDedicatedServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VRackDedicatedServer {
		return vs[0].(map[string]*VRackDedicatedServer)[vs[1].(string)]
	}).(VRackDedicatedServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VRackDedicatedServerInput)(nil)).Elem(), &VRackDedicatedServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*VRackDedicatedServerArrayInput)(nil)).Elem(), VRackDedicatedServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VRackDedicatedServerMapInput)(nil)).Elem(), VRackDedicatedServerMap{})
	pulumi.RegisterOutputType(VRackDedicatedServerOutput{})
	pulumi.RegisterOutputType(VRackDedicatedServerArrayOutput{})
	pulumi.RegisterOutputType(VRackDedicatedServerMapOutput{})
}