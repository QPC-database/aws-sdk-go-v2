// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestar/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the user profiles configured for your AWS account in AWS CodeStar.
func (c *Client) ListUserProfiles(ctx context.Context, params *ListUserProfilesInput, optFns ...func(*Options)) (*ListUserProfilesOutput, error) {
	if params == nil {
		params = &ListUserProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserProfiles", params, optFns, addOperationListUserProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUserProfilesInput struct {

	// The maximum number of results to return in a response.
	MaxResults *int32

	// The continuation token for the next set of results, if the results cannot be
	// returned in one response.
	NextToken *string
}

type ListUserProfilesOutput struct {

	// All the user profiles configured in AWS CodeStar for an AWS account.
	//
	// This member is required.
	UserProfiles []*types.UserProfileSummary

	// The continuation token to use when requesting the next set of results, if there
	// are more results to be returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUserProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUserProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUserProfiles{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListUserProfiles(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListUserProfiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "ListUserProfiles",
	}
}
