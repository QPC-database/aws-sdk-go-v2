// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the password for the specified IAM user, which terminates the user's
// ability to access AWS services through the AWS Management Console. Deleting a
// user's password does not prevent a user from accessing AWS through the command
// line interface or the API. To prevent all user access, you must also either make
// any access keys inactive or delete them. For more information about making keys
// inactive or deleting them, see UpdateAccessKey () and DeleteAccessKey ().
func (c *Client) DeleteLoginProfile(ctx context.Context, params *DeleteLoginProfileInput, optFns ...func(*Options)) (*DeleteLoginProfileOutput, error) {
	if params == nil {
		params = &DeleteLoginProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLoginProfile", params, optFns, addOperationDeleteLoginProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLoginProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLoginProfileInput struct {

	// The name of the user whose password you want to delete. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string
}

type DeleteLoginProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteLoginProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteLoginProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteLoginProfile{}, middleware.After)
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
	addOpDeleteLoginProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLoginProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteLoginProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteLoginProfile",
	}
}
