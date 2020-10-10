// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns descriptive information about an update against your Amazon EKS cluster
// or associated managed node group. When the status of the update is Succeeded,
// the update is complete. If an update fails, the status is Failed, and an error
// detail explains the reason for the failure.
func (c *Client) DescribeUpdate(ctx context.Context, params *DescribeUpdateInput, optFns ...func(*Options)) (*DescribeUpdateOutput, error) {
	if params == nil {
		params = &DescribeUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUpdate", params, optFns, addOperationDescribeUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUpdateInput struct {

	// The name of the Amazon EKS cluster associated with the update.
	//
	// This member is required.
	Name *string

	// The ID of the update to describe.
	//
	// This member is required.
	UpdateId *string

	// The name of the Amazon EKS node group associated with the update.
	NodegroupName *string
}

type DescribeUpdateOutput struct {

	// The full description of the specified update.
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeUpdate{}, middleware.After)
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
	addOpDescribeUpdateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUpdate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeUpdate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "DescribeUpdate",
	}
}
