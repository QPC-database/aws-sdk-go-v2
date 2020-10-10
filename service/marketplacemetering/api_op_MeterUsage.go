// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// API to emit metering records. For identical requests, the API is idempotent. It
// simply returns the metering record ID. MeterUsage is authenticated on the
// buyer's AWS account using credentials from the EC2 instance, ECS task, or EKS
// pod.
func (c *Client) MeterUsage(ctx context.Context, params *MeterUsageInput, optFns ...func(*Options)) (*MeterUsageOutput, error) {
	if params == nil {
		params = &MeterUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MeterUsage", params, optFns, addOperationMeterUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MeterUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MeterUsageInput struct {

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of a new
	// product.
	//
	// This member is required.
	ProductCode *string

	// Timestamp, in UTC, for which the usage is being reported. Your application can
	// meter usage for up to one hour in the past. Make sure the timestamp value is not
	// before the start of the software usage.
	//
	// This member is required.
	Timestamp *time.Time

	// It will be one of the fcp dimension name provided during the publishing of the
	// product.
	//
	// This member is required.
	UsageDimension *string

	// Checks whether you have the permissions required for the action, but does not
	// make the request. If you have the permissions, the request returns
	// DryRunOperation; otherwise, it returns UnauthorizedException. Defaults to false
	// if not specified.
	DryRun *bool

	// Consumption value for the hour. Defaults to 0 if not specified.
	UsageQuantity *int32
}

type MeterUsageOutput struct {

	// Metering record id.
	MeteringRecordId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationMeterUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpMeterUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpMeterUsage{}, middleware.After)
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
	addOpMeterUsageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opMeterUsage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opMeterUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "MeterUsage",
	}
}
