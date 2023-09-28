package grpc

import (
	"context"

	"github.com/Azamjon99/logistic-vehicle-service/src/application/services"
	pb "github.com/Azamjon99/logistic-vehicle-service/src/application/protos/logistics_vehicle"
)

type Server struct {
	pb.VehicleServiceServer
	vehicleApp services.VehicleApplicationService
}

func NewServer(vehicleApp services.VehicleApplicationService) *Server {
	return &Server{
		vehicleApp: vehicleApp,
	}
}

func (s *Server) CreateVehicle(ctx context.Context, r *pb.CreateVehicleRequest) (*pb.CreateVehicleResponse, error) {
	return s.vehicleApp.CreateVehicle(ctx, r)
}

func (s *Server) UpdateVehicle(ctx context.Context, r *pb.UpdateVehicleRequest) (*pb.UpdateVehicleResponse, error) {
	return s.vehicleApp.UpdateVehicle(ctx, r)
}

func (s *Server) DeleteVehicle(ctx context.Context, r *pb.DeleteVehicleRequest) (*pb.DeleteVehicleResponse, error) {
	return s.vehicleApp.DeleteVehicle(ctx, r)
}

func (s *Server) GetVehicle(ctx context.Context, r *pb.GetVehicleRequest) (*pb.GetVehicleResponse, error) {
	return s.vehicleApp.GetVehicle(ctx, r)
}

func (s *Server) ListVehicle(ctx context.Context, r *pb.ListVehicleRequest) (*pb.ListVehicleResponse, error) {
	return s.vehicleApp.ListVehicleByDriver(ctx, r)
}