// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves your request parameters, Savings Plan Recommendations Summary and
// Details.
func (c *Client) GetSavingsPlansPurchaseRecommendation(ctx context.Context, params *GetSavingsPlansPurchaseRecommendationInput, optFns ...func(*Options)) (*GetSavingsPlansPurchaseRecommendationOutput, error) {
	if params == nil {
		params = &GetSavingsPlansPurchaseRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSavingsPlansPurchaseRecommendation", params, optFns, addOperationGetSavingsPlansPurchaseRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSavingsPlansPurchaseRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSavingsPlansPurchaseRecommendationInput struct {

	// The lookback period used to generate the recommendation.
	//
	// This member is required.
	LookbackPeriodInDays types.LookbackPeriodInDays

	// The payment option used to generate these recommendations.
	//
	// This member is required.
	PaymentOption types.PaymentOption

	// The Savings Plans recommendation type requested.
	//
	// This member is required.
	SavingsPlansType types.SupportedSavingsPlansType

	// The savings plan recommendation term used to generate these recommendations.
	//
	// This member is required.
	TermInYears types.TermInYears

	// The account scope that you want your recommendations for. Amazon Web Services
	// calculates recommendations including the payer account and linked accounts if
	// the value is set to PAYER. If the value is LINKED, recommendations are
	// calculated for individual linked accounts only.
	AccountScope types.AccountScope

	// You can filter your recommendations by Account ID with the LINKED_ACCOUNT
	// dimension. To filter your recommendations by Account ID, specify Key as
	// LINKED_ACCOUNT and Value as the comma-separated Acount ID(s) for which you want
	// to see Savings Plans purchase recommendations. For
	// GetSavingsPlansPurchaseRecommendation, the Filter does not include
	// CostCategories or Tags. It only includes Dimensions. With Dimensions, Key must
	// be LINKED_ACCOUNT and Value can be a single Account ID or multiple
	// comma-separated Account IDs for which you want to see Savings Plans Purchase
	// Recommendations. AND and OR operators are not supported.
	Filter *types.Expression

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize *int32
}

type GetSavingsPlansPurchaseRecommendationOutput struct {

	// Information regarding this specific recommendation set.
	Metadata *types.SavingsPlansPurchaseRecommendationMetadata

	// The token for the next set of retrievable results. AWS provides the token when
	// the response from a previous call has more results than the maximum page size.
	NextPageToken *string

	// Contains your request parameters, Savings Plan Recommendations Summary, and
	// Details.
	SavingsPlansPurchaseRecommendation *types.SavingsPlansPurchaseRecommendation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSavingsPlansPurchaseRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSavingsPlansPurchaseRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSavingsPlansPurchaseRecommendation{}, middleware.After)
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
	addOpGetSavingsPlansPurchaseRecommendationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSavingsPlansPurchaseRecommendation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSavingsPlansPurchaseRecommendation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetSavingsPlansPurchaseRecommendation",
	}
}
