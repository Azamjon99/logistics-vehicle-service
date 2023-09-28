package dtos


import ("github.com/Azamjon99/logistic-vehicle-service/src/domain/models"
pb "github.com/Azamjon99/logistic-vehicle-service/src/application/protos/logistics_vehicle"
)
func ConvertToVehicleProtos(vehicles []*models.Vehicle) []*pb.Vehicle {
	vehicleProtos := make([]*pb.Vehicle, len(vehicles))
	for i, vehicle := range vehicles {
		vehicleProtos[i] = ConvertToVehicleProtoList(vehicle)
	}
	return vehicleProtos
}
func ConvertToVehicleProtoList(vehicle *models.Vehicle) *pb.Vehicle {
	return &pb.Vehicle{
		Id:        vehicle.ID,
		DriverId:  vehicle.DriverID,
		Model:     vehicle.Model,
		Make:      vehicle.Make,
	}
}
