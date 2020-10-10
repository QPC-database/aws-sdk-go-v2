// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata about a backup vault specified by its name.
func (c *Client) DescribeBackupVault(ctx context.Context, params *DescribeBackupVaultInput, optFns ...func(*Options)) (*DescribeBackupVaultOutput, error) {
	if params == nil {
		params = &DescribeBackupVaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBackupVault", params, optFns, addOperationDescribeBackupVaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBackupVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBackupVaultInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	//
	// This member is required.
	BackupVaultName *string
}

type DescribeBackupVaultOutput struct {

	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Region where they are created. They consist of lowercase letters, numbers, and
	// hyphens.
	BackupVaultName *string

	// The date and time that a backup vault is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of executing the operation twice.
	CreatorRequestId *string

	// The server-side encryption key that is used to protect your backups; for
	// example,
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	EncryptionKeyArn *string

	// The number of recovery points that are stored in a backup vault.
	NumberOfRecoveryPoints *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBackupVaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBackupVault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBackupVault{}, middleware.After)
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
	addOpDescribeBackupVaultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackupVault(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBackupVault(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeBackupVault",
	}
}
