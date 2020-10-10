// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates an Amazon EMR block public access configuration for your AWS
// account in the current Region. For more information see Configure Block Public
// Access for Amazon EMR
// (https://docs.aws.amazon.com/emr/latest/ManagementGuide/configure-block-public-access.html)
// in the Amazon EMR Management Guide.
func (c *Client) PutBlockPublicAccessConfiguration(ctx context.Context, params *PutBlockPublicAccessConfigurationInput, optFns ...func(*Options)) (*PutBlockPublicAccessConfigurationOutput, error) {
	if params == nil {
		params = &PutBlockPublicAccessConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBlockPublicAccessConfiguration", params, optFns, addOperationPutBlockPublicAccessConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBlockPublicAccessConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBlockPublicAccessConfigurationInput struct {

	// A configuration for Amazon EMR block public access. The configuration applies to
	// all clusters created in your account for the current Region. The configuration
	// specifies whether block public access is enabled. If block public access is
	// enabled, security groups associated with the cluster cannot have rules that
	// allow inbound traffic from 0.0.0.0/0 or ::/0 on a port, unless the port is
	// specified as an exception using PermittedPublicSecurityGroupRuleRanges in the
	// BlockPublicAccessConfiguration. By default, Port 22 (SSH) is an exception, and
	// public access is allowed on this port. You can change this by updating
	// BlockPublicSecurityGroupRules to remove the exception. For accounts that created
	// clusters in a Region before November 25, 2019, block public access is disabled
	// by default in that Region. To use this feature, you must manually enable and
	// configure it. For accounts that did not create an EMR cluster in a Region before
	// this date, block public access is enabled by default in that Region.
	//
	// This member is required.
	BlockPublicAccessConfiguration *types.BlockPublicAccessConfiguration
}

type PutBlockPublicAccessConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBlockPublicAccessConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutBlockPublicAccessConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutBlockPublicAccessConfiguration{}, middleware.After)
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
	addOpPutBlockPublicAccessConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBlockPublicAccessConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutBlockPublicAccessConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "PutBlockPublicAccessConfiguration",
	}
}
