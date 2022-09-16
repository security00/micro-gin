package reply

type UsersReply struct {
	Id *string `json:"id"` 
	Name *string `json:"name"` 
	Mobile *string `json:"mobile"` 
	Password *string `json:"password"` 
	CreatedAt *jsontime.JsonTime `json:"createdAt"` 
	UpdatedAt *jsontime.JsonTime `json:"updatedAt"` 
	DeletedAt *jsontime.JsonTime `json:"deletedAt"` 
}
