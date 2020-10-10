// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the budgets that are associated with an account. The Request Syntax
// section shows the BudgetLimit syntax. For PlannedBudgetLimits, see the Examples
// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_DescribeBudgets.html#API_DescribeBudgets_Examples)
// section.
func (c *Client) DescribeBudgets(ctx context.Context, params *DescribeBudgetsInput, optFns ...func(*Options)) (*DescribeBudgetsOutput, error) {
	if params == nil {
		params = &DescribeBudgetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBudgets", params, optFns, addOperationDescribeBudgetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBudgetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of DescribeBudgets
type DescribeBudgetsInput struct {

	// The accountId that is associated with the budgets that you want descriptions of.
	//
	// This member is required.
	AccountId *string

	// An optional integer that represents how many entries a paginated response
	// contains. The maximum is 100.
	MaxResults *int32

	// The pagination token that you include in your request to indicate the next set
	// of results that you want to retrieve.
	NextToken *string
}

// Response of DescribeBudgets
type DescribeBudgetsOutput struct {

	// A list of budgets.
	Budgets []*types.Budget

	// The pagination token in the service response that indicates the next set of
	// results that you can retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBudgetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBudgets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBudgets{}, middleware.After)
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
	addOpDescribeBudgetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBudgets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBudgets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DescribeBudgets",
	}
}
