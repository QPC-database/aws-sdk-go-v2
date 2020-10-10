// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the fleet ID that an alias is currently pointing to.
//
//     *
// CreateAlias ()
//
//     * ListAliases ()
//
//     * DescribeAlias ()
//
//     * UpdateAlias
// ()
//
//     * DeleteAlias ()
//
//     * ResolveAlias ()
func (c *Client) ResolveAlias(ctx context.Context, params *ResolveAliasInput, optFns ...func(*Options)) (*ResolveAliasOutput, error) {
	if params == nil {
		params = &ResolveAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResolveAlias", params, optFns, addOperationResolveAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResolveAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type ResolveAliasInput struct {

	// The unique identifier of the alias that you want to retrieve a fleet ID for. You
	// can use either the alias ID or ARN value.
	//
	// This member is required.
	AliasId *string
}

// Represents the returned data in response to a request action.
type ResolveAliasOutput struct {

	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html))
	// associated with the GameLift fleet resource that this alias points to.
	FleetArn *string

	// The fleet identifier that the alias is pointing to.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResolveAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResolveAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResolveAlias{}, middleware.After)
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
	addOpResolveAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResolveAlias(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opResolveAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ResolveAlias",
	}
}
