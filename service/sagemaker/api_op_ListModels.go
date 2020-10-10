// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists models created with the CreateModel () API.
func (c *Client) ListModels(ctx context.Context, params *ListModelsInput, optFns ...func(*Options)) (*ListModelsOutput, error) {
	if params == nil {
		params = &ListModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModels", params, optFns, addOperationListModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelsInput struct {

	// A filter that returns only models with a creation time greater than or equal to
	// the specified time (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only models created before the specified time (timestamp).
	CreationTimeBefore *time.Time

	// The maximum number of models to return in the response.
	MaxResults *int32

	// A string in the training job name. This filter returns only models in the
	// training job whose name contains the specified string.
	NameContains *string

	// If the response to a previous ListModels request was truncated, the response
	// includes a NextToken. To retrieve the next set of models, use the token in the
	// next request.
	NextToken *string

	// Sorts the list of results. The default is CreationTime.
	SortBy types.ModelSortKey

	// The sort order for results. The default is Descending.
	SortOrder types.OrderKey
}

type ListModelsOutput struct {

	// An array of ModelSummary objects, each of which lists a model.
	//
	// This member is required.
	Models []*types.ModelSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of models, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListModels{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListModels(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListModels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListModels",
	}
}
