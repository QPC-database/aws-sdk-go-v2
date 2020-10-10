// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Launches an AWS Elastic Beanstalk environment for the specified application
// using the specified configuration.
func (c *Client) CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) {
	if params == nil {
		params = &CreateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironment", params, optFns, addOperationCreateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateEnvironmentInput struct {

	// The name of the application that is associated with this environment.
	//
	// This member is required.
	ApplicationName *string

	// If specified, the environment attempts to use this value as the prefix for the
	// CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is
	// generated automatically by appending a random alphanumeric string to the
	// environment name.
	CNAMEPrefix *string

	// Your description for this environment.
	Description *string

	// A unique name for the environment. Constraint: Must be from 4 to 40 characters
	// in length. The name can contain only letters, numbers, and hyphens. It can't
	// start or end with a hyphen. This name must be unique within a region in your
	// account. If the specified name already exists in the region, Elastic Beanstalk
	// returns an InvalidParameterValue error. If you don't specify the CNAMEPrefix
	// parameter, the environment name becomes part of the CNAME, and therefore part of
	// the visible URL for your application.
	EnvironmentName *string

	// The name of the group to which the target environment belongs. Specify a group
	// name only if the environment's name is specified in an environment manifest and
	// not with the environment name parameter. See Environment Manifest (env.yaml)
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/environment-cfg-manifest.html)
	// for details.
	GroupName *string

	// The Amazon Resource Name (ARN) of an existing IAM role to be used as the
	// environment's operations role. If specified, Elastic Beanstalk uses the
	// operations role for permissions to downstream services during this call and
	// during subsequent calls acting on this environment. To specify an operations
	// role, you must have the iam:PassRole permission for the role. For more
	// information, see Operations roles
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-operationsrole.html)
	// in the AWS Elastic Beanstalk Developer Guide.
	OperationsRole *string

	// If specified, AWS Elastic Beanstalk sets the specified configuration options to
	// the requested value in the configuration set for the new environment. These
	// override the values obtained from the solution stack or the configuration
	// template.
	OptionSettings []*types.ConfigurationOptionSetting

	// A list of custom user-defined configuration options to remove from the
	// configuration set for this new environment.
	OptionsToRemove []*types.OptionSpecification

	// The Amazon Resource Name (ARN) of the custom platform to use with the
	// environment. For more information, see Custom Platforms
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html)
	// in the AWS Elastic Beanstalk Developer Guide.  <p>If you specify
	// <code>PlatformArn</code>, don't specify <code>SolutionStackName</code>.</p>
	// </note>
	PlatformArn *string

	// The name of an Elastic Beanstalk solution stack (platform version) to use with
	// the environment. If specified, Elastic Beanstalk sets the configuration values
	// to the default values associated with the specified solution stack. For a list
	// of current solution stacks, see Elastic Beanstalk Supported Platforms
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/platforms/platforms-supported.html)
	// in the AWS Elastic Beanstalk Platforms guide. If you specify SolutionStackName,
	// don't specify PlatformArn or TemplateName.
	SolutionStackName *string

	// Specifies the tags applied to resources in the environment.
	Tags []*types.Tag

	// The name of the Elastic Beanstalk configuration template to use with the
	// environment. If you specify TemplateName, then don't specify SolutionStackName.
	TemplateName *string

	// Specifies the tier to use in creating this environment. The environment tier
	// that you choose determines whether Elastic Beanstalk provisions resources to
	// support a web application that handles HTTP(S) requests or a web application
	// that handles background-processing tasks.
	Tier *types.EnvironmentTier

	// The name of the application version to deploy. Default: If not specified,
	// Elastic Beanstalk attempts to deploy the sample application.
	VersionLabel *string
}

// Describes the properties of an environment.
type CreateEnvironmentOutput struct {

	// Indicates if there is an in-progress environment configuration update or
	// application version deployment that you can cancel. true: There is an update in
	// progress. false: There are no updates currently in progress.
	AbortableOperationInProgress *bool

	// The name of the application associated with this environment.
	ApplicationName *string

	// The URL to the CNAME for this environment.
	CNAME *string

	// The creation date for this environment.
	DateCreated *time.Time

	// The last modified date for this environment.
	DateUpdated *time.Time

	// Describes this environment.
	Description *string

	// For load-balanced, autoscaling environments, the URL to the LoadBalancer. For
	// single-instance environments, the IP address of the instance.
	EndpointURL *string

	// The environment's Amazon Resource Name (ARN), which can be used in other API
	// requests that require an ARN.
	EnvironmentArn *string

	// The ID of this environment.
	EnvironmentId *string

	// A list of links to other environments in the same group.
	EnvironmentLinks []*types.EnvironmentLink

	// The name of this environment.
	EnvironmentName *string

	// Describes the health status of the environment. AWS Elastic Beanstalk indicates
	// the failure levels for a running environment:
	//
	//     * Red: Indicates the
	// environment is not responsive. Occurs when three or more consecutive failures
	// occur for an environment.
	//
	//     * Yellow: Indicates that something is wrong.
	// Occurs when two consecutive failures occur for an environment.
	//
	//     * Green:
	// Indicates the environment is healthy and fully functional.
	//
	//     * Grey: Default
	// health for a new environment. The environment is not fully launched and health
	// checks have not started or health checks are suspended during an
	// UpdateEnvironment or RestartEnvironment request.
	//
	// Default: Grey
	Health types.EnvironmentHealth

	// Returns the health status of the application running in your environment. For
	// more information, see Health Colors and Statuses
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced-status.html).
	HealthStatus types.EnvironmentHealthStatus

	// The Amazon Resource Name (ARN) of the environment's operations role. For more
	// information, see Operations roles
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-operationsrole.html)
	// in the AWS Elastic Beanstalk Developer Guide.
	OperationsRole *string

	// The ARN of the platform version.
	PlatformArn *string

	// The description of the AWS resources used by this environment.
	Resources *types.EnvironmentResourcesDescription

	// The name of the SolutionStack deployed with this environment.
	SolutionStackName *string

	// The current operational status of the environment:  <ul> <li> <p>
	// <code>Launching</code>: Environment is in the process of initial deployment.</p>
	// </li> <li> <p> <code>Updating</code>: Environment is in the process of updating
	// its configuration settings or application version.</p> </li> <li> <p>
	// <code>Ready</code>: Environment is available to have an action performed on it,
	// such as update or terminate.</p> </li> <li> <p> <code>Terminating</code>:
	// Environment is in the shut-down process.</p> </li> <li> <p>
	// <code>Terminated</code>: Environment is not running.</p> </li> </ul>
	Status types.EnvironmentStatus

	// The name of the configuration template used to originally launch this
	// environment.
	TemplateName *string

	// Describes the current tier of this environment.
	Tier *types.EnvironmentTier

	// The application version deployed in this environment.
	VersionLabel *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateEnvironment{}, middleware.After)
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
	addOpCreateEnvironmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateEnvironment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "CreateEnvironment",
	}
}
