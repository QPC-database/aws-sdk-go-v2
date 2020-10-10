// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Restores a backup to a server that is in a CONNECTION_LOST, HEALTHY, RUNNING,
// UNHEALTHY, or TERMINATED state. When you run RestoreServer, the server's EC2
// instance is deleted, and a new EC2 instance is configured. RestoreServer
// maintains the existing server endpoint, so configuration management of the
// server's client devices (nodes) should continue to work. Restoring from a backup
// is performed by creating a new EC2 instance. If restoration is successful, and
// the server is in a HEALTHY state, AWS OpsWorks CM switches traffic over to the
// new instance. After restoration is finished, the old EC2 instance is maintained
// in a Running or Stopped state, but is eventually terminated. This operation is
// asynchronous. An InvalidStateException is thrown when the server is not in a
// valid state. A ResourceNotFoundException is thrown when the server does not
// exist. A ValidationException is raised when parameters of the request are not
// valid.
func (c *Client) RestoreServer(ctx context.Context, params *RestoreServerInput, optFns ...func(*Options)) (*RestoreServerOutput, error) {
	if params == nil {
		params = &RestoreServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreServer", params, optFns, addOperationRestoreServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreServerInput struct {

	// The ID of the backup that you want to use to restore a server.
	//
	// This member is required.
	BackupId *string

	// The name of the server that you want to restore.
	//
	// This member is required.
	ServerName *string

	// The type of instance to restore. Valid values must be specified in the following
	// format: ^([cm][34]|t2).* For example, m5.large. Valid values are m5.large,
	// r5.xlarge, and r5.2xlarge. If you do not specify this parameter, RestoreServer
	// uses the instance type from the specified backup.
	InstanceType *string

	// The name of the key pair to set on the new EC2 instance. This can be helpful if
	// the administrator no longer has the SSH key.
	KeyPair *string
}

type RestoreServerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRestoreServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRestoreServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRestoreServer{}, middleware.After)
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
	addOpRestoreServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreServer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRestoreServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "RestoreServer",
	}
}
