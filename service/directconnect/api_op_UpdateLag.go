// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the attributes of the specified link aggregation group (LAG). You can
// update the following attributes:
//
//     * The name of the LAG.
//
//     * The value
// for the minimum number of connections that must be operational for the LAG
// itself to be operational.
//
// When you create a LAG, the default value for the
// minimum number of operational connections is zero (0). If you update this value
// and the number of operational connections falls below the specified value, the
// LAG automatically goes down to avoid over-utilization of the remaining
// connections. Adjust this value with care, as it could force the LAG down if it
// is set higher than the current number of operational connections.
func (c *Client) UpdateLag(ctx context.Context, params *UpdateLagInput, optFns ...func(*Options)) (*UpdateLagOutput, error) {
	if params == nil {
		params = &UpdateLagInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLag", params, optFns, addOperationUpdateLagMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLagOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLagInput struct {

	// The ID of the LAG.
	//
	// This member is required.
	LagId *string

	// The name of the LAG.
	LagName *string

	// The minimum number of physical connections that must be operational for the LAG
	// itself to be operational.
	MinimumLinks *int32
}

// Information about a link aggregation group (LAG).
type UpdateLagOutput struct {

	// Indicates whether the LAG can host other connections.
	AllowsHostedConnections *bool

	// The AWS Direct Connect endpoint that hosts the LAG.
	AwsDevice *string

	// The AWS Direct Connect endpoint that hosts the LAG.
	AwsDeviceV2 *string

	// The connections bundled by the LAG.
	Connections []*types.Connection

	// The individual bandwidth of the physical connections bundled by the LAG. The
	// possible values are 1Gbps and 10Gbps.
	ConnectionsBandwidth *string

	// Indicates whether the LAG supports a secondary BGP peer in the same address
	// family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The name of the LAG.
	LagName *string

	// The state of the LAG. The following are the possible values:
	//
	//     * requested:
	// The initial state of a LAG. The LAG stays in the requested state until the
	// Letter of Authorization (LOA) is available.
	//
	//     * pending: The LAG has been
	// approved and is being initialized.
	//
	//     * available: The network link is
	// established and the LAG is ready for use.
	//
	//     * down: The network link is
	// down.
	//
	//     * deleting: The LAG is being deleted.
	//
	//     * deleted: The LAG is
	// deleted.
	//
	//     * unknown: The state of the LAG is not available.
	LagState types.LagState

	// The location of the LAG.
	Location *string

	// The minimum number of physical connections that must be operational for the LAG
	// itself to be operational.
	MinimumLinks *int32

	// The number of physical connections bundled by the LAG, up to a maximum of 10.
	NumberOfConnections *int32

	// The ID of the AWS account that owns the LAG.
	OwnerAccount *string

	// The name of the service provider associated with the LAG.
	ProviderName *string

	// The AWS Region where the connection is located.
	Region *string

	// The tags associated with the LAG.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateLagMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLag{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLag{}, middleware.After)
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
	addOpUpdateLagValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLag(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateLag(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "UpdateLag",
	}
}
