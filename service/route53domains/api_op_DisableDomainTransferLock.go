// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation removes the transfer lock on the domain (specifically the
// clientTransferProhibited status) to allow domain transfers. We recommend you
// refrain from performing this action unless you intend to transfer the domain to
// a different registrar. Successful submission returns an operation ID that you
// can use to track the progress and completion of the action. If the request is
// not completed successfully, the domain registrant will be notified by email.
func (c *Client) DisableDomainTransferLock(ctx context.Context, params *DisableDomainTransferLockInput, optFns ...func(*Options)) (*DisableDomainTransferLockOutput, error) {
	if params == nil {
		params = &DisableDomainTransferLockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableDomainTransferLock", params, optFns, addOperationDisableDomainTransferLockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableDomainTransferLockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The DisableDomainTransferLock request includes the following element.
type DisableDomainTransferLockInput struct {

	// The name of the domain that you want to remove the transfer lock for.
	//
	// This member is required.
	DomainName *string
}

// The DisableDomainTransferLock response includes the following element.
type DisableDomainTransferLockOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	//
	// This member is required.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableDomainTransferLockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisableDomainTransferLock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableDomainTransferLock{}, middleware.After)
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
	addOpDisableDomainTransferLockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableDomainTransferLock(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisableDomainTransferLock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "DisableDomainTransferLock",
	}
}
