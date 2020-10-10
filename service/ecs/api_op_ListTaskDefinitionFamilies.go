// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of task definition families that are registered to your account
// (which may include task definition families that no longer have any ACTIVE task
// definition revisions). You can filter out task definition families that do not
// contain any ACTIVE task definition revisions by setting the status parameter to
// ACTIVE. You can also filter the results with the familyPrefix parameter.
func (c *Client) ListTaskDefinitionFamilies(ctx context.Context, params *ListTaskDefinitionFamiliesInput, optFns ...func(*Options)) (*ListTaskDefinitionFamiliesOutput, error) {
	if params == nil {
		params = &ListTaskDefinitionFamiliesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTaskDefinitionFamilies", params, optFns, addOperationListTaskDefinitionFamiliesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTaskDefinitionFamiliesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTaskDefinitionFamiliesInput struct {

	// The familyPrefix is a string that is used to filter the results of
	// ListTaskDefinitionFamilies. If you specify a familyPrefix, only task definition
	// family names that begin with the familyPrefix string are returned.
	FamilyPrefix *string

	// The maximum number of task definition family results returned by
	// ListTaskDefinitionFamilies in paginated output. When this parameter is used,
	// ListTaskDefinitions only returns maxResults results in a single page along with
	// a nextToken response element. The remaining results of the initial request can
	// be seen by sending another ListTaskDefinitionFamilies request with the returned
	// nextToken value. This value can be between 1 and 100. If this parameter is not
	// used, then ListTaskDefinitionFamilies returns up to 100 results and a nextToken
	// value if applicable.
	MaxResults *int32

	// The nextToken value returned from a ListTaskDefinitionFamilies request
	// indicating that more results are available to fulfill the request and further
	// calls will be needed. If maxResults was provided, it is possible the number of
	// results to be fewer than maxResults. This token should be treated as an opaque
	// identifier that is only used to retrieve the next items in a list and not for
	// other programmatic purposes.
	NextToken *string

	// The task definition family status with which to filter the
	// ListTaskDefinitionFamilies results. By default, both ACTIVE and INACTIVE task
	// definition families are listed. If this parameter is set to ACTIVE, only task
	// definition families that have an ACTIVE task definition revision are returned.
	// If this parameter is set to INACTIVE, only task definition families that do not
	// have any ACTIVE task definition revisions are returned. If you paginate the
	// resulting output, be sure to keep the status value constant in each subsequent
	// request.
	Status types.TaskDefinitionFamilyStatus
}

type ListTaskDefinitionFamiliesOutput struct {

	// The list of task definition family names that match the
	// ListTaskDefinitionFamilies request.
	Families []*string

	// The nextToken value to include in a future ListTaskDefinitionFamilies request.
	// When the results of a ListTaskDefinitionFamilies request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTaskDefinitionFamiliesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTaskDefinitionFamilies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTaskDefinitionFamilies{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTaskDefinitionFamilies(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTaskDefinitionFamilies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "ListTaskDefinitionFamilies",
	}
}
