// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the configuration recorder. After the configuration recorder is deleted,
// AWS Config will not record resource configuration changes until you create a new
// configuration recorder. This action does not delete the configuration
// information that was previously recorded. You will be able to access the
// previously recorded information by using the GetResourceConfigHistory action,
// but you will not be able to access this information in the AWS Config console
// until you create a new configuration recorder.
func (c *Client) DeleteConfigurationRecorder(ctx context.Context, params *DeleteConfigurationRecorderInput, optFns ...func(*Options)) (*DeleteConfigurationRecorderOutput, error) {
	if params == nil {
		params = &DeleteConfigurationRecorderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConfigurationRecorder", params, optFns, addOperationDeleteConfigurationRecorderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConfigurationRecorderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for the DeleteConfigurationRecorder action.
type DeleteConfigurationRecorderInput struct {

	// The name of the configuration recorder to be deleted. You can retrieve the name
	// of your configuration recorder by using the DescribeConfigurationRecorders
	// action.
	//
	// This member is required.
	ConfigurationRecorderName *string
}

type DeleteConfigurationRecorderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteConfigurationRecorderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteConfigurationRecorder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteConfigurationRecorder{}, middleware.After)
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
	addOpDeleteConfigurationRecorderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationRecorder(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteConfigurationRecorder(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteConfigurationRecorder",
	}
}
