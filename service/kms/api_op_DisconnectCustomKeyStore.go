// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disconnects the custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// from its associated AWS CloudHSM cluster. While a custom key store is
// disconnected, you can manage the custom key store and its customer master keys
// (CMKs), but you cannot create or use CMKs in the custom key store. You can
// reconnect the custom key store at any time. While a custom key store is
// disconnected, all attempts to create customer master keys (CMKs) in the custom
// key store or to use existing CMKs in cryptographic operations
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
// will fail. This action can prevent users from storing and accessing sensitive
// data. To find the connection state of a custom key store, use the
// DescribeCustomKeyStores () operation. To reconnect a custom key store, use the
// ConnectCustomKeyStore () operation. If the operation succeeds, it returns a JSON
// object with no properties. This operation is part of the Custom Key Store
// feature
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in AWS KMS, which combines the convenience and extensive integration of
// AWS KMS with the isolation and control of a single-tenant key store.
func (c *Client) DisconnectCustomKeyStore(ctx context.Context, params *DisconnectCustomKeyStoreInput, optFns ...func(*Options)) (*DisconnectCustomKeyStoreOutput, error) {
	if params == nil {
		params = &DisconnectCustomKeyStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisconnectCustomKeyStore", params, optFns, addOperationDisconnectCustomKeyStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisconnectCustomKeyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisconnectCustomKeyStoreInput struct {

	// Enter the ID of the custom key store you want to disconnect. To find the ID of a
	// custom key store, use the DescribeCustomKeyStores () operation.
	//
	// This member is required.
	CustomKeyStoreId *string
}

type DisconnectCustomKeyStoreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisconnectCustomKeyStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisconnectCustomKeyStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisconnectCustomKeyStore{}, middleware.After)
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
	addOpDisconnectCustomKeyStoreValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisconnectCustomKeyStore(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisconnectCustomKeyStore(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "DisconnectCustomKeyStore",
	}
}
