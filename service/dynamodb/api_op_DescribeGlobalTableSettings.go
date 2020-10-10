// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes Region-specific settings for a global table. This operation only
// applies to Version 2017.11.29
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html)
// of global tables.
func (c *Client) DescribeGlobalTableSettings(ctx context.Context, params *DescribeGlobalTableSettingsInput, optFns ...func(*Options)) (*DescribeGlobalTableSettingsOutput, error) {
	if params == nil {
		params = &DescribeGlobalTableSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGlobalTableSettings", params, optFns, addOperationDescribeGlobalTableSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGlobalTableSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGlobalTableSettingsInput struct {

	// The name of the global table to describe.
	//
	// This member is required.
	GlobalTableName *string
}

type DescribeGlobalTableSettingsOutput struct {

	// The name of the global table.
	GlobalTableName *string

	// The Region-specific settings for the global table.
	ReplicaSettings []*types.ReplicaSettingsDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeGlobalTableSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeGlobalTableSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeGlobalTableSettings{}, middleware.After)
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
	addOpDescribeGlobalTableSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGlobalTableSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)
	return nil
}

func newServiceMetadataMiddleware_opDescribeGlobalTableSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "DescribeGlobalTableSettings",
	}
}
