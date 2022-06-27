// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a user in a public cloud project.
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
// 		_, err := ovh.NewCloudProjectUser(ctx, "user1", &ovh.CloudProjectUserArgs{
// 			ServiceName: pulumi.String("XXX"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CloudProjectUser struct {
	pulumi.CustomResourceState

	// the date the user was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// A description associated with the user.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// a convenient map representing an openstackRc file.
	// Note: no password nor sensitive token is set in this map.
	OpenstackRc pulumi.MapOutput `pulumi:"openstackRc"`
	// (Sensitive) the password generated for the user. The password can
	// be used with the Openstack API. This attribute is sensitive and will only be
	// retrieve once during creation.
	Password pulumi.StringOutput `pulumi:"password"`
	// The name of a role. See `roleNames`.
	RoleName pulumi.StringPtrOutput `pulumi:"roleName"`
	// A list of role names. Values can be:
	// - administrator,
	// - aiTrainingOperator
	// - authentication
	// - backupOperator
	// - computeOperator
	// - imageOperator
	// - infrastructureSupervisor
	// - networkOperator
	// - networkSecurityOperator
	// - objectstoreOperator
	// - volume_operator
	RoleNames pulumi.StringArrayOutput `pulumi:"roleNames"`
	// A list of roles associated with the user.
	Roles CloudProjectUserRoleArrayOutput `pulumi:"roles"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// the status of the user. should be normally set to 'ok'.
	Status pulumi.StringOutput `pulumi:"status"`
	// the username generated for the user. This username can be used with
	// the Openstack API.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewCloudProjectUser registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectUser(ctx *pulumi.Context,
	name string, args *CloudProjectUserArgs, opts ...pulumi.ResourceOption) (*CloudProjectUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource CloudProjectUser
	err := ctx.RegisterResource("ovh:index/cloudProjectUser:CloudProjectUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectUser gets an existing CloudProjectUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectUserState, opts ...pulumi.ResourceOption) (*CloudProjectUser, error) {
	var resource CloudProjectUser
	err := ctx.ReadResource("ovh:index/cloudProjectUser:CloudProjectUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectUser resources.
type cloudProjectUserState struct {
	// the date the user was created.
	CreationDate *string `pulumi:"creationDate"`
	// A description associated with the user.
	Description *string `pulumi:"description"`
	// a convenient map representing an openstackRc file.
	// Note: no password nor sensitive token is set in this map.
	OpenstackRc map[string]interface{} `pulumi:"openstackRc"`
	// (Sensitive) the password generated for the user. The password can
	// be used with the Openstack API. This attribute is sensitive and will only be
	// retrieve once during creation.
	Password *string `pulumi:"password"`
	// The name of a role. See `roleNames`.
	RoleName *string `pulumi:"roleName"`
	// A list of role names. Values can be:
	// - administrator,
	// - aiTrainingOperator
	// - authentication
	// - backupOperator
	// - computeOperator
	// - imageOperator
	// - infrastructureSupervisor
	// - networkOperator
	// - networkSecurityOperator
	// - objectstoreOperator
	// - volume_operator
	RoleNames []string `pulumi:"roleNames"`
	// A list of roles associated with the user.
	Roles []CloudProjectUserRole `pulumi:"roles"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// the status of the user. should be normally set to 'ok'.
	Status *string `pulumi:"status"`
	// the username generated for the user. This username can be used with
	// the Openstack API.
	Username *string `pulumi:"username"`
}

type CloudProjectUserState struct {
	// the date the user was created.
	CreationDate pulumi.StringPtrInput
	// A description associated with the user.
	Description pulumi.StringPtrInput
	// a convenient map representing an openstackRc file.
	// Note: no password nor sensitive token is set in this map.
	OpenstackRc pulumi.MapInput
	// (Sensitive) the password generated for the user. The password can
	// be used with the Openstack API. This attribute is sensitive and will only be
	// retrieve once during creation.
	Password pulumi.StringPtrInput
	// The name of a role. See `roleNames`.
	RoleName pulumi.StringPtrInput
	// A list of role names. Values can be:
	// - administrator,
	// - aiTrainingOperator
	// - authentication
	// - backupOperator
	// - computeOperator
	// - imageOperator
	// - infrastructureSupervisor
	// - networkOperator
	// - networkSecurityOperator
	// - objectstoreOperator
	// - volume_operator
	RoleNames pulumi.StringArrayInput
	// A list of roles associated with the user.
	Roles CloudProjectUserRoleArrayInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// the status of the user. should be normally set to 'ok'.
	Status pulumi.StringPtrInput
	// the username generated for the user. This username can be used with
	// the Openstack API.
	Username pulumi.StringPtrInput
}

