// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Function object. A function is a reusable entity. Multiple functions
// can be used to compose the resolver logic.
func (c *Client) CreateFunction(ctx context.Context, params *CreateFunctionInput, optFns ...func(*Options)) (*CreateFunctionOutput, error) {
	if params == nil {
		params = &CreateFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFunction", params, optFns, addOperationCreateFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFunctionInput struct {

	// The GraphQL API ID.
	//
	// This member is required.
	ApiId *string

	// The FunctionDataSource name.
	//
	// This member is required.
	DataSourceName *string

	// The version of the request mapping template. Currently the supported value is
	// 2018-05-29.
	//
	// This member is required.
	FunctionVersion *string

	// The Function name. The function name does not have to be unique.
	//
	// This member is required.
	Name *string

	// The Function request mapping template. Functions support only the 2018-05-29
	// version of the request mapping template.
	//
	// This member is required.
	RequestMappingTemplate *string

	// The Function description.
	Description *string

	// The Function response mapping template.
	ResponseMappingTemplate *string
}

type CreateFunctionOutput struct {

	// The Function object.
	FunctionConfiguration *types.FunctionConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFunction{}, middleware.After)
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
	addOpCreateFunctionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFunction(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateFunction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "CreateFunction",
	}
}
