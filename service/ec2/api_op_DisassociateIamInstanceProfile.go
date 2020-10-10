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

// Disassociates an IAM instance profile from a running or stopped instance. Use
// DescribeIamInstanceProfileAssociations () to get the association ID.
func (c *Client) DisassociateIamInstanceProfile(ctx context.Context, params *DisassociateIamInstanceProfileInput, optFns ...func(*Options)) (*DisassociateIamInstanceProfileOutput, error) {
	if params == nil {
		params = &DisassociateIamInstanceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateIamInstanceProfile", params, optFns, addOperationDisassociateIamInstanceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateIamInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateIamInstanceProfileInput struct {

	// The ID of the IAM instance profile association.
	//
	// This member is required.
	AssociationId *string
}

type DisassociateIamInstanceProfileOutput struct {

	// Information about the IAM instance profile association.
	IamInstanceProfileAssociation *types.IamInstanceProfileAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateIamInstanceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisassociateIamInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateIamInstanceProfile{}, middleware.After)
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
	addOpDisassociateIamInstanceProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateIamInstanceProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateIamInstanceProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisassociateIamInstanceProfile",
	}
}
