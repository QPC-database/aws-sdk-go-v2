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

// Retrieves details for the specified Amazon Chime account, such as account type
// and supported licenses.
func (c *Client) GetAccount(ctx context.Context, params *GetAccountInput, optFns ...func(*Options)) (*GetAccountOutput, error) {
	if params == nil {
		params = &GetAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccount", params, optFns, addOperationGetAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccountInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string
}

type GetAccountOutput struct {

	// The Amazon Chime account details.
	Account *types.Account

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAccount{}, middleware.After)
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
	addOpGetAccountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetAccount",
	}
}
