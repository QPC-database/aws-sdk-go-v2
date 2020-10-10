// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the IAM roles that have the specified path prefix. If there are none, the
// operation returns an empty list. For more information about roles, go to Working
// with Roles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html). You
// can paginate the results using the MaxItems and Marker parameters.
func (c *Client) ListRoles(ctx context.Context, params *ListRolesInput, optFns ...func(*Options)) (*ListRolesOutput, error) {
	if params == nil {
		params = &ListRolesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRoles", params, optFns, addOperationListRolesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRolesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRolesInput struct {

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32

	// The path prefix for filtering the results. For example, the prefix
	// /application_abc/component_xyz/ gets all roles whose path starts with
	// /application_abc/component_xyz/. This parameter is optional. If it is not
	// included, it defaults to a slash (/), listing all roles. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of either a forward slash (/) by itself or a string that
	// must begin and end with forward slashes. In addition, it can contain any ASCII
	// character from the ! (\u0021) through the DEL character (\u007F), including most
	// punctuation characters, digits, and upper and lowercased letters.
	PathPrefix *string
}

// Contains the response to a successful ListRoles () request.
type ListRolesOutput struct {

	// A list of roles.
	//
	// This member is required.
	Roles []*types.Role

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you receive
	// all your results.
	IsTruncated *bool

	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRolesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListRoles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListRoles{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRoles(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListRoles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListRoles",
	}
}
