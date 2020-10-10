// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of tags associated with the specified stream. In the request, you
// must specify either the StreamName or the StreamARN.
func (c *Client) ListTagsForStream(ctx context.Context, params *ListTagsForStreamInput, optFns ...func(*Options)) (*ListTagsForStreamOutput, error) {
	if params == nil {
		params = &ListTagsForStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForStream", params, optFns, addOperationListTagsForStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsForStreamInput struct {

	// If you specify this parameter and the result of a ListTagsForStream call is
	// truncated, the response includes a token that you can use in the next request to
	// fetch the next batch of tags.
	NextToken *string

	// The Amazon Resource Name (ARN) of the stream that you want to list tags for.
	StreamARN *string

	// The name of the stream that you want to list tags for.
	StreamName *string
}

type ListTagsForStreamOutput struct {

	// If you specify this parameter and the result of a ListTags call is truncated,
	// the response includes a token that you can use in the next request to fetch the
	// next set of tags.
	NextToken *string

	// A map of tag keys and values associated with the specified stream.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsForStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTagsForStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTagsForStream{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForStream(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTagsForStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "ListTagsForStream",
	}
}
