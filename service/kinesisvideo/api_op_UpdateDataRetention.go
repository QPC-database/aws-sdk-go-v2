// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Increases or decreases the stream's data retention period by the value that you
// specify. To indicate whether you want to increase or decrease the data retention
// period, specify the Operation parameter in the request body. In the request, you
// must specify either the StreamName or the StreamARN. The retention period that
// you specify replaces the current value.  <p>This operation requires permission
// for the <code>KinesisVideo:UpdateDataRetention</code> action.</p> <p>Changing
// the data retention period affects the data in the stream as follows:</p> <ul>
// <li> <p>If the data retention period is increased, existing data is retained for
// the new retention period. For example, if the data retention period is increased
// from one hour to seven hours, all existing data is retained for seven hours.</p>
// </li> <li> <p>If the data retention period is decreased, existing data is
// retained for the new retention period. For example, if the data retention period
// is decreased from seven hours to one hour, all existing data is retained for one
// hour, and any data older than one hour is deleted immediately.</p> </li> </ul>
func (c *Client) UpdateDataRetention(ctx context.Context, params *UpdateDataRetentionInput, optFns ...func(*Options)) (*UpdateDataRetentionOutput, error) {
	if params == nil {
		params = &UpdateDataRetentionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataRetention", params, optFns, addOperationUpdateDataRetentionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataRetentionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataRetentionInput struct {

	// The version of the stream whose retention period you want to change. To get the
	// version, call either the DescribeStream or the ListStreams API.
	//
	// This member is required.
	CurrentVersion *string

	// The retention period, in hours. The value you specify replaces the current
	// value. The maximum value for this parameter is 87600 (ten years).
	//
	// This member is required.
	DataRetentionChangeInHours *int32

	// Indicates whether you want to increase or decrease the retention period.
	//
	// This member is required.
	Operation types.UpdateDataRetentionOperation

	// The Amazon Resource Name (ARN) of the stream whose retention period you want to
	// change.
	StreamARN *string

	// The name of the stream whose retention period you want to change.
	StreamName *string
}

type UpdateDataRetentionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDataRetentionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataRetention{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataRetention{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateDataRetentionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataRetention(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDataRetention(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "UpdateDataRetention",
	}
}
