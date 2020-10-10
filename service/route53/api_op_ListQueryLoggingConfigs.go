// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the configurations for DNS query logging that are associated with the
// current AWS account or the configuration that is associated with a specified
// hosted zone.  <p>For more information about DNS query logs, see <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html">CreateQueryLoggingConfig</a>.
// Additional information, including the format of DNS query logs, appears in <a
// href="https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html">Logging
// DNS Queries</a> in the <i>Amazon Route 53 Developer Guide</i>.</p>
func (c *Client) ListQueryLoggingConfigs(ctx context.Context, params *ListQueryLoggingConfigsInput, optFns ...func(*Options)) (*ListQueryLoggingConfigsOutput, error) {
	if params == nil {
		params = &ListQueryLoggingConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueryLoggingConfigs", params, optFns, addOperationListQueryLoggingConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueryLoggingConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueryLoggingConfigsInput struct {

	// (Optional) If you want to list the query logging configuration that is
	// associated with a hosted zone, specify the ID in HostedZoneId. If you don't
	// specify a hosted zone ID, ListQueryLoggingConfigs returns all of the
	// configurations that are associated with the current AWS account.
	HostedZoneId *string

	// (Optional) The maximum number of query logging configurations that you want
	// Amazon Route 53 to return in response to the current request. If the current AWS
	// account has more than MaxResults configurations, use the value of NextToken
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListQueryLoggingConfigs.html#API_ListQueryLoggingConfigs_RequestSyntax)
	// in the response to get the next page of results. If you don't specify a value
	// for MaxResults, Route 53 returns up to 100 configurations.
	MaxResults *string

	// (Optional) If the current AWS account has more than MaxResults query logging
	// configurations, use NextToken to get the second and subsequent pages of results.
	// For the first ListQueryLoggingConfigs request, omit this value. For the second
	// and subsequent requests, get the value of NextToken from the previous response
	// and specify that value for NextToken in the request.
	NextToken *string
}

type ListQueryLoggingConfigsOutput struct {

	// An array that contains one QueryLoggingConfig
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_QueryLoggingConfig.html)
	// element for each configuration for DNS query logging that is associated with the
	// current AWS account.
	//
	// This member is required.
	QueryLoggingConfigs []*types.QueryLoggingConfig

	// If a response includes the last of the query logging configurations that are
	// associated with the current AWS account, NextToken doesn't appear in the
	// response. If a response doesn't include the last of the configurations, you can
	// get more configurations by submitting another ListQueryLoggingConfigs
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListQueryLoggingConfigs.html)
	// request. Get the value of NextToken that Amazon Route 53 returned in the
	// previous response and include it in NextToken in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListQueryLoggingConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListQueryLoggingConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListQueryLoggingConfigs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListQueryLoggingConfigs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addSanitizeURLMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListQueryLoggingConfigs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListQueryLoggingConfigs",
	}
}
