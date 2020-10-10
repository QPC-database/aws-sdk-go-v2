// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Simulate the execution of an Authorizer () in your RestApi () with headers,
// parameters, and an incoming request body. Use Lambda Function as Authorizer
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html)Use
// Cognito User Pool as Authorizer
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html)
func (c *Client) TestInvokeAuthorizer(ctx context.Context, params *TestInvokeAuthorizerInput, optFns ...func(*Options)) (*TestInvokeAuthorizerOutput, error) {
	if params == nil {
		params = &TestInvokeAuthorizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestInvokeAuthorizer", params, optFns, addOperationTestInvokeAuthorizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestInvokeAuthorizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Make a request to simulate the execution of an Authorizer ().
type TestInvokeAuthorizerInput struct {

	// [Required] Specifies a test invoke authorizer request's Authorizer () ID.
	//
	// This member is required.
	AuthorizerId *string

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	// [Optional] A key-value map of additional context variables.
	AdditionalContext map[string]*string

	// [Optional] The simulated request body of an incoming invocation request.
	Body *string

	// [Required] A key-value map of headers to simulate an incoming invocation
	// request. This is where the incoming authorization token, or identity source,
	// should be specified.
	Headers map[string]*string

	// [Optional] The headers as a map from string to list of values to simulate an
	// incoming invocation request. This is where the incoming authorization token, or
	// identity source, may be specified.
	MultiValueHeaders map[string][]*string

	// [Optional] The URI path, including query string, of the simulated invocation
	// request. Use this to specify path parameters and query string parameters.
	PathWithQueryString *string

	// A key-value map of stage variables to simulate an invocation on a deployed Stage
	// ().
	StageVariables map[string]*string
}

// Represents the response of the test invoke request for a custom Authorizer ()
type TestInvokeAuthorizerOutput struct {
	Authorization map[string][]*string

	// The open identity claims
	// (https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims), with any
	// supported custom attributes, returned from the Cognito Your User Pool configured
	// for the API.
	Claims map[string]*string

	// The HTTP status code that the client would have received. Value is 0 if the
	// authorizer succeeded.
	ClientStatus *int32

	// The execution latency of the test authorizer request.
	Latency *int64

	// The API Gateway execution log for the test authorizer request.
	Log *string

	// The JSON policy document returned by the Authorizer ()
	Policy *string

	// The principal identity returned by the Authorizer ()
	PrincipalId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTestInvokeAuthorizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestInvokeAuthorizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestInvokeAuthorizer{}, middleware.After)
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
	addOpTestInvokeAuthorizerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestInvokeAuthorizer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opTestInvokeAuthorizer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "TestInvokeAuthorizer",
	}
}
