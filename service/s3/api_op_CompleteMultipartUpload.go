// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Completes a multipart upload by assembling previously uploaded parts. You first
// initiate the multipart upload and then upload all parts using the UploadPart ()
// operation. After successfully uploading all relevant parts of an upload, you
// call this operation to complete the upload. Upon receiving this request, Amazon
// S3 concatenates all the parts in ascending order by part number to create a new
// object. In the Complete Multipart Upload request, you must provide the parts
// list. You must ensure that the parts list is complete. This operation
// concatenates the parts that you provide in the list. For each part in the list,
// you must provide the part number and the ETag value, returned after that part
// was uploaded. Processing of a Complete Multipart Upload request could take
// several minutes to complete. After Amazon S3 begins processing the request, it
// sends an HTTP response header that specifies a 200 OK response. While processing
// is in progress, Amazon S3 periodically sends white space characters to keep the
// connection from timing out. Because a request could fail after the initial 200
// OK response has been sent, it is important that you check the response body to
// determine whether the request succeeded. Note that if CompleteMultipartUpload
// fails, applications should be prepared to retry the failed requests. For more
// information, see Amazon S3 Error Best Practices
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ErrorBestPractices.html). For
// more information about multipart uploads, see Uploading Objects Using Multipart
// Upload (https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html).
// For information about permissions required to use the multipart upload API, see
// Multipart Upload API and Permissions
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html).  <p>
// <code>GetBucketLifecycle</code> has the following special errors:</p> <ul> <li>
// <p>Error code: <code>EntityTooSmall</code> </p> <ul> <li> <p>Description: Your
// proposed upload is smaller than the minimum allowed object size. Each part must
// be at least 5 MB in size, except the last part.</p> </li> <li> <p>400 Bad
// Request</p> </li> </ul> </li> <li> <p>Error code: <code>InvalidPart</code> </p>
// <ul> <li> <p>Description: One or more of the specified parts could not be found.
// The part might not have been uploaded, or the specified entity tag might not
// have matched the part's entity tag.</p> </li> <li> <p>400 Bad Request</p> </li>
// </ul> </li> <li> <p>Error code: <code>InvalidPartOrder</code> </p> <ul> <li>
// <p>Description: The list of parts was not in ascending order. The parts list
// must be specified in order by part number.</p> </li> <li> <p>400 Bad Request</p>
// </li> </ul> </li> <li> <p>Error code: <code>NoSuchUpload</code> </p> <ul> <li>
// <p>Description: The specified multipart upload does not exist. The upload ID
// might be invalid, or the multipart upload might have been aborted or
// completed.</p> </li> <li> <p>404 Not Found</p> </li> </ul> </li> </ul> <p>The
// following operations are related to <code>CompleteMultipartUpload</code>:</p>
// <ul> <li> <p> <a>CreateMultipartUpload</a> </p> </li> <li> <p> <a>UploadPart</a>
// </p> </li> <li> <p> <a>AbortMultipartUpload</a> </p> </li> <li> <p>
// <a>ListParts</a> </p> </li> <li> <p> <a>ListMultipartUploads</a> </p> </li>
// </ul>
func (c *Client) CompleteMultipartUpload(ctx context.Context, params *CompleteMultipartUploadInput, optFns ...func(*Options)) (*CompleteMultipartUploadOutput, error) {
	if params == nil {
		params = &CompleteMultipartUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompleteMultipartUpload", params, optFns, addOperationCompleteMultipartUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompleteMultipartUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompleteMultipartUploadInput struct {

	// Name of the bucket to which the multipart upload was initiated.
	//
	// This member is required.
	Bucket *string

	// Object key for which the multipart upload was initiated.
	//
	// This member is required.
	Key *string

	// ID for the initiated multipart upload.
	//
	// This member is required.
	UploadId *string

	// The container for the multipart upload request information.
	MultipartUpload *types.CompletedMultipartUpload

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer
}

type CompleteMultipartUploadOutput struct {

	// The name of the bucket that contains the newly created object.
	Bucket *string

	// Entity tag that identifies the newly created object's data. Objects with
	// different object data will have different entity tags. The entity tag is an
	// opaque string. The entity tag may or may not be an MD5 digest of the object
	// data. If the entity tag is not an MD5 digest of the object data, it will contain
	// one or more nonhexadecimal characters and/or will consist of less than 32 or
	// more than 32 hexadecimal digits.
	ETag *string

	// If the object expiration is configured, this will contain the expiration date
	// (expiry-date) and rule ID (rule-id). The value of rule-id is URL encoded.
	Expiration *string

	// The object key of the newly created object.
	Key *string

	// The URI that identifies the newly created object.
	Location *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string

	// If you specified server-side encryption either with an Amazon S3-managed
	// encryption key or an AWS KMS customer master key (CMK) in your initiate
	// multipart upload request, the response includes this header. It confirms the
	// encryption algorithm that Amazon S3 used to encrypt the object.
	ServerSideEncryption types.ServerSideEncryption

	// Version ID of the newly created object, in case the bucket has versioning turned
	// on.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCompleteMultipartUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCompleteMultipartUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCompleteMultipartUpload{}, middleware.After)
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
	addOpCompleteMultipartUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteMultipartUpload(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	s3cust.HandleResponseErrorWith200Status(stack)
	return nil
}

func newServiceMetadataMiddleware_opCompleteMultipartUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CompleteMultipartUpload",
	}
}
