// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the standards that are currently enabled.
func (c *Client) GetEnabledStandards(ctx context.Context, params *GetEnabledStandardsInput, optFns ...func(*Options)) (*GetEnabledStandardsOutput, error) {
	if params == nil {
		params = &GetEnabledStandardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnabledStandards", params, optFns, addOperationGetEnabledStandardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnabledStandardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnabledStandardsInput struct {

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The token that is required for pagination. On your first call to the
	// GetEnabledStandards operation, set the value of this parameter to NULL. For
	// subsequent calls to the operation, to continue listing data, set the value of
	// this parameter to the value returned from the previous response.
	NextToken *string

	// The list of the standards subscription ARNs for the standards to retrieve.
	StandardsSubscriptionArns []*string
}

type GetEnabledStandardsOutput struct {

	// The pagination token to use to request the next page of results.
	NextToken *string

	// The list of StandardsSubscriptions objects that include information about the
	// enabled standards.
	StandardsSubscriptions []*types.StandardsSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEnabledStandardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEnabledStandards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEnabledStandards{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnabledStandards(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetEnabledStandards(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "GetEnabledStandards",
	}
}
