// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityContext is the golang structure of table activity_context for DAO operations like Where/Data.
type ActivityContext struct {
	g.Meta       `orm:"table:activity_context, do:true"`
	ActivityUuid interface{} // 活动ID
	Pic          []byte      // 图片内容
}
