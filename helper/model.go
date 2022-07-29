package helper

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
)

func ToFacilityResponse(facility domain.Facility) web.FacilityResponse {
	return web.FacilityResponse{
		Id:          facility.Id,
		Type:        facility.Type,
		Description: facility.Description,
	}
}

func ToFacilityResponses(facilitas []domain.Facility) []web.FacilityResponse {
	var facilityResponses []web.FacilityResponse
	for _, facility := range facilitas {
		facilityResponses = append(facilityResponses, ToFacilityResponse(facility))
	}
	return facilityResponses
}

func ToRoomFacilityResponse(roomfacility domain.RoomFacility) web.RoomFacilityResponse {
	return web.RoomFacilityResponse{
		Id:         roomfacility.Id,
		RoomId:     roomfacility.RoomId,
		FacilityId: roomfacility.FacilityId,
	}
}

func ToRoomFacilityResponses(roomfacilitas []domain.RoomFacility) []web.RoomFacilityResponse {
	var roomfacilityResponses []web.RoomFacilityResponse
	for _, roomfacility := range roomfacilitas {
		roomfacilityResponses = append(roomfacilityResponses, ToRoomFacilityResponse(roomfacility))
	}
	return roomfacilityResponses
}

func ToRoomResponse(room domain.Room) web.RoomResponse {
	return web.RoomResponse{
		Id:           room.Id,
		Name:         room.Name,
		Capacity:     room.Capacity,
		PricePerHour: room.PricePerHour,
		PricePerDay:  room.PricePerDay,
		Facility_id:  room.Facility_id,
	}
}

func ToRoomResponses(rooms []domain.Room) []web.RoomResponse {
	var roomResponses []web.RoomResponse
	for _, room := range rooms {
		roomResponses = append(roomResponses, ToRoomResponse(room))
	}
	return roomResponses
}

func ToUserProfileResponse(userProfile domain.UserProfile) web.UserProfileResponse {
	return web.UserProfileResponse{
		Id:       userProfile.Id,
		Email:    userProfile.Email,
		Password: userProfile.Password,
		Name:     userProfile.Name,
		Gender:   userProfile.Gender,
	}
}

func ToUserProfileResponses(userProfiles []domain.UserProfile) []web.UserProfileResponse {
	var userProfileResponses []web.UserProfileResponse
	for _, userProfile := range userProfiles {
		userProfileResponses = append(userProfileResponses, ToUserProfileResponse(userProfile))
	}
	return userProfileResponses
}

func ToUserHotelResponse(userHotel domain.UserHotel) web.UserHotelResponse {
	return web.UserHotelResponse{
		Id:            userHotel.Id,
		UserProfileId: userHotel.UserProfileId,
		HotelId:       userHotel.HotelId,
	}
}

func ToUserHotelResponses(userHotels []domain.UserHotel) []web.UserHotelResponse {
	var userHotelResponses []web.UserHotelResponse
	for _, userHotel := range userHotels {
		userHotelResponses = append(userHotelResponses, ToUserHotelResponse(userHotel))
	}
	return userHotelResponses
}

func ToHotelResponse(hotel domain.Hotel) web.HotelResponse {
	return web.HotelResponse{
		Id:       hotel.Id,
		Name:     hotel.Name,
		Address:  hotel.Address,
		Province: hotel.Province,
		City:     hotel.City,
		ZipCode:  hotel.ZipCode,
		Star:     hotel.Star,
	}
}

func ToHotel(hotel web.HotelResponse) domain.Hotel {
	return domain.Hotel{
		Id:       hotel.Id,
		Name:     hotel.Name,
		Address:  hotel.Address,
		Province: hotel.Province,
		City:     hotel.City,
		ZipCode:  hotel.ZipCode,
		Star:     hotel.Star,
	}
}

func ToHotelResponses(hotels []domain.Hotel) []web.HotelResponse {
	var hotelResponses []web.HotelResponse
	for _, hotel := range hotels {
		hotelResponses = append(hotelResponses, ToHotelResponse(hotel))
	}
	return hotelResponses
}

func ToFloorResponse(floor domain.Floor) web.FloorResponse {
	return web.FloorResponse{
		Id:      floor.Id,
		Number:  floor.Number,
		HotelId: floor.HotelId,
		RoomId:  floor.RoomId,
	}
}

func ToFloorResponses(floors []domain.Floor) []web.FloorResponse {
	var floorResponses []web.FloorResponse
	for _, floor := range floors {
		floorResponses = append(floorResponses, ToFloorResponse(floor))
	}
	return floorResponses
}

func ToEmployeeResponse(employee domain.Employee) web.EmployeeResponse {
	return web.EmployeeResponse{
		Id:       employee.Id,
		Name:     employee.Name,
		Gender:   employee.Gender,
		Email:    employee.Email,
		Password: employee.Password,
	}
}

func ToEmployee(employee web.EmployeeResponse) domain.Employee {
	return domain.Employee{
		Id:       employee.Id,
		Name:     employee.Name,
		Gender:   employee.Gender,
		Email:    employee.Email,
		Password: employee.Password,
	}
}

func ToEmployeeResponses(employee []domain.Employee) []web.EmployeeResponse {
	var employeeResponses []web.EmployeeResponse
	for _, employee := range employee {
		employeeResponses = append(employeeResponses, ToEmployeeResponse(employee))
	}
	return employeeResponses
}

func ToRoleResponse(role domain.Role) web.RoleResponse {
	return web.RoleResponse{
		Id:   role.Id,
		Type: role.Type,
	}
}

func ToRoleResponses(roles []domain.Role) []web.RoleResponse {
	var roleResponses []web.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}
	return roleResponses
}

func ToBookingResponse(b *domain.Booking) *web.BookingResponse {
	return &web.BookingResponse{
		Id:                 b.Id,
		Status:             b.Status,
		Room_id:            b.Room_id,
		Hotel_id:           b.Hotel_id,
		Pic_name:           b.Pic_name,
		Pic_Contact:        b.Pic_Contact,
		Event_start:        b.Event_start,
		Event_end:          b.Event_end,
		Invoice_number:     b.Invoice_number,
		Invoice_date:       b.Invoice_date,
		Invoice_grandtotal: b.Invoice_grandtotal,
		Discount_request:   b.Discount_request,
		Created_at:         b.Created_at,
		Updated_at:         b.Updated_at,
	}
}

func ToBooking(b web.BookingResponse) *domain.Booking {
	return &domain.Booking{
		Id:                 b.Id,
		Status:             b.Status,
		Room_id:            b.Room_id,
		Hotel_id:           b.Hotel_id,
		Pic_name:           b.Pic_name,
		Pic_Contact:        b.Pic_Contact,
		Event_start:        b.Event_start,
		Event_end:          b.Event_end,
		Invoice_number:     b.Invoice_number,
		Invoice_date:       b.Invoice_date,
		Invoice_grandtotal: b.Invoice_grandtotal,
		Discount_request:   b.Discount_request,
		Created_at:         b.Created_at,
		Updated_at:         b.Updated_at,
	}
}

func ToBookingResponses(b []domain.Booking) []web.BookingResponse {
	var bookingResponses []web.BookingResponse
	for _, booking := range b {
		bookingResponses = append(bookingResponses, *ToBookingResponse(&booking))
	}
	return bookingResponses
}
