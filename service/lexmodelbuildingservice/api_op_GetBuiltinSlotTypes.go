// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of built-in slot types that meet the specified criteria. For a list
// of built-in slot types, see Slot Type Reference
// (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/slot-type-reference)
// in the Alexa Skills Kit.  <p>This operation requires permission for the
// <code>lex:GetBuiltInSlotTypes</code> action.</p>
func (c *Client) GetBuiltinSlotTypes(ctx context.Context, params *GetBuiltinSlotTypesInput, optFns ...func(*Options)) (*GetBuiltinSlotTypesOutput, error) {
	if params == nil {
		params = &GetBuiltinSlotTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBuiltinSlotTypes", params, optFns, addOperationGetBuiltinSlotTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBuiltinSlotTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBuiltinSlotTypesInput struct {

	// A list of locales that the slot type supports.
	Locale types.Locale

	// The maximum number of slot types to return in the response. The default is 10.
	MaxResults *int32

	// A pagination token that fetches the next page of slot types. If the response to
	// this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of slot types, specify the pagination token in
	// the next request.
	NextToken *string

	// Substring to match in built-in slot type signatures. A slot type will be
	// returned if any part of its signature matches the substring. For example, "xyz"
	// matches both "xyzabc" and "abcxyz."
	SignatureContains *string
}

type GetBuiltinSlotTypesOutput struct {

	// If the response is truncated, the response includes a pagination token that you
	// can use in your next request to fetch the next page of slot types.
	NextToken *string

	// An array of BuiltInSlotTypeMetadata objects, one entry for each slot type
	// returned.
	SlotTypes []*types.BuiltinSlotTypeMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBuiltinSlotTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBuiltinSlotTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBuiltinSlotTypes{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBuiltinSlotTypes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBuiltinSlotTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBuiltinSlotTypes",
	}
}
