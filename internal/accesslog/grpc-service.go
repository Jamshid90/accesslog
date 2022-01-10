package accesslog

import (
	"context"
	"net/http"

	gp "github.com/MedHubUz/access-log/genproto/accesslog"
	"github.com/MedHubUz/access-log/internal/entity"
	grpcpkg "github.com/MedHubUz/access-log/internal/grpc"
	"github.com/MedHubUz/access-log/internal/validation"
	"go.uber.org/zap"
)

type accesslogService struct {
	logger           *zap.Logger
	accesslogUsecase entity.AccessLogUsecase
	gp.UnimplementedAccessLogServiceServer
}

func NewgRPC(logger *zap.Logger, accesslogUsecase entity.AccessLogUsecase) gp.AccessLogServiceServer {
	return &accesslogService{
		logger:           logger,
		accesslogUsecase: accesslogUsecase,
	}
}

func (s *accesslogService) Create(ctx context.Context, r *gp.AccessLogCreateRequest) (*gp.AccessLogCreateResponse, error) {
	cr := CreateRequest{
		ID:     r.Id,
		UserID: r.UserID,
		Method: r.Method,
		URL:    r.Url,
		Data:   r.Data,
	}

	if err := validation.Validator(&cr); err != nil {
		return nil, grpcpkg.Error(err)
	}

	ag := entity.AccessLog{
		ID:     cr.ID,
		UserID: cr.UserID,
		Method: cr.Method,
		URL:    cr.URL,
		Data:   cr.Data,
	}

	if err := s.accesslogUsecase.Create(ctx, &ag); err != nil {
		return nil, grpcpkg.Error(err)
	}

	return &gp.AccessLogCreateResponse{Code: http.StatusOK}, nil
}

func (s *accesslogService) List(ctx context.Context, r *gp.AccessLogListRequest) (*gp.AccessLogListResponse, error) {

	var ag gp.AccessLogListResponse

	list, err := s.accesslogUsecase.ReadAll(ctx, make(map[string][]string))
	if err != nil {
		return nil, grpcpkg.Error(err)
	}

	for _, v := range list {
		ag.Results = append(ag.Results, &gp.AccessLog{
			Id:        v.ID,
			UserID:    v.UserID,
			Method:    v.Method,
			Url:       v.URL,
			Data:      v.Data,
			CreatedAt: v.CreatedAt.String(),
		})
	}

	return &ag, err
}