func (CloudProjectUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectUserState)(nil)).Elem()
}

type cloudProjectUserArgs struct {
	// A description associated with the user.
	Description *string `pulumi:"description"`
	// a convenient map representing an openstackRc file.
	// Note: no password nor sensitive token is set in this map.
	OpenstackRc map[string]interface{} `pulumi:"openstackRc"`
	// The name of a role. See `roleNames`.
	RoleName *string `pulumi:"roleName"`
	// A list of role names. Values can be:
	// - administrator,
	// - aiTrainingOperator
	// - authentication
	// - backupOperator
	// - computeOperator
	// - imageOperator
	// - infrastructureSupervisor
	// - networkOperator
	// - networkSecurityOperator
	// - objectstoreOperator
	// - volume_operator
	RoleNames []string `pulumi:"roleNames"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectUser resource.
type CloudProjectUserArgs struct {
	// A description associated with the user.
	Description pulumi.StringPtrInput
	// a convenient map representing an openstackRc file.
	// Note: no password nor sensitive token is set in this map.
	OpenstackRc pulumi.MapInput
	// The name of a role. See `roleNames`.
	RoleName pulumi.StringPtrInput
	// A list of role names. Values can be:
	// - administrator,
	// - aiTrainingOperator
	// - authentication
	// - backupOperator
	// - computeOperator
	// - imageOperator
	// - infrastructureSupervisor
	// - networkOperator
	// - networkSecurityOperator
	// - objectstoreOperator
	// - volume_operator
	RoleNames pulumi.StringArrayInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (CloudProjectUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectUserArgs)(nil)).Elem()
}

type CloudProjectUserInput interface {
	pulumi.Input

	ToCloudProjectUserOutput() CloudProjectUserOutput
	ToCloudProjectUserOutputWithContext(ctx context.Context) CloudProjectUserOutput
}

func (*CloudProjectUser) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectUser)(nil)).Elem()
}

func (i *CloudProjectUser) ToCloudProjectUserOutput() CloudProjectUserOutput {
	return i.ToCloudProjectUserOutputWithContext(context.Background())
}

func (i *CloudProjectUser) ToCloudProjectUserOutputWithContext(ctx context.Context) CloudProjectUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectUserOutput)
}

// CloudProjectUserArrayInput is an input type that accepts CloudProjectUserArray and CloudProjectUserArrayOutput values.
// You can construct a concrete instance of `CloudProjectUserArrayInput` via:
//
//          CloudProjectUserArray{ CloudProjectUserArgs{...} }
type CloudProjectUserArrayInput interface {
	pulumi.Input

	ToCloudProjectUserArrayOutput() CloudProjectUserArrayOutput
	ToCloudProjectUserArrayOutputWithContext(context.Context) CloudProjectUserArrayOutput
}

type CloudProjectUserArray []CloudProjectUserInput

func (CloudProjectUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectUser)(nil)).Elem()
}

func (i CloudProjectUserArray) ToCloudProjectUserArrayOutput() CloudProjectUserArrayOutput {
	return i.ToCloudProjectUserArrayOutputWithContext(context.Background())
}

func (i CloudProjectUserArray) ToCloudProjectUserArrayOutputWithContext(ctx context.Context) CloudProjectUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectUserArrayOutput)
}

// CloudProjectUserMapInput is an input type that accepts CloudProjectUserMap and CloudProjectUserMapOutput values.
// You can construct a concrete instance of `CloudProjectUserMapInput` via:
//
//          CloudProjectUserMap{ "key": CloudProjectUserArgs{...} }
type CloudProjectUserMapInput interface {
	pulumi.Input

	ToCloudProjectUserMapOutput() CloudProjectUserMapOutput
	ToCloudProjectUserMapOutputWithContext(context.Context) CloudProjectUserMapOutput
}

type CloudProjectUserMap map[string]CloudProjectUserInput

func (CloudProjectUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectUser)(nil)).Elem()
}

func (i CloudProjectUserMap) ToCloudProjectUserMapOutput() CloudProjectUserMapOutput {
	return i.ToCloudProjectUserMapOutputWithContext(context.Background())
}

func (i CloudProjectUserMap) ToCloudProjectUserMapOutputWithContext(ctx context.Context) CloudProjectUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectUserMapOutput)
}

type CloudProjectUserOutput struct{ *pulumi.OutputState }

func (CloudProjectUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectUser)(nil)).Elem()
}

func (o CloudProjectUserOutput) ToCloudProjectUserOutput() CloudProjectUserOutput {
	return o
}

func (o CloudProjectUserOutput) ToCloudProjectUserOutputWithContext(ctx context.Context) CloudProjectUserOutput {
	return o
}

// the date the user was created.
func (o CloudProjectUserOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// A description associated with the user.
func (o CloudProjectUserOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// a convenient map representing an openstackRc file.
// Note: no password nor sensitive token is set in this map.
func (o CloudProjectUserOutput) OpenstackRc() pulumi.MapOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.MapOutput { return v.OpenstackRc }).(pulumi.MapOutput)
}

// (Sensitive) the password generated for the user. The password can
// be used with the Openstack API. This attribute is sensitive and will only be
// retrieve once during creation.
func (o CloudProjectUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The name of a role. See `roleNames`.
func (o CloudProjectUserOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringPtrOutput { return v.RoleName }).(pulumi.StringPtrOutput)
}

// A list of role names. Values can be:
// - administrator,
// - aiTrainingOperator
// - authentication
// - backupOperator
// - computeOperator
// - imageOperator
// - infrastructureSupervisor
// - networkOperator
// - networkSecurityOperator
// - objectstoreOperator
// - volume_operator
func (o CloudProjectUserOutput) RoleNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringArrayOutput { return v.RoleNames }).(pulumi.StringArrayOutput)
}

// A list of roles associated with the user.
func (o CloudProjectUserOutput) Roles() CloudProjectUserRoleArrayOutput {
	return o.ApplyT(func(v *CloudProjectUser) CloudProjectUserRoleArrayOutput { return v.Roles }).(CloudProjectUserRoleArrayOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o CloudProjectUserOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// the status of the user. should be normally set to 'ok'.
func (o CloudProjectUserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// the username generated for the user. This username can be used with
// the Openstack API.
func (o CloudProjectUserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type CloudProjectUserArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectUser)(nil)).Elem()
}

func (o CloudProjectUserArrayOutput) ToCloudProjectUserArrayOutput() CloudProjectUserArrayOutput {
	return o
}

func (o CloudProjectUserArrayOutput) ToCloudProjectUserArrayOutputWithContext(ctx context.Context) CloudProjectUserArrayOutput {
	return o
}

func (o CloudProjectUserArrayOutput) Index(i pulumi.IntInput) CloudProjectUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProjectUser {
		return vs[0].([]*CloudProjectUser)[vs[1].(int)]
	}).(CloudProjectUserOutput)
}

type CloudProjectUserMapOutput struct{ *pulumi.OutputState }

func (CloudProjectUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectUser)(nil)).Elem()
}

func (o CloudProjectUserMapOutput) ToCloudProjectUserMapOutput() CloudProjectUserMapOutput {
	return o
}

func (o CloudProjectUserMapOutput) ToCloudProjectUserMapOutputWithContext(ctx context.Context) CloudProjectUserMapOutput {
	return o
}

func (o CloudProjectUserMapOutput) MapIndex(k pulumi.StringInput) CloudProjectUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProjectUser {
		return vs[0].(map[string]*CloudProjectUser)[vs[1].(string)]
	}).(CloudProjectUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectUserInput)(nil)).Elem(), &CloudProjectUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectUserArrayInput)(nil)).Elem(), CloudProjectUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectUserMapInput)(nil)).Elem(), CloudProjectUserMap{})
	pulumi.RegisterOutputType(CloudProjectUserOutput{})
	pulumi.RegisterOutputType(CloudProjectUserArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectUserMapOutput{})
}