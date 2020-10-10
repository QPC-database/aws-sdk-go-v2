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

// Updates an existing Cost Category. Changes made to the Cost Category rules will
// be used to categorize the current month’s expenses and future expenses. This
// won’t change categorization for the previous months.
func (c *Client) UpdateCostCategoryDefinition(ctx context.Context, params *UpdateCostCategoryDefinitionInput, optFns ...func(*Options)) (*UpdateCostCategoryDefinitionOutput, error) {
	if params == nil {
		params = &UpdateCostCategoryDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCostCategoryDefinition", params, optFns, addOperationUpdateCostCategoryDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCostCategoryDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCostCategoryDefinitionInput struct {

	// The unique identifier for your Cost Category.
	//
	// This member is required.
	CostCategoryArn *string

	// The rule schema version in this particular Cost Category.
	//
	// This member is required.
	RuleVersion types.CostCategoryRuleVersion

	// The Expression object used to categorize costs. For more information, see
	// CostCategoryRule
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategoryRule.html).
	//
	// This member is required.
	Rules []*types.CostCategoryRule
}

type UpdateCostCategoryDefinitionOutput struct {

	// The unique identifier for your Cost Category.
	CostCategoryArn *string

	// The Cost Category's effective start date.
	EffectiveStart *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateCostCategoryDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCostCategoryDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCostCategoryDefinition{}, middleware.After)
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
	addOpUpdateCostCategoryDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCostCategoryDefinition(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateCostCategoryDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "UpdateCostCategoryDefinition",
	}
}
