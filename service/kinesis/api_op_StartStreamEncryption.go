// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables or updates server-side encryption using an AWS KMS key for a specified
// stream. Starting encryption is an asynchronous operation. Upon receiving the
// request, Kinesis Data Streams returns immediately and sets the status of the
// stream to UPDATING. After the update is complete, Kinesis Data Streams sets the
// status of the stream back to ACTIVE. Updating or applying encryption normally
// takes a few seconds to complete, but it can take minutes. You can continue to
// read and write data to your stream while its status is UPDATING. Once the status
// of the stream is ACTIVE, encryption begins for records written to the stream.
// API Limits: You can successfully apply a new AWS KMS key for server-side
// encryption 25 times in a rolling 24-hour period. Note: It can take up to 5
// seconds after the stream is in an ACTIVE status before all records written to
// the stream are encrypted. After you enable encryption, you can verify that
// encryption is applied by inspecting the API response from PutRecord or
// PutRecords.
func (c *Client) StartStreamEncryption(ctx context.Context, params *StartStreamEncryptionInput, optFns ...func(*Options)) (*StartStreamEncryptionOutput, error) {
	if params == nil {
		params = &StartStreamEncryptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartStreamEncryption", params, optFns, addOperationStartStreamEncryptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartStreamEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartStreamEncryptionInput struct {

	// The encryption type to use. The only valid value is KMS.
	//
	// This member is required.
	EncryptionType types.EncryptionType

	// The GUID for the customer-managed AWS KMS key to use for encryption. This value
	// can be a globally unique identifier, a fully specified Amazon Resource Name
	// (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You can
	// also use a master key owned by Kinesis Data Streams by specifying the alias
	// aws/kinesis.
	//
	//     * Key ARN example:
	// arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//
	// * Alias ARN example: arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	//     *
	// Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	//     *
	// Alias name example: alias/MyAliasName
	//
	//     * Master key owned by Kinesis Data
	// Streams: alias/aws/kinesis
	//
	// This member is required.
	KeyId *string

	// The name of the stream for which to start encrypting records.
	//
	// This member is required.
	StreamName *string
}

type StartStreamEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartStreamEncryptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartStreamEncryption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartStreamEncryption{}, middleware.After)
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
	addOpStartStreamEncryptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartStreamEncryption(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartStreamEncryption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "StartStreamEncryption",
	}
}
