// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// InferRxNorm detects medications as entities listed in a patient record and links
// to the normalized concept identifiers in the RxNorm database from the National
// Library of Medicine. Amazon Comprehend Medical only detects medical entities in
// English language texts.
func (c *Client) InferRxNorm(ctx context.Context, params *InferRxNormInput, optFns ...func(*Options)) (*InferRxNormOutput, error) {
	if params == nil {
		params = &InferRxNormInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InferRxNorm", params, optFns, addOperationInferRxNormMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InferRxNormOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InferRxNormInput struct {

	// The input text used for analysis. The input for InferRxNorm is a string from 1
	// to 10000 characters.
	//
	// This member is required.
	Text *string
}

type InferRxNormOutput struct {

	// The medication entities detected in the text linked to RxNorm concepts. If the
	// action is successful, the service sends back an HTTP 200 response, as well as
	// the entities detected.
	//
	// This member is required.
	Entities []*types.RxNormEntity

	// The version of the model used to analyze the documents, in the format n.n.n You
	// can use this information to track the model used for a particular batch of
	// documents.
	ModelVersion *string

	// If the result of the previous request to InferRxNorm was truncated, include the
	// PaginationToken to fetch the next page of medication entities.
	PaginationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInferRxNormMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInferRxNorm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInferRxNorm{}, middleware.After)
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
	addOpInferRxNormValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInferRxNorm(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opInferRxNorm(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "InferRxNorm",
	}
}
