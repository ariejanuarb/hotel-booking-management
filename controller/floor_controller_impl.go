package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type FloorControllerImpl struct {
	FloorService service.FloorService
}

func NewFloorController(floorService service.FloorService) FloorController {
	return &FloorControllerImpl{
		FloorService: floorService,
	}
}

func (controller *FloorControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	floorCreateRequest := web.FloorCreateRequest{}
	helper.ReadFromRequestBody(request, &floorCreateRequest)

	floorResponse := controller.FloorService.Create(request.Context(), floorCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   floorResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FloorControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	floorUpdateRequest := web.FloorUpdateRequest{}
	helper.ReadFromRequestBody(request, &floorUpdateRequest)

	floorId := params.ByName("floorId")
	id, err := strconv.Atoi(floorId)
	helper.PanicIfError(err)

	floorUpdateRequest.Id = id

	floorResponse := controller.FloorService.Update(request.Context(), floorUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   floorResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FloorControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	floorId := params.ByName("floorId")
	id, err := strconv.Atoi(floorId)
	helper.PanicIfError(err)

	controller.FloorService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FloorControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	floorId := params.ByName("floorId")
	id, err := strconv.Atoi(floorId)
	helper.PanicIfError(err)

	floorResponse := controller.FloorService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   floorResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FloorControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	floorResponses := controller.FloorService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   floorResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
