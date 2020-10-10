// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a deployment job.
func (c *Client) DescribeDeploymentJob(ctx context.Context, params *DescribeDeploymentJobInput, optFns ...func(*Options)) (*DescribeDeploymentJobOutput, error) {
	if params == nil {
		params = &DescribeDeploymentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDeploymentJob", params, optFns, addOperationDescribeDeploymentJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDeploymentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDeploymentJobInput struct {

	// The Amazon Resource Name (ARN) of the deployment job.
	//
	// This member is required.
	Job *string
}

type DescribeDeploymentJobOutput struct {

	// The Amazon Resource Name (ARN) of the deployment job.
	Arn *string

	// The time, in milliseconds since the epoch, when the deployment job was created.
	CreatedAt *time.Time

	// The deployment application configuration.
	DeploymentApplicationConfigs []*types.DeploymentApplicationConfig

	// The deployment configuration.
	DeploymentConfig *types.DeploymentConfig

	// The deployment job failure code.
	FailureCode types.DeploymentJobErrorCode

	// A short description of the reason why the deployment job failed.
	FailureReason *string

	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string

	// A list of robot deployment summaries.
	RobotDeploymentSummary []*types.RobotDeployment

	// The status of the deployment job.
	Status types.DeploymentStatus

	// The list of all tags added to the specified deployment job.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDeploymentJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDeploymentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDeploymentJob{}, middleware.After)
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
	addOpDescribeDeploymentJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDeploymentJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDeploymentJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "DescribeDeploymentJob",
	}
}
