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

// Lists the license configuration operations that failed.
func (c *Client) ListFailuresForLicenseConfigurationOperations(ctx context.Context, params *ListFailuresForLicenseConfigurationOperationsInput, optFns ...func(*Options)) (*ListFailuresForLicenseConfigurationOperationsOutput, error) {
	if params == nil {
		params = &ListFailuresForLicenseConfigurationOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFailuresForLicenseConfigurationOperations", params, optFns, addOperationListFailuresForLicenseConfigurationOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFailuresForLicenseConfigurationOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFailuresForLicenseConfigurationOperationsInput struct {

	// Amazon Resource Name of the license configuration.
	//
	// This member is required.
	LicenseConfigurationArn *string

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string
}

type ListFailuresForLicenseConfigurationOperationsOutput struct {

	// License configuration operations that failed.
	LicenseOperationFailureList []*types.LicenseOperationFailure

	// Token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFailuresForLicenseConfigurationOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFailuresForLicenseConfigurationOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFailuresForLicenseConfigurationOperations{}, middleware.After)
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
	addOpListFailuresForLicenseConfigurationOperationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFailuresForLicenseConfigurationOperations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListFailuresForLicenseConfigurationOperations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "ListFailuresForLicenseConfigurationOperations",
	}
}
