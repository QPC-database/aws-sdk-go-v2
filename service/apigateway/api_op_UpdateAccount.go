// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes information about the current Account () resource.
func (c *Client) UpdateAccount(ctx context.Context, params *UpdateAccountInput, optFns ...func(*Options)) (*UpdateAccountOutput, error) {
	if params == nil {
		params = &UpdateAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccount", params, optFns, addOperationUpdateAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to change information about the current Account ()
// resource.
type UpdateAccountInput struct {
	Name *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents an AWS account that is associated with API Gateway. To view the
// account info, call GET on this resource.
// Error Codes
//
// The following exception
// may be thrown when the request fails.
//
//     * UnauthorizedException
//
//     *
// NotFoundException
//
//     * TooManyRequestsException
//
// For detailed error code
// information, including the corresponding HTTP Status Codes, see API Gateway
// Error Codes
// (https://docs.aws.amazon.com/apigateway/api-reference/handling-errors/#api-error-codes)
// Example:
// Get the information about an account.
//
// Request
//
//     GET /account HTTP/1.1
// Content-Type: application/json Host: apigateway.us-east-1.amazonaws.com
// X-Amz-Date: 20160531T184618Z Authorization: AWS4-HMAC-SHA256
// Credential={access_key_ID}/us-east-1/apigateway/aws4_request,
// SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
//
// Response
//
// The
// successful response returns a 200 OK status code and a payload similar to the
// following: { "_links": { "curies": { "href":
// "https://docs.aws.amazon.com/apigateway/latest/developerguide/account-apigateway-{rel}.html",
// "name": "account", "templated": true }, "self": { "href": "/account" },
// "account:update": { "href": "/account" } }, "cloudwatchRoleArn":
// "arn:aws:iam::123456789012:role/apigAwsProxyRole", "throttleSettings": {
// "rateLimit": 500, "burstLimit": 1000 } }  In addition to making the REST API
// call directly, you can use the AWS CLI and an AWS SDK to access this resource.
// API Gateway Limits
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-limits.html)Developer
// Guide
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html), AWS
// CLI
// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-account.html)
type UpdateAccountOutput struct {

	// The version of the API keys used for the account.
	ApiKeyVersion *string

	// The ARN of an Amazon CloudWatch role for the current Account ().
	CloudwatchRoleArn *string

	// A list of features supported for the account. When usage plans are enabled, the
	// features list will include an entry of "UsagePlans".
	Features []*string

	// Specifies the API request limits configured for the current Account ().
	ThrottleSettings *types.ThrottleSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAccount{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateAccount",
	}
}
