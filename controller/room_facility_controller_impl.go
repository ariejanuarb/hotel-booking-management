package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type RoomFacilityControllerImpl struct {
	RoomFacilityService service.RoomFacilityService
}

// constructor Category
// dependency : repository , db ,validate
func NewRoomFacilityController(roomfacilityService service.RoomFacilityService) RoomFacilityController {
	return &RoomFacilityControllerImpl{
		RoomFacilityService: roomfacilityService,
	}
}

func (controller *RoomFacilityControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomfacilityCreateRequest := web.RoomFacilityCreateRequest{}
	helper.ReadFromRequestBody(request, &roomfacilityCreateRequest)

	roomfacilityResponse := controller.RoomFacilityService.Create(request.Context(), roomfacilityCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomfacilityResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoomFacilityControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomfacilityUpdateRequest := web.RoomFacilityUpdateRequest{}
	helper.ReadFromRequestBody(request, &roomfacilityUpdateRequest)

	roomfacilityId := params.ByName("roomfacilityId")
	id, err := strconv.Atoi(roomfacilityId)
	helper.PanicIfError(err)

	roomfacilityUpdateRequest.Id = id

	roomfacilityResponse := controller.RoomFacilityService.Update(request.Context(), roomfacilityUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomfacilityResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoomFacilityControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomfacilityId := params.ByName("roomfacilityId")
	id, err := strconv.Atoi(roomfacilityId)
	helper.PanicIfError(err)

	controller.RoomFacilityService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoomFacilityControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomfacilityId := params.ByName("roomfacilityId")
	id, err := strconv.Atoi(roomfacilityId)
	helper.PanicIfError(err)

	roomfacilityResponse := controller.RoomFacilityService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomfacilityResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoomFacilityControllerImpl) FindByAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomfacilityResponses := controller.RoomFacilityService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roomfacilityResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
