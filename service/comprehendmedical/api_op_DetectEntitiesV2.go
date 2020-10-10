// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Inspects the clinical text for a variety of medical entities and returns
// specific information about them such as entity category, location, and
// confidence score on that information. Amazon Comprehend Medical only detects
// medical entities in English language texts. The DetectEntitiesV2 operation
// replaces the DetectEntities () operation. This new action uses a different model
// for determining the entities in your medical text and changes the way that some
// entities are returned in the output. You should use the DetectEntitiesV2
// operation in all new applications. The DetectEntitiesV2 operation returns the
// Acuity and Direction entities as attributes instead of types.
func (c *Client) DetectEntitiesV2(ctx context.Context, params *DetectEntitiesV2Input, optFns ...func(*Options)) (*DetectEntitiesV2Output, error) {
	if params == nil {
		params = &DetectEntitiesV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectEntitiesV2", params, optFns, addOperationDetectEntitiesV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectEntitiesV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectEntitiesV2Input struct {

	// A UTF-8 string containing the clinical content being examined for entities. Each
	// string must contain fewer than 20,000 bytes of characters.
	//
	// This member is required.
	Text *string
}

type DetectEntitiesV2Output struct {

	// The collection of medical entities extracted from the input text and their
	// associated information. For each entity, the response provides the entity text,
	// the entity category, where the entity text begins and ends, and the level of
	// confidence in the detection and analysis. Attributes and traits of the entity
	// are also returned.
	//
	// This member is required.
	Entities []*types.Entity

	// The version of the model used to analyze the documents. The version number looks
	// like X.X.X. You can use this information to track the model used for a
	// particular batch of documents.
	//
	// This member is required.
	ModelVersion *string

	// If the result to the DetectEntitiesV2 operation was truncated, include the
	// PaginationToken to fetch the next page of entities.
	PaginationToken *string

	// Attributes extracted from the input text that couldn't be related to an entity.
	UnmappedAttributes []*types.UnmappedAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetectEntitiesV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectEntitiesV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectEntitiesV2{}, middleware.After)
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
	addOpDetectEntitiesV2ValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectEntitiesV2(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDetectEntitiesV2(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "DetectEntitiesV2",
	}
}
