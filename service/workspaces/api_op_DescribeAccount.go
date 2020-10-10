// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list that describes the configuration of Bring Your Own License
// (BYOL) for the specified account.
func (c *Client) DescribeAccount(ctx context.Context, params *DescribeAccountInput, optFns ...func(*Options)) (*DescribeAccountOutput, error) {
	if params == nil {
		params = &DescribeAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccount", params, optFns, addOperationDescribeAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountInput struct {
}

type DescribeAccountOutput struct {

	// The IP address range, specified as an IPv4 CIDR block, used for the management
	// network interface. The management network interface is connected to a secure
	// Amazon WorkSpaces management network. It is used for interactive streaming of
	// the WorkSpace desktop to Amazon WorkSpaces clients, and to allow Amazon
	// WorkSpaces to manage the WorkSpace.
	DedicatedTenancyManagementCidrRange *string

	// The status of BYOL (whether BYOL is enabled or disabled).
	DedicatedTenancySupport types.DedicatedTenancySupportResultEnum

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAccount{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeAccount",
	}
}
