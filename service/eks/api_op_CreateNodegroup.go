// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a managed node group for an Amazon EKS cluster. You can only create a
// node group for your cluster that is equal to the current Kubernetes version for
// the cluster. All node groups are created with the latest AMI release version for
// the respective minor Kubernetes version of the cluster, unless you deploy a
// custom AMI using a launch template. For more information about using launch
// templates, see Launch template support
// (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html). An
// Amazon EKS managed node group is an Amazon EC2 Auto Scaling group and associated
// Amazon EC2 instances that are managed by AWS for an Amazon EKS cluster. Each
// node group uses a version of the Amazon EKS optimized Amazon Linux 2 AMI. For
// more information, see Managed Node Groups
// (https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html) in
// the Amazon EKS User Guide.
func (c *Client) CreateNodegroup(ctx context.Context, params *CreateNodegroupInput, optFns ...func(*Options)) (*CreateNodegroupOutput, error) {
	if params == nil {
		params = &CreateNodegroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNodegroup", params, optFns, addOperationCreateNodegroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNodegroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNodegroupInput struct {

	// The name of the cluster to create the node group in.
	//
	// This member is required.
	ClusterName *string

	// The Amazon Resource Name (ARN) of the IAM role to associate with your node
	// group. The Amazon EKS worker node kubelet daemon makes calls to AWS APIs on your
	// behalf. Nodes receive permissions for these API calls through an IAM instance
	// profile and associated policies. Before you can launch nodes and register them
	// into a cluster, you must create an IAM role for those nodes to use when they are
	// launched. For more information, see Amazon EKS node IAM role
	// (https://docs.aws.amazon.com/eks/latest/userguide/worker_node_IAM_role.html) in
	// the Amazon EKS User Guide . If you specify launchTemplate, then don't specify
	// IamInstanceProfile
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html)
	// in your launch template, or the node group deployment will fail. For more
	// information about using launch templates with Amazon EKS, see Launch template
	// support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	//
	// This member is required.
	NodeRole *string

	// The unique name to give your node group.
	//
	// This member is required.
	NodegroupName *string

	// The subnets to use for the Auto Scaling group that is created for your node
	// group. If you specify launchTemplate, then don't specify SubnetId
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html)
	// in your launch template, or the node group deployment will fail. For more
	// information about using launch templates with Amazon EKS, see Launch template
	// support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	//
	// This member is required.
	Subnets []string

	// The AMI type for your node group. GPU instance types should use the
	// AL2_x86_64_GPU AMI type. Non-GPU instances should use the AL2_x86_64 AMI type.
	// Arm instances should use the AL2_ARM_64 AMI type. All types use the Amazon EKS
	// optimized Amazon Linux 2 AMI. If you specify launchTemplate, and your launch
	// template uses a custom AMI, then don't specify amiType, or the node group
	// deployment will fail. For more information about using launch templates with
	// Amazon EKS, see Launch template support
	// (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the
	// Amazon EKS User Guide.
	AmiType types.AMITypes

	// The capacity type for your node group.
	CapacityType types.CapacityTypes

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The root device disk size (in GiB) for your node group instances. The default
	// disk size is 20 GiB. If you specify launchTemplate, then don't specify diskSize,
	// or the node group deployment will fail. For more information about using launch
	// templates with Amazon EKS, see Launch template support
	// (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the
	// Amazon EKS User Guide.
	DiskSize *int32

	// Specify the instance types for a node group. If you specify a GPU instance type,
	// be sure to specify AL2_x86_64_GPU with the amiType parameter. If you specify
	// launchTemplate, then you can specify zero or one instance type in your launch
	// template or you can specify 0-20 instance types for instanceTypes. If however,
	// you specify an instance type in your launch template and specify any
	// instanceTypes, the node group deployment will fail. If you don't specify an
	// instance type in a launch template or for instanceTypes, then t3.medium is used,
	// by default. If you specify Spot for capacityType, then we recommend specifying
	// multiple values for instanceTypes. For more information, see Managed node group
	// capacity types
	// (https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types)
	// and Launch template support
	// (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the
	// Amazon EKS User Guide.
	InstanceTypes []string

	// The Kubernetes labels to be applied to the nodes in the node group when they are
	// created.
	Labels map[string]string

	// An object representing a node group's launch template specification. If
	// specified, then do not specify instanceTypes, diskSize, or remoteAccess and make
	// sure that the launch template meets the requirements in
	// launchTemplateSpecification.
	LaunchTemplate *types.LaunchTemplateSpecification

	// The AMI version of the Amazon EKS optimized AMI to use with your node group. By
	// default, the latest available AMI version for the node group's current
	// Kubernetes version is used. For more information, see Amazon EKS optimized
	// Amazon Linux 2 AMI versions
	// (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)
	// in the Amazon EKS User Guide. If you specify launchTemplate, and your launch
	// template uses a custom AMI, then don't specify releaseVersion, or the node group
	// deployment will fail. For more information about using launch templates with
	// Amazon EKS, see Launch template support
	// (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the
	// Amazon EKS User Guide.
	ReleaseVersion *string

	// The remote access (SSH) configuration to use with your node group. If you
	// specify launchTemplate, then don't specify remoteAccess, or the node group
	// deployment will fail. For more information about using launch templates with
	// Amazon EKS, see Launch template support
	// (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the
	// Amazon EKS User Guide.
	RemoteAccess *types.RemoteAccessConfig

	// The scaling configuration details for the Auto Scaling group that is created for
	// your node group.
	ScalingConfig *types.NodegroupScalingConfig

	// The metadata to apply to the node group to assist with categorization and
	// organization. Each tag consists of a key and an optional value, both of which
	// you define. Node group tags do not propagate to any other resources associated
	// with the node group, such as the Amazon EC2 instances or subnets.
	Tags map[string]string

	// The Kubernetes taints to be applied to the nodes in the node group.
	Taints []types.Taint

	// The Kubernetes version to use for your managed nodes. By default, the Kubernetes
	// version of the cluster is used, and this is the only accepted specified value.
	// If you specify launchTemplate, and your launch template uses a custom AMI, then
	// don't specify version, or the node group deployment will fail. For more
	// information about using launch templates with Amazon EKS, see Launch template
	// support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	Version *string
}

type CreateNodegroupOutput struct {

	// The full description of your new node group.
	Nodegroup *types.Nodegroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateNodegroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateNodegroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNodegroup{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateNodegroupMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateNodegroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNodegroup(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateNodegroup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNodegroup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNodegroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNodegroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNodegroupInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateNodegroupMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateNodegroup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNodegroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "CreateNodegroup",
	}
}
