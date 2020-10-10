// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a running batch build.
func (c *Client) StopBuildBatch(ctx context.Context, params *StopBuildBatchInput, optFns ...func(*Options)) (*StopBuildBatchOutput, error) {
	if params == nil {
		params = &StopBuildBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopBuildBatch", params, optFns, addOperationStopBuildBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopBuildBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopBuildBatchInput struct {

	// The identifier of the batch build to stop.
	//
	// This member is required.
	Id *string
}

type StopBuildBatchOutput struct {

	// Contains information about a batch build.
	BuildBatch *types.BuildBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopBuildBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopBuildBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopBuildBatch{}, middleware.After)
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
	addOpStopBuildBatchValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopBuildBatch(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopBuildBatch(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "StopBuildBatch",
	}
}
