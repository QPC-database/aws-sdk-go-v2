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

// Gets code to perform a specified mapping.
func (c *Client) GetPlan(ctx context.Context, params *GetPlanInput, optFns ...func(*Options)) (*GetPlanOutput, error) {
	if params == nil {
		params = &GetPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPlan", params, optFns, addOperationGetPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPlanInput struct {

	// The list of mappings from a source table to target tables.
	//
	// This member is required.
	Mapping []*types.MappingEntry

	// The source table.
	//
	// This member is required.
	Source *types.CatalogEntry

	// The programming language of the code to perform the mapping.
	Language types.Language

	// The parameters for the mapping.
	Location *types.Location

	// The target tables.
	Sinks []*types.CatalogEntry
}

type GetPlanOutput struct {

	// A Python script to perform the mapping.
	PythonScript *string

	// The Scala code to perform the mapping.
	ScalaCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPlan{}, middleware.After)
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
	addOpGetPlanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPlan(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetPlan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetPlan",
	}
}
