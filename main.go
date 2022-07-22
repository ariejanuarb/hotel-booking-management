package main

import (
	"booking-hotel/app"
	"booking-hotel/controller"
	"booking-hotel/helper"
	"booking-hotel/middleware"
	"booking-hotel/repository"
	"booking-hotel/service"
	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	userProfileRepository := repository.NewUserProfileRepository()
	userProfileService := service.NewUserProfileService(userProfileRepository, db, validate)
	userProfileController := controller.NewUserProfileController(userProfileService)

	userHotelRepository := repository.NewUserHotelRepository()
	userHotelService := service.NewUserHotelService(userHotelRepository, db, validate)
	userHotelController := controller.NewUserHotelController(userHotelService)

	hotelRepository := repository.NewHotelRepository()
	hotelService := service.NewHotelService(hotelRepository, db, validate)
	hotelController := controller.NewHotelController(hotelService)

	floorRepository := repository.NewFloorRepository()
	floorService := service.NewFloorService(floorRepository, db, validate)
	floorController := controller.NewFloorController(floorService)

	employeeRepository := repository.NewEmployeeRepository()
	employeeService := service.NewEmployeeService(employeeRepository, db, validate)
	employeeController := controller.NewEmployeeController(employeeService)

	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	bookingRepository := repository.NewBookingRepository()
	bookingService := service.NewBookingService(bookingRepository, db, validate)
	bookingController := controller.NewBookingController(bookingService)

	facilityRepostiory := repository.NewFacilityRepository()
	facilityService := service.NewFacilityService(facilityRepostiory, db, validate)
	facilityController := controller.NewFacilityController(facilityService)

	roomfacilityRepository := repository.NewRoomFacilityRepository()
	roomfacilityService := service.NewRoomFacilityService(roomfacilityRepository, db, validate)
	roomfacilityController := controller.NewRoomFacilityController(roomfacilityService)

	roomRepository := repository.NewRoomRepository()
	roomService := service.NewRoomService(roomRepository, db, validate)
	roomController := controller.NewRoomController(roomService)

	router := app.NewRouter(userProfileController, userHotelController, hotelController, floorController, employeeController, roleController, bookingController, facilityController, roomfacilityController, roomController)

	server := http.Server{
		Addr:    "localhost:3080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
