// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the properties of all entity recognizers that you created,
// including recognizers currently in training. Allows you to filter the list of
// recognizers based on criteria such as status and submission time. This call
// returns up to 500 entity recognizers in the list, with a default number of 100
// recognizers in the list. The results of this list are not in any particular
// order. Please get the list and sort locally if needed.
func (c *Client) ListEntityRecognizers(ctx context.Context, params *ListEntityRecognizersInput, optFns ...func(*Options)) (*ListEntityRecognizersOutput, error) {
	if params == nil {
		params = &ListEntityRecognizersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntityRecognizers", params, optFns, addOperationListEntityRecognizersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntityRecognizersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntityRecognizersInput struct {

	// Filters the list of entities returned. You can filter on Status,
	// SubmitTimeBefore, or SubmitTimeAfter. You can only set one filter at a time.
	Filter *types.EntityRecognizerFilter

	// The maximum number of results to return on each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string
}

type ListEntityRecognizersOutput struct {

	// The list of properties of an entity recognizer.
	EntityRecognizerPropertiesList []*types.EntityRecognizerProperties

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEntityRecognizersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEntityRecognizers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEntityRecognizers{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEntityRecognizers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListEntityRecognizers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListEntityRecognizers",
	}
}
