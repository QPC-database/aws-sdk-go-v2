// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the specified data catalog.
func (c *Client) GetDataCatalog(ctx context.Context, params *GetDataCatalogInput, optFns ...func(*Options)) (*GetDataCatalogOutput, error) {
	if params == nil {
		params = &GetDataCatalogInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataCatalog", params, optFns, addOperationGetDataCatalogMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataCatalogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataCatalogInput struct {

	// The name of the data catalog to return.
	//
	// This member is required.
	Name *string
}

type GetDataCatalogOutput struct {

	// The data catalog returned.
	DataCatalog *types.DataCatalog

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDataCatalogMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDataCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDataCatalog{}, middleware.After)
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
	addOpGetDataCatalogValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataCatalog(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDataCatalog(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "GetDataCatalog",
	}
}
