// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data
// Analytics API V2 Documentation (). Adds a CloudWatch log stream to monitor
// application configuration errors. For more information about using CloudWatch
// log streams with Amazon Kinesis Analytics applications, see Working with Amazon
// CloudWatch Logs
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/cloudwatch-logs.html).
func (c *Client) AddApplicationCloudWatchLoggingOption(ctx context.Context, params *AddApplicationCloudWatchLoggingOptionInput, optFns ...func(*Options)) (*AddApplicationCloudWatchLoggingOptionOutput, error) {
	if params == nil {
		params = &AddApplicationCloudWatchLoggingOptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationCloudWatchLoggingOption", params, optFns, addOperationAddApplicationCloudWatchLoggingOptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationCloudWatchLoggingOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationCloudWatchLoggingOptionInput struct {

	// The Kinesis Analytics application name.
	//
	// This member is required.
	ApplicationName *string

	// Provides the CloudWatch log stream Amazon Resource Name (ARN) and the IAM role
	// ARN. Note: To write application messages to CloudWatch, the IAM role that is
	// used must have the PutLogEvents policy action enabled.
	//
	// This member is required.
	CloudWatchLoggingOption *types.CloudWatchLoggingOption

	// The version ID of the Kinesis Analytics application.
	//
	// This member is required.
	CurrentApplicationVersionId *int64
}

type AddApplicationCloudWatchLoggingOptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddApplicationCloudWatchLoggingOptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationCloudWatchLoggingOption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationCloudWatchLoggingOption{}, middleware.After)
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
	addOpAddApplicationCloudWatchLoggingOptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationCloudWatchLoggingOption(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAddApplicationCloudWatchLoggingOption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "AddApplicationCloudWatchLoggingOption",
	}
}
