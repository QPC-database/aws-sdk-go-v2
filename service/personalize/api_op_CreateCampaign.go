// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a campaign by deploying a solution version. When a client calls the
// GetRecommendations
// (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html)
// and GetPersonalizedRanking
// (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetPersonalizedRanking.html)
// APIs, a campaign is specified in the request.  <p> <b>Minimum Provisioned TPS
// and Auto-Scaling</b> </p> <p>A transaction is a single
// <code>GetRecommendations</code> or <code>GetPersonalizedRanking</code> call.
// Transactions per second (TPS) is the throughput and unit of billing for Amazon
// Personalize. The minimum provisioned TPS (<code>minProvisionedTPS</code>)
// specifies the baseline throughput provisioned by Amazon Personalize, and thus,
// the minimum billing charge. If your TPS increases beyond
// <code>minProvisionedTPS</code>, Amazon Personalize auto-scales the provisioned
// capacity up and down, but never below <code>minProvisionedTPS</code>, to
// maintain a 70% utilization. There's a short time delay while the capacity is
// increased that might cause loss of transactions. It's recommended to start with
// a low <code>minProvisionedTPS</code>, track your usage using Amazon CloudWatch
// metrics, and then increase the <code>minProvisionedTPS</code> as necessary.</p>
// <p> <b>Status</b> </p> <p>A campaign can be in one of the following states:</p>
// <ul> <li> <p>CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED</p>
// </li> <li> <p>DELETE PENDING > DELETE IN_PROGRESS</p> </li> </ul> <p>To get the
// campaign status, call <a>DescribeCampaign</a>.</p> <note> <p>Wait until the
// <code>status</code> of the campaign is <code>ACTIVE</code> before asking the
// campaign for recommendations.</p> </note> <p class="title"> <b>Related APIs</b>
// </p> <ul> <li> <p> <a>ListCampaigns</a> </p> </li> <li> <p>
// <a>DescribeCampaign</a> </p> </li> <li> <p> <a>UpdateCampaign</a> </p> </li>
// <li> <p> <a>DeleteCampaign</a> </p> </li> </ul>
func (c *Client) CreateCampaign(ctx context.Context, params *CreateCampaignInput, optFns ...func(*Options)) (*CreateCampaignOutput, error) {
	if params == nil {
		params = &CreateCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCampaign", params, optFns, addOperationCreateCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCampaignInput struct {

	// Specifies the requested minimum provisioned transactions (recommendations) per
	// second that Amazon Personalize will support.
	//
	// This member is required.
	MinProvisionedTPS *int32

	// A name for the new campaign. The campaign name must be unique within your
	// account.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the solution version to deploy.
	//
	// This member is required.
	SolutionVersionArn *string
}

type CreateCampaignOutput struct {

	// The Amazon Resource Name (ARN) of the campaign.
	CampaignArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCampaign{}, middleware.After)
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
	addOpCreateCampaignValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCampaign(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateCampaign(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateCampaign",
	}
}
