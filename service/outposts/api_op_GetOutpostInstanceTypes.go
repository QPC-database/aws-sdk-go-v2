// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the instance types for the specified Outpost.
func (c *Client) GetOutpostInstanceTypes(ctx context.Context, params *GetOutpostInstanceTypesInput, optFns ...func(*Options)) (*GetOutpostInstanceTypesOutput, error) {
	if params == nil {
		params = &GetOutpostInstanceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOutpostInstanceTypes", params, optFns, addOperationGetOutpostInstanceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOutpostInstanceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOutpostInstanceTypesInput struct {

	// The ID of the Outpost.
	//
	// This member is required.
	OutpostId *string

	// The maximum page size.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type GetOutpostInstanceTypesOutput struct {

	// Information about the instance types.
	InstanceTypes []*types.InstanceTypeItem

	// The pagination token.
	NextToken *string

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string

	// The ID of the Outpost.
	OutpostId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetOutpostInstanceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetOutpostInstanceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetOutpostInstanceTypes{}, middleware.After)
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
	addOpGetOutpostInstanceTypesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOutpostInstanceTypes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetOutpostInstanceTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "outposts",
		OperationName: "GetOutpostInstanceTypes",
	}
}
