// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the principal permissions on the resource, filtered by the
// permissions of the caller. For example, if you are granted an ALTER permission,
// you are able to see only the principal permissions for ALTER. This operation
// returns only those permissions that have been explicitly granted. For
// information about permissions, see Security and Access Control to Metadata and
// Data
// (https://docs-aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
func (c *Client) ListPermissions(ctx context.Context, params *ListPermissionsInput, optFns ...func(*Options)) (*ListPermissionsOutput, error) {
	if params == nil {
		params = &ListPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissions", params, optFns, addOperationListPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionsInput struct {

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	// Specifies a principal to filter the permissions returned.
	Principal *types.DataLakePrincipal

	// A resource where you will get a list of the principal permissions. This
	// operation does not support getting privileges on a table with columns. Instead,
	// call this operation on the table, and the operation returns the table and the
	// table w columns.
	Resource *types.Resource

	// Specifies a resource type to filter the permissions returned.
	ResourceType types.DataLakeResourceType
}

type ListPermissionsOutput struct {

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	// A list of principals and their permissions on the resource for the specified
	// principal and resource types.
	PrincipalResourcePermissions []*types.PrincipalResourcePermissions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPermissions{}, middleware.After)
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
	addOpListPermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListPermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "ListPermissions",
	}
}
