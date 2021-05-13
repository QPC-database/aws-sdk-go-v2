// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/commander/types"
	"time"
)

func ExampleAction_outputUsage() {
	var union types.Action
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ActionMemberSsmAutomation:
		_ = v.Value // Value is types.SsmAutomation

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SsmAutomation

func ExampleAttributeValueList_outputUsage() {
	var union types.AttributeValueList
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AttributeValueListMemberIntegerValues:
		_ = v.Value // Value is []int32

	case *types.AttributeValueListMemberStringValues:
		_ = v.Value // Value is []string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []string
var _ []int32

func ExampleAutomationExecution_outputUsage() {
	var union types.AutomationExecution
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AutomationExecutionMemberSsmExecutionArn:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleChatChannel_outputUsage() {
	var union types.ChatChannel
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ChatChannelMemberChatbotSns:
		_ = v.Value // Value is []string

	case *types.ChatChannelMemberEmpty:
		_ = v.Value // Value is types.EmptyChatChannel

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EmptyChatChannel
var _ []string

func ExampleCondition_outputUsage() {
	var union types.Condition
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConditionMemberAfter:
		_ = v.Value // Value is time.Time

	case *types.ConditionMemberBefore:
		_ = v.Value // Value is time.Time

	case *types.ConditionMemberEquals:
		_ = v.Value // Value is types.AttributeValueList

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *time.Time
var _ types.AttributeValueList

func ExampleItemValue_outputUsage() {
	var union types.ItemValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ItemValueMemberArn:
		_ = v.Value // Value is string

	case *types.ItemValueMemberMetricDefinition:
		_ = v.Value // Value is string

	case *types.ItemValueMemberUrl:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string
var _ *string

func ExampleNotificationTargetItem_outputUsage() {
	var union types.NotificationTargetItem
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.NotificationTargetItemMemberSnsTopicArn:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleRelatedItemsUpdate_outputUsage() {
	var union types.RelatedItemsUpdate
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RelatedItemsUpdateMemberItemToAdd:
		_ = v.Value // Value is types.RelatedItem

	case *types.RelatedItemsUpdateMemberItemToRemove:
		_ = v.Value // Value is types.ItemIdentifier

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.RelatedItem
var _ *types.ItemIdentifier

func ExampleUpdateReplicationSetAction_outputUsage() {
	var union types.UpdateReplicationSetAction
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.UpdateReplicationSetActionMemberAddRegionAction:
		_ = v.Value // Value is types.AddRegionAction

	case *types.UpdateReplicationSetActionMemberDeleteRegionAction:
		_ = v.Value // Value is types.DeleteRegionAction

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.DeleteRegionAction
var _ *types.AddRegionAction
