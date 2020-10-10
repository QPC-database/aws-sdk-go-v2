// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a robot.
func (c *Client) CreateRobot(ctx context.Context, params *CreateRobotInput, optFns ...func(*Options)) (*CreateRobotOutput, error) {
	if params == nil {
		params = &CreateRobotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRobot", params, optFns, addOperationCreateRobotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRobotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRobotInput struct {

	// The target architecture of the robot.
	//
	// This member is required.
	Architecture types.Architecture

	// The Greengrass group id.
	//
	// This member is required.
	GreengrassGroupId *string

	// The name for the robot.
	//
	// This member is required.
	Name *string

	// A map that contains tag keys and tag values that are attached to the robot.
	Tags map[string]*string
}

type CreateRobotOutput struct {

	// The target architecture of the robot.
	Architecture types.Architecture

	// The Amazon Resource Name (ARN) of the robot.
	Arn *string

	// The time, in milliseconds since the epoch, when the robot was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the Greengrass group associated with the
	// robot.
	GreengrassGroupId *string

	// The name of the robot.
	Name *string

	// The list of all tags added to the robot.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRobotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRobot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRobot{}, middleware.After)
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
	addOpCreateRobotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRobot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateRobot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateRobot",
	}
}
