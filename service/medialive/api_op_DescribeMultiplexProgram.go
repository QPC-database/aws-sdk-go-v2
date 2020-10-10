// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the details for a program in a multiplex.
func (c *Client) DescribeMultiplexProgram(ctx context.Context, params *DescribeMultiplexProgramInput, optFns ...func(*Options)) (*DescribeMultiplexProgramOutput, error) {
	if params == nil {
		params = &DescribeMultiplexProgramInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMultiplexProgram", params, optFns, addOperationDescribeMultiplexProgramMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMultiplexProgramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeMultiplexProgramRequest
type DescribeMultiplexProgramInput struct {

	// The ID of the multiplex that the program belongs to.
	//
	// This member is required.
	MultiplexId *string

	// The name of the program.
	//
	// This member is required.
	ProgramName *string
}

// Placeholder documentation for DescribeMultiplexProgramResponse
type DescribeMultiplexProgramOutput struct {

	// The MediaLive channel associated with the program.
	ChannelId *string

	// The settings for this multiplex program.
	MultiplexProgramSettings *types.MultiplexProgramSettings

	// The packet identifier map for this multiplex program.
	PacketIdentifiersMap *types.MultiplexProgramPacketIdentifiersMap

	// The name of the multiplex program.
	ProgramName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMultiplexProgramMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMultiplexProgram{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMultiplexProgram{}, middleware.After)
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
	addOpDescribeMultiplexProgramValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMultiplexProgram(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeMultiplexProgram(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeMultiplexProgram",
	}
}
