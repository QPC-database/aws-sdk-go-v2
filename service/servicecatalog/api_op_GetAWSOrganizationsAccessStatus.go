// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the Access Status for AWS Organization portfolio share feature. This API can
// only be called by the master account in the organization or by a delegated
// admin.
func (c *Client) GetAWSOrganizationsAccessStatus(ctx context.Context, params *GetAWSOrganizationsAccessStatusInput, optFns ...func(*Options)) (*GetAWSOrganizationsAccessStatusOutput, error) {
	if params == nil {
		params = &GetAWSOrganizationsAccessStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAWSOrganizationsAccessStatus", params, optFns, addOperationGetAWSOrganizationsAccessStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAWSOrganizationsAccessStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAWSOrganizationsAccessStatusInput struct {
}

type GetAWSOrganizationsAccessStatusOutput struct {

	// The status of the portfolio share feature.
	AccessStatus types.AccessStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAWSOrganizationsAccessStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAWSOrganizationsAccessStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAWSOrganizationsAccessStatus{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAWSOrganizationsAccessStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAWSOrganizationsAccessStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "GetAWSOrganizationsAccessStatus",
	}
}
