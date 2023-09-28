package dtos

import (

	pb "github.com/Azamjon99/logistic-vehicle-service/src/application/protos/logistics_vehicle"
	"github.com/Azamjon99/logistic-vehicle-service/src/domain/models"
)

func convertToVehicleProto(vehicle *models.Vehicle) *pb.Vehicle {
	return &pb.Vehicle{
		Id:        vehicle.ID,
		DriverId:  vehicle.DriverID,
		Model:     vehicle.Model,
		Make:      vehicle.Make,
	}
}