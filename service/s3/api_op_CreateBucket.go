// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new bucket. To create a bucket, you must register with Amazon S3 and
// have a valid AWS Access Key ID to authenticate requests. Anonymous requests are
// never allowed to create buckets. By creating the bucket, you become the bucket
// owner. Not every string is an acceptable bucket name. For information on bucket
// naming restrictions, see Working with Amazon S3 Buckets
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html). By default,
// the bucket is created in the US East (N. Virginia) Region. You can optionally
// specify a Region in the request body. You might choose a Region to optimize
// latency, minimize costs, or address regulatory requirements. For example, if you
// reside in Europe, you will probably find it advantageous to create buckets in
// the Europe (Ireland) Region. For more information, see How to Select a Region
// for Your Buckets
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro).
// If you send your create bucket request to the s3.amazonaws.com endpoint, the
// request goes to the us-east-1 Region. Accordingly, the signature calculations in
// Signature Version 4 must use us-east-1 as the Region, even if the location
// constraint in the request specifies another Region where the bucket is to be
// created. If you create a bucket in a Region other than US East (N. Virginia),
// your application must be able to handle 307 redirect. For more information, see
// Virtual Hosting of Buckets
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html). When
// creating a bucket using this operation, you can optionally specify the accounts
// or groups that should be granted specific permissions on the bucket. There are
// two ways to grant the appropriate permissions using the request headers.
//
//     *
// Specify a canned ACL using the x-amz-acl request header. Amazon S3 supports a
// set of predefined ACLs, known as canned ACLs. Each canned ACL has a predefined
// set of grantees and permissions. For more information, see Canned ACL
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
//
//
// * Specify access permissions explicitly using the x-amz-grant-read,
// x-amz-grant-write, x-amz-grant-read-acp, x-amz-grant-write-acp, and
// x-amz-grant-full-control headers. These headers map to the set of permissions
// Amazon S3 supports in an ACL. For more information, see Access Control List
// (ACL) Overview
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html). You specify
// each grantee as a type=value pair, where the type is one of the following:
//
//
// * id – if the value specified is the canonical user ID of an AWS account
//
//
// * uri – if you are granting permissions to a predefined group
//
//         *
// emailAddress – if the value specified is the email address of an AWS account
// Using email addresses to specify a grantee is only supported in the following
// AWS Regions:
//
//             * US East (N. Virginia)
//
//             * US West (N.
// California)
//
//             * US West (Oregon)
//
//             * Asia Pacific
// (Singapore)
//
//             * Asia Pacific (Sydney)
//
//             * Asia Pacific
// (Tokyo)
//
//             * Europe (Ireland)
//
//             * South America (São
// Paulo)
//
//         For a list of all the Amazon S3 supported Regions and endpoints,
// see Regions and Endpoints
// (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) in the AWS
// General Reference.
//
//     For example, the following x-amz-grant-read header
// grants the AWS accounts identified by account IDs permissions to read object
// data and its metadata: x-amz-grant-read: id="11112222333", id="444455556666"
// </li> </ul> <note> <p>You can use either a canned ACL or specify access
// permissions explicitly. You cannot do both.</p> </note> <p>The following
// operations are related to <code>CreateBucket</code>:</p> <ul> <li> <p>
// <a>PutObject</a> </p> </li> <li> <p> <a>DeleteBucket</a> </p> </li> </ul>
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

	// The name of the bucket to create.
	//
	// This member is required.
	Bucket *string

	// The canned ACL to apply to the bucket.
	ACL types.BucketCannedACL

	// The configuration information for the bucket.
	CreateBucketConfiguration *types.CreateBucketConfiguration

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket.
	GrantRead *string

	// Allows grantee to read the bucket ACL.
	GrantReadACP *string

	// Allows grantee to create, overwrite, and delete any object in the bucket.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string

	// Specifies whether you want S3 Object Lock to be enabled for the new bucket.
	ObjectLockEnabledForBucket *bool
}

type CreateBucketOutput struct {

	// Specifies the Region where the bucket will be created. If you are creating a
	// bucket on the US East (N. Virginia) Region (us-east-1), you do not need to
	// specify the location.
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
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
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
