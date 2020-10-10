// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an HSM configuration that contains the information required by an Amazon
// Redshift cluster to store and use database encryption keys in a Hardware
// Security Module (HSM). After creating the HSM configuration, you can specify it
// as a parameter when creating a cluster. The cluster will then store its
// encryption keys in the HSM. In addition to creating an HSM configuration, you
// must also create an HSM client certificate. For more information, go to Hardware
// Security Modules
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-HSM.html) in the
// Amazon Redshift Cluster Management Guide.
func (c *Client) CreateHsmConfiguration(ctx context.Context, params *CreateHsmConfigurationInput, optFns ...func(*Options)) (*CreateHsmConfigurationOutput, error) {
	if params == nil {
		params = &CreateHsmConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHsmConfiguration", params, optFns, addOperationCreateHsmConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHsmConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateHsmConfigurationInput struct {

	// A text description of the HSM configuration to be created.
	//
	// This member is required.
	Description *string

	// The identifier to be assigned to the new Amazon Redshift HSM configuration.
	//
	// This member is required.
	HsmConfigurationIdentifier *string

	// The IP address that the Amazon Redshift cluster must use to access the HSM.
	//
	// This member is required.
	HsmIpAddress *string

	// The name of the partition in the HSM where the Amazon Redshift clusters will
	// store their database encryption keys.
	//
	// This member is required.
	HsmPartitionName *string

	// The password required to access the HSM partition.
	//
	// This member is required.
	HsmPartitionPassword *string

	// The HSMs public certificate file. When using Cloud HSM, the file name is
	// server.pem.
	//
	// This member is required.
	HsmServerPublicCertificate *string

	// A list of tag instances.
	Tags []*types.Tag
}

type CreateHsmConfigurationOutput struct {

	// Returns information about an HSM configuration, which is an object that
	// describes to Amazon Redshift clusters the information they require to connect to
	// an HSM where they can store database encryption keys.
	HsmConfiguration *types.HsmConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateHsmConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateHsmConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateHsmConfiguration{}, middleware.After)
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
	addOpCreateHsmConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHsmConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateHsmConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateHsmConfiguration",
	}
}
