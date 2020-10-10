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

// Adds tags to on-premises instances.
func (c *Client) AddTagsToOnPremisesInstances(ctx context.Context, params *AddTagsToOnPremisesInstancesInput, optFns ...func(*Options)) (*AddTagsToOnPremisesInstancesOutput, error) {
	if params == nil {
		params = &AddTagsToOnPremisesInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddTagsToOnPremisesInstances", params, optFns, addOperationAddTagsToOnPremisesInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddTagsToOnPremisesInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of, and adds tags to, an on-premises instance operation.
type AddTagsToOnPremisesInstancesInput struct {

	// The names of the on-premises instances to which to add tags.
	//
	// This member is required.
	InstanceNames []*string

	// The tag key-value pairs to add to the on-premises instances. Keys and values are
	// both required. Keys cannot be null or empty strings. Value-only tags are not
	// allowed.
	//
	// This member is required.
	Tags []*types.Tag
}

type AddTagsToOnPremisesInstancesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddTagsToOnPremisesInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddTagsToOnPremisesInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddTagsToOnPremisesInstances{}, middleware.After)
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
	addOpAddTagsToOnPremisesInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddTagsToOnPremisesInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAddTagsToOnPremisesInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "AddTagsToOnPremisesInstances",
	}
}
