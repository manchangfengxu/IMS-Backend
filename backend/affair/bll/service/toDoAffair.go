package service

import (
	"context"
	"ims-server/affair/bll/pack"
	"ims-server/affair/dal/model"
	"ims-server/affair/dal/repo"
	"ims-server/affair/param"
	egoerror "ims-server/pkg/error"
	"ims-server/pkg/util"
)

type affairService struct {
}

func NewAffairService() *affairService {
	return &affairService{}
}

func (a *affairService) CreateAffair(ctx context.Context, req *param.ToDoAffairRequest) (*param.CreateAffairResponse, error) {
	affair := &model.ToDoAffair{
		Title:   req.Title,
		Content: req.Content,
		State:   req.State,
	}

	err := repo.NewAffairRepo().Create(ctx, affair)
	if err != nil {
		return nil, egoerror.ErrInvalidParam
	}

	resp := pack.ToAffairResponse(affair)
	return &param.CreateAffairResponse{
		ToDoAffairResponse: resp,
	}, nil
}

func (a *affairService) GetAffairByID(ctx context.Context, req *param.GetAffairByIDRequest) (*param.GetAffairByIDResponse, error) {
	id := req.ID
	affair, err := repo.NewAffairRepo().Get(ctx, id)
	if err != nil {
		return nil, egoerror.ErrNotFound
	}

	resp := pack.ToAffairResponse(affair)
	return &param.GetAffairByIDResponse{
		ToDoAffairResponse: resp,
	}, nil
}

func (a *affairService) MGetAffairByIDs(ctx context.Context, req *param.MGetAffairByIDsRequest) (*param.MGetAffairByIDsResponse, error) {
	res, err := repo.NewAffairRepo().MGet(ctx, req.IDs)
	if err != nil {
		return nil, egoerror.ErrNotFound
	}

	resp := []param.ToDoAffairResponse{}
	for _, affair := range res {
		info := pack.ToAffairResponse(&affair)
		resp = append(resp, info)
	}

	return &param.MGetAffairByIDsResponse{
		List: resp,
	}, nil
}

func (a *affairService) UpdateAffairByID(ctx context.Context, req *param.UpdateAffairByIDRequest) (*param.UpdateAffairByIDResponse, error) {
	_, err1 := repo.NewAffairRepo().Get(ctx, req.ID)
	if err1 != nil {
		return nil, egoerror.ErrNotFound // 如果用户不存在，则返回错误信息
	}

	affairMap := util.RequestToSnakeMapWithIgnoreZeroValueAndIDKey(req)
	
	update, err2 := repo.NewAffairRepo().Update(ctx, req.ID, affairMap)
	if err2 != nil {
		return nil, egoerror.ErrInvalidParam // 如果更新失败，则返回无效参数的错误信息
	}

	resp := pack.ToAffairResponse(update)
	return &param.UpdateAffairByIDResponse{
		ToDoAffairResponse: resp,
	}, nil
}

func (a *affairService)DeleteAffairByID (ctx context.Context, req *param.DeleteAffairByIDRequest)(*param.DeleteAffairByIDResponse, error) {
	_, err1 := repo.NewAffairRepo().Get(ctx, req.ID, "id")
	if err1 != nil {
		return nil, egoerror.ErrNotFound
	}
	err2 := repo.NewAffairRepo().Delete(ctx, req.ID)
	if err2 != nil {
		return nil, egoerror.ErrNotFound
	}
	return &param.DeleteAffairByIDResponse{}, nil
}
