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

// Gets information about uploads, given an AWS Device Farm project ARN.
func (c *Client) ListUploads(ctx context.Context, params *ListUploadsInput, optFns ...func(*Options)) (*ListUploadsOutput, error) {
	if params == nil {
		params = &ListUploadsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUploads", params, optFns, addOperationListUploadsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUploadsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the list uploads operation.
type ListUploadsInput struct {

	// The Amazon Resource Name (ARN) of the project for which you want to list
	// uploads.
	//
	// This member is required.
	Arn *string

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// The type of upload. Must be one of the following values:
	//
	//     * ANDROID_APP
	//
	//
	// * IOS_APP
	//
	//     * WEB_APP
	//
	//     * EXTERNAL_DATA
	//
	//     *
	// APPIUM_JAVA_JUNIT_TEST_PACKAGE
	//
	//     * APPIUM_JAVA_TESTNG_TEST_PACKAGE
	//
	//     *
	// APPIUM_PYTHON_TEST_PACKAGE
	//
	//     * APPIUM_NODE_TEST_PACKAGE
	//
	//     *
	// APPIUM_RUBY_TEST_PACKAGE
	//
	//     * APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE
	//
	//     *
	// APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE
	//
	//     * APPIUM_WEB_PYTHON_TEST_PACKAGE
	//
	//     *
	// APPIUM_WEB_NODE_TEST_PACKAGE
	//
	//     * APPIUM_WEB_RUBY_TEST_PACKAGE
	//
	//     *
	// CALABASH_TEST_PACKAGE
	//
	//     * INSTRUMENTATION_TEST_PACKAGE
	//
	//     *
	// UIAUTOMATION_TEST_PACKAGE
	//
	//     * UIAUTOMATOR_TEST_PACKAGE
	//
	//     *
	// XCTEST_TEST_PACKAGE
	//
	//     * XCTEST_UI_TEST_PACKAGE
	//
	//     *
	// APPIUM_JAVA_JUNIT_TEST_SPEC
	//
	//     * APPIUM_JAVA_TESTNG_TEST_SPEC
	//
	//     *
	// APPIUM_PYTHON_TEST_SPEC
	//
	//     * APPIUM_NODE_TEST_SPEC
	//
	//     *
	// APPIUM_RUBY_TEST_SPEC
	//
	//     * APPIUM_WEB_JAVA_JUNIT_TEST_SPEC
	//
	//     *
	// APPIUM_WEB_JAVA_TESTNG_TEST_SPEC
	//
	//     * APPIUM_WEB_PYTHON_TEST_SPEC
	//
	//     *
	// APPIUM_WEB_NODE_TEST_SPEC
	//
	//     * APPIUM_WEB_RUBY_TEST_SPEC
	//
	//     *
	// INSTRUMENTATION_TEST_SPEC
	//
	//     * XCTEST_UI_TEST_SPEC
	Type types.UploadType
}

// Represents the result of a list uploads request.
type ListUploadsOutput struct {

	// If the number of items that are returned is significantly large, this is an
	// identifier that is also returned. It can be used in a subsequent call to this
	// operation to return the next set of items in the list.
	NextToken *string

	// Information about the uploads.
	Uploads []*types.Upload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUploadsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUploads{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUploads{}, middleware.After)
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
	addOpListUploadsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListUploads(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListUploads(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListUploads",
	}
}
