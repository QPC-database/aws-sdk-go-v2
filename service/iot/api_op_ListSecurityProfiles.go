// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the Device Defender security profiles you have created. You can use
// filters to list only those security profiles associated with a thing group or
// only those associated with your account.
func (c *Client) ListSecurityProfiles(ctx context.Context, params *ListSecurityProfilesInput, optFns ...func(*Options)) (*ListSecurityProfilesOutput, error) {
	if params == nil {
		params = &ListSecurityProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityProfiles", params, optFns, addOperationListSecurityProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityProfilesInput struct {

	// A filter to limit results to the security profiles that use the defined
	// dimension.
	DimensionName *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string
}

type ListSecurityProfilesOutput struct {

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// A list of security profile identifiers (names and ARNs).
	SecurityProfileIdentifiers []*types.SecurityProfileIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSecurityProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSecurityProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSecurityProfiles{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityProfiles(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListSecurityProfiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListSecurityProfiles",
	}
}
