// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates a replication configuration for an application.
func (c *Client) PutAppReplicationConfiguration(ctx context.Context, params *PutAppReplicationConfigurationInput, optFns ...func(*Options)) (*PutAppReplicationConfigurationOutput, error) {
	if params == nil {
		params = &PutAppReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAppReplicationConfiguration", params, optFns, addOperationPutAppReplicationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAppReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAppReplicationConfigurationInput struct {

	// ID of the application tassociated with the replication configuration.
	AppId *string

	// Replication configurations for server groups in the application.
	ServerGroupReplicationConfigurations []*types.ServerGroupReplicationConfiguration
}

type PutAppReplicationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutAppReplicationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAppReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAppReplicationConfiguration{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutAppReplicationConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutAppReplicationConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "PutAppReplicationConfiguration",
	}
}
