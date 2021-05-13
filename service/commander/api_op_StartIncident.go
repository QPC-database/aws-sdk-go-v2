// Code generated by smithy-go-codegen DO NOT EDIT.

package commander

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/commander/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used to start an incident from CloudWatch alarms, EventBridge events, or
// manually.
func (c *Client) StartIncident(ctx context.Context, params *StartIncidentInput, optFns ...func(*Options)) (*StartIncidentOutput, error) {
	if params == nil {
		params = &StartIncidentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartIncident", params, optFns, addOperationStartIncidentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartIncidentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartIncidentInput struct {

	// The Amazon Resource Name (ARN) of the response plan that pre-defines summary,
	// chat channels, SNS topics, runbooks, title, and impact of the incident.
	//
	// This member is required.
	ResponsePlanArn *string

	// A token ensuring that the action is called only once with the specified details.
	ClientToken *string

	// Defines the impact to the customers. Providing an impact overwrites the impact
	// provided by a response plan. Possible impacts:
	//
	// * 1 - Critical impact, this
	// typically relates to full application failure that impacts many to all
	// customers.
	//
	// * 2 - High impact, partial application failure with impact to many
	// customers.
	//
	// * 3 - Medium impact, the application is providing reduced service to
	// customers.
	//
	// * 4 - Low impact, customer might aren't impacted by the problem
	// yet.
	//
	// * 5 - No impact, customers aren't currently impacted but urgent action is
	// needed to avoid impact.
	Impact *int32

	// Add related items to the incident for other responders to use. Related items are
	// AWS resources, external links, or files uploaded to an S3 bucket.
	RelatedItems []types.RelatedItem

	// Provide a title for the incident. Providing a title overwrites the title
	// provided by the response plan.
	Title *string

	// Details of what created the incident record in Incident Manager.
	TriggerDetails *types.TriggerDetails
}

type StartIncidentOutput struct {

	// The ARN of the newly created incident record.
	//
	// This member is required.
	IncidentRecordArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartIncidentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartIncident{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartIncident{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opStartIncidentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartIncidentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartIncident(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpStartIncident struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartIncident) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartIncident) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartIncidentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartIncidentInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartIncidentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartIncident{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartIncident(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-incidents",
		OperationName: "StartIncident",
	}
}
