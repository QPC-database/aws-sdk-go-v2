// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists schema major versions applied to a directory. If SchemaArn is provided,
// lists the minor version.
func (c *Client) ListAppliedSchemaArns(ctx context.Context, params *ListAppliedSchemaArnsInput, optFns ...func(*Options)) (*ListAppliedSchemaArnsOutput, error) {
	if params == nil {
		params = &ListAppliedSchemaArnsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppliedSchemaArns", params, optFns, addOperationListAppliedSchemaArnsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppliedSchemaArnsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppliedSchemaArnsInput struct {

	// The ARN of the directory you are listing.
	//
	// This member is required.
	DirectoryArn *string

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	// The response for ListAppliedSchemaArns when this parameter is used will list all
	// minor version ARNs for a major version.
	SchemaArn *string
}

type ListAppliedSchemaArnsOutput struct {

	// The pagination token.
	NextToken *string

	// The ARNs of schemas that are applied to the directory.
	SchemaArns []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAppliedSchemaArnsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppliedSchemaArns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppliedSchemaArns{}, middleware.After)
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
	addOpListAppliedSchemaArnsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAppliedSchemaArns(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListAppliedSchemaArns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListAppliedSchemaArns",
	}
}
