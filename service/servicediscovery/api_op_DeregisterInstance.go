// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the Amazon Route 53 DNS records and health check, if any, that AWS Cloud
// Map created for the specified instance.
func (c *Client) DeregisterInstance(ctx context.Context, params *DeregisterInstanceInput, optFns ...func(*Options)) (*DeregisterInstanceOutput, error) {
	if params == nil {
		params = &DeregisterInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterInstance", params, optFns, addOperationDeregisterInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterInstanceInput struct {

	// The value that you specified for Id in the RegisterInstance
	// (https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html)
	// request.
	//
	// This member is required.
	InstanceId *string

	// The ID of the service that the instance is associated with.
	//
	// This member is required.
	ServiceId *string
}

type DeregisterInstanceOutput struct {

	// A value that you can use to determine whether the request completed
	// successfully. For more information, see GetOperation
	// (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterInstance{}, middleware.After)
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
	addOpDeregisterInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeregisterInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "DeregisterInstance",
	}
}
