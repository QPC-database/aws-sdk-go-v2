// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves all the fields and values of a single log event. All fields are
// retrieved, even if the original query that produced the logRecordPointer
// retrieved only a subset of fields. Fields are returned as field name/field value
// pairs. Additionally, the entire unparsed log event is returned within @message.
func (c *Client) GetLogRecord(ctx context.Context, params *GetLogRecordInput, optFns ...func(*Options)) (*GetLogRecordOutput, error) {
	if params == nil {
		params = &GetLogRecordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLogRecord", params, optFns, addOperationGetLogRecordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLogRecordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLogRecordInput struct {

	// The pointer corresponding to the log event record you want to retrieve. You get
	// this from the response of a GetQueryResults operation. In that response, the
	// value of the @ptr field for a log event is the value to use as logRecordPointer
	// to retrieve that complete log event record.
	//
	// This member is required.
	LogRecordPointer *string
}

type GetLogRecordOutput struct {

	// The requested log event, as a JSON string.
	LogRecord map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLogRecordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLogRecord{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLogRecord{}, middleware.After)
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
	addOpGetLogRecordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLogRecord(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetLogRecord(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "GetLogRecord",
	}
}
