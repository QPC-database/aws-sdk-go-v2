// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes read and write permissions on a template.
func (c *Client) DescribeTemplatePermissions(ctx context.Context, params *DescribeTemplatePermissionsInput, optFns ...func(*Options)) (*DescribeTemplatePermissionsOutput, error) {
	if params == nil {
		params = &DescribeTemplatePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTemplatePermissions", params, optFns, addOperationDescribeTemplatePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTemplatePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTemplatePermissionsInput struct {

	// The ID of the AWS account that contains the template that you're describing.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the template.
	//
	// This member is required.
	TemplateId *string
}

type DescribeTemplatePermissionsOutput struct {

	// A list of resource permissions to be set on the template.
	Permissions []*types.ResourcePermission

	// The AWS request ID for this operation.
	RequestId *string

	// The Amazon Resource Name (ARN) of the template.
	TemplateArn *string

	// The ID for the template.
	TemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTemplatePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTemplatePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTemplatePermissions{}, middleware.After)
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
	addOpDescribeTemplatePermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTemplatePermissions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTemplatePermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeTemplatePermissions",
	}
}
