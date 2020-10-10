// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the log pattern sets in the specific application.
func (c *Client) ListLogPatternSets(ctx context.Context, params *ListLogPatternSetsInput, optFns ...func(*Options)) (*ListLogPatternSetsOutput, error) {
	if params == nil {
		params = &ListLogPatternSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLogPatternSets", params, optFns, addOperationListLogPatternSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLogPatternSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLogPatternSetsInput struct {

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string
}

type ListLogPatternSetsOutput struct {

	// The list of log pattern sets.
	LogPatternSets []*string

	// The token used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The name of the resource group.
	ResourceGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListLogPatternSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLogPatternSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLogPatternSets{}, middleware.After)
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
	addOpListLogPatternSetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListLogPatternSets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListLogPatternSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "ListLogPatternSets",
	}
}
