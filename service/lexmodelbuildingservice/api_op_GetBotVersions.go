// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about all of the versions of a bot. The GetBotVersions
// operation returns a BotMetadata object for each version of a bot. For example,
// if a bot has three numbered versions, the GetBotVersions operation returns four
// BotMetadata objects in the response, one for each numbered version and one for
// the $LATEST version. The GetBotVersions operation always returns at least one
// version, the $LATEST version. This operation requires permissions for the
// lex:GetBotVersions action.
func (c *Client) GetBotVersions(ctx context.Context, params *GetBotVersionsInput, optFns ...func(*Options)) (*GetBotVersionsOutput, error) {
	if params == nil {
		params = &GetBotVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBotVersions", params, optFns, addOperationGetBotVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBotVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotVersionsInput struct {

	// The name of the bot for which versions should be returned.
	//
	// This member is required.
	Name *string

	// The maximum number of bot versions to return in the response. The default is 10.
	MaxResults *int32

	// A pagination token for fetching the next page of bot versions. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string
}

type GetBotVersionsOutput struct {

	// An array of BotMetadata objects, one for each numbered version of the bot plus
	// one for the $LATEST version.
	Bots []*types.BotMetadata

	// A pagination token for fetching the next page of bot versions. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBotVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBotVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBotVersions{}, middleware.After)
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
	addOpGetBotVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBotVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBotVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBotVersions",
	}
}
