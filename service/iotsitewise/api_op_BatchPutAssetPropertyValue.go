// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends a list of asset property values to AWS IoT SiteWise. Each value is a
// timestamp-quality-value (TQV) data point. For more information, see Ingesting
// Data Using the API
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/ingest-api.html) in
// the AWS IoT SiteWise User Guide. To identify an asset property, you must specify
// one of the following:
//
//     * The assetId and propertyId of an asset property.
//
//
// * A propertyAlias, which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature). To define an asset property's alias,
// see UpdateAssetProperty
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html).
//
// With
// respect to Unix epoch time, AWS IoT SiteWise accepts only TQVs that have a
// timestamp of no more than 15 minutes in the past and no more than 5 minutes in
// the future. AWS IoT SiteWise rejects timestamps outside of the inclusive range
// of [-15, +5] minutes and returns a TimestampOutOfRangeException error. For each
// asset property, AWS IoT SiteWise overwrites TQVs with duplicate timestamps
// unless the newer TQV has a different quality. For example, if you store a TQV
// {T1, GOOD, V1}, then storing {T1, GOOD, V2} replaces the existing TQV.
func (c *Client) BatchPutAssetPropertyValue(ctx context.Context, params *BatchPutAssetPropertyValueInput, optFns ...func(*Options)) (*BatchPutAssetPropertyValueOutput, error) {
	if params == nil {
		params = &BatchPutAssetPropertyValueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchPutAssetPropertyValue", params, optFns, addOperationBatchPutAssetPropertyValueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchPutAssetPropertyValueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutAssetPropertyValueInput struct {

	// The list of asset property value entries for the batch put request. You can
	// specify up to 10 entries per request.
	//
	// This member is required.
	Entries []*types.PutAssetPropertyValueEntry
}

type BatchPutAssetPropertyValueOutput struct {

	// A list of the errors (if any) associated with the batch put request. Each error
	// entry contains the entryId of the entry that failed.
	//
	// This member is required.
	ErrorEntries []*types.BatchPutAssetPropertyErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchPutAssetPropertyValueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchPutAssetPropertyValue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchPutAssetPropertyValue{}, middleware.After)
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
	addOpBatchPutAssetPropertyValueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutAssetPropertyValue(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchPutAssetPropertyValue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "BatchPutAssetPropertyValue",
	}
}
