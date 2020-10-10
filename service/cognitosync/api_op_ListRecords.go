// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets paginated records, optionally changed after a particular sync count for a
// dataset and identity. With Amazon Cognito Sync, each identity has access only to
// its own data. Thus, the credentials used to make this API call need to have
// access to the identity data. ListRecords can be called with temporary user
// credentials provided by Cognito Identity or with developer credentials. You
// should use Cognito Identity credentials to make this API call.
func (c *Client) ListRecords(ctx context.Context, params *ListRecordsInput, optFns ...func(*Options)) (*ListRecordsOutput, error) {
	if params == nil {
		params = &ListRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecords", params, optFns, addOperationListRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for a list of records.
type ListRecordsInput struct {

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	//
	// This member is required.
	DatasetName *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string

	// The last server sync count for this record.
	LastSyncCount *int64

	// The maximum number of results to be returned.
	MaxResults *int32

	// A pagination token for obtaining the next page of results.
	NextToken *string

	// A token containing a session ID, identity ID, and expiration.
	SyncSessionToken *string
}

// Returned for a successful ListRecordsRequest.
type ListRecordsOutput struct {

	// Total number of records.
	Count *int32

	// A boolean value specifying whether to delete the dataset locally.
	DatasetDeletedAfterRequestedSyncCount *bool

	// Indicates whether the dataset exists.
	DatasetExists *bool

	// Server sync count for this dataset.
	DatasetSyncCount *int64

	// The user/device that made the last change to this record.
	LastModifiedBy *string

	// Names of merged datasets.
	MergedDatasetNames []*string

	// A pagination token for obtaining the next page of results.
	NextToken *string

	// A list of all records.
	Records []*types.Record

	// A token containing a session ID, identity ID, and expiration.
	SyncSessionToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecords{}, middleware.After)
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
	addOpListRecordsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRecords(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListRecords(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "ListRecords",
	}
}
