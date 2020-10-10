// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the PublicAccessBlock configuration for an AWS account. For more
// information, see  Using Amazon S3 block public access
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).
// Related actions include:
//
//     * GetPublicAccessBlock
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetPublicAccessBlock.html)
//
//
// * PutPublicAccessBlock
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutPublicAccessBlock.html)
func (c *Client) DeletePublicAccessBlock(ctx context.Context, params *DeletePublicAccessBlockInput, optFns ...func(*Options)) (*DeletePublicAccessBlockOutput, error) {
	if params == nil {
		params = &DeletePublicAccessBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePublicAccessBlock", params, optFns, addOperationDeletePublicAccessBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePublicAccessBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePublicAccessBlockInput struct {

	// The account ID for the AWS account whose PublicAccessBlock configuration you
	// want to remove.
	//
	// This member is required.
	AccountId *string
}

type DeletePublicAccessBlockOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeletePublicAccessBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeletePublicAccessBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeletePublicAccessBlock{}, middleware.After)
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
	addOpDeletePublicAccessBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePublicAccessBlock(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeletePublicAccessBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeletePublicAccessBlock",
	}
}
