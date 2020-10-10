// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the default engine and system parameter information for the specified
// database engine.
func (c *Client) DescribeEngineDefaultParameters(ctx context.Context, params *DescribeEngineDefaultParametersInput, optFns ...func(*Options)) (*DescribeEngineDefaultParametersOutput, error) {
	if params == nil {
		params = &DescribeEngineDefaultParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEngineDefaultParameters", params, optFns, addOperationDescribeEngineDefaultParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEngineDefaultParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEngineDefaultParametersInput struct {

	// The name of the DB parameter group family.
	//
	// This member is required.
	DBParameterGroupFamily *string

	// Not currently supported.
	Filters []*types.Filter

	// An optional pagination token provided by a previous
	// DescribeEngineDefaultParameters request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

type DescribeEngineDefaultParametersOutput struct {

	// Contains the result of a successful invocation of the
	// DescribeEngineDefaultParameters () action.
	EngineDefaults *types.EngineDefaults

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEngineDefaultParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEngineDefaultParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEngineDefaultParameters{}, middleware.After)
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
	addOpDescribeEngineDefaultParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEngineDefaultParameters(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEngineDefaultParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeEngineDefaultParameters",
	}
}
