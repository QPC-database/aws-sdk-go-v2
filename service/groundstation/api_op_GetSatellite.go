// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a satellite.
func (c *Client) GetSatellite(ctx context.Context, params *GetSatelliteInput, optFns ...func(*Options)) (*GetSatelliteOutput, error) {
	if params == nil {
		params = &GetSatelliteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSatellite", params, optFns, addOperationGetSatelliteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSatelliteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type GetSatelliteInput struct {

	// UUID of a satellite.
	//
	// This member is required.
	SatelliteId *string
}

//
type GetSatelliteOutput struct {

	// A list of ground stations to which the satellite is on-boarded.
	GroundStations []*string

	// NORAD satellite ID number.
	NoradSatelliteID *int32

	// ARN of a satellite.
	SatelliteArn *string

	// UUID of a satellite.
	SatelliteId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSatelliteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSatellite{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSatellite{}, middleware.After)
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
	addOpGetSatelliteValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
