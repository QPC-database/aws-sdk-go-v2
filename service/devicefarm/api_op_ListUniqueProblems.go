// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about unique problems, such as exceptions or crashes. Unique
// problems are defined as a single instance of an error across a run, job, or
// suite. For example, if a call in your application consistently raises an
// exception (OutOfBoundsException in MyActivity.java:386), ListUniqueProblems
// returns a single entry instead of many individual entries for that exception.
func (c *Client) ListUniqueProblems(ctx context.Context, params *ListUniqueProblemsInput, optFns ...func(*Options)) (*ListUniqueProblemsOutput, error) {
	if params == nil {
		params = &ListUniqueProblemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUniqueProblems", params, optFns, addOperationListUniqueProblemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUniqueProblemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the list unique problems operation.
type ListUniqueProblemsInput struct {

	// The unique problems' ARNs.
	//
	// This member is required.
	Arn *string

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Represents the result of a list unique problems request.
type ListUniqueProblemsOutput struct {

	// If the number of items that are returned is significantly large, this is an
	// identifier that is also returned. It can be used in a subsequent call to this
	// operation to return the next set of items in the list.
	NextToken *string

	// Information about the unique problems. Allowed values include:
	//
	//     * PENDING
	//
	//
	// * PASSED
	//
	//     * WARNED
	//
	//     * FAILED
	//
	//     * SKIPPED
	//
	//     * ERRORED
	//
	//     *
	// STOPPED
	UniqueProblems map[string][]*types.UniqueProblem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUniqueProblemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUniqueProblems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUniqueProblems{}, middleware.After)
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
	addOpListUniqueProblemsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListUniqueProblems(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListUniqueProblems(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListUniqueProblems",
	}
}
