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

// Returns a list of the available DB engines.
func (c *Client) DescribeDBEngineVersions(ctx context.Context, params *DescribeDBEngineVersionsInput, optFns ...func(*Options)) (*DescribeDBEngineVersionsOutput, error) {
	if params == nil {
		params = &DescribeDBEngineVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBEngineVersions", params, optFns, addOperationDescribeDBEngineVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBEngineVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBEngineVersionsInput struct {

	// The name of a specific DB parameter group family to return details for.
	// Constraints:
	//
	//     * If supplied, must match an existing DBParameterGroupFamily.
	DBParameterGroupFamily *string

	// Indicates that only the default version of the specified engine or engine and
	// major version combination is returned.
	DefaultOnly *bool

	// The database engine to return.
	Engine *string

	// The database engine version to return. Example: 5.1.49
	EngineVersion *string

	// Not currently supported.
	Filters []*types.Filter

	// If this parameter is specified and the requested engine supports the
	// CharacterSetName parameter for CreateDBInstance, the response includes a list of
	// supported character sets for each engine version.
	ListSupportedCharacterSets *bool

	// If this parameter is specified and the requested engine supports the TimeZone
	// parameter for CreateDBInstance, the response includes a list of supported time
	// zones for each engine version.
	ListSupportedTimezones *bool

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included in
	// the response so that the following results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

type DescribeDBEngineVersionsOutput struct {

	// A list of DBEngineVersion elements.
	DBEngineVersions []*types.DBEngineVersion

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBEngineVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBEngineVersions{}, middleware.After)
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
	addOpDescribeDBEngineVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBEngineVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDBEngineVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBEngineVersions",
	}
}
