// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a progress update stream which is an AWS resource used for access
// control as well as a namespace for migration task names that is implicitly
// linked to your AWS account. It must uniquely identify the migration tool as it
// is used for all updates made by the tool; however, it does not need to be unique
// for each AWS account because it is scoped to the AWS account.
func (c *Client) CreateProgressUpdateStream(ctx context.Context, params *CreateProgressUpdateStreamInput, optFns ...func(*Options)) (*CreateProgressUpdateStreamOutput, error) {
	if params == nil {
		params = &CreateProgressUpdateStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProgressUpdateStream", params, optFns, addOperationCreateProgressUpdateStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProgressUpdateStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProgressUpdateStreamInput struct {

	// The name of the ProgressUpdateStream. Do not store personal data in this field.
	//
	// This member is required.
	ProgressUpdateStreamName *string

	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun *bool
}

type CreateProgressUpdateStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateProgressUpdateStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProgressUpdateStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProgressUpdateStream{}, middleware.After)
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
	addOpCreateProgressUpdateStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProgressUpdateStream(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateProgressUpdateStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "CreateProgressUpdateStream",
	}
}
