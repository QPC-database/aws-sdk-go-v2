// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all of the configuration sets associated with your Amazon Pinpoint account
// in the current region.
func (c *Client) ListConfigurationSets(ctx context.Context, params *ListConfigurationSetsInput, optFns ...func(*Options)) (*ListConfigurationSetsOutput, error) {
	if params == nil {
		params = &ListConfigurationSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurationSets", params, optFns, addOperationListConfigurationSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationSetsInput struct {

	// A token returned from a previous call to the API that indicates the position in
	// the list of results.
	NextToken *string

	// Used to specify the number of items that should be returned in the response.
	PageSize *string
}

// An object that contains information about the configuration sets for your
// account in the current region.
type ListConfigurationSetsOutput struct {

	// An object that contains a list of configuration sets for your account in the
	// current region.
	ConfigurationSets []*string

	// A token returned from a previous call to ListConfigurationSets to indicate the
	// position in the list of configuration sets.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListConfigurationSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfigurationSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfigurationSets{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurationSets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListConfigurationSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "ListConfigurationSets",
	}
}
