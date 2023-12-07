package repo

import (
	"ims-server/affair/dal/model"
	ioginx "ims-server/pkg/ginx"
)

type affairRepo struct {
	ioginx.IRepo[model.ToDoAffair]
}

func NewAffairRepo() *affairRepo {
	return &affairRepo{}
}

//TODO:其他种类的“增删改查”