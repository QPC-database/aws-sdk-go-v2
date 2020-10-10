// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used by the AWS X-Ray daemon to upload telemetry.
func (c *Client) PutTelemetryRecords(ctx context.Context, params *PutTelemetryRecordsInput, optFns ...func(*Options)) (*PutTelemetryRecordsOutput, error) {
	if params == nil {
		params = &PutTelemetryRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutTelemetryRecords", params, optFns, addOperationPutTelemetryRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutTelemetryRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutTelemetryRecordsInput struct {

	//
	//
	// This member is required.
	TelemetryRecords []*types.TelemetryRecord

	//
	EC2InstanceId *string

	//
	Hostname *string

	//
	ResourceARN *string
}

type PutTelemetryRecordsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutTelemetryRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutTelemetryRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutTelemetryRecords{}, middleware.After)
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
	addOpPutTelemetryRecordsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutTelemetryRecords(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutTelemetryRecords(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "PutTelemetryRecords",
	}
}
