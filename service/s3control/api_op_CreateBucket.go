// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This API operation creates an Amazon S3 on Outposts bucket. To create an S3
// bucket, see Create Bucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html) in the
// Amazon Simple Storage Service API. Creates a new Outposts bucket. By creating
// the bucket, you become the bucket owner. To create an Outposts bucket, you must
// have S3 on Outposts. For more information, see Using Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in Amazon
// Simple Storage Service Developer Guide. Not every string is an acceptable bucket
// name. For information on bucket naming restrictions, see Working with Amazon S3
// Buckets
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html#bucketnamingrules).
// <p>S3 on Outposts buckets do not support </p> <ul> <li> <p>ACLs. Instead,
// configure access point policies to manage access to buckets.</p> </li> <li>
// <p>Public access. </p> </li> <li> <p>Object Lock</p> </li> <li> <p>Bucket
// Location constraint</p> </li> </ul> <p>For an example of the request syntax for
// Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and
// outpost-id in your API request, see the <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_CreateBucket.html#API_control_CreateBucket_Examples">
// Example</a> section below.</p> <p>The following actions are related to
// <code>CreateBucket</code> for Amazon S3 on Outposts:</p> <ul> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html">PutObject</a>
// </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_GetBucket.html">GetBucket</a>
// </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_DeleteBucket.html">DeleteBucket</a>
// </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_CreateAccessPoint.html">CreateAccessPoint</a>
// </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_PutAccessPointPolicy.html">PutAccessPointPolicy</a>
// </p> </li> </ul>
func (c *Client) CreateBucket(ctx context.Context, params *CreateBucketInput, optFns ...func(*Options)) (*CreateBucketOutput, error) {
	if params == nil {
		params = &CreateBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBucket", params, optFns, addOperationCreateBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBucketInput struct {

	// The name of the bucket.
	//
	// This member is required.
	Bucket *string

	// The canned ACL to apply to the bucket. This is not supported by Amazon S3 on
	// Outposts buckets.
	ACL types.BucketCannedACL

	// The configuration information for the bucket. This is not supported by Amazon S3
	// on Outposts buckets.
	CreateBucketConfiguration *types.CreateBucketConfiguration

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket. This is not supported by Amazon S3 on Outposts buckets.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket. This is not supported by
	// Amazon S3 on Outposts buckets.
	GrantRead *string

	// Allows grantee to read the bucket ACL. This is not supported by Amazon S3 on
	// Outposts buckets.
	GrantReadACP *string

	// Allows grantee to create, overwrite, and delete any object in the bucket. This
	// is not supported by Amazon S3 on Outposts buckets.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket. This is not supported
	// by Amazon S3 on Outposts buckets.
	GrantWriteACP *string

	// Specifies whether you want S3 Object Lock to be enabled for the new bucket. This
	// is not supported by Amazon S3 on Outposts buckets.
	ObjectLockEnabledForBucket *bool

	// The ID of the Outposts where the bucket is being created. This is required by
	// Amazon S3 on Outposts buckets.
	OutpostId *string
}

type CreateBucketOutput struct {

	// The Amazon Resource Name (ARN) of the bucket. For Amazon S3 on Outposts specify
	// the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/. For example, to access the bucket
	// reports through outpost my-outpost owned by account 123456789012 in Region
	// us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	BucketArn *string

	// The location of the bucket.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateBucket{}, middleware.After)
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
	addOpCreateBucketValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBucket(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateBucket(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CreateBucket",
	}
}
