// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type InstanceMetadataEndpointState string

// Enum values for InstanceMetadataEndpointState
const (
	InstanceMetadataEndpointStateDisabled InstanceMetadataEndpointState = "disabled"
	InstanceMetadataEndpointStateEnabled  InstanceMetadataEndpointState = "enabled"
)

type InstanceMetadataHttpTokensState string

// Enum values for InstanceMetadataHttpTokensState
const (
	InstanceMetadataHttpTokensStateOptional InstanceMetadataHttpTokensState = "optional"
	InstanceMetadataHttpTokensStateRequired InstanceMetadataHttpTokensState = "required"
)

type InstanceRefreshStatus string

// Enum values for InstanceRefreshStatus
const (
	InstanceRefreshStatusPending    InstanceRefreshStatus = "Pending"
	InstanceRefreshStatusInprogress InstanceRefreshStatus = "InProgress"
	InstanceRefreshStatusSuccessful InstanceRefreshStatus = "Successful"
	InstanceRefreshStatusFailed     InstanceRefreshStatus = "Failed"
	InstanceRefreshStatusCancelling InstanceRefreshStatus = "Cancelling"
	InstanceRefreshStatusCancelled  InstanceRefreshStatus = "Cancelled"
)

type LifecycleState string

// Enum values for LifecycleState
const (
	LifecycleStatePending             LifecycleState = "Pending"
	LifecycleStatePending_wait        LifecycleState = "Pending:Wait"
	LifecycleStatePending_proceed     LifecycleState = "Pending:Proceed"
	LifecycleStateQuarantined         LifecycleState = "Quarantined"
	LifecycleStateIn_service          LifecycleState = "InService"
	LifecycleStateTerminating         LifecycleState = "Terminating"
	LifecycleStateTerminating_wait    LifecycleState = "Terminating:Wait"
	LifecycleStateTerminating_proceed LifecycleState = "Terminating:Proceed"
	LifecycleStateTerminated          LifecycleState = "Terminated"
	LifecycleStateDetaching           LifecycleState = "Detaching"
	LifecycleStateDetached            LifecycleState = "Detached"
	LifecycleStateEntering_standby    LifecycleState = "EnteringStandby"
	LifecycleStateStandby             LifecycleState = "Standby"
)

type MetricStatistic string

// Enum values for MetricStatistic
const (
	MetricStatisticAverage     MetricStatistic = "Average"
	MetricStatisticMinimum     MetricStatistic = "Minimum"
	MetricStatisticMaximum     MetricStatistic = "Maximum"
	MetricStatisticSamplecount MetricStatistic = "SampleCount"
	MetricStatisticSum         MetricStatistic = "Sum"
)

type MetricType string

// Enum values for MetricType
const (
	MetricTypeAsgaveragecpuutilization MetricType = "ASGAverageCPUUtilization"
	MetricTypeAsgaveragenetworkin      MetricType = "ASGAverageNetworkIn"
	MetricTypeAsgaveragenetworkout     MetricType = "ASGAverageNetworkOut"
	MetricTypeAlbrequestcountpertarget MetricType = "ALBRequestCountPerTarget"
)

type RefreshStrategy string

// Enum values for RefreshStrategy
const (
	RefreshStrategyRolling RefreshStrategy = "Rolling"
)

type ScalingActivityStatusCode string

// Enum values for ScalingActivityStatusCode
const (
	ScalingActivityStatusCodePendingspotbidplacement         ScalingActivityStatusCode = "PendingSpotBidPlacement"
	ScalingActivityStatusCodeWaitingforspotinstancerequestid ScalingActivityStatusCode = "WaitingForSpotInstanceRequestId"
	ScalingActivityStatusCodeWaitingforspotinstanceid        ScalingActivityStatusCode = "WaitingForSpotInstanceId"
	ScalingActivityStatusCodeWaitingforinstanceid            ScalingActivityStatusCode = "WaitingForInstanceId"
	ScalingActivityStatusCodePreinservice                    ScalingActivityStatusCode = "PreInService"
	ScalingActivityStatusCodeInprogress                      ScalingActivityStatusCode = "InProgress"
	ScalingActivityStatusCodeWaitingforelbconnectiondraining ScalingActivityStatusCode = "WaitingForELBConnectionDraining"
	ScalingActivityStatusCodeMidlifecycleaction              ScalingActivityStatusCode = "MidLifecycleAction"
	ScalingActivityStatusCodeWaitingforinstancewarmup        ScalingActivityStatusCode = "WaitingForInstanceWarmup"
	ScalingActivityStatusCodeSuccessful                      ScalingActivityStatusCode = "Successful"
	ScalingActivityStatusCodeFailed                          ScalingActivityStatusCode = "Failed"
	ScalingActivityStatusCodeCancelled                       ScalingActivityStatusCode = "Cancelled"
)