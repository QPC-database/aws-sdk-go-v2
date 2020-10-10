// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all of the smart home appliances associated with a room.
func (c *Client) ListSmartHomeAppliances(ctx context.Context, params *ListSmartHomeAppliancesInput, optFns ...func(*Options)) (*ListSmartHomeAppliancesOutput, error) {
	if params == nil {
		params = &ListSmartHomeAppliancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSmartHomeAppliances", params, optFns, addOperationListSmartHomeAppliancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSmartHomeAppliancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSmartHomeAppliancesInput struct {

	// The room that the appliances are associated with.
	//
	// This member is required.
	RoomArn *string

	// The maximum number of appliances to be returned, per paginated calls.
	MaxResults *int32

	// The tokens used for pagination.
	NextToken *string
}

type ListSmartHomeAppliancesOutput struct {

	// The tokens used for pagination.
	NextToken *string

	// The smart home appliances.
	SmartHomeAppliances []*types.SmartHomeAppliance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSmartHomeAppliancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSmartHomeAppliances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSmartHomeAppliances{}, middleware.After)
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
	addOpListSmartHomeAppliancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSmartHomeAppliances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListSmartHomeAppliances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ListSmartHomeAppliances",
	}
}
