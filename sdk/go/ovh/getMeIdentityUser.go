// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an identity user.
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
// 		_, err := ovh.LookupMeIdentityUser(ctx, &GetMeIdentityUserArgs{
// 			User: "my_user_login",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupMeIdentityUser(ctx *pulumi.Context, args *LookupMeIdentityUserArgs, opts ...pulumi.InvokeOption) (*LookupMeIdentityUserResult, error) {
	var rv LookupMeIdentityUserResult
	err := ctx.Invoke("ovh:index/getMeIdentityUser:getMeIdentityUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMeIdentityUser.
type LookupMeIdentityUserArgs struct {
	// User's login.
	User string `pulumi:"user"`
}

// A collection of values returned by getMeIdentityUser.
type LookupMeIdentityUserResult struct {
	// Creation date of this user.
	Creation string `pulumi:"creation"`
	// User description.
	Description string `pulumi:"description"`
	// User's email.
	Email string `pulumi:"email"`
	// User's group.
	Group string `pulumi:"group"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Last update of this user.
	LastUpdate string `pulumi:"lastUpdate"`
	// User's login suffix.
	Login string `pulumi:"login"`
	// When the user changed his password for the last time.
	PasswordLastUpdate string `pulumi:"passwordLastUpdate"`
	// Current user's status.
	Status string `pulumi:"status"`
	User   string `pulumi:"user"`
}

func LookupMeIdentityUserOutput(ctx *pulumi.Context, args LookupMeIdentityUserOutputArgs, opts ...pulumi.InvokeOption) LookupMeIdentityUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMeIdentityUserResult, error) {
			args := v.(LookupMeIdentityUserArgs)
			r, err := LookupMeIdentityUser(ctx, &args, opts...)
			var s LookupMeIdentityUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMeIdentityUserResultOutput)
}

// A collection of arguments for invoking getMeIdentityUser.
type LookupMeIdentityUserOutputArgs struct {
	// User's login.
	User pulumi.StringInput `pulumi:"user"`
}

func (LookupMeIdentityUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMeIdentityUserArgs)(nil)).Elem()
}

// A collection of values returned by getMeIdentityUser.
type LookupMeIdentityUserResultOutput struct{ *pulumi.OutputState }

func (LookupMeIdentityUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMeIdentityUserResult)(nil)).Elem()
}

func (o LookupMeIdentityUserResultOutput) ToLookupMeIdentityUserResultOutput() LookupMeIdentityUserResultOutput {
	return o
}

func (o LookupMeIdentityUserResultOutput) ToLookupMeIdentityUserResultOutputWithContext(ctx context.Context) LookupMeIdentityUserResultOutput {
	return o
}

// Creation date of this user.
func (o LookupMeIdentityUserResultOutput) Creation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeIdentityUserResult) string { return v.Creation }).(pulumi.StringOutput)
}

// User description.
func (o LookupMeIdentityUserResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeIdentityUserResult) string { return v.Description }).(pulumi.StringOutput)
}

// User's email.
func (o LookupMeIdentityUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeIdentityUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// User's group.
func (o LookupMeIdentityUserResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeIdentityUserResult) string { return v.Group }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMeIdentityUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeIdentityUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Last update of this user.
func (o LookupMeIdentityUserResultOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeIdentityUserResult) string { return v.LastUpdate }).(pulumi.StringOutput)
}

// User's login suffix.
func (o LookupMeIdentityUserResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeIdentityUserResult) string { return v.Login }).(pulumi.StringOutput)
}

// When the user changed his password for the last time.
func (o LookupMeIdentityUserResultOutput) PasswordLastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeIdentityUserResult) string { return v.PasswordLastUpdate }).(pulumi.StringOutput)
}

// Current user's status.
func (o LookupMeIdentityUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeIdentityUserResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupMeIdentityUserResultOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeIdentityUserResult) string { return v.User }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMeIdentityUserResultOutput{})
}