// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a trail. This operation must be called from the region in which the
// trail was created. DeleteTrail cannot be called on the shadow trails (replicated
// trails in other regions) of a trail that is enabled in all regions.
func (c *Client) DeleteTrail(ctx context.Context, params *DeleteTrailInput, optFns ...func(*Options)) (*DeleteTrailOutput, error) {
	if params == nil {
		params = &DeleteTrailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTrail", params, optFns, addOperationDeleteTrailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTrailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request that specifies the name of a trail to delete.
type DeleteTrailInput struct {

	// Specifies the name or the CloudTrail ARN of the trail to be deleted. The format
	// of a trail ARN is: arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// This member is required.
	Name *string
}

// Returns the objects or data listed below if successful. Otherwise, returns an
// error.
type DeleteTrailOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTrailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTrail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTrail{}, middleware.After)
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
	addOpDeleteTrailValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTrail(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteTrail(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "DeleteTrail",
	}
}
