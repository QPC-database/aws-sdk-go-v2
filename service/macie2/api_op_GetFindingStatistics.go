// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves (queries) aggregated statistical data about findings.
func (c *Client) GetFindingStatistics(ctx context.Context, params *GetFindingStatisticsInput, optFns ...func(*Options)) (*GetFindingStatisticsOutput, error) {
	if params == nil {
		params = &GetFindingStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFindingStatistics", params, optFns, addOperationGetFindingStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFindingStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFindingStatisticsInput struct {

	// The finding property to use to group the query results. Valid values are:
	//
	//     *
	// classificationDetails.jobId - The unique identifier for the classification job
	// that produced the finding.
	//
	//     * resourcesAffected.s3Bucket.name - The name of
	// the S3 bucket that the finding applies to.
	//
	//     * severity.description - The
	// severity of the finding, such as High or Medium.
	//
	//     * type - The type of
	// finding, such as Policy:IAMUser/S3BucketPublic and
	// SensitiveData:S3Object/Personal.
	//
	// This member is required.
	GroupBy types.GroupBy

	// The criteria to use to filter the query results.
	FindingCriteria *types.FindingCriteria

	// The maximum number of items to include in each page of the response.
	Size *int32

	// The criteria to use to sort the query results.
	SortCriteria *types.FindingStatisticsSortCriteria
}

type GetFindingStatisticsOutput struct {

	// An array of objects, one for each group of findings that meet the filter
	// criteria specified in the request.
	CountsByGroup []*types.GroupCount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetFindingStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFindingStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindingStatistics{}, middleware.After)
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
	addOpGetFindingStatisticsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindingStatistics(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetFindingStatistics(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetFindingStatistics",
	}
}
