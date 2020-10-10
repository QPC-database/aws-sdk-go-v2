// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an endpoint configuration that Amazon SageMaker hosting services uses to
// deploy models. In the configuration, you identify one or more models, created
// using the CreateModel API, to deploy and the resources that you want Amazon
// SageMaker to provision. Then you call the CreateEndpoint () API. Use this API if
// you want to use Amazon SageMaker hosting services to deploy models into
// production. In the request, you define a ProductionVariant, for each model that
// you want to deploy. Each ProductionVariant parameter also describes the
// resources that you want Amazon SageMaker to provision. This includes the number
// and type of ML compute instances to deploy. If you are hosting multiple models,
// you also assign a VariantWeight to specify how much traffic you want to allocate
// to each model. For example, suppose that you want to host two models, A and B,
// and you assign traffic weight 2 for model A and 1 for model B. Amazon SageMaker
// distributes two-thirds of the traffic to Model A, and one-third to model B. For
// an example that calls this method when deploying a model to Amazon SageMaker
// hosting services, see Deploy the Model to Amazon SageMaker Hosting Services (AWS
// SDK for Python (Boto 3)).
// (https://docs.aws.amazon.com/sagemaker/latest/dg/ex1-deploy-model.html#ex1-deploy-model-boto)
// When you call CreateEndpoint (), a load call is made to DynamoDB to verify that
// your endpoint configuration exists. When you read data from a DynamoDB table
// supporting Eventually Consistent Reads
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html),
// the response might not reflect the results of a recently completed write
// operation. The response might include some stale data. If the dependent entities
// are not yet in DynamoDB, this causes a validation error. If you repeat your read
// request after a short time, the response should return the latest data. So retry
// logic is recommended to handle these possible issues. We also recommend that
// customers call DescribeEndpointConfig () before calling CreateEndpoint () to
// minimize the potential impact of a DynamoDB eventually consistent read.
func (c *Client) CreateEndpointConfig(ctx context.Context, params *CreateEndpointConfigInput, optFns ...func(*Options)) (*CreateEndpointConfigOutput, error) {
	if params == nil {
		params = &CreateEndpointConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEndpointConfig", params, optFns, addOperationCreateEndpointConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEndpointConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEndpointConfigInput struct {

	// The name of the endpoint configuration. You specify this name in a
	// CreateEndpoint () request.
	//
	// This member is required.
	EndpointConfigName *string

	// An list of ProductionVariant objects, one for each model that you want to host
	// at this endpoint.
	//
	// This member is required.
	ProductionVariants []*types.ProductionVariant

	//
	DataCaptureConfig *types.DataCaptureConfig

	// The Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon
	// SageMaker uses to encrypt data on the storage volume attached to the ML compute
	// instance that hosts the endpoint. The KmsKeyId can be any of the following
	// formats:
	//
	//     * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//
	// * Alias name: alias/ExampleAlias
	//
	//     * Alias name ARN:
	// arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias
	//
	// The KMS key policy must
	// grant permission to the IAM role that you specify in your CreateEndpoint,
	// UpdateEndpoint requests. For more information, refer to the AWS Key Management
	// Service section Using Key Policies in AWS KMS
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)
	// Certain Nitro-based instances include local storage, dependent on the instance
	// type. Local storage volumes are encrypted using a hardware module on the
	// instance. You can't request a KmsKeyId when using an instance type with local
	// storage. If any of the models that you specify in the ProductionVariants
	// parameter use nitro-based instances with local storage, do not specify a value
	// for the KmsKeyId parameter. If you specify a value for KmsKeyId when using any
	// nitro-based instances with local storage, the call to CreateEndpointConfig
	// fails. For a list of instance types that support local instance storage, see
	// Instance Store Volumes
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes).
	// For more information about local instance storage encryption, see SSD Instance
	// Store Volumes
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html).
	KmsKeyId *string

	// A list of key-value pairs. For more information, see Using Cost Allocation Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the AWS Billing and Cost Management User Guide.
	Tags []*types.Tag
}

type CreateEndpointConfigOutput struct {

	// The Amazon Resource Name (ARN) of the endpoint configuration.
	//
	// This member is required.
	EndpointConfigArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEndpointConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEndpointConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEndpointConfig{}, middleware.After)
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
	addOpCreateEndpointConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpointConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateEndpointConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateEndpointConfig",
	}
}
