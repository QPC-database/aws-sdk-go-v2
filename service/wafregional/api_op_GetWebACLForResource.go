// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic Regional documentation. For more information, see AWS
// WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Returns the web ACL for the specified resource, either an
// application load balancer or Amazon API Gateway stage.
func (c *Client) GetWebACLForResource(ctx context.Context, params *GetWebACLForResourceInput, optFns ...func(*Options)) (*GetWebACLForResourceOutput, error) {
	if params == nil {
		params = &GetWebACLForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWebACLForResource", params, optFns, addOperationGetWebACLForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWebACLForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWebACLForResourceInput struct {

	// The ARN (Amazon Resource Name) of the resource for which to get the web ACL,
	// either an application load balancer or Amazon API Gateway stage. The ARN should
	// be in one of the following formats:
	//
	//     * For an Application Load Balancer:
	// arn:aws:elasticloadbalancing:region:account-id:loadbalancer/app/load-balancer-name/load-balancer-id
	//
	//
	// * For an Amazon API Gateway stage:
	// arn:aws:apigateway:region::/restapis/api-id/stages/stage-name
	//
	// This member is required.
	ResourceArn *string
}

type GetWebACLForResourceOutput struct {

	// Information about the web ACL that you specified in the GetWebACLForResource
	// request. If there is no associated resource, a null WebACLSummary is returned.
	WebACLSummary *types.WebACLSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetWebACLForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetWebACLForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetWebACLForResource{}, middleware.After)
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
	addOpGetWebACLForResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetWebACLForResource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetWebACLForResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetWebACLForResource",
	}
}
