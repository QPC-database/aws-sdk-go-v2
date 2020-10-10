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

// Returns the backup policy for the specified EFS file system.
func (c *Client) DescribeBackupPolicy(ctx context.Context, params *DescribeBackupPolicyInput, optFns ...func(*Options)) (*DescribeBackupPolicyOutput, error) {
	if params == nil {
		params = &DescribeBackupPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBackupPolicy", params, optFns, addOperationDescribeBackupPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBackupPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBackupPolicyInput struct {

	// Specifies which EFS file system to retrieve the BackupPolicy for.
	//
	// This member is required.
	FileSystemId *string
}

type DescribeBackupPolicyOutput struct {

	// Describes the file system's backup policy, indicating whether automatic backups
	// are turned on or off..
	BackupPolicy *types.BackupPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBackupPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBackupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBackupPolicy{}, middleware.After)
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
	addOpDescribeBackupPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackupPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBackupPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeBackupPolicy",
	}
}
