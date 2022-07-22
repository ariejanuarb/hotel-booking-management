package app

import (
	"booking-hotel/controller"
	"booking-hotel/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userProfileController controller.UserProfileController, userHotelController controller.UserHotelController, hotelController controller.HotelController, floorController controller.FloorController, employeeController controller.EmployeeController, roleController controller.RoleController, bookingController controller.BookingController, facilityController controller.FacilityController, roomfacilityController controller.RoomFacilityController, roomController controller.RoomController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/view-all-user-profile", userProfileController.FindAll)
	router.GET("/api/view-user-profile-id/:userProfileId", userProfileController.FindById)
	router.POST("/api/create-owner-profile", userProfileController.Create)
	router.PUT("/api/update-user-profile/:userProfileId", userProfileController.Update)
	router.DELETE("/api/delete-user-profile/:userProfileId", userProfileController.Delete)

	router.GET("/api/view-all-user-hotel", userHotelController.FindAll)
	router.GET("/api/view-user-hotel-id/:userHotelId", userHotelController.FindById)
	router.POST("/api/assign-employee-to-hotel", userHotelController.Create)
	router.PUT("/api/move-employee-to-another-hotel/:userHotelId", userHotelController.Update)
	router.DELETE("/api/delete-user-hotel-id/:userHotelId", userHotelController.Delete)

	router.GET("/api/view-all-hotel", hotelController.FindAll)
	router.GET("/api/view-hotel-id/:hotelId", hotelController.FindById)
	router.POST("/api/create-hotel", hotelController.Create)
	router.PUT("/api/update-hotel-id/:hotelId", hotelController.Update)
	router.DELETE("/api/delete-hotel-id/:hotelId", hotelController.Delete)

	router.GET("/api/view-all-floor", floorController.FindAll)
	router.GET("/api/view-floor-id/:floorId", floorController.FindById)
	router.POST("/api/create-floor", floorController.Create)
	router.PUT("/api/update-floor-id/:floorId", floorController.Update)
	router.DELETE("/api/delete-floor-id/:floorId", floorController.Delete)

	router.GET("/api/view-all-employee", employeeController.FindAll)
	router.GET("/api/view-employee-id/:employeeId", employeeController.FindById)
	router.POST("/api/create-new-employee", employeeController.Create)
	router.PUT("/api/update-employee-id/:employeeId", employeeController.Update)
	router.DELETE("/api/delete-employee-id/:employeeId", employeeController.Delete)

	router.GET("/api/view-all-role", roleController.FindAll)
	router.GET("/api/view-role-id/:roleId", roleController.FindById)
	router.POST("/api/create-new-role", roleController.Create)
	router.PUT("/api/update-role-id/:roleId", roleController.Update)
	router.DELETE("/api/delete-role-id/:roleId", roleController.Delete)

	router.POST("/api/booking", bookingController.Create)
	router.PUT("/api/bookingStatus/:bookingId", bookingController.Cancel)
	router.GET("/api/booking/:bookingId", bookingController.FindById)
	router.GET("/api/booking", bookingController.FindAll)

	router.PUT("/api/bookingDiscount/:bookingId", bookingController.Discount)
	router.PUT("/api/responseDiscount/:bookingId", bookingController.ResponseDiscount)

	router.GET("/api/bookingDiscount", bookingController.FindAllDiscount)

	router.GET("/api/hotelRevenue", bookingController.HotelRevenue)
	router.GET("/api/Revenue", bookingController.Revenue)

	router.GET("/api/facilitas", facilityController.FindByAll)
	router.GET("/api/facilitas/:facilityId", facilityController.FindById)
	router.POST("/api/facilitas", facilityController.Create)
	router.PUT("/api/facilitas/:facilityId", facilityController.Update)
	router.DELETE("/api/facilitas/:facilityId", facilityController.Delete)

	router.GET("/api/roomfacilitas", roomfacilityController.FindByAll)
	router.GET("/api/roomfacilitas/:roomfacilityId", roomfacilityController.FindById)
	router.POST("/api/roomfacilitas", roomfacilityController.Create)
	router.PUT("/api/roomfacilitas/:roomfacilityId", roomfacilityController.Update)
	router.DELETE("/api/roomfacilitas/:roomfacilityId", roomfacilityController.Delete)

	router.GET("/api/rooms", roomController.FindByAll)
	router.GET("/api/rooms/:roomId", roomController.FindById)
	router.POST("/api/rooms", roomController.Create)
	router.PUT("/api/rooms/:roomId", roomController.Update)
	router.DELETE("/api/rooms/:roomId", roomController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
