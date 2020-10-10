// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates the definition for a dimension. You cannot change the type of a
// dimension after it is created (you can delete it and re-create it).
func (c *Client) UpdateDimension(ctx context.Context, params *UpdateDimensionInput, optFns ...func(*Options)) (*UpdateDimensionOutput, error) {
	if params == nil {
		params = &UpdateDimensionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDimension", params, optFns, addOperationUpdateDimensionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDimensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDimensionInput struct {

	// A unique identifier for the dimension. Choose something that describes the type
	// and value to make it easy to remember what it does.
	//
	// This member is required.
	Name *string

	// Specifies the value or list of values for the dimension. For TOPIC_FILTER
	// dimensions, this is a pattern used to match the MQTT topic (for example,
	// "admin/#").
	//
	// This member is required.
	StringValues []*string
}

type UpdateDimensionOutput struct {

	// The ARN (Amazon resource name) of the created dimension.
	Arn *string

	// The date and time, in milliseconds since epoch, when the dimension was initially
	// created.
	CreationDate *time.Time

	// The date and time, in milliseconds since epoch, when the dimension was most
	// recently updated.
	LastModifiedDate *time.Time

	// A unique identifier for the dimension.
	Name *string

	// The value or list of values used to scope the dimension. For example, for topic
	// filters, this is the pattern used to match the MQTT topic name.
	StringValues []*string

	// The type of the dimension.
	Type types.DimensionType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDimensionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDimension{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDimension{}, middleware.After)
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
	addOpUpdateDimensionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDimension(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDimension(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateDimension",
	}
}
