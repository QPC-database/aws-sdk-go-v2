// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all field-level encryption configurations that have been created in
// CloudFront for this account.
func (c *Client) ListFieldLevelEncryptionConfigs(ctx context.Context, params *ListFieldLevelEncryptionConfigsInput, optFns ...func(*Options)) (*ListFieldLevelEncryptionConfigsOutput, error) {
	if params == nil {
		params = &ListFieldLevelEncryptionConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFieldLevelEncryptionConfigs", params, optFns, addOperationListFieldLevelEncryptionConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFieldLevelEncryptionConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFieldLevelEncryptionConfigsInput struct {

	// Use this when paginating results to indicate where to begin in your list of
	// configurations. The results include configurations in the list that occur after
	// the marker. To get the next page of results, set the Marker to the value of the
	// NextMarker from the current page's response (which is also the ID of the last
	// configuration on that page).
	Marker *string

	// The maximum number of field-level encryption configurations you want in the
	// response body.
	MaxItems *string
}

type ListFieldLevelEncryptionConfigsOutput struct {

	// Returns a list of all field-level encryption configurations that have been
	// created in CloudFront for this account.
	FieldLevelEncryptionList *types.FieldLevelEncryptionList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFieldLevelEncryptionConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListFieldLevelEncryptionConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListFieldLevelEncryptionConfigs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFieldLevelEncryptionConfigs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListFieldLevelEncryptionConfigs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListFieldLevelEncryptionConfigs",
	}
}
