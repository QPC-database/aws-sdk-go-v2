// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a new infrastructure configuration. An infrastructure configuration
// defines the environment in which your image will be built and tested.
func (c *Client) UpdateInfrastructureConfiguration(ctx context.Context, params *UpdateInfrastructureConfigurationInput, optFns ...func(*Options)) (*UpdateInfrastructureConfigurationOutput, error) {
	if params == nil {
		params = &UpdateInfrastructureConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInfrastructureConfiguration", params, optFns, addOperationUpdateInfrastructureConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInfrastructureConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInfrastructureConfigurationInput struct {

	// The idempotency token used to make this request idempotent.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the infrastructure configuration that you want
	// to update.
	//
	// This member is required.
	InfrastructureConfigurationArn *string

	// The instance profile to associate with the instance used to customize your EC2
	// AMI.
	//
	// This member is required.
	InstanceProfileName *string

	// The description of the infrastructure configuration.
	Description *string

	// The instance types of the infrastructure configuration. You can specify one or
	// more instance types to use for this build. The service will pick one of these
	// instance types based on availability.
	InstanceTypes []*string

	// The key pair of the infrastructure configuration. This can be used to log on to
	// and debug the instance used to create your image.
	KeyPair *string

	// The logging configuration of the infrastructure configuration.
	Logging *types.Logging

	// The tags attached to the resource created by Image Builder.
	ResourceTags map[string]*string

	// The security group IDs to associate with the instance used to customize your EC2
	// AMI.
	SecurityGroupIds []*string

	// The SNS topic on which to send image build events.
	SnsTopicArn *string

	// The subnet ID to place the instance used to customize your EC2 AMI in.
	SubnetId *string

	// The terminate instance on failure setting of the infrastructure configuration.
	// Set to false if you want Image Builder to retain the instance used to configure
	// your AMI if the build or test phase of your workflow fails.
	TerminateInstanceOnFailure *bool
}

type UpdateInfrastructureConfigurationOutput struct {

	// The idempotency token used to make this request idempotent.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the infrastructure configuration that was
	// updated by this request.
	InfrastructureConfigurationArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateInfrastructureConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateInfrastructureConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateInfrastructureConfiguration{}, middleware.After)
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
	addIdempotencyToken_opUpdateInfrastructureConfigurationMiddleware(stack, options)
	addOpUpdateInfrastructureConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInfrastructureConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpUpdateInfrastructureConfiguration struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateInfrastructureConfiguration) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateInfrastructureConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateInfrastructureConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateInfrastructureConfigurationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateInfrastructureConfigurationMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateInfrastructureConfiguration{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateInfrastructureConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "UpdateInfrastructureConfiguration",
	}
}
