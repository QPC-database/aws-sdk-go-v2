// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an association between a configuration set and a custom domain for open
// and click event tracking. By default, images and links used for tracking open
// and click events are hosted on domains operated by Amazon SES. You can configure
// a subdomain of your own to handle these events. For information about using
// custom domains, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/configure-custom-open-click-domains.html).
func (c *Client) CreateConfigurationSetTrackingOptions(ctx context.Context, params *CreateConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*CreateConfigurationSetTrackingOptionsOutput, error) {
	if params == nil {
		params = &CreateConfigurationSetTrackingOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfigurationSetTrackingOptions", params, optFns, addOperationCreateConfigurationSetTrackingOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationSetTrackingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to create an open and click tracking option object in a
// configuration set.
type CreateConfigurationSetTrackingOptionsInput struct {

	// The name of the configuration set that the tracking options should be associated
	// with.
	//
	// This member is required.
	ConfigurationSetName *string

	// A domain that is used to redirect email recipients to an Amazon SES-operated
	// domain. This domain captures open and click events generated by Amazon SES
	// emails. For more information, see Configuring Custom Domains to Handle Open and
	// Click Tracking
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/configure-custom-open-click-domains.html)
	// in the Amazon SES Developer Guide.
	//
	// This member is required.
	TrackingOptions *types.TrackingOptions
}

// An empty element returned on a successful request.
type CreateConfigurationSetTrackingOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateConfigurationSetTrackingOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateConfigurationSetTrackingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateConfigurationSetTrackingOptions{}, middleware.After)
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
	addOpCreateConfigurationSetTrackingOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfigurationSetTrackingOptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateConfigurationSetTrackingOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateConfigurationSetTrackingOptions",
	}
}
