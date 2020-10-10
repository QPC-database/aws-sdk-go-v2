// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata about a self-managed object storage server location.
func (c *Client) DescribeLocationObjectStorage(ctx context.Context, params *DescribeLocationObjectStorageInput, optFns ...func(*Options)) (*DescribeLocationObjectStorageOutput, error) {
	if params == nil {
		params = &DescribeLocationObjectStorageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationObjectStorage", params, optFns, addOperationDescribeLocationObjectStorageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationObjectStorageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeLocationObjectStorageRequest
type DescribeLocationObjectStorageInput struct {

	// The Amazon Resource Name (ARN) of the self-managed object storage server
	// location that was described.
	//
	// This member is required.
	LocationArn *string
}

// DescribeLocationObjectStorageResponse
type DescribeLocationObjectStorageOutput struct {

	// Optional. The access key is used if credentials are required to access the
	// self-managed object storage server.
	AccessKey *string

	// The Amazon Resource Name (ARN) of the agents associated with the self-managed
	// object storage server location.
	AgentArns []*string

	// The time that the self-managed object storage server agent was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the self-managed object storage server
	// location to describe.
	LocationArn *string

	// The URL of the source self-managed object storage server location that was
	// described.
	LocationUri *string

	// The port that your self-managed object storage server accepts inbound network
	// traffic on. The server port is set by default to TCP 80 (HTTP) or TCP 443
	// (HTTPS).
	ServerPort *int32

	// The protocol that the object storage server uses to communicate. Valid values
	// are HTTP or HTTPS.
	ServerProtocol types.ObjectStorageServerProtocol

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLocationObjectStorageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationObjectStorage{}, middleware.After)
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
	addOpDescribeLocationObjectStorageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationObjectStorage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeLocationObjectStorage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeLocationObjectStorage",
	}
}
