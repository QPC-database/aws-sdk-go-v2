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

// Lists all IAM users, groups, and roles that the specified managed policy is
// attached to. You can use the optional EntityFilter parameter to limit the
// results to a particular type of entity (users, groups, or roles). For example,
// to list only the roles that are attached to the specified policy, set
// EntityFilter to Role. You can paginate the results using the MaxItems and Marker
// parameters.
func (c *Client) ListEntitiesForPolicy(ctx context.Context, params *ListEntitiesForPolicyInput, optFns ...func(*Options)) (*ListEntitiesForPolicyOutput, error) {
	if params == nil {
		params = &ListEntitiesForPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntitiesForPolicy", params, optFns, addOperationListEntitiesForPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntitiesForPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitiesForPolicyInput struct {

	// The Amazon Resource Name (ARN) of the IAM policy for which you want the
	// versions. For more information about ARNs, see Amazon Resource Names (ARNs) and
	// AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	PolicyArn *string

	// The entity type to use for filtering the results. For example, when EntityFilter
	// is Role, only the roles that are attached to the specified policy are returned.
	// This parameter is optional. If it is not included, all attached entities (users,
	// groups, and roles) are returned. The argument for this parameter must be one of
	// the valid values listed below.
	EntityFilter types.EntityType

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

	// The path prefix for filtering the results. This parameter is optional. If it is
	// not included, it defaults to a slash (/), listing all entities. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of either a forward slash (/) by itself or a string that
	// must begin and end with forward slashes. In addition, it can contain any ASCII
	// character from the ! (\u0021) through the DEL character (\u007F), including most
	// punctuation characters, digits, and upper and lowercased letters.
	PathPrefix *string

	// The policy usage method to use for filtering the results. To list only
	// permissions policies, set PolicyUsageFilter to PermissionsPolicy. To list only
	// the policies used to set permissions boundaries, set the value to
	// PermissionsBoundary. This parameter is optional. If it is not included, all
	// policies are returned.
	PolicyUsageFilter types.PolicyUsageType
}

// Contains the response to a successful ListEntitiesForPolicy () request.
type ListEntitiesForPolicyOutput struct {

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

	// A list of IAM groups that the policy is attached to.
	PolicyGroups []*types.PolicyGroup

	// A list of IAM roles that the policy is attached to.
	PolicyRoles []*types.PolicyRole

	// A list of IAM users that the policy is attached to.
	PolicyUsers []*types.PolicyUser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEntitiesForPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListEntitiesForPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListEntitiesForPolicy{}, middleware.After)
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
	addOpListEntitiesForPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEntitiesForPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListEntitiesForPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListEntitiesForPolicy",
	}
}
