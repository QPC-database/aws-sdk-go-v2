// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes an existing virtual node.
func (c *Client) DescribeVirtualNode(ctx context.Context, params *DescribeVirtualNodeInput, optFns ...func(*Options)) (*DescribeVirtualNodeOutput, error) {
	if params == nil {
		params = &DescribeVirtualNodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVirtualNode", params, optFns, addOperationDescribeVirtualNodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVirtualNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeVirtualNodeInput struct {

	// The name of the service mesh that the virtual node resides in.
	//
	// This member is required.
	MeshName *string

	// The name of the virtual node to describe.
	//
	// This member is required.
	VirtualNodeName *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

//
type DescribeVirtualNodeOutput struct {

	// The full description of your virtual node.
	//
	// This member is required.
	VirtualNode *types.VirtualNodeData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVirtualNodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeVirtualNode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeVirtualNode{}, middleware.After)
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
	addOpDescribeVirtualNodeValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
