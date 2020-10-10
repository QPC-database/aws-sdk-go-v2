// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the file system's backup policy. Use this action to start or stop
// automatic backups of the file system.
func (c *Client) PutBackupPolicy(ctx context.Context, params *PutBackupPolicyInput, optFns ...func(*Options)) (*PutBackupPolicyOutput, error) {
	if params == nil {
		params = &PutBackupPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBackupPolicy", params, optFns, addOperationPutBackupPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBackupPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBackupPolicyInput struct {

	// The backup policy included in the PutBackupPolicy request.
	//
	// This member is required.
	BackupPolicy *types.BackupPolicy

	// Specifies which EFS file system to update the backup policy for.
	//
	// This member is required.
	FileSystemId *string
}

type PutBackupPolicyOutput struct {

	// Describes the file system's backup policy, indicating whether automatic backups
	// are turned on or off..
	BackupPolicy *types.BackupPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBackupPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutBackupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutBackupPolicy{}, middleware.After)
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
	addOpPutBackupPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBackupPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutBackupPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "PutBackupPolicy",
	}
}
