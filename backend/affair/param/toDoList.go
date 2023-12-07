package param

type ToDoAffairResponse struct {
	Title   string `json:"title" form:"title"`
	State   bool   `json:"state" form:"state"`
	Content string `json:"content" form:"content"`
}

type ToDoAffairRequest struct {
	Title   string `json:"title" form:"title"`
	State   bool   `json:"state" form:"state"`
	Content string `json:"content" form:"content"`
}

type CreateAffairResponse struct {
	ToDoAffairResponse
}

type GetAffairByIDRequest struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

type GetAffairByIDResponse struct {
	ToDoAffairResponse
}

type MGetAffairByIDsRequest struct {
	IDs []uint `json:"ids" form:"ids" binding:"required"`
}

type MGetAffairByIDsResponse struct {
	List []ToDoAffairResponse `json:"list"`
}

type UpdateAffairByIDRequest struct {
	ID uint `json:"id" form:"id" binding:"required"`

	Title   string `json:"title" form:"title" binding:"required"`
	State   bool   `json:"state" form:"state"`
	Content string `json:"content" form:"content"`
}

type UpdateAffairByIDResponse struct {
	ToDoAffairResponse
}

type DeleteAffairByIDRequest struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

type DeleteAffairByIDResponse struct {
}
