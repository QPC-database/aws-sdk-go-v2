// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Copy an AWS CloudHSM cluster backup to a different region.
func (c *Client) CopyBackupToRegion(ctx context.Context, params *CopyBackupToRegionInput, optFns ...func(*Options)) (*CopyBackupToRegionOutput, error) {
	if params == nil {
		params = &CopyBackupToRegionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyBackupToRegion", params, optFns, addOperationCopyBackupToRegionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyBackupToRegionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyBackupToRegionInput struct {

	// The ID of the backup that will be copied to the destination region.
	//
	// This member is required.
	BackupId *string

	// The AWS region that will contain your copied CloudHSM cluster backup.
	//
	// This member is required.
	DestinationRegion *string

	// Tags to apply to the destination backup during creation. If you specify tags,
	// only these tags will be applied to the destination backup. If you do not specify
	// tags, the service copies tags from the source backup to the destination backup.
	TagList []*types.Tag
}

type CopyBackupToRegionOutput struct {

	// Information on the backup that will be copied to the destination region,
	// including CreateTimestamp, SourceBackup, SourceCluster, and Source Region.
	// CreateTimestamp of the destination backup will be the same as that of the source
	// backup. You will need to use the sourceBackupID returned in this operation to
	// use the DescribeBackups () operation on the backup that will be copied to the
	// destination region.
	DestinationBackup *types.DestinationBackup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCopyBackupToRegionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCopyBackupToRegion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCopyBackupToRegion{}, middleware.After)
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
	addOpCopyBackupToRegionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyBackupToRegion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCopyBackupToRegion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "CopyBackupToRegion",
	}
}
