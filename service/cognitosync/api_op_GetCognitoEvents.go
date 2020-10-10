// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the events and the corresponding Lambda functions associated with an
// identity pool. This API can only be called with developer credentials. You
// cannot call this API with the temporary user credentials provided by Cognito
// Identity.
func (c *Client) GetCognitoEvents(ctx context.Context, params *GetCognitoEventsInput, optFns ...func(*Options)) (*GetCognitoEventsOutput, error) {
	if params == nil {
		params = &GetCognitoEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCognitoEvents", params, optFns, addOperationGetCognitoEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCognitoEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for a list of the configured Cognito Events
type GetCognitoEventsInput struct {

	// The Cognito Identity Pool ID for the request
	//
	// This member is required.
	IdentityPoolId *string
}

// The response from the GetCognitoEvents request
type GetCognitoEventsOutput struct {

	// The Cognito Events returned from the GetCognitoEvents request
	Events map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCognitoEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCognitoEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCognitoEvents{}, middleware.After)
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
	addOpGetCognitoEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCognitoEvents(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCognitoEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "GetCognitoEvents",
	}
}
