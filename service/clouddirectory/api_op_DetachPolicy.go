// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches a policy from an object.
func (c *Client) DetachPolicy(ctx context.Context, params *DetachPolicyInput, optFns ...func(*Options)) (*DetachPolicyOutput, error) {
	if params == nil {
		params = &DetachPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachPolicy", params, optFns, addOperationDetachPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachPolicyInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory () where
	// both objects reside. For more information, see arns ().
	//
	// This member is required.
	DirectoryArn *string

	// Reference that identifies the object whose policy object will be detached.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// Reference that identifies the policy object.
	//
	// This member is required.
	PolicyReference *types.ObjectReference
}

type DetachPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetachPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDetachPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDetachPolicy{}, middleware.After)
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
	addOpDetachPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDetachPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "DetachPolicy",
	}
}
