// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the definition of a trigger.
func (c *Client) GetTrigger(ctx context.Context, params *GetTriggerInput, optFns ...func(*Options)) (*GetTriggerOutput, error) {
	if params == nil {
		params = &GetTriggerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTrigger", params, optFns, addOperationGetTriggerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTriggerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTriggerInput struct {

	// The name of the trigger to retrieve.
	//
	// This member is required.
	Name *string
}

type GetTriggerOutput struct {

	// The requested trigger definition.
	Trigger *types.Trigger

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTriggerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTrigger{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTrigger{}, middleware.After)
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
	addOpGetTriggerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTrigger(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetTrigger(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetTrigger",
	}
}
