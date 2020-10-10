// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates stack instances for the specified accounts, within the specified
// Regions. A stack instance refers to a stack in a specific account and Region.
// You must specify at least one value for either Accounts or DeploymentTargets,
// and you must specify at least one value for Regions.
func (c *Client) CreateStackInstances(ctx context.Context, params *CreateStackInstancesInput, optFns ...func(*Options)) (*CreateStackInstancesOutput, error) {
	if params == nil {
		params = &CreateStackInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStackInstances", params, optFns, addOperationCreateStackInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStackInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStackInstancesInput struct {

	// The names of one or more Regions where you want to create stack instances using
	// the specified AWS account(s).
	//
	// This member is required.
	Regions []*string

	// The name or unique ID of the stack set that you want to create stack instances
	// from.
	//
	// This member is required.
	StackSetName *string

	// [Self-managed permissions] The names of one or more AWS accounts that you want
	// to create stack instances in the specified Region(s) for. You can specify
	// Accounts or DeploymentTargets, but not both.
	Accounts []*string

	// [Service-managed permissions] The AWS Organizations accounts for which to create
	// stack instances in the specified Regions. You can specify Accounts or
	// DeploymentTargets, but not both.
	DeploymentTargets *types.DeploymentTargets

	// The unique identifier for this stack set operation. The operation ID also
	// functions as an idempotency token, to ensure that AWS CloudFormation performs
	// the stack set operation only once, even if you retry the request multiple times.
	// You might retry stack set operation requests to ensure that AWS CloudFormation
	// successfully received them. If you don't specify an operation ID, the SDK
	// generates one automatically. Repeating this stack set operation with a new
	// operation ID retries all stack instances whose status is OUTDATED.
	OperationId *string

	// Preferences for how AWS CloudFormation performs this stack set operation.
	OperationPreferences *types.StackSetOperationPreferences

	// A list of stack set parameters whose values you want to override in the selected
	// stack instances. Any overridden parameter values will be applied to all stack
	// instances in the specified accounts and Regions. When specifying parameters and
	// their values, be aware of how AWS CloudFormation sets parameter values during
	// stack instance operations:
	//
	//     * To override the current value for a parameter,
	// include the parameter and specify its value.
	//
	//     * To leave a parameter set to
	// its present value, you can do one of the following:
	//
	//         * Do not include
	// the parameter in the list.
	//
	//         * Include the parameter and specify
	// UsePreviousValue as true. (You cannot specify both a value and set
	// UsePreviousValue to true.)
	//
	//     * To set all overridden parameter back to the
	// values specified in the stack set, specify a parameter list but do not include
	// any parameters.
	//
	//     * To leave all parameters set to their present values, do
	// not specify this property at all.
	//
	// During stack set updates, any parameter
	// values overridden for a stack instance are not updated, but retain their
	// overridden value. You can only override the parameter values that are specified
	// in the stack set; to add or delete a parameter itself, use UpdateStackSet
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
	// to update the stack set template.
	ParameterOverrides []*types.Parameter
}

type CreateStackInstancesOutput struct {

	// The unique identifier for this stack set operation.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateStackInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateStackInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateStackInstances{}, middleware.After)
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
	addIdempotencyToken_opCreateStackInstancesMiddleware(stack, options)
	addOpCreateStackInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStackInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateStackInstances struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateStackInstances) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateStackInstances) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateStackInstancesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateStackInstancesInput ")
	}

	if input.OperationId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.OperationId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateStackInstancesMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateStackInstances{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateStackInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "CreateStackInstances",
	}
}
