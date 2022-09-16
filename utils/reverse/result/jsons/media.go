package reply

type MediaReply struct {
	Id *string `json:"id"` 
	DiskType *string `json:"diskType"` 
	SrcType *int8 `json:"srcType"` 
	Src *string `json:"src"` 
	CreatedAt *jsontime.JsonTime `json:"createdAt"` 
	UpdatedAt *jsontime.JsonTime `json:"updatedAt"` 
}
