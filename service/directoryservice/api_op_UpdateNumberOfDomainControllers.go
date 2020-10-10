// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or removes domain controllers to or from the directory. Based on the
// difference between current value and new value (provided through this API call),
// domain controllers will be added or removed. It may take up to 45 minutes for
// any new domain controllers to become fully active once the requested number of
// domain controllers is updated. During this time, you cannot make another update
// request.
func (c *Client) UpdateNumberOfDomainControllers(ctx context.Context, params *UpdateNumberOfDomainControllersInput, optFns ...func(*Options)) (*UpdateNumberOfDomainControllersOutput, error) {
	if params == nil {
		params = &UpdateNumberOfDomainControllersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNumberOfDomainControllers", params, optFns, addOperationUpdateNumberOfDomainControllersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNumberOfDomainControllersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNumberOfDomainControllersInput struct {

	// The number of domain controllers desired in the directory.
	//
	// This member is required.
	DesiredNumber *int32

	// Identifier of the directory to which the domain controllers will be added or
	// removed.
	//
	// This member is required.
	DirectoryId *string
}

type UpdateNumberOfDomainControllersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateNumberOfDomainControllersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNumberOfDomainControllers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNumberOfDomainControllers{}, middleware.After)
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
	addOpUpdateNumberOfDomainControllersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNumberOfDomainControllers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateNumberOfDomainControllers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "UpdateNumberOfDomainControllers",
	}
}
