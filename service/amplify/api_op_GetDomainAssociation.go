// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the domain information for an Amplify app.
func (c *Client) GetDomainAssociation(ctx context.Context, params *GetDomainAssociationInput, optFns ...func(*Options)) (*GetDomainAssociationOutput, error) {
	if params == nil {
		params = &GetDomainAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainAssociation", params, optFns, addOperationGetDomainAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the get domain association request.
type GetDomainAssociationInput struct {

	// The unique id for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of the domain.
	//
	// This member is required.
	DomainName *string
}

// The result structure for the get domain association request.
type GetDomainAssociationOutput struct {

	// Describes the structure of a domain association, which associates a custom
	// domain with an Amplify app.
	//
	// This member is required.
	DomainAssociation *types.DomainAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDomainAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDomainAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomainAssociation{}, middleware.After)
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
	addOpGetDomainAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainAssociation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDomainAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "GetDomainAssociation",
	}
}
