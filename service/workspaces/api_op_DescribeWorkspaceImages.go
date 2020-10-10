// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list that describes one or more specified images, if the image
// identifiers are provided. Otherwise, all images in the account are described.
func (c *Client) DescribeWorkspaceImages(ctx context.Context, params *DescribeWorkspaceImagesInput, optFns ...func(*Options)) (*DescribeWorkspaceImagesOutput, error) {
	if params == nil {
		params = &DescribeWorkspaceImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkspaceImages", params, optFns, addOperationDescribeWorkspaceImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkspaceImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceImagesInput struct {

	// The identifier of the image.
	ImageIds []*string

	// The type (owned or shared) of the image.
	ImageType types.ImageType

	// The maximum number of items to return.
	MaxResults *int32

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string
}

type DescribeWorkspaceImagesOutput struct {

	// Information about the images.
	Images []*types.WorkspaceImage

	// The token to use to retrieve the next set of results, or null if no more results
	// are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeWorkspaceImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaceImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaceImages{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceImages(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeWorkspaceImages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaceImages",
	}
}
