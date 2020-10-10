// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the AWS Organizations master account that is associated with AWS
// Firewall Manager as the AWS Firewall Manager administrator.
func (c *Client) GetAdminAccount(ctx context.Context, params *GetAdminAccountInput, optFns ...func(*Options)) (*GetAdminAccountOutput, error) {
	if params == nil {
		params = &GetAdminAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAdminAccount", params, optFns, addOperationGetAdminAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAdminAccountInput struct {
}

type GetAdminAccountOutput struct {

	// The AWS account that is set as the AWS Firewall Manager administrator.
	AdminAccount *string

	// The status of the AWS account that you set as the AWS Firewall Manager
	// administrator.
	RoleStatus types.AccountRoleStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAdminAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAdminAccount{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAdminAccount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAdminAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "GetAdminAccount",
	}
}
