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

// Lists all default service quotas for the specified AWS service or all AWS
// services. ListAWSDefaultServiceQuotas is similar to ListServiceQuotas () except
// for the Value object. The Value object returned by ListAWSDefaultServiceQuotas
// is the default value assigned by AWS. This request returns a list of all service
// quotas for the specified service. The listing of each you'll see the default
// values are the values that AWS provides for the quotas. Always check the
// NextToken response parameter when calling any of the List* operations. These
// operations can return an unexpected list of results, even when there are more
// results available. When this happens, the NextToken response parameter contains
// a value to pass the next call to the same API to request the next part of the
// list.
func (c *Client) ListAWSDefaultServiceQuotas(ctx context.Context, params *ListAWSDefaultServiceQuotasInput, optFns ...func(*Options)) (*ListAWSDefaultServiceQuotasOutput, error) {
	if params == nil {
		params = &ListAWSDefaultServiceQuotasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAWSDefaultServiceQuotas", params, optFns, addOperationListAWSDefaultServiceQuotasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAWSDefaultServiceQuotasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAWSDefaultServiceQuotasInput struct {

	// Specifies the service that you want to use.
	//
	// This member is required.
	ServiceCode *string

	// (Optional) Limits the number of results that you want to include in the
	// response. If you don't include this parameter, the response defaults to a value
	// that's specific to the operation. If additional items exist beyond the specified
	// maximum, the NextToken element is present and has a value (isn't null). Include
	// that value as the NextToken request parameter in the call to the operation to
	// get the next part of the results. You should check NextToken after every
	// operation to ensure that you receive all of the results.
	MaxResults *int32

	// (Optional) Use this parameter in a request if you receive a NextToken response
	// in a previous request that indicates that there's more output available. In a
	// subsequent call, set it to the value of the previous call's NextToken response
	// to indicate where the output should continue from. If additional items exist
	// beyond the specified maximum, the NextToken element is present and has a value
	// (isn't null). Include that value as the NextToken request parameter in the call
	// to the operation to get the next part of the results. You should check NextToken
	// after every operation to ensure that you receive all of the results.
	NextToken *string
}

type ListAWSDefaultServiceQuotasOutput struct {

	// (Optional) Use this parameter in a request if you receive a NextToken response
	// in a previous request that indicates that there's more output available. In a
	// subsequent call, set it to the value of the previous call's NextToken response
	// to indicate where the output should continue from.
	NextToken *string

	// A list of the quotas in the account with the AWS default values.
	Quotas []*types.ServiceQuota

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAWSDefaultServiceQuotasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAWSDefaultServiceQuotas{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAWSDefaultServiceQuotas{}, middleware.After)
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
	addOpListAWSDefaultServiceQuotasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAWSDefaultServiceQuotas(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListAWSDefaultServiceQuotas(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListAWSDefaultServiceQuotas",
	}
}
