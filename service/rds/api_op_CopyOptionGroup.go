// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Copies the specified option group.
func (c *Client) CopyOptionGroup(ctx context.Context, params *CopyOptionGroupInput, optFns ...func(*Options)) (*CopyOptionGroupOutput, error) {
	if params == nil {
		params = &CopyOptionGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyOptionGroup", params, optFns, addOperationCopyOptionGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyOptionGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CopyOptionGroupInput struct {

	// The identifier or ARN for the source option group. For information about
	// creating an ARN, see  Constructing an ARN for Amazon RDS
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)
	// in the Amazon RDS User Guide. Constraints:
	//
	//     * Must specify a valid option
	// group.
	//
	//     * If the source option group is in the same AWS Region as the copy,
	// specify a valid option group identifier, for example my-option-group, or a valid
	// ARN.
	//
	//     * If the source option group is in a different AWS Region than the
	// copy, specify a valid option group ARN, for example
	// arn:aws:rds:us-west-2:123456789012:og:special-options.
	//
	// This member is required.
	SourceOptionGroupIdentifier *string

	// The description for the copied option group.
	//
	// This member is required.
	TargetOptionGroupDescription *string

	// The identifier for the copied option group. Constraints:
	//
	//     * Can't be null,
	// empty, or blank
	//
	//     * Must contain from 1 to 255 letters, numbers, or hyphens
	//
	//
	// * First character must be a letter
	//
	//     * Can't end with a hyphen or contain two
	// consecutive hyphens
	//
	// Example: my-option-group
	//
	// This member is required.
	TargetOptionGroupIdentifier *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []*types.Tag
}

type CopyOptionGroupOutput struct {

	//
	OptionGroup *types.OptionGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCopyOptionGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyOptionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyOptionGroup{}, middleware.After)
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
	addOpCopyOptionGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyOptionGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCopyOptionGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CopyOptionGroup",
	}
}
