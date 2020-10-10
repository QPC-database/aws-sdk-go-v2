// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a profile that can be applied to one or more private fleet device
// instances.
func (c *Client) CreateInstanceProfile(ctx context.Context, params *CreateInstanceProfileInput, optFns ...func(*Options)) (*CreateInstanceProfileOutput, error) {
	if params == nil {
		params = &CreateInstanceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstanceProfile", params, optFns, addOperationCreateInstanceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstanceProfileInput struct {

	// The name of your instance profile.
	//
	// This member is required.
	Name *string

	// The description of your instance profile.
	Description *string

	// An array of strings that specifies the list of app packages that should not be
	// cleaned up from the device after a test run. The list of packages is considered
	// only if you set packageCleanup to true.
	ExcludeAppPackagesFromCleanup []*string

	// When set to true, Device Farm removes app packages after a test run. The default
	// value is false for private devices.
	PackageCleanup *bool

	// When set to true, Device Farm reboots the instance after a test run. The default
	// value is true.
	RebootAfterUse *bool
}

type CreateInstanceProfileOutput struct {

	// An object that contains information about your instance profile.
	InstanceProfile *types.InstanceProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateInstanceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateInstanceProfile{}, middleware.After)
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
	addOpCreateInstanceProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstanceProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateInstanceProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateInstanceProfile",
	}
}
