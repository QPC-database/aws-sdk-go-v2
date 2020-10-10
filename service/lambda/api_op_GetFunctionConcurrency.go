// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns details about the reserved concurrency configuration for a function. To
// set a concurrency limit for a function, use PutFunctionConcurrency ().
func (c *Client) GetFunctionConcurrency(ctx context.Context, params *GetFunctionConcurrencyInput, optFns ...func(*Options)) (*GetFunctionConcurrencyOutput, error) {
	if params == nil {
		params = &GetFunctionConcurrencyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFunctionConcurrency", params, optFns, addOperationGetFunctionConcurrencyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFunctionConcurrencyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFunctionConcurrencyInput struct {

	// The name of the Lambda function. Name formats
	//
	//     * Function name -
	// my-function.
	//
	//     * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//     * Partial ARN -
	// 123456789012:function:my-function.
	//
	// The length constraint applies only to the
	// full ARN. If you specify only the function name, it is limited to 64 characters
	// in length.
	//
	// This member is required.
	FunctionName *string
}

type GetFunctionConcurrencyOutput struct {

	// The number of simultaneous executions that are reserved for the function.
	ReservedConcurrentExecutions *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetFunctionConcurrencyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFunctionConcurrency{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFunctionConcurrency{}, middleware.After)
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
	addOpGetFunctionConcurrencyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFunctionConcurrency(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetFunctionConcurrency(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetFunctionConcurrency",
	}
}
