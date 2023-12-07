package pack

import (
	"ims-server/affair/dal/model"
	"ims-server/affair/param"
)

func ToAffairResponse(t *model.ToDoAffair) param.ToDoAffairResponse {
	res := param.ToDoAffairResponse{
		Title: t.Title,
		State: t.State,
		Content: t.Content,
	}
	return res
}
