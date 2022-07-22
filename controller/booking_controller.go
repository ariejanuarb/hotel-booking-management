package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BookingController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Cancel(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Discount(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ResponseDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	HotelRevenue(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Revenue(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
