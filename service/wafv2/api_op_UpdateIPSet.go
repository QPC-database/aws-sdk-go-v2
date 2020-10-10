// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Updates the specified IPSet ().
func (c *Client) UpdateIPSet(ctx context.Context, params *UpdateIPSetInput, optFns ...func(*Options)) (*UpdateIPSetOutput, error) {
	if params == nil {
		params = &UpdateIPSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIPSet", params, optFns, addOperationUpdateIPSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIPSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIPSetInput struct {

	// Contains an array of strings that specify one or more IP addresses or blocks of
	// IP addresses in Classless Inter-Domain Routing (CIDR) notation. AWS WAF supports
	// all address ranges for IP versions IPv4 and IPv6. Examples:
	//
	//     * To configure
	// AWS WAF to allow, block, or count requests that originated from the IP address
	// 192.0.2.44, specify 192.0.2.44/32.
	//
	//     * To configure AWS WAF to allow, block,
	// or count requests that originated from IP addresses from 192.0.2.0 to
	// 192.0.2.255, specify 192.0.2.0/24.
	//
	//     * To configure AWS WAF to allow, block,
	// or count requests that originated from the IP address
	// 1111:0000:0000:0000:0000:0000:0000:0111, specify
	// 1111:0000:0000:0000:0000:0000:0000:0111/128.
	//
	//     * To configure AWS WAF to
	// allow, block, or count requests that originated from IP addresses
	// 1111:0000:0000:0000:0000:0000:0000:0000 to
	// 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify
	// 1111:0000:0000:0000:0000:0000:0000:0000/64.
	//
	// For more information about CIDR
	// notation, see the Wikipedia entry Classless Inter-Domain Routing
	// (https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).
	//
	// This member is required.
	Addresses []*string

	// A unique identifier for the set. This ID is returned in the responses to create
	// and list commands. You provide it to operations like update and delete.
	//
	// This member is required.
	Id *string

	// A token used for optimistic locking. AWS WAF returns a token to your get and
	// list requests, to mark the state of the entity at the time of the request. To
	// make changes to the entity associated with the token, you provide the token to
	// operations like update and delete. AWS WAF uses the token to ensure that no
	// changes have been made to the entity since you last retrieved it. If a change
	// has been made, the update fails with a WAFOptimisticLockException. If this
	// happens, perform another get, and use the new token returned by that operation.
	//
	// This member is required.
	LockToken *string

	// The name of the IP set. You cannot change the name of an IPSet after you create
	// it.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB) or
	// an API Gateway stage. To work with CloudFront, you must also specify the Region
	// US East (N. Virginia) as follows:
	//
	//     * CLI - Specify the Region when you use
	// the CloudFront scope: --scope=CLOUDFRONT --region=us-east-1.
	//
	//     * API and SDKs
	// - For all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// A description of the IP set that helps with identification. You cannot change
	// the description of an IP set after you create it.
	Description *string
}

type UpdateIPSetOutput struct {

	// A token used for optimistic locking. AWS WAF returns this token to your update
	// requests. You use NextLockToken in the same manner as you use LockToken.
	NextLockToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateIPSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateIPSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateIPSet{}, middleware.After)
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
	addOpUpdateIPSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIPSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateIPSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "UpdateIPSet",
	}
}
