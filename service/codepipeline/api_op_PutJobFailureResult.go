// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Represents the failure of a job as returned to the pipeline by a job worker.
// Used for custom actions only.
func (c *Client) PutJobFailureResult(ctx context.Context, params *PutJobFailureResultInput, optFns ...func(*Options)) (*PutJobFailureResultOutput, error) {
	if params == nil {
		params = &PutJobFailureResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutJobFailureResult", params, optFns, addOperationPutJobFailureResultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutJobFailureResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutJobFailureResult action.
type PutJobFailureResultInput struct {

	// The details about the failure of a job.
	//
	// This member is required.
	FailureDetails *types.FailureDetails

	// The unique system-generated ID of the job that failed. This is the same ID
	// returned from PollForJobs.
	//
	// This member is required.
	JobId *string
}

type PutJobFailureResultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutJobFailureResultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutJobFailureResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutJobFailureResult{}, middleware.After)
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
	addOpPutJobFailureResultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutJobFailureResult(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutJobFailureResult(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "PutJobFailureResult",
	}
}
