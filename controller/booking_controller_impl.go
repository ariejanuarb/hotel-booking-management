package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookingControllerImpl struct {
	BookingService service.BookingService
}

func NewBookingController(BookingService *service.BookingServiceImpl) BookingController {
	return &BookingControllerImpl{
		BookingService: BookingService,
	}
}

func (controller *BookingControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var webResponse web.WebResponse
	bookingCreateRequest := web.BookingCreateRequest{}
	helper.ReadFromRequestBody(request, &bookingCreateRequest)

	bookingResponse, err := controller.BookingService.Create(request.Context(), &bookingCreateRequest)
	if err != nil {
		webResponse = web.WebResponse{
			Code:   400,
			Status: "Error",
			Data:   err.Error(),
		}
	} else {
		webResponse = web.WebResponse{
			Code:   200,
			Status: "Ok",
			Data:   bookingResponse,
		}
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) Cancel(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var webResponse web.WebResponse
	updateStatusRequest := web.UpdateRequest{}

	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	updateStatusRequest.Id = id

	bookingResponse, err := controller.BookingService.Cancel(request.Context(), &updateStatusRequest)
	if err != nil {
		webResponse = web.WebResponse{
			Code:   400,
			Status: "Error",
			Data:   err.Error(),
		}
	} else {
		webResponse = web.WebResponse{
			Code:   200,
			Status: "Ok",
			Data:   bookingResponse,
		}
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) Discount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var webResponse web.WebResponse
	updateDiscountRequest := web.Discount{}

	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	updateDiscountRequest.Id = id

	bookingResponse, err := controller.BookingService.Discount(request.Context(), &updateDiscountRequest)
	if err != nil {
		webResponse = web.WebResponse{
			Code:   400,
			Status: "Error",
			Data:   err.Error(),
		}
	} else {
		webResponse = web.WebResponse{
			Code:   200,
			Status: "Ok",
			Data:   bookingResponse,
		}
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) ResponseDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var webResponse web.WebResponse
	responseDiscountRequest := web.ResponseDiscount{}
	helper.ReadFromRequestBody(request, &responseDiscountRequest)

	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	responseDiscountRequest.Id = id

	bookingResponse, err := controller.BookingService.ResponseDiscount(request.Context(), &responseDiscountRequest)
	if err != nil {
		webResponse = web.WebResponse{
			Code:   400,
			Status: "Error",
			Data:   err.Error(),
		}
	} else {
		webResponse = web.WebResponse{
			Code:   200,
			Status: "Ok",
			Data:   bookingResponse,
		}
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) FindAllDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingResponse := controller.BookingService.FindAllDiscount(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	bookingResponse := controller.BookingService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingResponse := controller.BookingService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) HotelRevenue(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	query := request.URL.Query()
	hotelId := query.Get("hotelId")
	invoiceDate := query.Get("invoiceDate")
	id, err := strconv.Atoi(hotelId)
	helper.PanicIfError(err)

	bookingResponse := controller.BookingService.HotelRevenue(request.Context(), id, invoiceDate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) Revenue(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	query := request.URL.Query()
	invoiceDate := query.Get("invoiceDate")

	bookingResponse := controller.BookingService.Revenue(request.Context(), invoiceDate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
