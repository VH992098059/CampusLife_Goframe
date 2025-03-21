package backapi

/*请求通用数据*/
type CommonPaginationReq struct {
	Page int `json:"page" in:"query" d:"1"  v:"min:1#分页号码错误"     dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"10" v:"min:1|max:10#分页数量最小1条|分页数量最大10条" dc:"分页数量，默认最大10"`
}

/*相应通用数据*/
type CommonPaginationRes struct {
	List   interface{} `dc:"列表数据"`
	Total  int         `dc:"总数"`
	Page   int         `dc:"分页号码"`
	Size   int         `dc:"分页数量"`
	Offset int         `json:"offset" dc:"初始容量"`
	Limit  int         `json:"limit" dc:"最大容量"`
}
