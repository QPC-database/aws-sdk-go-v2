// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of DC2 ReservedNodeOfferings that matches the payment type,
// term, and usage price of the given DC1 reserved node.
func (c *Client) GetReservedNodeExchangeOfferings(ctx context.Context, params *GetReservedNodeExchangeOfferingsInput, optFns ...func(*Options)) (*GetReservedNodeExchangeOfferingsOutput, error) {
	if params == nil {
		params = &GetReservedNodeExchangeOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReservedNodeExchangeOfferings", params, optFns, addOperationGetReservedNodeExchangeOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReservedNodeExchangeOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type GetReservedNodeExchangeOfferingsInput struct {

	// A string representing the node identifier for the DC1 Reserved Node to be
	// exchanged.
	//
	// This member is required.
	ReservedNodeId *string

	// A value that indicates the starting point for the next set of
	// ReservedNodeOfferings.
	Marker *string

	// An integer setting the maximum number of ReservedNodeOfferings to retrieve.
	MaxRecords *int32
}

type GetReservedNodeExchangeOfferingsOutput struct {

	// An optional parameter that specifies the starting point for returning a set of
	// response records. When the results of a GetReservedNodeExchangeOfferings request
	// exceed the value specified in MaxRecords, Amazon Redshift returns a value in the
	// marker field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the marker parameter and retrying the
	// request.
	Marker *string

	// Returns an array of ReservedNodeOffering () objects.
	ReservedNodeOfferings []*types.ReservedNodeOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetReservedNodeExchangeOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetReservedNodeExchangeOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetReservedNodeExchangeOfferings{}, middleware.After)
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
	addOpGetReservedNodeExchangeOfferingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservedNodeExchangeOfferings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetReservedNodeExchangeOfferings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "GetReservedNodeExchangeOfferings",
	}
}
