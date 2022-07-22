package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type UserProfileControllerImpl struct {
	UserProfileService service.UserProfileService
}

func NewUserProfileController(userProfileService service.UserProfileService) UserProfileController {
	return &UserProfileControllerImpl{
		UserProfileService: userProfileService,
	}
}

func (controller *UserProfileControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userProfileCreateRequest := web.UserProfileCreateRequest{}
	helper.ReadFromRequestBody(request, &userProfileCreateRequest)

	userProfileResponse := controller.UserProfileService.Create(request.Context(), userProfileCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userProfileResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserProfileControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userProfileUpdateRequest := web.UserProfileUpdateRequest{}
	helper.ReadFromRequestBody(request, &userProfileUpdateRequest)

	userProfileId := params.ByName("userProfileId")
	id, err := strconv.Atoi(userProfileId)
	helper.PanicIfError(err)

	userProfileUpdateRequest.Id = id

	userProfileResponse := controller.UserProfileService.Update(request.Context(), userProfileUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userProfileResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserProfileControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userProfileId := params.ByName("userProfileId")
	id, err := strconv.Atoi(userProfileId)
	helper.PanicIfError(err)

	controller.UserProfileService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserProfileControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userProfileId := params.ByName("userProfileId")
	id, err := strconv.Atoi(userProfileId)
	helper.PanicIfError(err)

	userProfileResponse := controller.UserProfileService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userProfileResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserProfileControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userProfileResponses := controller.UserProfileService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userProfileResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
