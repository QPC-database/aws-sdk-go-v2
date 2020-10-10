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

// Lists the analytics configurations for the bucket. You can have up to 1,000
// analytics configurations per bucket.  <p>This operation supports list pagination
// and does not return more than 100 configurations at a time. You should always
// check the <code>IsTruncated</code> element in the response. If there are no more
// configurations to list, <code>IsTruncated</code> is set to false. If there are
// more configurations to list, <code>IsTruncated</code> is set to true, and there
// will be a value in <code>NextContinuationToken</code>. You use the
// <code>NextContinuationToken</code> value to continue the pagination of the list
// by passing the value in continuation-token in the request to <code>GET</code>
// the next page.</p> <p>To use this operation, you must have permissions to
// perform the <code>s3:GetAnalyticsConfiguration</code> action. The bucket owner
// has this permission by default. The bucket owner can grant this permission to
// others. For more information about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources">Permissions
// Related to Bucket Subresource Operations</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p>For information about
// Amazon S3 analytics feature, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html">Amazon
// S3 Analytics – Storage Class Analysis</a>. </p> <p>The following operations are
// related to <code>ListBucketAnalyticsConfigurations</code>:</p> <ul> <li> <p>
// <a>GetBucketAnalyticsConfiguration</a> </p> </li> <li> <p>
// <a>DeleteBucketAnalyticsConfiguration</a> </p> </li> <li> <p>
// <a>PutBucketAnalyticsConfiguration</a> </p> </li> </ul>
func (c *Client) ListBucketAnalyticsConfigurations(ctx context.Context, params *ListBucketAnalyticsConfigurationsInput, optFns ...func(*Options)) (*ListBucketAnalyticsConfigurationsOutput, error) {
	if params == nil {
		params = &ListBucketAnalyticsConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBucketAnalyticsConfigurations", params, optFns, addOperationListBucketAnalyticsConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBucketAnalyticsConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBucketAnalyticsConfigurationsInput struct {

	// The name of the bucket from which analytics configurations are retrieved.
	//
	// This member is required.
	Bucket *string

	// The ContinuationToken that represents a placeholder from where this request
	// should begin.
	ContinuationToken *string
}

type ListBucketAnalyticsConfigurationsOutput struct {

	// The list of analytics configurations for a bucket.
	AnalyticsConfigurationList []*types.AnalyticsConfiguration

	// The marker that is used as a starting point for this analytics configuration
	// list response. This value is present if it was sent in the request.
	ContinuationToken *string

	// Indicates whether the returned list of analytics configurations is complete. A
	// value of true indicates that the list is not complete and the
	// NextContinuationToken will be provided for a subsequent request.
	IsTruncated *bool

	// NextContinuationToken is sent when isTruncated is true, which indicates that
	// there are more analytics configurations to list. The next request must include
	// this NextContinuationToken. The token is obfuscated and is not a usable value.
	NextContinuationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBucketAnalyticsConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListBucketAnalyticsConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListBucketAnalyticsConfigurations{}, middleware.After)
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
	addOpListBucketAnalyticsConfigurationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBucketAnalyticsConfigurations(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opListBucketAnalyticsConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListBucketAnalyticsConfigurations",
	}
}
