// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes Challenge-Handshake Authentication Protocol (CHAP) credentials for a
// specified iSCSI target and initiator pair. This operation is supported in volume
// and tape gateway types.
func (c *Client) DeleteChapCredentials(ctx context.Context, params *DeleteChapCredentialsInput, optFns ...func(*Options)) (*DeleteChapCredentialsOutput, error) {
	if params == nil {
		params = &DeleteChapCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteChapCredentials", params, optFns, addOperationDeleteChapCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteChapCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:  <ul> <li> <p>
// <a>DeleteChapCredentialsInput$InitiatorName</a> </p> </li> <li> <p>
// <a>DeleteChapCredentialsInput$TargetARN</a> </p> </li> </ul>
type DeleteChapCredentialsInput struct {

	// The iSCSI initiator that connects to the target.
	//
	// This member is required.
	InitiatorName *string

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the
	// DescribeStorediSCSIVolumes () operation to return to retrieve the TargetARN for
	// specified VolumeARN.
	//
	// This member is required.
	TargetARN *string
}

// A JSON object containing the following fields:
type DeleteChapCredentialsOutput struct {

	// The iSCSI initiator that connects to the target.
	InitiatorName *string

	// The Amazon Resource Name (ARN) of the target.
	TargetARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteChapCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteChapCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteChapCredentials{}, middleware.After)
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
	addOpDeleteChapCredentialsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteChapCredentials(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteChapCredentials(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DeleteChapCredentials",
	}
}
