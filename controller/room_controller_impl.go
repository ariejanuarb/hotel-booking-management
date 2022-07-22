package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type RoomControllerImpl struct {
	RoomService service.RoomService
}

func NewRoomController(roomService service.RoomService) RoomController {
	return &RoomControllerImpl{
		RoomService: roomService,
	}
}

func (controller *RoomControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomCreateRequest := web.RoomCreateRequest{}
	helper.ReadFromRequestBody(request, &roomCreateRequest)

	roomResponse := controller.RoomService.Create(request.Context(), roomCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoomControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomUpdateRequest := web.RoomUpdateRequest{}
	helper.ReadFromRequestBody(request, &roomUpdateRequest)

	roomId := params.ByName("roomId")
	id, err := strconv.Atoi(roomId)
	helper.PanicIfError(err)

	roomUpdateRequest.Id = id

	roomResponse := controller.RoomService.Update(request.Context(), roomUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoomControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomId := params.ByName("roomId")
	id, err := strconv.Atoi(roomId)
	helper.PanicIfError(err)

	controller.RoomService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoomControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomId := params.ByName("roomId")
	id, err := strconv.Atoi(roomId)
	helper.PanicIfError(err)

	roomResponse := controller.RoomService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoomControllerImpl) FindByAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomResponses := controller.RoomService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
