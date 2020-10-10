// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the celebrity recognition results for a Amazon Rekognition Video analysis
// started by StartCelebrityRecognition (). Celebrity recognition in a video is an
// asynchronous operation. Analysis is started by a call to
// StartCelebrityRecognition () which returns a job identifier (JobId). When the
// celebrity recognition operation finishes, Amazon Rekognition Video publishes a
// completion status to the Amazon Simple Notification Service topic registered in
// the initial call to StartCelebrityRecognition. To get the results of the
// celebrity recognition analysis, first check that the status value published to
// the Amazon SNS topic is SUCCEEDED. If so, call GetCelebrityDetection and pass
// the job identifier (JobId) from the initial call to StartCelebrityDetection.
// <p>For more information, see Working With Stored Videos in the Amazon
// Rekognition Developer Guide.</p> <p> <code>GetCelebrityRecognition</code>
// returns detected celebrities and the time(s) they are detected in an array
// (<code>Celebrities</code>) of <a>CelebrityRecognition</a> objects. Each
// <code>CelebrityRecognition</code> contains information about the celebrity in a
// <a>CelebrityDetail</a> object and the time, <code>Timestamp</code>, the
// celebrity was detected. </p> <note> <p> <code>GetCelebrityRecognition</code>
// only returns the default facial attributes (<code>BoundingBox</code>,
// <code>Confidence</code>, <code>Landmarks</code>, <code>Pose</code>, and
// <code>Quality</code>). The other facial attributes listed in the
// <code>Face</code> object of the following response syntax are not returned. For
// more information, see FaceDetail in the Amazon Rekognition Developer Guide. </p>
// </note> <p>By default, the <code>Celebrities</code> array is sorted by time
// (milliseconds from the start of the video). You can also sort the array by
// celebrity by specifying the value <code>ID</code> in the <code>SortBy</code>
// input parameter.</p> <p>The <code>CelebrityDetail</code> object includes the
// celebrity identifer and additional information urls. If you don't store the
// additional information urls, you can get them later by calling
// <a>GetCelebrityInfo</a> with the celebrity identifer.</p> <p>No information is
// returned for faces not recognized as celebrities.</p> <p>Use MaxResults
// parameter to limit the number of labels returned. If there are more results than
// specified in <code>MaxResults</code>, the value of <code>NextToken</code> in the
// operation response contains a pagination token for getting the next set of
// results. To get the next page of results, call
// <code>GetCelebrityDetection</code> and populate the <code>NextToken</code>
// request parameter with the token value returned from the previous call to
// <code>GetCelebrityRecognition</code>.</p>
func (c *Client) GetCelebrityRecognition(ctx context.Context, params *GetCelebrityRecognitionInput, optFns ...func(*Options)) (*GetCelebrityRecognitionOutput, error) {
	if params == nil {
		params = &GetCelebrityRecognitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCelebrityRecognition", params, optFns, addOperationGetCelebrityRecognitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCelebrityRecognitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCelebrityRecognitionInput struct {

	// Job identifier for the required celebrity recognition analysis. You can get the
	// job identifer from a call to StartCelebrityRecognition.
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there is more recognized
	// celebrities to retrieve), Amazon Rekognition Video returns a pagination token in
	// the response. You can use this pagination token to retrieve the next set of
	// celebrities.
	NextToken *string

	// Sort to use for celebrities returned in Celebrities field. Specify ID to sort by
	// the celebrity identifier, specify TIMESTAMP to sort by the time the celebrity
	// was recognized.
	SortBy types.CelebrityRecognitionSortBy
}

type GetCelebrityRecognitionOutput struct {

	// Array of celebrities recognized in the video.
	Celebrities []*types.CelebrityRecognition

	// The current status of the celebrity recognition job.
	JobStatus types.VideoJobStatus

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of celebrities.
	NextToken *string

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition Video
	// operation.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCelebrityRecognitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCelebrityRecognition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCelebrityRecognition{}, middleware.After)
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
	addOpGetCelebrityRecognitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCelebrityRecognition(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCelebrityRecognition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetCelebrityRecognition",
	}
}
