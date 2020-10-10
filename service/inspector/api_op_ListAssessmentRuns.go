// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the assessment runs that correspond to the assessment templates that are
// specified by the ARNs of the assessment templates.
func (c *Client) ListAssessmentRuns(ctx context.Context, params *ListAssessmentRunsInput, optFns ...func(*Options)) (*ListAssessmentRunsOutput, error) {
	if params == nil {
		params = &ListAssessmentRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssessmentRuns", params, optFns, addOperationListAssessmentRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssessmentRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssessmentRunsInput struct {

	// The ARNs that specify the assessment templates whose assessment runs you want to
	// list.
	AssessmentTemplateArns []*string

	// You can use this parameter to specify a subset of data to be included in the
	// action's response. For a record to match a filter, all specified filter
	// attributes must match. When multiple values are specified for a filter
	// attribute, any of the values can match.
	Filter *types.AssessmentRunFilter

	// You can use this parameter to indicate the maximum number of items that you want
	// in the response. The default value is 10. The maximum value is 500.
	MaxResults *int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListAssessmentRuns action.
	// Subsequent calls to the action fill nextToken in the request with the value of
	// NextToken from the previous response to continue listing data.
	NextToken *string
}

type ListAssessmentRunsOutput struct {

	// A list of ARNs that specifies the assessment runs that are returned by the
	// action.
	//
	// This member is required.
	AssessmentRunArns []*string

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to be
	// listed, this parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAssessmentRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAssessmentRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssessmentRuns{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAssessmentRuns(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListAssessmentRuns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "ListAssessmentRuns",
	}
}
