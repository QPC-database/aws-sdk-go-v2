// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified cluster subnet group.
func (c *Client) DeleteClusterSubnetGroup(ctx context.Context, params *DeleteClusterSubnetGroupInput, optFns ...func(*Options)) (*DeleteClusterSubnetGroupOutput, error) {
	if params == nil {
		params = &DeleteClusterSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteClusterSubnetGroup", params, optFns, addOperationDeleteClusterSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteClusterSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteClusterSubnetGroupInput struct {

	// The name of the cluster subnet group name to be deleted.
	//
	// This member is required.
	ClusterSubnetGroupName *string
}

type DeleteClusterSubnetGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteClusterSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteClusterSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteClusterSubnetGroup{}, middleware.After)
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
	addOpDeleteClusterSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteClusterSubnetGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteClusterSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DeleteClusterSubnetGroup",
	}
}
