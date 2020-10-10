// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a predictor created using the CreatePredictor () operation. In
// addition to listing the properties provided in the CreatePredictor request, this
// operation lists the following properties:
//
//     * DatasetImportJobArns - The
// dataset import jobs used to import training data.
//
//     * AutoMLAlgorithmArns -
// If AutoML is performed, the algorithms that were evaluated.
//
//     *
// CreationTime
//
//     * LastModificationTime
//
//     * Status
//
//     * Message - If an
// error occurred, information about the error.
func (c *Client) DescribePredictor(ctx context.Context, params *DescribePredictorInput, optFns ...func(*Options)) (*DescribePredictorOutput, error) {
	if params == nil {
		params = &DescribePredictorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePredictor", params, optFns, addOperationDescribePredictorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePredictorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePredictorInput struct {

	// The Amazon Resource Name (ARN) of the predictor that you want information about.
	//
	// This member is required.
	PredictorArn *string
}

type DescribePredictorOutput struct {

	// The Amazon Resource Name (ARN) of the algorithm used for model training.
	AlgorithmArn *string

	// When PerformAutoML is specified, the ARN of the chosen algorithm.
	AutoMLAlgorithmArns []*string

	// When the model training task was created.
	CreationTime *time.Time

	// An array of the ARNs of the dataset import jobs used to import training data for
	// the predictor.
	DatasetImportJobArns []*string

	// An AWS Key Management Service (KMS) key and the AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig *types.EncryptionConfig

	// Used to override the default evaluation parameters of the specified algorithm.
	// Amazon Forecast evaluates a predictor by splitting a dataset into training data
	// and testing data. The evaluation parameters define how to perform the split and
	// the number of iterations.
	EvaluationParameters *types.EvaluationParameters

	// The featurization configuration.
	FeaturizationConfig *types.FeaturizationConfig

	// The number of time-steps of the forecast. The forecast horizon is also called
	// the prediction length.
	ForecastHorizon *int32

	// The hyperparameter override values for the algorithm.
	HPOConfig *types.HyperParameterTuningJobConfig

	// Describes the dataset group that contains the data to use to train the
	// predictor.
	InputDataConfig *types.InputDataConfig

	// Initially, the same as CreationTime (when the status is CREATE_PENDING). This
	// value is updated when training starts (when the status changes to
	// CREATE_IN_PROGRESS), and when training has completed (when the status changes to
	// ACTIVE) or fails (when the status changes to CREATE_FAILED).
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// Whether the predictor is set to perform AutoML.
	PerformAutoML *bool

	// Whether the predictor is set to perform hyperparameter optimization (HPO).
	PerformHPO *bool

	// The ARN of the predictor.
	PredictorArn *string

	// Details on the the status and results of the backtests performed to evaluate the
	// accuracy of the predictor. You specify the number of backtests to perform when
	// you call the operation.
	PredictorExecutionDetails *types.PredictorExecutionDetails

	// The name of the predictor.
	PredictorName *string

	// The status of the predictor. States include:
	//
	//     * ACTIVE
	//
	//     *
	// CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//     * DELETE_PENDING,
	// DELETE_IN_PROGRESS, DELETE_FAILED
	//
	//     * UPDATE_PENDING, UPDATE_IN_PROGRESS,
	// UPDATE_FAILED
	//
	// The Status of the predictor must be ACTIVE before you can use the
	// predictor to create a forecast.
	Status *string

	// The default training parameters or overrides selected during model training. If
	// using the AutoML algorithm or if HPO is turned on while using the DeepAR+
	// algorithms, the optimized values for the chosen hyperparameters are returned.
	// For more information, see aws-forecast-choosing-recipes ().
	TrainingParameters map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePredictorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePredictor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePredictor{}, middleware.After)
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
	addOpDescribePredictorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePredictor(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribePredictor(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribePredictor",
	}
}
