// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SchoolInfo is the golang structure for table school_info.
type SchoolInfo struct {
	SchoolId   int    `json:"schoolId"   orm:"school_id"   description:"学校编号"`
	SchoolName string `json:"schoolName" orm:"school_name" description:"学校名称"`
	Province   string `json:"province"   orm:"province"    description:"学校省份"`
	City       string `json:"city"       orm:"city"        description:"学校城市"`
	County     string `json:"county"     orm:"county"      description:"学号市区"`
	Address    string `json:"address"    orm:"address"     description:"学校详细地址"`
}
