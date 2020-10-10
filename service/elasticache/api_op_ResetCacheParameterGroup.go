// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the parameters of a cache parameter group to the engine or system
// default value. You can reset specific parameters by submitting a list of
// parameter names. To reset the entire cache parameter group, specify the
// ResetAllParameters and CacheParameterGroupName parameters.
func (c *Client) ResetCacheParameterGroup(ctx context.Context, params *ResetCacheParameterGroupInput, optFns ...func(*Options)) (*ResetCacheParameterGroupOutput, error) {
	if params == nil {
		params = &ResetCacheParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetCacheParameterGroup", params, optFns, addOperationResetCacheParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetCacheParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ResetCacheParameterGroup operation.
type ResetCacheParameterGroupInput struct {

	// The name of the cache parameter group to reset.
	//
	// This member is required.
	CacheParameterGroupName *string

	// An array of parameter names to reset to their default values. If
	// ResetAllParameters is true, do not use ParameterNameValues. If
	// ResetAllParameters is false, you must specify the name of at least one parameter
	// to reset.
	ParameterNameValues []*types.ParameterNameValue

	// If true, all parameters in the cache parameter group are reset to their default
	// values. If false, only the parameters listed by ParameterNameValues are reset to
	// their default values. Valid values: true | false
	ResetAllParameters *bool
}

// Represents the output of one of the following operations:
//
//     *
// ModifyCacheParameterGroup
//
//     * ResetCacheParameterGroup
type ResetCacheParameterGroupOutput struct {

	// The name of the cache parameter group.
	CacheParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResetCacheParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpResetCacheParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpResetCacheParameterGroup{}, middleware.After)
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
	addOpResetCacheParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetCacheParameterGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opResetCacheParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ResetCacheParameterGroup",
	}
}
