// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the namespaces for the specified AWS account.
func (c *Client) ListNamespaces(ctx context.Context, params *ListNamespacesInput, optFns ...func(*Options)) (*ListNamespacesOutput, error) {
	if params == nil {
		params = &ListNamespacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNamespaces", params, optFns, addOperationListNamespacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNamespacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNamespacesInput struct {

	// The ID for the AWS account that contains the QuickSight namespaces that you want
	// to list.
	//
	// This member is required.
	AwsAccountId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string
}

type ListNamespacesOutput struct {

	// The information about the namespaces in this AWS account. The response includes
	// the namespace ARN, name, AWS Region, notification email address, creation
	// status, and identity store.
	Namespaces []*types.NamespaceInfoV2

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListNamespacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNamespaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNamespaces{}, middleware.After)
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
	addOpListNamespacesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListNamespaces(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListNamespaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListNamespaces",
	}
}
