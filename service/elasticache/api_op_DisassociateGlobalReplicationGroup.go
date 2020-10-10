// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove a secondary cluster from the Global Datastore using the Global Datastore
// name. The secondary cluster will no longer receive updates from the primary
// cluster, but will remain as a standalone cluster in that AWS region.
func (c *Client) DisassociateGlobalReplicationGroup(ctx context.Context, params *DisassociateGlobalReplicationGroupInput, optFns ...func(*Options)) (*DisassociateGlobalReplicationGroupOutput, error) {
	if params == nil {
		params = &DisassociateGlobalReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateGlobalReplicationGroup", params, optFns, addOperationDisassociateGlobalReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateGlobalReplicationGroupInput struct {

	// The name of the Global Datastore
	//
	// This member is required.
	GlobalReplicationGroupId *string

	// The name of the secondary cluster you wish to remove from the Global Datastore
	//
	// This member is required.
	ReplicationGroupId *string

	// The AWS region of secondary cluster you wish to remove from the Global Datastore
	//
	// This member is required.
	ReplicationGroupRegion *string
}

type DisassociateGlobalReplicationGroupOutput struct {

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.  <ul> <li> <p>The <b>GlobalReplicationGroupIdSuffix</b>
	// represents the name of the Global Datastore, which is what you use to associate
	// a secondary cluster.</p> </li> </ul>
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateGlobalReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDisassociateGlobalReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDisassociateGlobalReplicationGroup{}, middleware.After)
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
	addOpDisassociateGlobalReplicationGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateGlobalReplicationGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateGlobalReplicationGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DisassociateGlobalReplicationGroup",
	}
}
