// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the list of available network profiles.
func (c *Client) ListNetworkProfiles(ctx context.Context, params *ListNetworkProfilesInput, optFns ...func(*Options)) (*ListNetworkProfilesOutput, error) {
	if params == nil {
		params = &ListNetworkProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNetworkProfiles", params, optFns, addOperationListNetworkProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNetworkProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNetworkProfilesInput struct {

	// The Amazon Resource Name (ARN) of the project for which you want to list network
	// profiles.
	//
	// This member is required.
	Arn *string

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// The type of network profile to return information about. Valid values are listed
	// here.
	Type types.NetworkProfileType
}

type ListNetworkProfilesOutput struct {

	// A list of the available network profiles.
	NetworkProfiles []*types.NetworkProfile

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListNetworkProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListNetworkProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListNetworkProfiles{}, middleware.After)
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
	addOpListNetworkProfilesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListNetworkProfiles(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListNetworkProfiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListNetworkProfiles",
	}
}
