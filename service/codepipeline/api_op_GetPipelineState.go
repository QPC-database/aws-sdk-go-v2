// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about the state of a pipeline, including the stages and
// actions. Values returned in the revisionId and revisionUrl fields indicate the
// source revision information, such as the commit ID, for the current state.
func (c *Client) GetPipelineState(ctx context.Context, params *GetPipelineStateInput, optFns ...func(*Options)) (*GetPipelineStateOutput, error) {
	if params == nil {
		params = &GetPipelineStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPipelineState", params, optFns, addOperationGetPipelineStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPipelineStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetPipelineState action.
type GetPipelineStateInput struct {

	// The name of the pipeline about which you want to get information.
	//
	// This member is required.
	Name *string
}

// Represents the output of a GetPipelineState action.
type GetPipelineStateOutput struct {

	// The date and time the pipeline was created, in timestamp format.
	Created *time.Time

	// The name of the pipeline for which you want to get the state.
	PipelineName *string

	// The version number of the pipeline. A newly created pipeline is always assigned
	// a version number of 1.
	PipelineVersion *int32

	// A list of the pipeline stage output information, including stage name, state,
	// most recent run details, whether the stage is disabled, and other data.
	StageStates []*types.StageState

	// The date and time the pipeline was last updated, in timestamp format.
	Updated *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPipelineStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPipelineState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPipelineState{}, middleware.After)
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
	addOpGetPipelineStateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPipelineState(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetPipelineState(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "GetPipelineState",
	}
}
