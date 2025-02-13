// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ActivityContext is the golang structure for table activity_context.
type ActivityContext struct {
	ActivityUuid string `json:"activityUuid" orm:"activity_uuid" description:"活动ID"`
	Pic          []byte `json:"pic"          orm:"pic"           description:"图片内容"`
}
