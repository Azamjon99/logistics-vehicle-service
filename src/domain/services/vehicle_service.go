package services

import (
	"context"
	"time"

	"github.com/Azamjon99/logistic-vehicle-service/src/domain/models"
	repositories "github.com/Azamjon99/logistic-vehicle-service/src/domain/repository"
)

type VehicleService interface {
	CreateVehicle(ctx context.Context, driverID, model, make string) (*models.Vehicle, error)
	UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) error
	DeleteVehicle(ctx context.Context, vehicleID string) error
	GetVehicle(ctx context.Context, vehicleID string) (*models.Vehicle, error)
	ListVehicleByDriver(ctx context.Context, driverID string) ([]*models.Vehicle, error)
	}

type vehicleSvcImpl struct {
	vehicleRepo repositories.VehicleRepository
}

func NewVehicleSvc(vehicleRepo  repositories.VehicleRepository) VehicleService {
	return &vehicleSvcImpl{
		vehicleRepo: vehicleRepo,
	}
}
func (s *vehicleSvcImpl) CreateVehicle(ctx context.Context,  driverID string, model , make string)(*models.Vehicle, error){
	vehicle :=  &models.Vehicle{
		DriverID: driverID,
		Model: model,
		Make: make,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err:= s.vehicleRepo.SaveVehicle(ctx, vehicle)

	if err != nil {
		return nil, err
	}

	return vehicle, nil;
}


func (s *vehicleSvcImpl) UpdateVehicle(ctx context.Context,  vehicle *models.Vehicle )( error){


	err:= s.vehicleRepo.UpdateVehicle(ctx, vehicle)

	if err != nil {
		return  err
	}

	return  nil;
}


func (s *vehicleSvcImpl) DeleteVehicle(ctx context.Context,  vehicleID string )( error){


	err:= s.vehicleRepo.DeleteVehicle(ctx, vehicleID)

	if err != nil {
		return  err
	}

	return nil;
}

func (s *vehicleSvcImpl) GetVehicle(ctx context.Context,  vehicleID string )( *models.Vehicle, error){


	vehicle, err:= s.vehicleRepo.GetVehicle(ctx, vehicleID)

	if err != nil {
		return nil, err
	}

	return vehicle , nil;
}

func (s *vehicleSvcImpl) ListVehicleByDriver(ctx context.Context,  vehicleID string )( []*models.Vehicle, error){


	vehicle, err:= s.vehicleRepo.ListVehicleByDriver(ctx, vehicleID)

	if err != nil {
		return nil, err
	}

	return vehicle , nil;
}