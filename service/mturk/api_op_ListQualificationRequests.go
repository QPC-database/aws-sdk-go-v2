// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListQualificationRequests operation retrieves requests for Qualifications of
// a particular Qualification type. The owner of the Qualification type calls this
// operation to poll for pending requests, and accepts them using the
// AcceptQualification operation.
func (c *Client) ListQualificationRequests(ctx context.Context, params *ListQualificationRequestsInput, optFns ...func(*Options)) (*ListQualificationRequestsOutput, error) {
	if params == nil {
		params = &ListQualificationRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQualificationRequests", params, optFns, addOperationListQualificationRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQualificationRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQualificationRequestsInput struct {

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The ID of the QualificationType.
	QualificationTypeId *string
}

type ListQualificationRequestsOutput struct {

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of Qualification requests on this page in the filtered results list,
	// equivalent to the number of Qualification requests being returned by this call.
	NumResults *int32

	// The Qualification request. The response includes one QualificationRequest
	// element for each Qualification request returned by the query.
	QualificationRequests []*types.QualificationRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListQualificationRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListQualificationRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListQualificationRequests{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListQualificationRequests(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListQualificationRequests(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListQualificationRequests",
	}
}
