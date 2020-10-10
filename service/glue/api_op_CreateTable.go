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

// Creates a new table definition in the Data Catalog.
func (c *Client) CreateTable(ctx context.Context, params *CreateTableInput, optFns ...func(*Options)) (*CreateTableOutput, error) {
	if params == nil {
		params = &CreateTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTable", params, optFns, addOperationCreateTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTableInput struct {

	// The catalog database in which to create the new table. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// The TableInput object that defines the metadata table to create in the catalog.
	//
	// This member is required.
	TableInput *types.TableInput

	// The ID of the Data Catalog in which to create the Table. If none is supplied,
	// the AWS account ID is used by default.
	CatalogId *string
}

type CreateTableOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTable{}, middleware.After)
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
	addOpCreateTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTable(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateTable",
	}
}
