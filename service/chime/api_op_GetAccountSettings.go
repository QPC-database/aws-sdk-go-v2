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

// Retrieves account settings for the specified Amazon Chime account ID, such as
// remote control and dial out settings. For more information about these settings,
// see Use the Policies Page
// (https://docs.aws.amazon.com/chime/latest/ag/policies.html) in the Amazon Chime
// Administration Guide.
func (c *Client) GetAccountSettings(ctx context.Context, params *GetAccountSettingsInput, optFns ...func(*Options)) (*GetAccountSettingsOutput, error) {
	if params == nil {
		params = &GetAccountSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccountSettings", params, optFns, addOperationGetAccountSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccountSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccountSettingsInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string
}

type GetAccountSettingsOutput struct {

	// The Amazon Chime account settings.
	AccountSettings *types.AccountSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAccountSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAccountSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAccountSettings{}, middleware.After)
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
	addOpGetAccountSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccountSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAccountSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetAccountSettings",
	}
}
