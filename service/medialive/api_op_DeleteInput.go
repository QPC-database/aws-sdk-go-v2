// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the input end point
func (c *Client) DeleteInput(ctx context.Context, params *DeleteInputInput, optFns ...func(*Options)) (*DeleteInputOutput, error) {
	if params == nil {
		params = &DeleteInputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteInput", params, optFns, addOperationDeleteInputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteInputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DeleteInputRequest
type DeleteInputInput struct {

	// Unique ID of the input
	//
	// This member is required.
	InputId *string
}

// Placeholder documentation for DeleteInputResponse
type DeleteInputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteInputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteInput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteInput{}, middleware.After)
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
	addOpDeleteInputValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInput(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteInput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DeleteInput",
	}
}
