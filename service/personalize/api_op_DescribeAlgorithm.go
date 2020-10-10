// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the given algorithm.
func (c *Client) DescribeAlgorithm(ctx context.Context, params *DescribeAlgorithmInput, optFns ...func(*Options)) (*DescribeAlgorithmOutput, error) {
	if params == nil {
		params = &DescribeAlgorithmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAlgorithm", params, optFns, addOperationDescribeAlgorithmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAlgorithmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAlgorithmInput struct {

	// The Amazon Resource Name (ARN) of the algorithm to describe.
	//
	// This member is required.
	AlgorithmArn *string
}

type DescribeAlgorithmOutput struct {

	// A listing of the properties of the algorithm.
	Algorithm *types.Algorithm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAlgorithmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAlgorithm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAlgorithm{}, middleware.After)
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
	addOpDescribeAlgorithmValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAlgorithm(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAlgorithm(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DescribeAlgorithm",
	}
}
