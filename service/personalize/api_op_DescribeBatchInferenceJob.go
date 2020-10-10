// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the properties of a batch inference job including name, Amazon Resource
// Name (ARN), status, input and output configurations, and the ARN of the solution
// version used to generate the recommendations.
func (c *Client) DescribeBatchInferenceJob(ctx context.Context, params *DescribeBatchInferenceJobInput, optFns ...func(*Options)) (*DescribeBatchInferenceJobOutput, error) {
	if params == nil {
		params = &DescribeBatchInferenceJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBatchInferenceJob", params, optFns, addOperationDescribeBatchInferenceJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBatchInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBatchInferenceJobInput struct {

	// The ARN of the batch inference job to describe.
	//
	// This member is required.
	BatchInferenceJobArn *string
}

type DescribeBatchInferenceJobOutput struct {

	// Information on the specified batch inference job.
	BatchInferenceJob *types.BatchInferenceJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBatchInferenceJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBatchInferenceJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBatchInferenceJob{}, middleware.After)
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
	addOpDescribeBatchInferenceJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBatchInferenceJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBatchInferenceJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DescribeBatchInferenceJob",
	}
}
