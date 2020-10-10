// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the list of available instance images, or blueprints. You can use a
// blueprint to create a new instance already running a specific operating system,
// as well as a preinstalled app or development stack. The software each instance
// is running depends on the blueprint image you choose. Use active blueprints when
// creating new instances. Inactive blueprints are listed to support customers with
// existing instances and are not necessarily available to create new instances.
// Blueprints are marked inactive when they become outdated due to operating system
// updates or new application releases.
func (c *Client) GetBlueprints(ctx context.Context, params *GetBlueprintsInput, optFns ...func(*Options)) (*GetBlueprintsOutput, error) {
	if params == nil {
		params = &GetBlueprintsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBlueprints", params, optFns, addOperationGetBlueprintsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBlueprintsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBlueprintsInput struct {

	// A Boolean value indicating whether to include inactive results in your request.
	IncludeInactive *bool

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetBlueprints request. If your results are
	// paginated, the response will return a next page token that you can specify as
	// the page token in a subsequent request.
	PageToken *string
}

type GetBlueprintsOutput struct {

	// An array of key-value pairs that contains information about the available
	// blueprints.
	Blueprints []*types.Blueprint

	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetBlueprints request and specify the next page
	// token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBlueprintsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetBlueprints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetBlueprints{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBlueprints(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBlueprints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetBlueprints",
	}
}
