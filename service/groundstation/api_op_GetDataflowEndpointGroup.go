// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the dataflow endpoint group.
func (c *Client) GetDataflowEndpointGroup(ctx context.Context, params *GetDataflowEndpointGroupInput, optFns ...func(*Options)) (*GetDataflowEndpointGroupOutput, error) {
	if params == nil {
		params = &GetDataflowEndpointGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataflowEndpointGroup", params, optFns, addOperationGetDataflowEndpointGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataflowEndpointGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type GetDataflowEndpointGroupInput struct {

	// UUID of a dataflow endpoint group.
	//
	// This member is required.
	DataflowEndpointGroupId *string
}

//
type GetDataflowEndpointGroupOutput struct {

	// ARN of a dataflow endpoint group.
	DataflowEndpointGroupArn *string

	// UUID of a dataflow endpoint group.
	DataflowEndpointGroupId *string

	// Details of a dataflow endpoint.
	EndpointsDetails []*types.EndpointDetails

	// Tags assigned to a dataflow endpoint group.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDataflowEndpointGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataflowEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataflowEndpointGroup{}, middleware.After)
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
	addOpGetDataflowEndpointGroupValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
