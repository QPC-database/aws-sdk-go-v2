// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the version of the theme that the specified theme alias points to. If
// you provide a specific alias, you delete the version of the theme that the alias
// points to.
func (c *Client) DeleteThemeAlias(ctx context.Context, params *DeleteThemeAliasInput, optFns ...func(*Options)) (*DeleteThemeAliasOutput, error) {
	if params == nil {
		params = &DeleteThemeAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteThemeAlias", params, optFns, addOperationDeleteThemeAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteThemeAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteThemeAliasInput struct {

	// The unique name for the theme alias to delete.
	//
	// This member is required.
	AliasName *string

	// The ID of the AWS account that contains the theme alias to delete.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the theme that the specified alias is for.
	//
	// This member is required.
	ThemeId *string
}

type DeleteThemeAliasOutput struct {

	// The name for the theme alias.
	AliasName *string

	// The Amazon Resource Name (ARN) of the theme resource using the deleted alias.
	Arn *string

	// The AWS request ID for this operation.
	RequestId *string

	// An ID for the theme associated with the deletion.
	ThemeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteThemeAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteThemeAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteThemeAlias{}, middleware.After)
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
	addOpDeleteThemeAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteThemeAlias(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteThemeAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteThemeAlias",
	}
}
