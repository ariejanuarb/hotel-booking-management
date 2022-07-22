package service

import (
	"booking-hotel/model/web"
	"context"
)

type RoomFacilityService interface {
	Create(ctx context.Context, request web.RoomFacilityCreateRequest) web.RoomFacilityResponse
	Update(ctx context.Context, request web.RoomFacilityUpdateRequest) web.RoomFacilityResponse
	Delete(ctx context.Context, roomfacilityId int)
	FindById(ctx context.Context, roomfacilityId int) web.RoomFacilityResponse
	FindByAll(ctx context.Context) []web.RoomFacilityResponse
}
