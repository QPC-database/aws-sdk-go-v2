// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes tags from an Amazon EMR resource. Tags make it easier to associate
// clusters in various ways, such as grouping clusters to track your Amazon EMR
// resource allocation costs. For more information, see Tag Clusters
// (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-tags.html). The
// following example removes the stack tag with value Prod from a cluster:
func (c *Client) RemoveTags(ctx context.Context, params *RemoveTagsInput, optFns ...func(*Options)) (*RemoveTagsOutput, error) {
	if params == nil {
		params = &RemoveTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveTags", params, optFns, addOperationRemoveTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input identifies a cluster and a list of tags to remove.
type RemoveTagsInput struct {

	// The Amazon EMR resource identifier from which tags will be removed. This value
	// must be a cluster identifier.
	//
	// This member is required.
	ResourceId *string

	// A list of tag keys to remove from a resource.
	//
	// This member is required.
	TagKeys []*string
}

// This output indicates the result of removing tags from a resource.
type RemoveTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveTags{}, middleware.After)
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
	addOpRemoveTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTags(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRemoveTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "RemoveTags",
	}
}
