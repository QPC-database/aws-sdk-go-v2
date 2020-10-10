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

// Creates a new Amazon ElastiCache cache parameter group. An ElastiCache cache
// parameter group is a collection of parameters and their values that are applied
// to all of the nodes in any cluster or replication group using the
// CacheParameterGroup. A newly created CacheParameterGroup is an exact duplicate
// of the default parameter group for the CacheParameterGroupFamily. To customize
// the newly created CacheParameterGroup you can change the values of specific
// parameters. For more information, see:
//
//     * ModifyCacheParameterGroup
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheParameterGroup.html)
// in the ElastiCache API Reference.
//
//     * Parameters and Parameter Groups
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/ParameterGroups.html)
// in the ElastiCache User Guide.
func (c *Client) CreateCacheParameterGroup(ctx context.Context, params *CreateCacheParameterGroupInput, optFns ...func(*Options)) (*CreateCacheParameterGroupOutput, error) {
	if params == nil {
		params = &CreateCacheParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCacheParameterGroup", params, optFns, addOperationCreateCacheParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCacheParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCacheParameterGroup operation.
type CreateCacheParameterGroupInput struct {

	// The name of the cache parameter group family that the cache parameter group can
	// be used with. Valid values are: memcached1.4 | memcached1.5 | redis2.6 |
	// redis2.8 | redis3.2 | redis4.0 | redis5.0 |
	//
	// This member is required.
	CacheParameterGroupFamily *string

	// A user-specified name for the cache parameter group.
	//
	// This member is required.
	CacheParameterGroupName *string

	// A user-specified description for the cache parameter group.
	//
	// This member is required.
	Description *string
}

type CreateCacheParameterGroupOutput struct {

	// Represents the output of a CreateCacheParameterGroup operation.
	CacheParameterGroup *types.CacheParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCacheParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateCacheParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCacheParameterGroup{}, middleware.After)
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
	addOpCreateCacheParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCacheParameterGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateCacheParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateCacheParameterGroup",
	}
}
