// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a resource group with the specified name and description. You can
// optionally include a resource query, or a service configuration.
func (c *Client) CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) {
	if params == nil {
		params = &CreateGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGroup", params, optFns, addOperationCreateGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGroupInput struct {

	// The name of the group, which is the identifier of the group in other operations.
	// You can't change the name of a resource group after you create it. A resource
	// group name can consist of letters, numbers, hyphens, periods, and underscores.
	// The name cannot start with AWS or aws; these are reserved. A resource group name
	// must be unique within each AWS Region in your AWS account.
	//
	// This member is required.
	Name *string

	// A configuration associates the resource group with an AWS service and specifies
	// how the service can interact with the resources in the group. A configuration is
	// an array of GroupConfigurationItem () elements. You can specify either a
	// Configuration or a ResourceQuery in a group, but not both.
	Configuration []*types.GroupConfigurationItem

	// The description of the resource group. Descriptions can consist of letters,
	// numbers, hyphens, underscores, periods, and spaces.
	Description *string

	// The resource query that determines which AWS resources are members of this
	// group. You can specify either a ResourceQuery or a Configuration, but not both.
	ResourceQuery *types.ResourceQuery

	// The tags to add to the group. A tag is key-value pair string.
	Tags map[string]*string
}

type CreateGroupOutput struct {

	// The description of the resource group.
	Group *types.Group

	// The service configuration associated with the resource group. AWS Resource
	// Groups supports adding service configurations for the following resource group
	// types:
	//
	//     * AWS::EC2::CapacityReservationPool - Amazon EC2 capacity
	// reservation pools. For more information, see Working with capacity reservation
	// groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/capacity-reservations-using.html#create-cr-group)
	// in the EC2 Users Guide.
	GroupConfiguration *types.GroupConfiguration

	// The resource query associated with the group.
	ResourceQuery *types.ResourceQuery

	// The tags associated with the group.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGroup{}, middleware.After)
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
	addOpCreateGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "CreateGroup",
	}
}
