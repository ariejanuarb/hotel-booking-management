package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type UserHotelControllerImpl struct {
	UserHotelService service.UserHotelService
}

func NewUserHotelController(userHotelService service.UserHotelService) UserHotelController {
	return &UserHotelControllerImpl{
		UserHotelService: userHotelService,
	}
}

func (controller *UserHotelControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userHotelCreateRequest := web.UserHotelCreateRequest{}
	helper.ReadFromRequestBody(request, &userHotelCreateRequest)

	userHotelResponse := controller.UserHotelService.Create(request.Context(), userHotelCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userHotelResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserHotelControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userHotelUpdateRequest := web.UserHotelUpdateRequest{}
	helper.ReadFromRequestBody(request, &userHotelUpdateRequest)

	userHotelId := params.ByName("userHotelId")
	id, err := strconv.Atoi(userHotelId)
	helper.PanicIfError(err)

	userHotelUpdateRequest.Id = id

	userHotelResponse := controller.UserHotelService.Update(request.Context(), userHotelUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userHotelResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserHotelControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userHotelId := params.ByName("userHotelId")
	id, err := strconv.Atoi(userHotelId)
	helper.PanicIfError(err)

	controller.UserHotelService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserHotelControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userHotelId := params.ByName("userHotelId")
	id, err := strconv.Atoi(userHotelId)
	helper.PanicIfError(err)

	userHotelResponse := controller.UserHotelService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userHotelResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserHotelControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userHotelResponses := controller.UserHotelService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userHotelResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
