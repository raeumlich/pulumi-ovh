// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a DBaas logs output graylog stream.
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
// 		_, err := ovh.LookupDBaaSLogsOutputGraylogStream(ctx, &GetDBaaSLogsOutputGraylogStreamArgs{
// 			ServiceName: "XXXXXX",
// 			Title:       "my stream",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDBaaSLogsOutputGraylogStream(ctx *pulumi.Context, args *LookupDBaaSLogsOutputGraylogStreamArgs, opts ...pulumi.InvokeOption) (*LookupDBaaSLogsOutputGraylogStreamResult, error) {
	var rv LookupDBaaSLogsOutputGraylogStreamResult
	err := ctx.Invoke("ovh:index/getDBaaSLogsOutputGraylogStream:getDBaaSLogsOutputGraylogStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDBaaSLogsOutputGraylogStream.
type LookupDBaaSLogsOutputGraylogStreamArgs struct {
	// The service name
	ServiceName string `pulumi:"serviceName"`
	// Stream description
	Title string `pulumi:"title"`
}

// A collection of values returned by getDBaaSLogsOutputGraylogStream.
type LookupDBaaSLogsOutputGraylogStreamResult struct {
	CanAlert bool `pulumi:"canAlert"`
	// Cold storage compression method
	ColdStorageCompression string `pulumi:"coldStorageCompression"`
	// ColdStorage content
	ColdStorageContent string `pulumi:"coldStorageContent"`
	// Is Cold storage enabled?
	ColdStorageEnabled bool `pulumi:"coldStorageEnabled"`
	// Notify on new Cold storage archive
	ColdStorageNotifyEnabled bool `pulumi:"coldStorageNotifyEnabled"`
	// Cold storage retention in year
	ColdStorageRetention int `pulumi:"coldStorageRetention"`
	// ColdStorage destination
	ColdStorageTarget string `pulumi:"coldStorageTarget"`
	// Stream creation
	CreatedAt string `pulumi:"createdAt"`
	// Stream description
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable ES indexing
	IndexingEnabled bool `pulumi:"indexingEnabled"`
	// Maximum indexing size (in GB)
	IndexingMaxSize int `pulumi:"indexingMaxSize"`
	// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
	IndexingNotifyEnabled bool `pulumi:"indexingNotifyEnabled"`
	// Indicates if you are allowed to edit entry
	IsEditable bool `pulumi:"isEditable"`
	// Indicates if you are allowed to share entry
	IsShareable bool `pulumi:"isShareable"`
	// Number of alert condition
	NbAlertCondition int `pulumi:"nbAlertCondition"`
	// Number of coldstored archives
	NbArchive int `pulumi:"nbArchive"`
	// Parent stream ID
	ParentStreamId string `pulumi:"parentStreamId"`
	// If set, pause indexing when maximum size is reach
	PauseIndexingOnMaxSize bool `pulumi:"pauseIndexingOnMaxSize"`
	// Retention ID
	RetentionId string `pulumi:"retentionId"`
	ServiceName string `pulumi:"serviceName"`
	// Stream ID
	StreamId string `pulumi:"streamId"`
	Title    string `pulumi:"title"`
	// Stream last update
	UpdatedAt string `pulumi:"updatedAt"`
	// Enable Websocket
	WebSocketEnabled string `pulumi:"webSocketEnabled"`
}

func LookupDBaaSLogsOutputGraylogStreamOutput(ctx *pulumi.Context, args LookupDBaaSLogsOutputGraylogStreamOutputArgs, opts ...pulumi.InvokeOption) LookupDBaaSLogsOutputGraylogStreamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDBaaSLogsOutputGraylogStreamResult, error) {
			args := v.(LookupDBaaSLogsOutputGraylogStreamArgs)
			r, err := LookupDBaaSLogsOutputGraylogStream(ctx, &args, opts...)
			var s LookupDBaaSLogsOutputGraylogStreamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDBaaSLogsOutputGraylogStreamResultOutput)
}

// A collection of arguments for invoking getDBaaSLogsOutputGraylogStream.
type LookupDBaaSLogsOutputGraylogStreamOutputArgs struct {
	// The service name
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Stream description
	Title pulumi.StringInput `pulumi:"title"`
}

func (LookupDBaaSLogsOutputGraylogStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBaaSLogsOutputGraylogStreamArgs)(nil)).Elem()
}

// A collection of values returned by getDBaaSLogsOutputGraylogStream.
type LookupDBaaSLogsOutputGraylogStreamResultOutput struct{ *pulumi.OutputState }

func (LookupDBaaSLogsOutputGraylogStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBaaSLogsOutputGraylogStreamResult)(nil)).Elem()
}

func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) ToLookupDBaaSLogsOutputGraylogStreamResultOutput() LookupDBaaSLogsOutputGraylogStreamResultOutput {
	return o
}

func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) ToLookupDBaaSLogsOutputGraylogStreamResultOutputWithContext(ctx context.Context) LookupDBaaSLogsOutputGraylogStreamResultOutput {
	return o
}

func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) CanAlert() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) bool { return v.CanAlert }).(pulumi.BoolOutput)
}

// Cold storage compression method
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) ColdStorageCompression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.ColdStorageCompression }).(pulumi.StringOutput)
}

// ColdStorage content
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) ColdStorageContent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.ColdStorageContent }).(pulumi.StringOutput)
}

// Is Cold storage enabled?
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) ColdStorageEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) bool { return v.ColdStorageEnabled }).(pulumi.BoolOutput)
}

// Notify on new Cold storage archive
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) ColdStorageNotifyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) bool { return v.ColdStorageNotifyEnabled }).(pulumi.BoolOutput)
}

// Cold storage retention in year
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) ColdStorageRetention() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) int { return v.ColdStorageRetention }).(pulumi.IntOutput)
}

// ColdStorage destination
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) ColdStorageTarget() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.ColdStorageTarget }).(pulumi.StringOutput)
}

// Stream creation
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Stream description
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable ES indexing
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) IndexingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) bool { return v.IndexingEnabled }).(pulumi.BoolOutput)
}

// Maximum indexing size (in GB)
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) IndexingMaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) int { return v.IndexingMaxSize }).(pulumi.IntOutput)
}

// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) IndexingNotifyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) bool { return v.IndexingNotifyEnabled }).(pulumi.BoolOutput)
}

// Indicates if you are allowed to edit entry
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) IsEditable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) bool { return v.IsEditable }).(pulumi.BoolOutput)
}

// Indicates if you are allowed to share entry
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) IsShareable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) bool { return v.IsShareable }).(pulumi.BoolOutput)
}

// Number of alert condition
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) NbAlertCondition() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) int { return v.NbAlertCondition }).(pulumi.IntOutput)
}

// Number of coldstored archives
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) NbArchive() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) int { return v.NbArchive }).(pulumi.IntOutput)
}

// Parent stream ID
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) ParentStreamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.ParentStreamId }).(pulumi.StringOutput)
}

// If set, pause indexing when maximum size is reach
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) PauseIndexingOnMaxSize() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) bool { return v.PauseIndexingOnMaxSize }).(pulumi.BoolOutput)
}

// Retention ID
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) RetentionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.RetentionId }).(pulumi.StringOutput)
}

func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Stream ID
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.StreamId }).(pulumi.StringOutput)
}

func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.Title }).(pulumi.StringOutput)
}

// Stream last update
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Enable Websocket
func (o LookupDBaaSLogsOutputGraylogStreamResultOutput) WebSocketEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDBaaSLogsOutputGraylogStreamResult) string { return v.WebSocketEnabled }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDBaaSLogsOutputGraylogStreamResultOutput{})
}
