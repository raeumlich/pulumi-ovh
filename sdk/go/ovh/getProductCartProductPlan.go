// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information of order cart product plan.
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
// 		mycart, err := ovh.GetOrderCart(ctx, &GetOrderCartArgs{
// 			OvhSubsidiary: "fr",
// 			Description:   pulumi.StringRef("my cart"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ovh.GetProductCartProductPlan(ctx, &GetProductCartProductPlanArgs{
// 			CartId:        mycart.Id,
// 			PriceCapacity: "renew",
// 			Product:       "cloud",
// 			PlanCode:      "project",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetProductCartProductPlan(ctx *pulumi.Context, args *GetProductCartProductPlanArgs, opts ...pulumi.InvokeOption) (*GetProductCartProductPlanResult, error) {
	var rv GetProductCartProductPlanResult
	err := ctx.Invoke("ovh:index/getProductCartProductPlan:getProductCartProductPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProductCartProductPlan.
type GetProductCartProductPlanArgs struct {
	// Cart identifier
	CartId string `pulumi:"cartId"`
	// Catalog name
	CatalogName *string `pulumi:"catalogName"`
	// Product offer identifier
	PlanCode string `pulumi:"planCode"`
	// Capacity of the pricing (type of pricing)
	PriceCapacity string `pulumi:"priceCapacity"`
	// Product
	Product string `pulumi:"product"`
}

// A collection of values returned by getProductCartProductPlan.
type GetProductCartProductPlanResult struct {
	CartId      string  `pulumi:"cartId"`
	CatalogName *string `pulumi:"catalogName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Product offer identifier
	PlanCode      string `pulumi:"planCode"`
	PriceCapacity string `pulumi:"priceCapacity"`
	// Prices of the product offer
	Prices  []GetProductCartProductPlanPrice `pulumi:"prices"`
	Product string                           `pulumi:"product"`
	// Name of the product
	ProductName string `pulumi:"productName"`
	// Product type
	ProductType string `pulumi:"productType"`
	// Selected Price according to capacity
	SelectedPrices []GetProductCartProductPlanSelectedPrice `pulumi:"selectedPrices"`
}

func GetProductCartProductPlanOutput(ctx *pulumi.Context, args GetProductCartProductPlanOutputArgs, opts ...pulumi.InvokeOption) GetProductCartProductPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProductCartProductPlanResult, error) {
			args := v.(GetProductCartProductPlanArgs)
			r, err := GetProductCartProductPlan(ctx, &args, opts...)
			var s GetProductCartProductPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProductCartProductPlanResultOutput)
}

// A collection of arguments for invoking getProductCartProductPlan.
type GetProductCartProductPlanOutputArgs struct {
	// Cart identifier
	CartId pulumi.StringInput `pulumi:"cartId"`
	// Catalog name
	CatalogName pulumi.StringPtrInput `pulumi:"catalogName"`
	// Product offer identifier
	PlanCode pulumi.StringInput `pulumi:"planCode"`
	// Capacity of the pricing (type of pricing)
	PriceCapacity pulumi.StringInput `pulumi:"priceCapacity"`
	// Product
	Product pulumi.StringInput `pulumi:"product"`
}

func (GetProductCartProductPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductCartProductPlanArgs)(nil)).Elem()
}

// A collection of values returned by getProductCartProductPlan.
type GetProductCartProductPlanResultOutput struct{ *pulumi.OutputState }

func (GetProductCartProductPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductCartProductPlanResult)(nil)).Elem()
}

func (o GetProductCartProductPlanResultOutput) ToGetProductCartProductPlanResultOutput() GetProductCartProductPlanResultOutput {
	return o
}

func (o GetProductCartProductPlanResultOutput) ToGetProductCartProductPlanResultOutputWithContext(ctx context.Context) GetProductCartProductPlanResultOutput {
	return o
}

func (o GetProductCartProductPlanResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductCartProductPlanResult) string { return v.CartId }).(pulumi.StringOutput)
}

func (o GetProductCartProductPlanResultOutput) CatalogName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductCartProductPlanResult) *string { return v.CatalogName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProductCartProductPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductCartProductPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

// Product offer identifier
func (o GetProductCartProductPlanResultOutput) PlanCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductCartProductPlanResult) string { return v.PlanCode }).(pulumi.StringOutput)
}

func (o GetProductCartProductPlanResultOutput) PriceCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductCartProductPlanResult) string { return v.PriceCapacity }).(pulumi.StringOutput)
}

// Prices of the product offer
func (o GetProductCartProductPlanResultOutput) Prices() GetProductCartProductPlanPriceArrayOutput {
	return o.ApplyT(func(v GetProductCartProductPlanResult) []GetProductCartProductPlanPrice { return v.Prices }).(GetProductCartProductPlanPriceArrayOutput)
}

func (o GetProductCartProductPlanResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductCartProductPlanResult) string { return v.Product }).(pulumi.StringOutput)
}

// Name of the product
func (o GetProductCartProductPlanResultOutput) ProductName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductCartProductPlanResult) string { return v.ProductName }).(pulumi.StringOutput)
}

// Product type
func (o GetProductCartProductPlanResultOutput) ProductType() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductCartProductPlanResult) string { return v.ProductType }).(pulumi.StringOutput)
}

// Selected Price according to capacity
func (o GetProductCartProductPlanResultOutput) SelectedPrices() GetProductCartProductPlanSelectedPriceArrayOutput {
	return o.ApplyT(func(v GetProductCartProductPlanResult) []GetProductCartProductPlanSelectedPrice {
		return v.SelectedPrices
	}).(GetProductCartProductPlanSelectedPriceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProductCartProductPlanResultOutput{})
}
