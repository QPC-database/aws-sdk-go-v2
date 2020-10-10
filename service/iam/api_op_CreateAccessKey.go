// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new AWS secret access key and corresponding AWS access key ID for the
// specified user. The default status for new keys is Active. If you do not specify
// a user name, IAM determines the user name implicitly based on the AWS access key
// ID signing the request. This operation works for access keys under the AWS
// account. Consequently, you can use this operation to manage AWS account root
// user credentials. This is true even if the AWS account has no associated users.
// The number and size of IAM resources in an AWS account are limited. For more
// information, see IAM and STS Quotas
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in
// the IAM User Guide. To ensure the security of your AWS account, the secret
// access key is accessible only during key and user creation. You must save the
// key (for example, in a text file) if you want to be able to access it again. If
// a secret key is lost, you can delete the access keys for the associated user and
// then create new keys.
func (c *Client) CreateAccessKey(ctx context.Context, params *CreateAccessKeyInput, optFns ...func(*Options)) (*CreateAccessKeyOutput, error) {
	if params == nil {
		params = &CreateAccessKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessKey", params, optFns, addOperationCreateAccessKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessKeyInput struct {

	// The name of the IAM user that the new key will belong to. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	UserName *string
}

// Contains the response to a successful CreateAccessKey () request.
type CreateAccessKeyOutput struct {

	// A structure with details about the access key.
	//
	// This member is required.
	AccessKey *types.AccessKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAccessKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateAccessKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateAccessKey{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateAccessKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "CreateAccessKey",
	}
}
