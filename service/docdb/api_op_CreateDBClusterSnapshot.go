// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a snapshot of a cluster.
func (c *Client) CreateDBClusterSnapshot(ctx context.Context, params *CreateDBClusterSnapshotInput, optFns ...func(*Options)) (*CreateDBClusterSnapshotOutput, error) {
	if params == nil {
		params = &CreateDBClusterSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBClusterSnapshot", params, optFns, addOperationCreateDBClusterSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of CreateDBClusterSnapshot ().
type CreateDBClusterSnapshotInput struct {

	// The identifier of the cluster to create a snapshot for. This parameter is not
	// case sensitive. Constraints:
	//
	//     * Must match the identifier of an existing
	// DBCluster.
	//
	// Example: my-cluster
	//
	// This member is required.
	DBClusterIdentifier *string

	// The identifier of the cluster snapshot. This parameter is stored as a lowercase
	// string. Constraints:
	//
	//     * Must contain from 1 to 63 letters, numbers, or
	// hyphens.
	//
	//     * The first character must be a letter.
	//
	//     * Cannot end with a
	// hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster-snapshot1
	//
	// This member is required.
	DBClusterSnapshotIdentifier *string

	// The tags to be assigned to the cluster snapshot.
	Tags []*types.Tag
}

type CreateDBClusterSnapshotOutput struct {

	// Detailed information about a cluster snapshot.
	DBClusterSnapshot *types.DBClusterSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDBClusterSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBClusterSnapshot{}, middleware.After)
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
	addOpCreateDBClusterSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBClusterSnapshot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDBClusterSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBClusterSnapshot",
	}
}
