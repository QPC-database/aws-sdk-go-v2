// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of tags that are applied to the specified AWS OpsWorks for Chef
// Automate or AWS OpsWorks for Puppet Enterprise servers or backups.
func (c *Client) ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) {
	if params == nil {
		params = &ListTagsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForResource", params, optFns, addOperationListTagsForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsForResourceInput struct {

	// The Amazon Resource Number (ARN) of an AWS OpsWorks for Chef Automate or AWS
	// OpsWorks for Puppet Enterprise server for which you want to show applied tags.
	// For example,
	// arn:aws:opsworks-cm:us-west-2:123456789012:server/test-owcm-server/EXAMPLE-66b0-4196-8274-d1a2bEXAMPLE.
	//
	// This member is required.
	ResourceArn *string

	// To receive a paginated response, use this parameter to specify the maximum
	// number of results to be returned with a single call. If the number of available
	// results exceeds this maximum, the response includes a NextToken value that you
	// can assign to the NextToken request parameter to get the next set of results.
	MaxResults *int32

	// NextToken is a string that is returned in some command responses. It indicates
	// that not all entries have been returned, and that you must run at least one more
	// request to get remaining items. To get remaining results, call
	// ListTagsForResource again, and assign the token from the previous results as the
	// value of the nextToken parameter. If there are no more results, the response
	// object's nextToken parameter value is null. Setting a nextToken value that was
	// not returned in your previous results causes an InvalidNextTokenException to
	// occur.
	NextToken *string
}

type ListTagsForResourceOutput struct {

	// A token that you can use as the value of NextToken in subsequent calls to the
	// API to show more results.
	NextToken *string

	// Tags that have been applied to the resource.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTagsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTagsForResource{}, middleware.After)
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
	addOpListTagsForResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForResource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTagsForResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "ListTagsForResource",
	}
}
