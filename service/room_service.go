package service

import (
	"booking-hotel/model/web"
	"context"
)

type RoomService interface {
	Create(ctx context.Context, request web.RoomCreateRequest) web.RoomResponse
	Update(ctx context.Context, request web.RoomUpdateRequest) web.RoomResponse
	Delete(ctx context.Context, roomId int)
	FindById(ctx context.Context, roomId int) web.RoomResponse
	FindByAll(ctx context.Context) []web.RoomResponse
}
