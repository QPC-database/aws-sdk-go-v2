// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the access key pair associated with the specified IAM user. If you do
// not specify a user name, IAM determines the user name implicitly based on the
// AWS access key ID signing the request. This operation works for access keys
// under the AWS account. Consequently, you can use this operation to manage AWS
// account root user credentials even if the AWS account has no associated users.
func (c *Client) DeleteAccessKey(ctx context.Context, params *DeleteAccessKeyInput, optFns ...func(*Options)) (*DeleteAccessKeyOutput, error) {
	if params == nil {
		params = &DeleteAccessKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccessKey", params, optFns, addOperationDeleteAccessKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccessKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccessKeyInput struct {

	// The access key ID for the access key ID and secret access key you want to
	// delete. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters that can consist of
	// any upper or lowercased letter or digit.
	//
	// This member is required.
	AccessKeyId *string

	// The name of the user whose access key pair you want to delete. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	UserName *string
}

type DeleteAccessKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAccessKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteAccessKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteAccessKey{}, middleware.After)
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
	addOpDeleteAccessKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccessKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAccessKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteAccessKey",
	}
}
