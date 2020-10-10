// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes metadata tags from an Amazon Neptune resource.
func (c *Client) RemoveTagsFromResource(ctx context.Context, params *RemoveTagsFromResourceInput, optFns ...func(*Options)) (*RemoveTagsFromResourceOutput, error) {
	if params == nil {
		params = &RemoveTagsFromResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveTagsFromResource", params, optFns, addOperationRemoveTagsFromResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveTagsFromResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveTagsFromResourceInput struct {

	// The Amazon Neptune resource that the tags are removed from. This value is an
	// Amazon Resource Name (ARN). For information about creating an ARN, see
	// Constructing an Amazon Resource Name (ARN)
	// (https://docs.aws.amazon.com/neptune/latest/UserGuide/tagging.ARN.html#tagging.ARN.Constructing).
	//
	// This member is required.
	ResourceName *string

	// The tag key (name) of the tag to be removed.
	//
	// This member is required.
	TagKeys []*string
}

type RemoveTagsFromResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveTagsFromResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRemoveTagsFromResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveTagsFromResource{}, middleware.After)
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
	addOpRemoveTagsFromResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTagsFromResource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRemoveTagsFromResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RemoveTagsFromResource",
	}
}
