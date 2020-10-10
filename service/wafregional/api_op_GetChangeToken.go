// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. When you want to create, update, or delete AWS WAF objects, get a
// change token and include the change token in the create, update, or delete
// request. Change tokens ensure that your application doesn't submit conflicting
// requests to AWS WAF. Each create, update, or delete request must use a unique
// change token. If your application submits a GetChangeToken request and then
// submits a second GetChangeToken request before submitting a create, update, or
// delete request, the second GetChangeToken request returns the same value as the
// first GetChangeToken request. When you use a change token in a create, update,
// or delete request, the status of the change token changes to PENDING, which
// indicates that AWS WAF is propagating the change to all AWS WAF servers. Use
// GetChangeTokenStatus to determine the status of your change token.
func (c *Client) GetChangeToken(ctx context.Context, params *GetChangeTokenInput, optFns ...func(*Options)) (*GetChangeTokenOutput, error) {
	if params == nil {
		params = &GetChangeTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChangeToken", params, optFns, addOperationGetChangeTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChangeTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetChangeTokenInput struct {
}

type GetChangeTokenOutput struct {

	// The ChangeToken that you used in the request. Use this value in a
	// GetChangeTokenStatus request to get the current status of the request.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetChangeTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetChangeToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetChangeToken{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetChangeToken(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetChangeToken(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetChangeToken",
	}
}
