// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns configuration information about the specified access point.  <p>All
// Amazon S3 on Outposts REST API requests for this action require an additional
// parameter of outpost-id to be passed with the request and an S3 on Outposts
// endpoint hostname prefix instead of s3-control. For an example of the request
// syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname
// prefix and the outpost-id derived using the access point ARN, see the <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_GetAccessPoint.html#API_control_GetAccessPoint_Examples">
// Example</a> section below.</p> <p>The following actions are related to
// <code>GetAccessPoint</code>:</p> <ul> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html">CreateAccessPoint</a>
// </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html">DeleteAccessPoint</a>
// </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html">ListAccessPoints</a>
// </p> </li> </ul>
func (c *Client) GetAccessPoint(ctx context.Context, params *GetAccessPointInput, optFns ...func(*Options)) (*GetAccessPointOutput, error) {
	if params == nil {
		params = &GetAccessPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessPoint", params, optFns, addOperationGetAccessPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessPointInput struct {

	// The account ID for the account that owns the specified access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point whose configuration information you want to
	// retrieve. For Amazon S3 on Outposts specify the ARN of the access point accessed
	// in the format arn:aws:s3-outposts:::outpost//accesspoint/. For example, to
	// access the access point reports-ap through outpost my-outpost owned by account
	// 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap.
	// The value must be URL encoded.
	//
	// This member is required.
	Name *string
}

type GetAccessPointOutput struct {

	// The name of the bucket associated with the specified access point.
	Bucket *string

	// The date and time when the specified access point was created.
	CreationDate *time.Time

	// The name of the specified access point.
	Name *string

	// Indicates whether this access point allows access from the public internet. If
	// VpcConfiguration is specified for this access point, then NetworkOrigin is VPC,
	// and the access point doesn't allow access from the public internet. Otherwise,
	// NetworkOrigin is Internet, and the access point allows access from the public
	// internet, subject to the access point and bucket access policies. This will
	// always be true for an Amazon S3 on Outposts access point
	NetworkOrigin types.NetworkOrigin

	// The PublicAccessBlock configuration that you want to apply to this Amazon S3
	// bucket. You can enable the configuration options in any combination. For more
	// information about when Amazon S3 considers a bucket or object public, see The
	// Meaning of "Public"
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status)
	// in the Amazon Simple Storage Service Developer Guide. This is not supported for
	// Amazon S3 on Outposts.
	PublicAccessBlockConfiguration *types.PublicAccessBlockConfiguration

	// Contains the virtual private cloud (VPC) configuration for the specified access
	// point.
	VpcConfiguration *types.VpcConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAccessPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessPoint{}, middleware.After)
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
	addOpGetAccessPointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessPoint(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAccessPoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetAccessPoint",
	}
}
