package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type HotelControllerImpl struct {
	HotelService service.HotelService
}

func NewHotelController(hotelService service.HotelService) HotelController {
	return &HotelControllerImpl{
		HotelService: hotelService,
	}
}

func (controller *HotelControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	hotelCreateRequest := web.HotelCreateRequest{}
	helper.ReadFromRequestBody(request, &hotelCreateRequest)

	hotelResponse := controller.HotelService.Create(request.Context(), hotelCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   hotelResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *HotelControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	hotelUpdateRequest := web.HotelUpdateRequest{}
	helper.ReadFromRequestBody(request, &hotelUpdateRequest)

	hotelId := params.ByName("hotelId")
	id, err := strconv.Atoi(hotelId)
	helper.PanicIfError(err)

	hotelUpdateRequest.Id = id

	hotelResponse := controller.HotelService.Update(request.Context(), hotelUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   hotelResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *HotelControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	hotelId := params.ByName("hotelId")
	id, err := strconv.Atoi(hotelId)
	helper.PanicIfError(err)

	controller.HotelService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *HotelControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	hotelId := params.ByName("hotelId")
	id, err := strconv.Atoi(hotelId)
	helper.PanicIfError(err)

	hotelResponse := controller.HotelService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   hotelResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *HotelControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	hotelResponses := controller.HotelService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   hotelResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
