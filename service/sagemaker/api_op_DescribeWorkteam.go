// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a specific work team. You can see information such as the
// create date, the last updated date, membership information, and the work team's
// Amazon Resource Name (ARN).
func (c *Client) DescribeWorkteam(ctx context.Context, params *DescribeWorkteamInput, optFns ...func(*Options)) (*DescribeWorkteamOutput, error) {
	if params == nil {
		params = &DescribeWorkteamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkteam", params, optFns, addOperationDescribeWorkteamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkteamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkteamInput struct {

	// The name of the work team to return a description of.
	//
	// This member is required.
	WorkteamName *string
}

type DescribeWorkteamOutput struct {

	// A Workteam instance that contains information about the work team.
	//
	// This member is required.
	Workteam *types.Workteam

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeWorkteamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkteam{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkteam{}, middleware.After)
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
	addOpDescribeWorkteamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkteam(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeWorkteam(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeWorkteam",
	}
}
