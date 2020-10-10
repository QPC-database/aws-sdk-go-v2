// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the EvaluationName of an Evaluation. You can use the GetEvaluation
// operation to view the contents of the updated data element.
func (c *Client) UpdateEvaluation(ctx context.Context, params *UpdateEvaluationInput, optFns ...func(*Options)) (*UpdateEvaluationOutput, error) {
	if params == nil {
		params = &UpdateEvaluationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEvaluation", params, optFns, addOperationUpdateEvaluationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEvaluationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEvaluationInput struct {

	// The ID assigned to the Evaluation during creation.
	//
	// This member is required.
	EvaluationId *string

	// A new user-supplied name or description of the Evaluation that will replace the
	// current content.
	//
	// This member is required.
	EvaluationName *string
}

// Represents the output of an UpdateEvaluation operation. You can see the updated
// content by using the GetEvaluation operation.
type UpdateEvaluationOutput struct {

	// The ID assigned to the Evaluation during creation. This value should be
	// identical to the value of the Evaluation in the request.
	EvaluationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateEvaluationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEvaluation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEvaluation{}, middleware.After)
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
	addOpUpdateEvaluationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEvaluation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateEvaluation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "UpdateEvaluation",
	}
}
