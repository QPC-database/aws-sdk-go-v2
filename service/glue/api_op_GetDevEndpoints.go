// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves all the development endpoints in this AWS account. When you create a
// development endpoint in a virtual private cloud (VPC), AWS Glue returns only a
// private IP address and the public IP address field is not populated. When you
// create a non-VPC development endpoint, AWS Glue returns only a public IP
// address.
func (c *Client) GetDevEndpoints(ctx context.Context, params *GetDevEndpointsInput, optFns ...func(*Options)) (*GetDevEndpointsOutput, error) {
	if params == nil {
		params = &GetDevEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDevEndpoints", params, optFns, addOperationGetDevEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDevEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDevEndpointsInput struct {

	// The maximum size of information to return.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string
}

type GetDevEndpointsOutput struct {

	// A list of DevEndpoint definitions.
	DevEndpoints []*types.DevEndpoint

	// A continuation token, if not all DevEndpoint definitions have yet been returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDevEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDevEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDevEndpoints{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDevEndpoints(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDevEndpoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetDevEndpoints",
	}
}
