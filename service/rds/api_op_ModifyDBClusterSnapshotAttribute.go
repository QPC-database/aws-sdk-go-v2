// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an attribute and values to, or removes an attribute and values from, a
// manual DB cluster snapshot. To share a manual DB cluster snapshot with other AWS
// accounts, specify restore as the AttributeName and use the ValuesToAdd parameter
// to add a list of IDs of the AWS accounts that are authorized to restore the
// manual DB cluster snapshot. Use the value all to make the manual DB cluster
// snapshot public, which means that it can be copied or restored by all AWS
// accounts. Don't add the all value for any manual DB cluster snapshots that
// contain private information that you don't want available to all AWS accounts.
// If a manual DB cluster snapshot is encrypted, it can be shared, but only by
// specifying a list of authorized AWS account IDs for the ValuesToAdd parameter.
// You can't use all as a value for that parameter in this case. To view which AWS
// accounts have access to copy or restore a manual DB cluster snapshot, or whether
// a manual DB cluster snapshot is public or private, use the
// DescribeDBClusterSnapshotAttributes () API action. The accounts are returned as
// values for the restore attribute. This action only applies to Aurora DB
// clusters.
func (c *Client) ModifyDBClusterSnapshotAttribute(ctx context.Context, params *ModifyDBClusterSnapshotAttributeInput, optFns ...func(*Options)) (*ModifyDBClusterSnapshotAttributeOutput, error) {
	if params == nil {
		params = &ModifyDBClusterSnapshotAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBClusterSnapshotAttribute", params, optFns, addOperationModifyDBClusterSnapshotAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterSnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyDBClusterSnapshotAttributeInput struct {

	// The name of the DB cluster snapshot attribute to modify. To manage authorization
	// for other AWS accounts to copy or restore a manual DB cluster snapshot, set this
	// value to restore. To view the list of attributes available to modify, use the
	// DescribeDBClusterSnapshotAttributes () API action.
	//
	// This member is required.
	AttributeName *string

	// The identifier for the DB cluster snapshot to modify the attributes for.
	//
	// This member is required.
	DBClusterSnapshotIdentifier *string

	// A list of DB cluster snapshot attributes to add to the attribute specified by
	// AttributeName. To authorize other AWS accounts to copy or restore a manual DB
	// cluster snapshot, set this list to include one or more AWS account IDs, or all
	// to make the manual DB cluster snapshot restorable by any AWS account. Do not add
	// the all value for any manual DB cluster snapshots that contain private
	// information that you don't want available to all AWS accounts.
	ValuesToAdd []*string

	// A list of DB cluster snapshot attributes to remove from the attribute specified
	// by AttributeName. To remove authorization for other AWS accounts to copy or
	// restore a manual DB cluster snapshot, set this list to include one or more AWS
	// account identifiers, or all to remove authorization for any AWS account to copy
	// or restore the DB cluster snapshot. If you specify all, an AWS account whose
	// account ID is explicitly added to the restore attribute can still copy or
	// restore a manual DB cluster snapshot.
	ValuesToRemove []*string
}

type ModifyDBClusterSnapshotAttributeOutput struct {

	// Contains the results of a successful call to the
	// DescribeDBClusterSnapshotAttributes API action. Manual DB cluster snapshot
	// attributes are used to authorize other AWS accounts to copy or restore a manual
	// DB cluster snapshot. For more information, see the
	// ModifyDBClusterSnapshotAttribute API action.
	DBClusterSnapshotAttributesResult *types.DBClusterSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyDBClusterSnapshotAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterSnapshotAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterSnapshotAttribute{}, middleware.After)
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
	addOpModifyDBClusterSnapshotAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterSnapshotAttribute(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyDBClusterSnapshotAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBClusterSnapshotAttribute",
	}
}
