// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the default service quotas values. The Value returned for each quota
// is the AWS default value, even if the quotas have been increased..
func (c *Client) GetAWSDefaultServiceQuota(ctx context.Context, params *GetAWSDefaultServiceQuotaInput, optFns ...func(*Options)) (*GetAWSDefaultServiceQuotaOutput, error) {
	if params == nil {
		params = &GetAWSDefaultServiceQuotaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAWSDefaultServiceQuota", params, optFns, addOperationGetAWSDefaultServiceQuotaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAWSDefaultServiceQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAWSDefaultServiceQuotaInput struct {

	// Identifies the service quota you want to select.
	//
	// This member is required.
	QuotaCode *string

	// Specifies the service that you want to use.
	//
	// This member is required.
	ServiceCode *string
}

type GetAWSDefaultServiceQuotaOutput struct {

	// Returns the ServiceQuota () object which contains all values for a quota.
	Quota *types.ServiceQuota

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAWSDefaultServiceQuotaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAWSDefaultServiceQuota{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAWSDefaultServiceQuota{}, middleware.After)
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
	addOpGetAWSDefaultServiceQuotaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAWSDefaultServiceQuota(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAWSDefaultServiceQuota(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "GetAWSDefaultServiceQuota",
	}
}
