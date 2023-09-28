package services

import (
	"context"

	"github.com/Azamjon99/logistic-vehicle-service/src/application/dtos"
	pb "github.com/Azamjon99/logistic-vehicle-service/src/application/protos/logistics_vehicle"
	"github.com/Azamjon99/logistic-vehicle-service/src/domain/models"
	vehicleservices "github.com/Azamjon99/logistic-vehicle-service/src/domain/services"
)

type VehicleApplicationService interface {
	CreateVehicle(ctx context.Context, req *pb.CreateVehicleRequest) (*pb.CreateVehicleResponse, error)
	UpdateVehicle(ctx context.Context, req *pb.UpdateVehicleRequest) (*pb.UpdateVehicleResponse, error)
	DeleteVehicle(ctx context.Context, req *pb.DeleteVehicleRequest) (*pb.DeleteVehicleResponse, error)
	GetVehicle(ctx context.Context, req *pb.GetVehicleRequest) (*pb.GetVehicleResponse, error)
	ListVehicleByDriver(ctx context.Context, req *pb.ListVehicleRequest) (*pb.ListVehicleResponse, error)
}

type vehicleAppServiceImpl struct {
	vehicleSvc vehicleservices.VehicleService
}

func NewVehicleApplicationService(vehicleSvc vehicleservices.VehicleService) VehicleApplicationService {
	return &vehicleAppServiceImpl{
		vehicleSvc: vehicleSvc,
	}
}

func (s *vehicleAppServiceImpl) CreateVehicle(ctx context.Context, req *pb.CreateVehicleRequest) (*pb.CreateVehicleResponse, error) {


	 result, err := s.vehicleSvc.CreateVehicle(ctx, req.DriverId, req.Model, req.Make); 
	if err != nil {
		return nil, err
	}

    response, err :=dtos.ConvertToCreateVehicleResponse(result)
    if err != nil {
        return nil, err
    }
    return response, nil
}

func (s *vehicleAppServiceImpl) UpdateVehicle(ctx context.Context, req *pb.UpdateVehicleRequest) (*pb.UpdateVehicleResponse, error) {
	vehicle := &models.Vehicle{
		ID:      req.Vehicle.Id,
		DriverID: req.Vehicle.DriverId,
		Model:    req.Vehicle.Model,
		Make:     req.Vehicle.Make,
	}

	 err := s.vehicleSvc.UpdateVehicle(ctx, vehicle); 
   if err != nil {
	   return nil, err
   }

 
   response := &pb.UpdateVehicleResponse{}
	return response, nil

}



func (s *vehicleAppServiceImpl) DeleteVehicle(ctx context.Context, req *pb.DeleteVehicleRequest) (*pb.DeleteVehicleResponse, error) {
	if err := s.vehicleSvc.DeleteVehicle(ctx, req.VehicleId); err != nil {
		return nil, err
	}

	response := &pb.DeleteVehicleResponse{}
	return response, nil
}

func (s *vehicleAppServiceImpl) GetVehicle(ctx context.Context, req *pb.GetVehicleRequest) (*pb.GetVehicleResponse, error) {
	result, err := s.vehicleSvc.GetVehicle(ctx, req.VehicleId)
	if err != nil {
		return nil, err
	}
    response, err :=dtos.ConvertToGetVehicleResponse(result)


	return response, nil
}

func (s *vehicleAppServiceImpl) ListVehicleByDriver(ctx context.Context, req *pb.ListVehicleRequest) (*pb.ListVehicleResponse, error) {
	vehicles, err := s.vehicleSvc.ListVehicleByDriver(ctx, req.DriverId)
	if err != nil {
		return nil, err
	}
	response := &pb.ListVehicleResponse{
		Vehicles: dtos.ConvertToVehicleProtos(vehicles),
	}	
	return response, nil
}


