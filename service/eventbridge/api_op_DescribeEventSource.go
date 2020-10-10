// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This operation lists details about a partner event source that is shared with
// your account.
func (c *Client) DescribeEventSource(ctx context.Context, params *DescribeEventSourceInput, optFns ...func(*Options)) (*DescribeEventSourceOutput, error) {
	if params == nil {
		params = &DescribeEventSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventSource", params, optFns, addOperationDescribeEventSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventSourceInput struct {

	// The name of the partner event source to display the details of.
	//
	// This member is required.
	Name *string
}

type DescribeEventSourceOutput struct {

	// The ARN of the partner event source.
	Arn *string

	// The name of the SaaS partner that created the event source.
	CreatedBy *string

	// The date and time that the event source was created.
	CreationTime *time.Time

	// The date and time that the event source will expire if you do not create a
	// matching event bus.
	ExpirationTime *time.Time

	// The name of the partner event source.
	Name *string

	// The state of the event source. If it is ACTIVE, you have already created a
	// matching event bus for this event source, and that event bus is active. If it is
	// PENDING, either you haven't yet created a matching event bus, or that event bus
	// is deactivated. If it is DELETED, you have created a matching event bus, but the
	// event source has since been deleted.
	State types.EventSourceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventSource{}, middleware.After)
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
	addOpDescribeEventSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventSource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEventSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DescribeEventSource",
	}
}
