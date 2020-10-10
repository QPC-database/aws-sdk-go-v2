// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns an Evaluation that includes metadata as well as the current status of
// the Evaluation.
func (c *Client) GetEvaluation(ctx context.Context, params *GetEvaluationInput, optFns ...func(*Options)) (*GetEvaluationOutput, error) {
	if params == nil {
		params = &GetEvaluationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEvaluation", params, optFns, addOperationGetEvaluationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEvaluationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEvaluationInput struct {

	// The ID of the Evaluation to retrieve. The evaluation of each MLModel is recorded
	// and cataloged. The ID provides the means to access the information.
	//
	// This member is required.
	EvaluationId *string
}

// Represents the output of a GetEvaluation operation and describes an Evaluation.
type GetEvaluationOutput struct {

	// The approximate CPU time in milliseconds that Amazon Machine Learning spent
	// processing the Evaluation, normalized and scaled on computation resources.
	// ComputeTime is only available if the Evaluation is in the COMPLETED state.
	ComputeTime *int64

	// The time that the Evaluation was created. The time is expressed in epoch time.
	CreatedAt *time.Time

	// The AWS user account that invoked the evaluation. The account type can be either
	// an AWS root account or an AWS Identity and Access Management (IAM) user account.
	CreatedByIamUser *string

	// The DataSource used for this evaluation.
	EvaluationDataSourceId *string

	// The evaluation ID which is same as the EvaluationId in the request.
	EvaluationId *string

	// The epoch time when Amazon Machine Learning marked the Evaluation as COMPLETED
	// or FAILED. FinishedAt is only available when the Evaluation is in the COMPLETED
	// or FAILED state.
	FinishedAt *time.Time

	// The location of the data file or directory in Amazon Simple Storage Service
	// (Amazon S3).
	InputDataLocationS3 *string

	// The time of the most recent edit to the Evaluation. The time is expressed in
	// epoch time.
	LastUpdatedAt *time.Time

	// A link to the file that contains logs of the CreateEvaluation operation.
	LogUri *string

	// The ID of the MLModel that was the focus of the evaluation.
	MLModelId *string

	// A description of the most recent details about evaluating the MLModel.
	Message *string

	// A user-supplied name or description of the Evaluation.
	Name *string

	// Measurements of how well the MLModel performed using observations referenced by
	// the DataSource. One of the following metric is returned based on the type of the
	// MLModel:
	//
	//     * BinaryAUC: A binary MLModel uses the Area Under the Curve (AUC)
	// technique to measure performance.
	//
	//     * RegressionRMSE: A regression MLModel
	// uses the Root Mean Square Error (RMSE) technique to measure performance. RMSE
	// measures the difference between predicted and actual values for a single
	// variable.
	//
	//     * MulticlassAvgFScore: A multiclass MLModel uses the F1 score
	// technique to measure performance.
	//
	// For more information about performance
	// metrics, please see the Amazon Machine Learning Developer Guide
	// (https://docs.aws.amazon.com/machine-learning/latest/dg).
	PerformanceMetrics *types.PerformanceMetrics

	// The epoch time when Amazon Machine Learning marked the Evaluation as INPROGRESS.
	// StartedAt isn't available if the Evaluation is in the PENDING state.
	StartedAt *time.Time

	// The status of the evaluation. This element can have one of the following
	// values:
	//
	//     * PENDING - Amazon Machine Language (Amazon ML) submitted a request
	// to evaluate an MLModel.
	//
	//     * INPROGRESS - The evaluation is underway.
	//
	//     *
	// FAILED - The request to evaluate an MLModel did not run to completion. It is not
	// usable.
	//
	//     * COMPLETED - The evaluation process completed successfully.
	//
	//     *
	// DELETED - The Evaluation is marked as deleted. It is not usable.
	Status types.EntityStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEvaluationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetEvaluation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEvaluation{}, middleware.After)
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
	addOpGetEvaluationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEvaluation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetEvaluation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "GetEvaluation",
	}
}
