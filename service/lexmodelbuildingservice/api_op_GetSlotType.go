// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about a specific version of a slot type. In addition to
// specifying the slot type name, you must specify the slot type version. This
// operation requires permissions for the lex:GetSlotType action.
func (c *Client) GetSlotType(ctx context.Context, params *GetSlotTypeInput, optFns ...func(*Options)) (*GetSlotTypeOutput, error) {
	if params == nil {
		params = &GetSlotTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSlotType", params, optFns, addOperationGetSlotTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSlotTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSlotTypeInput struct {

	// The name of the slot type. The name is case sensitive.
	//
	// This member is required.
	Name *string

	// The version of the slot type.
	//
	// This member is required.
	Version *string
}

type GetSlotTypeOutput struct {

	// Checksum of the $LATEST version of the slot type.
	Checksum *string

	// The date that the slot type was created.
	CreatedDate *time.Time

	// A description of the slot type.
	Description *string

	// A list of EnumerationValue objects that defines the values that the slot type
	// can take.
	EnumerationValues []*types.EnumerationValue

	// The date that the slot type was updated. When you create a resource, the
	// creation date and last update date are the same.
	LastUpdatedDate *time.Time

	// The name of the slot type.
	Name *string

	// The built-in slot type used as a parent for the slot type.
	ParentSlotTypeSignature *string

	// Configuration information that extends the parent built-in slot type.
	SlotTypeConfigurations []*types.SlotTypeConfiguration

	// The strategy that Amazon Lex uses to determine the value of the slot. For more
	// information, see PutSlotType ().
	ValueSelectionStrategy types.SlotValueSelectionStrategy

	// The version of the slot type.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSlotTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSlotType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSlotType{}, middleware.After)
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
	addOpGetSlotTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSlotType(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSlotType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetSlotType",
	}
}
