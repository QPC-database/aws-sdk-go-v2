// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a unique key that you can distribute to clients who are executing your
// API.
func (c *Client) CreateApiKey(ctx context.Context, params *CreateApiKeyInput, optFns ...func(*Options)) (*CreateApiKeyOutput, error) {
	if params == nil {
		params = &CreateApiKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApiKey", params, optFns, addOperationCreateApiKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApiKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApiKeyInput struct {

	// The ID for your GraphQL API.
	//
	// This member is required.
	ApiId *string

	// A description of the purpose of the API key.
	Description *string

	// The time from creation time after which the API key expires. The date is
	// represented as seconds since the epoch, rounded down to the nearest hour. The
	// default value for this parameter is 7 days from creation time. For more
	// information, see .
	Expires *int64
}

type CreateApiKeyOutput struct {

	// The API key.
	ApiKey *types.ApiKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateApiKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApiKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApiKey{}, middleware.After)
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
	addOpCreateApiKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApiKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateApiKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "CreateApiKey",
	}
}
