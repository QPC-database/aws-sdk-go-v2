// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all license usage records for a license configuration, displaying license
// consumption details by resource at a selected point in time. Use this action to
// audit the current license consumption for any license inventory and
// configuration.
func (c *Client) ListUsageForLicenseConfiguration(ctx context.Context, params *ListUsageForLicenseConfigurationInput, optFns ...func(*Options)) (*ListUsageForLicenseConfigurationOutput, error) {
	if params == nil {
		params = &ListUsageForLicenseConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsageForLicenseConfiguration", params, optFns, addOperationListUsageForLicenseConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsageForLicenseConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUsageForLicenseConfigurationInput struct {

	// Amazon Resource Name (ARN) of the license configuration.
	//
	// This member is required.
	LicenseConfigurationArn *string

	// Filters to scope the results. The following filters and logical operators are
	// supported:
	//
	//     * resourceArn - The ARN of the license configuration resource.
	// Logical operators are EQUALS | NOT_EQUALS.
	//
	//     * resourceType - The resource
	// type (EC2_INSTANCE | EC2_HOST | EC2_AMI | SYSTEMS_MANAGER_MANAGED_INSTANCE).
	// Logical operators are EQUALS | NOT_EQUALS.
	//
	//     * resourceAccount - The ID of
	// the account that owns the resource. Logical operators are EQUALS | NOT_EQUALS.
	Filters []*types.Filter

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string
}

type ListUsageForLicenseConfigurationOutput struct {

	// Information about the license configurations.
	LicenseConfigurationUsageList []*types.LicenseConfigurationUsage

	// Token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUsageForLicenseConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUsageForLicenseConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUsageForLicenseConfiguration{}, middleware.After)
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
	addOpListUsageForLicenseConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListUsageForLicenseConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListUsageForLicenseConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "ListUsageForLicenseConfiguration",
	}
}
