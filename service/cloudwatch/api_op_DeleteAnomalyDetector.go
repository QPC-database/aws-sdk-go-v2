// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified anomaly detection model from your account.
func (c *Client) DeleteAnomalyDetector(ctx context.Context, params *DeleteAnomalyDetectorInput, optFns ...func(*Options)) (*DeleteAnomalyDetectorOutput, error) {
	if params == nil {
		params = &DeleteAnomalyDetectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAnomalyDetector", params, optFns, addOperationDeleteAnomalyDetectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAnomalyDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAnomalyDetectorInput struct {

	// The metric name associated with the anomaly detection model to delete.
	//
	// This member is required.
	MetricName *string

	// The namespace associated with the anomaly detection model to delete.
	//
	// This member is required.
	Namespace *string

	// The statistic associated with the anomaly detection model to delete.
	//
	// This member is required.
	Stat *string

	// The metric dimensions associated with the anomaly detection model to delete.
	Dimensions []*types.Dimension
}

type DeleteAnomalyDetectorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAnomalyDetectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteAnomalyDetector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteAnomalyDetector{}, middleware.After)
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
	addOpDeleteAnomalyDetectorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAnomalyDetector(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAnomalyDetector(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "DeleteAnomalyDetector",
	}
}
