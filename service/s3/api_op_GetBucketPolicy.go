// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the policy of a specified bucket. If you are using an identity other
// than the root user of the AWS account that owns the bucket, the calling identity
// must have the GetBucketPolicy permissions on the specified bucket and belong to
// the bucket owner's account in order to use this operation.  <p>If you don't have
// <code>GetBucketPolicy</code> permissions, Amazon S3 returns a <code>403 Access
// Denied</code> error. If you have the correct permissions, but you're not using
// an identity that belongs to the bucket owner's account, Amazon S3 returns a
// <code>405 Method Not Allowed</code> error.</p> <important> <p>As a security
// precaution, the root user of the AWS account that owns a bucket can always use
// this operation, even if the policy explicitly denies the root user the ability
// to perform this action.</p> </important> <p>For more information about bucket
// policies, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html">Using
// Bucket Policies and User Policies</a>.</p> <p>The following operation is related
// to <code>GetBucketPolicy</code>:</p> <ul> <li> <p> <a>GetObject</a> </p> </li>
// </ul>
func (c *Client) GetBucketPolicy(ctx context.Context, params *GetBucketPolicyInput, optFns ...func(*Options)) (*GetBucketPolicyOutput, error) {
	if params == nil {
		params = &GetBucketPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketPolicy", params, optFns, addOperationGetBucketPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketPolicyInput struct {

	// The bucket name for which to get the bucket policy.
	//
	// This member is required.
	Bucket *string
}

type GetBucketPolicyOutput struct {

	// The bucket policy as a JSON document.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketPolicy{}, middleware.After)
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
	addOpGetBucketPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketPolicy(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBucketPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketPolicy",
	}
}
