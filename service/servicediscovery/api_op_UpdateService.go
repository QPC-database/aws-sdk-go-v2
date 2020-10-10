// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Submits a request to perform the following operations:
//
//     * Update the TTL
// setting for existing DnsRecords configurations
//
//     * Add, update, or delete
// HealthCheckConfig for a specified service You can't add, update, or delete a
// HealthCheckCustomConfig configuration.
//
// For public and private DNS namespaces,
// note the following:
//
//     * If you omit any existing DnsRecords or
// HealthCheckConfig configurations from an UpdateService request, the
// configurations are deleted from the service.
//
//     * If you omit an existing
// HealthCheckCustomConfig configuration from an UpdateService request, the
// configuration is not deleted from the service.
//
// When you update settings for a
// service, AWS Cloud Map also updates the corresponding settings in all the
// records and health checks that were created by using the specified service.
func (c *Client) UpdateService(ctx context.Context, params *UpdateServiceInput, optFns ...func(*Options)) (*UpdateServiceOutput, error) {
	if params == nil {
		params = &UpdateServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateService", params, optFns, addOperationUpdateServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceInput struct {

	// The ID of the service that you want to update.
	//
	// This member is required.
	Id *string

	// A complex type that contains the new settings for the service.
	//
	// This member is required.
	Service *types.ServiceChange
}

type UpdateServiceOutput struct {

	// A value that you can use to determine whether the request completed
	// successfully. To get the status of the operation, see GetOperation
	// (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateService{}, middleware.After)
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
	addOpUpdateServiceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateService(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateService(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "UpdateService",
	}
}
