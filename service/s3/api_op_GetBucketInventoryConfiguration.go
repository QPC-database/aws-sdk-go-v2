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

// Returns an inventory configuration (identified by the inventory configuration
// ID) from the bucket.  <p>To use this operation, you must have permissions to
// perform the <code>s3:GetInventoryConfiguration</code> action. The bucket owner
// has this permission by default and can grant this permission to others. For more
// information about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources">Permissions
// Related to Bucket Subresource Operations</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p>For information about
// the Amazon S3 inventory feature, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html">Amazon
// S3 Inventory</a>.</p> <p>The following operations are related to
// <code>GetBucketInventoryConfiguration</code>:</p> <ul> <li> <p>
// <a>DeleteBucketInventoryConfiguration</a> </p> </li> <li> <p>
// <a>ListBucketInventoryConfigurations</a> </p> </li> <li> <p>
// <a>PutBucketInventoryConfiguration</a> </p> </li> </ul>
func (c *Client) GetBucketInventoryConfiguration(ctx context.Context, params *GetBucketInventoryConfigurationInput, optFns ...func(*Options)) (*GetBucketInventoryConfigurationOutput, error) {
	if params == nil {
		params = &GetBucketInventoryConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketInventoryConfiguration", params, optFns, addOperationGetBucketInventoryConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketInventoryConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketInventoryConfigurationInput struct {

	// The name of the bucket containing the inventory configuration to retrieve.
	//
	// This member is required.
	Bucket *string

	// The ID used to identify the inventory configuration.
	//
	// This member is required.
	Id *string
}

type GetBucketInventoryConfigurationOutput struct {

	// Specifies the inventory configuration.
	InventoryConfiguration *types.InventoryConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketInventoryConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketInventoryConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketInventoryConfiguration{}, middleware.After)
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
	addOpGetBucketInventoryConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketInventoryConfiguration(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBucketInventoryConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketInventoryConfiguration",
	}
}
