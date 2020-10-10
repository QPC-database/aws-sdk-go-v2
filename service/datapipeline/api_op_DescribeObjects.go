// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the object definitions for a set of objects associated with the pipeline.
// Object definitions are composed of a set of fields that define the properties of
// the object.
func (c *Client) DescribeObjects(ctx context.Context, params *DescribeObjectsInput, optFns ...func(*Options)) (*DescribeObjectsOutput, error) {
	if params == nil {
		params = &DescribeObjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeObjects", params, optFns, addOperationDescribeObjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeObjects.
type DescribeObjectsInput struct {

	// The IDs of the pipeline objects that contain the definitions to be described.
	// You can pass as many as 25 identifiers in a single call to DescribeObjects.
	//
	// This member is required.
	ObjectIds []*string

	// The ID of the pipeline that contains the object definitions.
	//
	// This member is required.
	PipelineId *string

	// Indicates whether any expressions in the object should be evaluated when the
	// object descriptions are returned.
	EvaluateExpressions *bool

	// The starting point for the results to be returned. For the first call, this
	// value should be empty. As long as there are more results, continue to call
	// DescribeObjects with the marker value from the previous call to retrieve the
	// next set of results.
	Marker *string
}

// Contains the output of DescribeObjects.
type DescribeObjectsOutput struct {

	// An array of object definitions.
	//
	// This member is required.
	PipelineObjects []*types.PipelineObject

	// Indicates whether there are more results to return.
	HasMoreResults *bool

	// The starting point for the next page of results. To view the next page of
	// results, call DescribeObjects again with this marker value. If the value is
	// null, there are no more results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeObjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeObjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeObjects{}, middleware.After)
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
	addOpDescribeObjectsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeObjects(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeObjects(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "DescribeObjects",
	}
}
