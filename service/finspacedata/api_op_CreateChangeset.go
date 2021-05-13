// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new changeset in a FinSpace dataset.
func (c *Client) CreateChangeset(ctx context.Context, params *CreateChangesetInput, optFns ...func(*Options)) (*CreateChangesetOutput, error) {
	if params == nil {
		params = &CreateChangesetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChangeset", params, optFns, addOperationCreateChangesetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChangesetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChangesetInput struct {

	// Option to indicate how a changeset will be applied to a dataset.
	//
	// * REPLACE -
	// Changeset will be considered as a replacement to all prior loaded changesets.
	//
	// *
	// APPEND - Changeset will be considered as an addition to the end of all prior
	// loaded changesets.
	//
	// This member is required.
	ChangeType types.ChangeType

	// The unique identifier for the FinSpace dataset in which the changeset will be
	// created.
	//
	// This member is required.
	DatasetId *string

	// Source path from which the files to create the changeset will be sourced.
	//
	// This member is required.
	SourceParams map[string]string

	// Type of the data source from which the files to create the changeset will be
	// sourced.
	//
	// * S3 - Amazon S3.
	//
	// This member is required.
	SourceType types.SourceType

	// Options that define the structure of the source file(s).
	FormatParams map[string]string

	// Format type of the input files being loaded into the changeset.
	FormatType types.FormatType

	// Metadata tags to apply to this changeset.
	Tags map[string]string
}

type CreateChangesetOutput struct {

	// Returns the changeset details.
	Changeset *types.ChangesetInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateChangesetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChangeset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChangeset{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateChangesetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChangeset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateChangeset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "CreateChangeset",
	}
}
