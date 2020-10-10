// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the specified Lightsail load balancer.
func (c *Client) GetLoadBalancer(ctx context.Context, params *GetLoadBalancerInput, optFns ...func(*Options)) (*GetLoadBalancerOutput, error) {
	if params == nil {
		params = &GetLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLoadBalancer", params, optFns, addOperationGetLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoadBalancerInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string
}

type GetLoadBalancerOutput struct {

	// An object containing information about your load balancer.
	LoadBalancer *types.LoadBalancer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLoadBalancer{}, middleware.After)
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
	addOpGetLoadBalancerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLoadBalancer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetLoadBalancer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetLoadBalancer",
	}
}
