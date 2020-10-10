// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about an on-premises instance.
func (c *Client) GetOnPremisesInstance(ctx context.Context, params *GetOnPremisesInstanceInput, optFns ...func(*Options)) (*GetOnPremisesInstanceOutput, error) {
	if params == nil {
		params = &GetOnPremisesInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOnPremisesInstance", params, optFns, addOperationGetOnPremisesInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOnPremisesInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetOnPremisesInstance operation.
type GetOnPremisesInstanceInput struct {

	// The name of the on-premises instance about which to get information.
	//
	// This member is required.
	InstanceName *string
}

// Represents the output of a GetOnPremisesInstance operation.
type GetOnPremisesInstanceOutput struct {

	// Information about the on-premises instance.
	InstanceInfo *types.InstanceInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetOnPremisesInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOnPremisesInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOnPremisesInstance{}, middleware.After)
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
	addOpGetOnPremisesInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOnPremisesInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetOnPremisesInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "GetOnPremisesInstance",
	}
}
