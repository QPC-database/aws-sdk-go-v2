// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the specified TagOption.
func (c *Client) DescribeTagOption(ctx context.Context, params *DescribeTagOptionInput, optFns ...func(*Options)) (*DescribeTagOptionOutput, error) {
	if params == nil {
		params = &DescribeTagOptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTagOption", params, optFns, addOperationDescribeTagOptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTagOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTagOptionInput struct {

	// The TagOption identifier.
	//
	// This member is required.
	Id *string
}

type DescribeTagOptionOutput struct {

	// Information about the TagOption.
	TagOptionDetail *types.TagOptionDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTagOptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTagOption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTagOption{}, middleware.After)
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
	addOpDescribeTagOptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTagOption(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTagOption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribeTagOption",
	}
}
