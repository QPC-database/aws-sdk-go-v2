// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified attribute of the specified Amazon FPGA Image (AFI).
func (c *Client) DescribeFpgaImageAttribute(ctx context.Context, params *DescribeFpgaImageAttributeInput, optFns ...func(*Options)) (*DescribeFpgaImageAttributeOutput, error) {
	if params == nil {
		params = &DescribeFpgaImageAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFpgaImageAttribute", params, optFns, addOperationDescribeFpgaImageAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFpgaImageAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFpgaImageAttributeInput struct {

	// The AFI attribute.
	//
	// This member is required.
	Attribute types.FpgaImageAttributeName

	// The ID of the AFI.
	//
	// This member is required.
	FpgaImageId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DescribeFpgaImageAttributeOutput struct {

	// Information about the attribute.
	FpgaImageAttribute *types.FpgaImageAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFpgaImageAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeFpgaImageAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeFpgaImageAttribute{}, middleware.After)
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
	addOpDescribeFpgaImageAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFpgaImageAttribute(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeFpgaImageAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeFpgaImageAttribute",
	}
}
