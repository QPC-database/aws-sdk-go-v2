// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates phone number product types or calling names. You can update one
// attribute at a time for each UpdatePhoneNumberRequestItem. For example, you can
// update either the product type or the calling name. For product types, choose
// from Amazon Chime Business Calling and Amazon Chime Voice Connector. For
// toll-free numbers, you must use the Amazon Chime Voice Connector product type.
// Updates to outbound calling names can take up to 72 hours to complete. Pending
// updates to outbound calling names must be complete before you can request
// another update.
func (c *Client) BatchUpdatePhoneNumber(ctx context.Context, params *BatchUpdatePhoneNumberInput, optFns ...func(*Options)) (*BatchUpdatePhoneNumberOutput, error) {
	if params == nil {
		params = &BatchUpdatePhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdatePhoneNumber", params, optFns, addOperationBatchUpdatePhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdatePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdatePhoneNumberInput struct {

	// The request containing the phone number IDs and product types or calling names
	// to update.
	//
	// This member is required.
	UpdatePhoneNumberRequestItems []*types.UpdatePhoneNumberRequestItem
}

type BatchUpdatePhoneNumberOutput struct {

	// If the action fails for one or more of the phone numbers in the request, a list
	// of the phone numbers is returned, along with error codes and error messages.
	PhoneNumberErrors []*types.PhoneNumberError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchUpdatePhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdatePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdatePhoneNumber{}, middleware.After)
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
	addOpBatchUpdatePhoneNumberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdatePhoneNumber(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchUpdatePhoneNumber(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchUpdatePhoneNumber",
	}
}
