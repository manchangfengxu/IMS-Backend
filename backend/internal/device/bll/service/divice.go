package service

import (
	"context"
	"ims-server/internal/device/bll/pack"

	"ims-server/internal/device/param"

	"ims-server/internal/device/dal/repo"
	egoerror "ims-server/pkg/error"
	"ims-server/pkg/util"
)

type sensorService struct {
}

func (s *sensorService) GetMqttDateByID(ctx context.Context, req *param.GetMqttDateByID) (*param.GetMqttDateByIDResponse, error) {
	id := req.ID
	sensorDate, err := repo.NewSensorRepo().Get(ctx, id)
	if err != nil {
		return nil, egoerror.ErrNotFound
	}

	resp := pack.ToMqttDateResponse(sensorDate)
	return &param.GetMqttDateByIDResponse{
		MqttDateResponse: resp,
	}, nil
}

func (s *sensorService) MGetSensorByIDs(ctx context.Context, req *param.MGetMqttDateByIDsRequest) (*param.MGetMqttDateByIDsResponse, error) {
	res, err := repo.NewSensorRepo().MGet(ctx, req.IDs)
	if err != nil {
		return nil, egoerror.ErrNotFound
	}

	resp := []param.GetMqttDateByIDResponse{}
	for _, sensorDate := range res {
		info := pack.ToMqttDateResponse(&sensorDate)
		resp = append(resp, param.GetMqttDateByIDResponse{
			MqttDateResponse: info,
		})
	}

	return &param.MGetMqttDateByIDsResponse{
		List: resp,
	}, nil
}

func (s *sensorService) UpdateMqDateByID(ctx context.Context, req *param.UpdateMqttDateByIDRequest) (*param.UpdateMqttDateByIDResponse, error) {
	_, err1 := repo.NewSensorRepo().Get(ctx, req.ID)
	if err1 != nil {
		return nil, egoerror.ErrNotFound
	}

	mqDateMap := util.RequestToSnakeMapWithIgnoreZeroValueAndIDKey(req)

	update, err2 := repo.NewSensorRepo().Update(ctx, req.ID, mqDateMap)
	if err2 != nil {
		return nil, egoerror.ErrInvalidParam
	}

	resp := pack.ToMqttDateResponse(update)
	return &param.UpdateMqttDateByIDResponse{
		MqttDateResponse: resp,
	}, nil
}

func (s *sensorService) DeleteMqDaterByID(ctx context.Context, req *param.DeleteMqttDateByIDRequest) (*param.DeleteMqttDateByIDResponse, error) {
	_, err := repo.NewSensorRepo().Get(ctx, req.ID)
	if err != nil {
		return nil, egoerror.ErrNotFound
	}
	err = repo.NewSensorRepo().Delete(ctx, req.ID)
	if err != nil {
		return nil, egoerror.ErrNotFound
	}
	return &param.DeleteMqttDateByIDResponse{}, nil
}
