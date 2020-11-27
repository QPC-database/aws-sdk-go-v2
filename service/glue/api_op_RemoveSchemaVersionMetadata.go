// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a key value pair from the schema version metadata for the specified
// schema version ID.
func (c *Client) RemoveSchemaVersionMetadata(ctx context.Context, params *RemoveSchemaVersionMetadataInput, optFns ...func(*Options)) (*RemoveSchemaVersionMetadataOutput, error) {
	if params == nil {
		params = &RemoveSchemaVersionMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveSchemaVersionMetadata", params, optFns, addOperationRemoveSchemaVersionMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveSchemaVersionMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveSchemaVersionMetadataInput struct {

	// The value of the metadata key.
	//
	// This member is required.
	MetadataKeyValue *types.MetadataKeyValuePair

	// A wrapper structure that may contain the schema name and Amazon Resource Name
	// (ARN).
	SchemaId *types.SchemaId

	// The unique version ID of the schema version.
	SchemaVersionId *string

	// The version number of the schema.
	SchemaVersionNumber *types.SchemaVersionNumber
}

type RemoveSchemaVersionMetadataOutput struct {

	// The latest version of the schema.
	LatestVersion bool

	// The metadata key.
	MetadataKey *string

	// The value of the metadata key.
	MetadataValue *string

	// The name of the registry.
	RegistryName *string

	// The Amazon Resource Name (ARN) of the schema.
	SchemaArn *string

	// The name of the schema.
	SchemaName *string

	// The version ID for the schema version.
	SchemaVersionId *string

	// The version number of the schema.
	VersionNumber int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveSchemaVersionMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveSchemaVersionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveSchemaVersionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRemoveSchemaVersionMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveSchemaVersionMetadata(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opRemoveSchemaVersionMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "RemoveSchemaVersionMetadata",
	}
}