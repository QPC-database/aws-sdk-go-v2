// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a log stream for the specified log group. There is no limit on the
// number of log streams that you can create for a log group. There is a limit of
// 50 TPS on CreateLogStream operations, after which transactions are throttled.
// You must use the following guidelines when naming a log stream:
//
//     * Log
// stream names must be unique within the log group.
//
//     * Log stream names can be
// between 1 and 512 characters long.
//
//     * The ':' (colon) and '*' (asterisk)
// characters are not allowed.
func (c *Client) CreateLogStream(ctx context.Context, params *CreateLogStreamInput, optFns ...func(*Options)) (*CreateLogStreamOutput, error) {
	if params == nil {
		params = &CreateLogStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLogStream", params, optFns, addOperationCreateLogStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLogStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLogStreamInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The name of the log stream.
	//
	// This member is required.
	LogStreamName *string
}

type CreateLogStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLogStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLogStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLogStream{}, middleware.After)
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
	addOpCreateLogStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLogStream(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateLogStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "CreateLogStream",
	}
}
