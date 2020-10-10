// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes pending authorization requests for a specified aggregator account in a
// specified region.
func (c *Client) DeletePendingAggregationRequest(ctx context.Context, params *DeletePendingAggregationRequestInput, optFns ...func(*Options)) (*DeletePendingAggregationRequestOutput, error) {
	if params == nil {
		params = &DeletePendingAggregationRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePendingAggregationRequest", params, optFns, addOperationDeletePendingAggregationRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePendingAggregationRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePendingAggregationRequestInput struct {

	// The 12-digit account ID of the account requesting to aggregate data.
	//
	// This member is required.
	RequesterAccountId *string

	// The region requesting to aggregate data.
	//
	// This member is required.
	RequesterAwsRegion *string
}

type DeletePendingAggregationRequestOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeletePendingAggregationRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePendingAggregationRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePendingAggregationRequest{}, middleware.After)
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
	addOpDeletePendingAggregationRequestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePendingAggregationRequest(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeletePendingAggregationRequest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeletePendingAggregationRequest",
	}
}
