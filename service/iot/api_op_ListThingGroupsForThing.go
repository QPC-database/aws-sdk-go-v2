// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the thing groups to which the specified thing belongs.
func (c *Client) ListThingGroupsForThing(ctx context.Context, params *ListThingGroupsForThingInput, optFns ...func(*Options)) (*ListThingGroupsForThingOutput, error) {
	if params == nil {
		params = &ListThingGroupsForThingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingGroupsForThing", params, optFns, addOperationListThingGroupsForThingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingGroupsForThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingGroupsForThingInput struct {

	// The thing name.
	//
	// This member is required.
	ThingName *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string
}

type ListThingGroupsForThingOutput struct {

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The thing groups.
	ThingGroups []*types.GroupNameAndArn

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListThingGroupsForThingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingGroupsForThing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingGroupsForThing{}, middleware.After)
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
	addOpListThingGroupsForThingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListThingGroupsForThing(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListThingGroupsForThing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingGroupsForThing",
	}
}
