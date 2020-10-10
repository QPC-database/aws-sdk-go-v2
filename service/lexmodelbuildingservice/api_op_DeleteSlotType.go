// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes all versions of the slot type, including the $LATEST version. To delete
// a specific version of the slot type, use the DeleteSlotTypeVersion () operation.
// You can delete a version of a slot type only if it is not referenced. To delete
// a slot type that is referred to in one or more intents, you must remove those
// references first. If you get the ResourceInUseException exception, the exception
// provides an example reference that shows the intent where the slot type is
// referenced. To remove the reference to the slot type, either update the intent
// or delete it. If you get the same exception when you attempt to delete the slot
// type again, repeat until the slot type has no references and the DeleteSlotType
// call is successful. This operation requires permission for the
// lex:DeleteSlotType action.
func (c *Client) DeleteSlotType(ctx context.Context, params *DeleteSlotTypeInput, optFns ...func(*Options)) (*DeleteSlotTypeOutput, error) {
	if params == nil {
		params = &DeleteSlotTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSlotType", params, optFns, addOperationDeleteSlotTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSlotTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSlotTypeInput struct {

	// The name of the slot type. The name is case sensitive.
	//
	// This member is required.
	Name *string
}

type DeleteSlotTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteSlotTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSlotType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSlotType{}, middleware.After)
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
	addOpDeleteSlotTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSlotType(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteSlotType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteSlotType",
	}
}
