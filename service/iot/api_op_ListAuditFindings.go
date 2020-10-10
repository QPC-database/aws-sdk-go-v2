// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the findings (results) of a Device Defender audit or of the audits
// performed during a specified time period. (Findings are retained for 180 days.)
func (c *Client) ListAuditFindings(ctx context.Context, params *ListAuditFindingsInput, optFns ...func(*Options)) (*ListAuditFindingsOutput, error) {
	if params == nil {
		params = &ListAuditFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAuditFindings", params, optFns, addOperationListAuditFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAuditFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAuditFindingsInput struct {

	// A filter to limit results to the findings for the specified audit check.
	CheckName *string

	// A filter to limit results to those found before the specified time. You must
	// specify either the startTime and endTime or the taskId, but not both.
	EndTime *time.Time

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// Information identifying the noncompliant resource.
	ResourceIdentifier *types.ResourceIdentifier

	// A filter to limit results to those found after the specified time. You must
	// specify either the startTime and endTime or the taskId, but not both.
	StartTime *time.Time

	// A filter to limit results to the audit with the specified ID. You must specify
	// either the taskId or the startTime and endTime, but not both.
	TaskId *string
}

type ListAuditFindingsOutput struct {

	// The findings (results) of the audit.
	Findings []*types.AuditFinding

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAuditFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAuditFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAuditFindings{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAuditFindings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListAuditFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListAuditFindings",
	}
}
