// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// You can use the GetMetricWidgetImage API to retrieve a snapshot graph of one or
// more Amazon CloudWatch metrics as a bitmap image. You can then embed this image
// into your services and products, such as wiki pages, reports, and documents. You
// could also retrieve images regularly, such as every minute, and create your own
// custom live dashboard.  <p>The graph you retrieve can include all CloudWatch
// metric graph features, including metric math and horizontal and vertical
// annotations.</p> <p>There is a limit of 20 transactions per second for this API.
// Each <code>GetMetricWidgetImage</code> action has the following limits:</p> <ul>
// <li> <p>As many as 100 metrics in the graph.</p> </li> <li> <p>Up to 100 KB
// uncompressed payload.</p> </li> </ul>
func (c *Client) GetMetricWidgetImage(ctx context.Context, params *GetMetricWidgetImageInput, optFns ...func(*Options)) (*GetMetricWidgetImageOutput, error) {
	if params == nil {
		params = &GetMetricWidgetImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetricWidgetImage", params, optFns, addOperationGetMetricWidgetImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetricWidgetImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetricWidgetImageInput struct {

	// A JSON string that defines the bitmap graph to be retrieved. The string includes
	// the metrics to include in the graph, statistics, annotations, title, axis
	// limits, and so on. You can include only one MetricWidget parameter in each
	// GetMetricWidgetImage call. For more information about the syntax of MetricWidget
	// see GetMetricWidgetImage: Metric Widget Structure and Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Metric-Widget-Structure.html).
	// <p>If any metric on the graph could not load all the requested data points, an
	// orange triangle with an exclamation point appears next to the graph legend.</p>
	//
	// This member is required.
	MetricWidget *string

	// The format of the resulting image. Only PNG images are supported.  <p>The
	// default is <code>png</code>. If you specify <code>png</code>, the API returns an
	// HTTP response with the content-type set to <code>text/xml</code>. The image data
	// is in a <code>MetricWidgetImage</code> field. For example:</p> <p> <code>
	// <GetMetricWidgetImageResponse xmlns=<URLstring>></code> </p> <p> <code>
	// <GetMetricWidgetImageResult></code> </p> <p> <code> <MetricWidgetImage></code>
	// </p> <p> <code> iVBORw0KGgoAAAANSUhEUgAAAlgAAAGQEAYAAAAip...</code> </p> <p>
	// <code> </MetricWidgetImage></code> </p> <p> <code>
	// </GetMetricWidgetImageResult></code> </p> <p> <code> <ResponseMetadata></code>
	// </p> <p> <code>
	// <RequestId>6f0d4192-4d42-11e8-82c1-f539a07e0e3b</RequestId></code> </p> <p>
	// <code> </ResponseMetadata></code> </p> <p>
	// <code></GetMetricWidgetImageResponse></code> </p> <p>The <code>image/png</code>
	// setting is intended only for custom HTTP requests. For most use cases, and all
	// actions using an AWS SDK, you should use <code>png</code>. If you specify
	// <code>image/png</code>, the HTTP response has a content-type set to
	// <code>image/png</code>, and the body of the response is a PNG image. </p>
	OutputFormat *string
}

type GetMetricWidgetImageOutput struct {

	// The image of the graph, in the output format specified. The output is
	// base64-encoded.
	MetricWidgetImage []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMetricWidgetImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetMetricWidgetImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetMetricWidgetImage{}, middleware.After)
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
	addOpGetMetricWidgetImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetricWidgetImage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetMetricWidgetImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "GetMetricWidgetImage",
	}
}
