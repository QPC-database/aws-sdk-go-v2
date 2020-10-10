// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Instructs the specified agents or connectors to stop collecting data.
func (c *Client) StopDataCollectionByAgentIds(ctx context.Context, params *StopDataCollectionByAgentIdsInput, optFns ...func(*Options)) (*StopDataCollectionByAgentIdsOutput, error) {
	if params == nil {
		params = &StopDataCollectionByAgentIdsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDataCollectionByAgentIds", params, optFns, addOperationStopDataCollectionByAgentIdsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDataCollectionByAgentIdsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDataCollectionByAgentIdsInput struct {

	// The IDs of the agents or connectors from which to stop collecting data.
	//
	// This member is required.
	AgentIds []*string
}

type StopDataCollectionByAgentIdsOutput struct {

	// Information about the agents or connector that were instructed to stop
	// collecting data. Information includes the agent/connector ID, a description of
	// the operation performed, and whether the agent/connector configuration was
	// updated.
	AgentsConfigurationStatus []*types.AgentConfigurationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopDataCollectionByAgentIdsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopDataCollectionByAgentIds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDataCollectionByAgentIds{}, middleware.After)
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
	addOpStopDataCollectionByAgentIdsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDataCollectionByAgentIds(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopDataCollectionByAgentIds(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "StopDataCollectionByAgentIds",
	}
}
