// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Turns on notifications on a backup vault for the specified topic and events.
func (c *Client) PutBackupVaultNotifications(ctx context.Context, params *PutBackupVaultNotificationsInput, optFns ...func(*Options)) (*PutBackupVaultNotificationsOutput, error) {
	if params == nil {
		params = &PutBackupVaultNotificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBackupVaultNotifications", params, optFns, addOperationPutBackupVaultNotificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBackupVaultNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBackupVaultNotificationsInput struct {

	// An array of events that indicate the status of jobs to back up resources to the
	// backup vault.
	//
	// This member is required.
	BackupVaultEvents []types.BackupVaultEvent

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s
	// events; for example, arn:aws:sns:us-west-2:111122223333:MyVaultTopic.
	//
	// This member is required.
	SNSTopicArn *string
}

type PutBackupVaultNotificationsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBackupVaultNotificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutBackupVaultNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutBackupVaultNotifications{}, middleware.After)
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
	addOpPutBackupVaultNotificationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBackupVaultNotifications(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutBackupVaultNotifications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "PutBackupVaultNotifications",
	}
}
