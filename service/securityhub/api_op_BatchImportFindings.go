// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Imports security findings generated from an integrated third-party product into
// Security Hub. This action is requested by the integrated product to import its
// findings into Security Hub. The maximum allowed size for a finding is 240 Kb. An
// error is returned for any finding larger than 240 Kb. After a finding is
// created, BatchImportFindings cannot be used to update the following finding
// fields and objects, which Security Hub customers use to manage their
// investigation workflow.
//
//     * Confidence
//
//     * Criticality
//
//     * Note
//
//     *
// RelatedFindings
//
//     * Severity
//
//     * Types
//
//     * UserDefinedFields
//
//     *
// VerificationState
//
//     * Workflow
func (c *Client) BatchImportFindings(ctx context.Context, params *BatchImportFindingsInput, optFns ...func(*Options)) (*BatchImportFindingsOutput, error) {
	if params == nil {
		params = &BatchImportFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchImportFindings", params, optFns, addOperationBatchImportFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchImportFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchImportFindingsInput struct {

	// A list of findings to import. To successfully import a finding, it must follow
	// the AWS Security Finding Format
	// (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format.html).
	// Maximum of 100 findings per request.
	//
	// This member is required.
	Findings []*types.AwsSecurityFinding
}

type BatchImportFindingsOutput struct {

	// The number of findings that failed to import.
	//
	// This member is required.
	FailedCount *int32

	// The number of findings that were successfully imported.
	//
	// This member is required.
	SuccessCount *int32

	// The list of findings that failed to import.
	FailedFindings []*types.ImportFindingsError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchImportFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchImportFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchImportFindings{}, middleware.After)
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
	addOpBatchImportFindingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchImportFindings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchImportFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchImportFindings",
	}
}
