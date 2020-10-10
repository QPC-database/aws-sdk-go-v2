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

// Updates the criteria and other settings for a findings filter.
func (c *Client) UpdateFindingsFilter(ctx context.Context, params *UpdateFindingsFilterInput, optFns ...func(*Options)) (*UpdateFindingsFilterOutput, error) {
	if params == nil {
		params = &UpdateFindingsFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFindingsFilter", params, optFns, addOperationUpdateFindingsFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFindingsFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFindingsFilterInput struct {

	// The unique identifier for the Amazon Macie resource or account that the request
	// applies to.
	//
	// This member is required.
	Id *string

	// The action to perform on findings that meet the filter criteria
	// (findingCriteria). Valid values are: ARCHIVE, suppress (automatically archive)
	// the findings; and, NOOP, don't perform any action on the findings.
	Action types.FindingsFilterAction

	// A custom description of the filter. The description can contain as many as 512
	// characters. We strongly recommend that you avoid including any sensitive data in
	// the description of a filter. Other users might be able to see the filter's
	// description, depending on the actions that they're allowed to perform in Amazon
	// Macie.
	Description *string

	// The criteria to use to filter findings.
	FindingCriteria *types.FindingCriteria

	// A custom name for the filter. The name must contain at least 3 characters and
	// can contain as many as 64 characters. We strongly recommend that you avoid
	// including any sensitive data in the name of a filter. Other users might be able
	// to see the filter's name, depending on the actions that they're allowed to
	// perform in Amazon Macie.
	Name *string

	// The position of the filter in the list of saved filters on the Amazon Macie
	// console. This value also determines the order in which the filter is applied to
	// findings, relative to other filters that are also applied to the findings.
	Position *int32
}

type UpdateFindingsFilterOutput struct {

	// The Amazon Resource Name (ARN) of the filter that was updated.
	Arn *string

	// The unique identifier for the filter that was updated.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateFindingsFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFindingsFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFindingsFilter{}, middleware.After)
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
	addOpUpdateFindingsFilterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFindingsFilter(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateFindingsFilter(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "UpdateFindingsFilter",
	}
}
