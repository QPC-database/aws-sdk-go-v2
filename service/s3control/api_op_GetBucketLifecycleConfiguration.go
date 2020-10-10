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

// This API operation gets an Amazon S3 on Outposts bucket's lifecycle
// configuration. To get an S3 bucket's lifecycle configuration, see
// GetBucketLifecycleConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html)
// in the Amazon Simple Storage Service API. Returns the lifecycle configuration
// information set on the Outposts bucket. For more information, see Using Amazon
// S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) and for
// information about lifecycle configuration, see  Object Lifecycle Management
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in
// Amazon Simple Storage Service Developer Guide.  <p>To use this operation, you
// must have permission to perform the
// <code>s3outposts:GetLifecycleConfiguration</code> action. The Outposts bucket
// owner has this permission, by default. The bucket owner can grant this
// permission to others. For more information about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources">Permissions
// Related to Bucket Subresource Operations</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p>All Amazon S3 on
// Outposts REST API requests for this action require an additional parameter of
// outpost-id to be passed with the request and an S3 on Outposts endpoint hostname
// prefix instead of s3-control. For an example of the request syntax for Amazon S3
// on Outposts that uses the S3 on Outposts endpoint hostname prefix and the
// outpost-id derived using the access point ARN, see the <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_GetBucketLifecycleConfiguration.html#API_control_GetBucketLifecycleConfiguration_Examples">
// Example</a> section below.</p> <p> <code>GetBucketLifecycleConfiguration</code>
// has the following special error:</p> <ul> <li> <p>Error code:
// <code>NoSuchLifecycleConfiguration</code> </p> <ul> <li> <p>Description: The
// lifecycle configuration does not exist.</p> </li> <li> <p>HTTP Status Code: 404
// Not Found</p> </li> <li> <p>SOAP Fault Code Prefix: Client</p> </li> </ul> </li>
// </ul> <p>The following actions are related to
// <code>GetBucketLifecycleConfiguration</code>:</p> <ul> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html">PutBucketLifecycleConfiguration</a>
// </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html">DeleteBucketLifecycleConfiguration</a>
// </p> </li> </ul>
func (c *Client) GetBucketLifecycleConfiguration(ctx context.Context, params *GetBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*GetBucketLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &GetBucketLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketLifecycleConfiguration", params, optFns, addOperationGetBucketLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketLifecycleConfigurationInput struct {

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The Amazon Resource Name (ARN) of the bucket. For Amazon S3 on Outposts specify
	// the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/. For example, to access the bucket
	// reports through outpost my-outpost owned by account 123456789012 in Region
	// us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string
}

type GetBucketLifecycleConfigurationOutput struct {

	// Container for the lifecycle rule of the Outposts bucket.
	Rules []*types.LifecycleRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketLifecycleConfiguration{}, middleware.After)
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
	addOpGetBucketLifecycleConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketLifecycleConfiguration(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBucketLifecycleConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketLifecycleConfiguration",
	}
}
