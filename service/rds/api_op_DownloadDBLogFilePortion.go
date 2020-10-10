// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Downloads all or a portion of the specified log file, up to 1 MB in size.
func (c *Client) DownloadDBLogFilePortion(ctx context.Context, params *DownloadDBLogFilePortionInput, optFns ...func(*Options)) (*DownloadDBLogFilePortionOutput, error) {
	if params == nil {
		params = &DownloadDBLogFilePortionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DownloadDBLogFilePortion", params, optFns, addOperationDownloadDBLogFilePortionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DownloadDBLogFilePortionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DownloadDBLogFilePortionInput struct {

	// The customer-assigned name of the DB instance that contains the log files you
	// want to list. Constraints:
	//
	//     * Must match the identifier of an existing
	// DBInstance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The name of the log file to be downloaded.
	//
	// This member is required.
	LogFileName *string

	// The pagination token provided in the previous request or "0". If the Marker
	// parameter is specified the response includes only records beyond the marker
	// until the end of the file or up to NumberOfLines.
	Marker *string

	// The number of lines to download. If the number of lines specified results in a
	// file over 1 MB in size, the file is truncated at 1 MB in size. If the
	// NumberOfLines parameter is specified, then the block of lines returned can be
	// from the beginning or the end of the log file, depending on the value of the
	// Marker parameter.
	//
	//     * If neither Marker or NumberOfLines are specified, the
	// entire log file is returned up to a maximum of 10000 lines, starting with the
	// most recent log entries first.
	//
	//     * If NumberOfLines is specified and Marker
	// isn't specified, then the most recent lines from the end of the log file are
	// returned.
	//
	//     * If Marker is specified as "0", then the specified number of
	// lines from the beginning of the log file are returned.
	//
	//     * You can download
	// the log file in blocks of lines by specifying the size of the block using the
	// NumberOfLines parameter, and by specifying a value of "0" for the Marker
	// parameter in your first request. Include the Marker value returned in the
	// response as the Marker value for the next request, continuing until the
	// AdditionalDataPending response element returns false.
	NumberOfLines *int32
}

// This data type is used as a response element to DownloadDBLogFilePortion.
type DownloadDBLogFilePortionOutput struct {

	// Boolean value that if true, indicates there is more data to be downloaded.
	AdditionalDataPending *bool

	// Entries from the specified log file.
	LogFileData *string

	// A pagination token that can be used in a later DownloadDBLogFilePortion request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDownloadDBLogFilePortionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDownloadDBLogFilePortion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDownloadDBLogFilePortion{}, middleware.After)
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
	addOpDownloadDBLogFilePortionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDownloadDBLogFilePortion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDownloadDBLogFilePortion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DownloadDBLogFilePortion",
	}
}
