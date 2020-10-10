// Code generated by smithy-go-codegen DO NOT EDIT.

package savingsplans

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified Savings Plans offerings.
func (c *Client) DescribeSavingsPlansOfferings(ctx context.Context, params *DescribeSavingsPlansOfferingsInput, optFns ...func(*Options)) (*DescribeSavingsPlansOfferingsOutput, error) {
	if params == nil {
		params = &DescribeSavingsPlansOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSavingsPlansOfferings", params, optFns, addOperationDescribeSavingsPlansOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSavingsPlansOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSavingsPlansOfferingsInput struct {

	// The currencies.
	Currencies []types.CurrencyCode

	// The descriptions.
	Descriptions []*string

	// The durations, in seconds.
	Durations []*int64

	// The filters.
	Filters []*types.SavingsPlanOfferingFilterElement

	// The maximum number of results to return with a single call. To retrieve
	// additional results, make another call with the returned token value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the offerings.
	OfferingIds []*string

	// The specific AWS operation for the line item in the billing report.
	Operations []*string

	// The payment options.
	PaymentOptions []types.SavingsPlanPaymentOption

	// The plan type.
	PlanTypes []types.SavingsPlanType

	// The product type.
	ProductType types.SavingsPlanProductType

	// The services.
	ServiceCodes []*string

	// The usage details of the line item in the billing report.
	UsageTypes []*string
}

type DescribeSavingsPlansOfferingsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the Savings Plans offerings.
	SearchResults []*types.SavingsPlanOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSavingsPlansOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSavingsPlansOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSavingsPlansOfferings{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSavingsPlansOfferings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeSavingsPlansOfferings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "savingsplans",
		OperationName: "DescribeSavingsPlansOfferings",
	}
}
