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

// Lists the API keys for a given API. API keys are deleted automatically sometime
// after they expire. However, they may still be included in the response until
// they have actually been deleted. You can safely call DeleteApiKey to manually
// delete a key before it's automatically deleted.
func (c *Client) ListApiKeys(ctx context.Context, params *ListApiKeysInput, optFns ...func(*Options)) (*ListApiKeysOutput, error) {
	if params == nil {
		params = &ListApiKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApiKeys", params, optFns, addOperationListApiKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApiKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApiKeysInput struct {

	// The API ID.
	//
	// This member is required.
	ApiId *string

	// The maximum number of results you want the request to return.
	MaxResults *int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

type ListApiKeysOutput struct {

	// The ApiKey objects.
	ApiKeys []*types.ApiKey

	// An identifier to be passed in the next request to this operation to return the
	// next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListApiKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListApiKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListApiKeys{}, middleware.After)
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
	addOpListApiKeysValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListApiKeys(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListApiKeys(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "ListApiKeys",
	}
}
