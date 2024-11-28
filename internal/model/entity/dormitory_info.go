// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// DormitoryInfo is the golang structure for table dormitory_info.
type DormitoryInfo struct {
	DormitoryId     int    `json:"dormitoryId"     orm:"dormitory_id"     description:""`
	Floor           int    `json:"floor"           orm:"floor"            description:""`
	DormitoryNumber string `json:"dormitoryNumber" orm:"dormitory_number" description:""`
	DormitoryType   string `json:"dormitoryType"   orm:"dormitory_type"   description:""`
}
