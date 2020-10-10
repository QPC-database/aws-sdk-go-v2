// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a default subnet with a size /20 IPv4 CIDR block in the specified
// Availability Zone in your default VPC. You can have only one default subnet per
// Availability Zone. For more information, see Creating a Default Subnet
// (https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html#create-default-subnet)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) CreateDefaultSubnet(ctx context.Context, params *CreateDefaultSubnetInput, optFns ...func(*Options)) (*CreateDefaultSubnetOutput, error) {
	if params == nil {
		params = &CreateDefaultSubnetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDefaultSubnet", params, optFns, addOperationCreateDefaultSubnetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDefaultSubnetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDefaultSubnetInput struct {

	// The Availability Zone in which to create the default subnet.
	//
	// This member is required.
	AvailabilityZone *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type CreateDefaultSubnetOutput struct {

	// Information about the subnet.
	Subnet *types.Subnet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDefaultSubnetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateDefaultSubnet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateDefaultSubnet{}, middleware.After)
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
	addOpCreateDefaultSubnetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDefaultSubnet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDefaultSubnet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateDefaultSubnet",
	}
}
