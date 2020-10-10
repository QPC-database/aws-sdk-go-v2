// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an asynchronous job to transcribe speech to text.
func (c *Client) StartTranscriptionJob(ctx context.Context, params *StartTranscriptionJobInput, optFns ...func(*Options)) (*StartTranscriptionJobOutput, error) {
	if params == nil {
		params = &StartTranscriptionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTranscriptionJob", params, optFns, addOperationStartTranscriptionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTranscriptionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTranscriptionJobInput struct {

	// The language code for the language used in the input media file.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// An object that describes the input media for a transcription job.
	//
	// This member is required.
	Media *types.Media

	// The name of the job. Note that you can't use the strings "." or ".." by
	// themselves as the job name. The name must also be unique within an AWS account.
	// If you try to create a transcription job with the same name as a previous
	// transcription job you will receive a ConflictException error.
	//
	// This member is required.
	TranscriptionJobName *string

	// An object that contains the request parameters for content redaction.
	ContentRedaction *types.ContentRedaction

	// Provides information about how a transcription job is executed. Use this field
	// to indicate that the job can be queued for deferred execution if the concurrency
	// limit is reached and there are no slots available to immediately run the job.
	JobExecutionSettings *types.JobExecutionSettings

	// The format of the input media file.
	MediaFormat types.MediaFormat

	// The sample rate, in Hertz, of the audio track in the input media file. If you do
	// not specify the media sample rate, Amazon Transcribe determines the sample rate.
	// If you specify the sample rate, it must match the sample rate detected by Amazon
	// Transcribe. In most cases, you should leave the MediaSampleRateHertz field blank
	// and let Amazon Transcribe determine the sample rate.
	MediaSampleRateHertz *int32

	// The location where the transcription is stored. If you set the OutputBucketName,
	// Amazon Transcribe puts the transcript in the specified S3 bucket. When you call
	// the GetTranscriptionJob () operation, the operation returns this location in the
	// TranscriptFileUri field. If you enable content redaction, the redacted
	// transcript appears in RedactedTranscriptFileUri. If you enable content redaction
	// and choose to output an unredacted transcript, that transcript's location still
	// appears in the TranscriptFileUri. The S3 bucket must have permissions that allow
	// Amazon Transcribe to put files in the bucket. For more information, see
	// Permissions Required for IAM User Roles
	// (https://docs.aws.amazon.com/transcribe/latest/dg/security_iam_id-based-policy-examples.html#auth-role-iam-user).
	// You can specify an AWS Key Management Service (KMS) key to encrypt the output of
	// your transcription using the OutputEncryptionKMSKeyId parameter. If you don't
	// specify a KMS key, Amazon Transcribe uses the default Amazon S3 key for
	// server-side encryption of transcripts that are placed in your S3 bucket. If you
	// don't set the OutputBucketName, Amazon Transcribe generates a pre-signed URL, a
	// shareable URL that provides secure access to your transcription, and returns it
	// in the TranscriptFileUri field. Use this URL to download the transcription.
	OutputBucketName *string

	// The Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key used
	// to encrypt the output of the transcription job. The user calling the
	// StartTranscriptionJob operation must have permission to use the specified KMS
	// key. You can use either of the following to identify a KMS key in the current
	// account:
	//
	//     * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	//     * KMS
	// Key Alias: "alias/ExampleAlias"
	//
	// You can use either of the following to identify
	// a KMS key in the current account or another account:
	//
	//     * Amazon Resource Name
	// (ARN) of a KMS Key: "arn:aws:kms:region:account
	// ID:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	//     * ARN of a KMS Key Alias:
	// "arn:aws:kms:region:account ID:alias/ExampleAlias"
	//
	// If you don't specify an
	// encryption key, the output of the transcription job is encrypted with the
	// default Amazon S3 key (SSE-S3). If you specify a KMS key to encrypt your output,
	// you must also specify an output location in the OutputBucketName parameter.
	OutputEncryptionKMSKeyId *string

	// A Settings object that provides optional settings for a transcription job.
	Settings *types.Settings
}

type StartTranscriptionJobOutput struct {

	// An object containing details of the asynchronous transcription job.
	TranscriptionJob *types.TranscriptionJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartTranscriptionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTranscriptionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTranscriptionJob{}, middleware.After)
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
	addOpStartTranscriptionJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartTranscriptionJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartTranscriptionJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "StartTranscriptionJob",
	}
}
