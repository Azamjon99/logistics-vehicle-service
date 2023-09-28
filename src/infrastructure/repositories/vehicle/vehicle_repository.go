package vehicle

import (
	"context"
	"github.com/Azamjon99/logistic-vehicle-service/src/domain/models"
	repositories "github.com/Azamjon99/logistic-vehicle-service/src/domain/repository"

	"gorm.io/gorm"
)

const(

  vehicleTable = "support.vehicles"
)

type vehiclerepoImpl struct {
	db *gorm.DB
}
 
func NewVehicleRepository(db *gorm.DB) repositories.VehicleRepository{
	return &vehiclerepoImpl{
		db:db,
	}
}

func (r *vehiclerepoImpl)SaveVehicle(ctx context.Context, vehicle *models.Vehicle) error{
	result := r.db.WithContext(ctx).Table(vehicleTable).Create(&vehicle)

	if result.Error != nil{
		return result.Error
	}

	return nil;
}

func (r *vehiclerepoImpl)UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) error{
	result := r.db.WithContext(ctx).Table(vehicleTable).Save(&vehicle)

	if result.Error != nil{
		return result.Error
	}

	return nil;
}



func (r *vehiclerepoImpl) DeleteVehicle(ctx context.Context, vehicleID string) error {
	result := r.db.WithContext(ctx).Table(vehicleTable).Where("vehicle_id = ?", vehicleID).Delete(&models.Vehicle{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *vehiclerepoImpl)ListVehicleByDriver(ctx context.Context, driverID string)([]*models.Vehicle, error){
	var vehicle []*models.Vehicle
	result := r.db.WithContext(ctx).Table(vehicleTable).First(&vehicle, "order_id = ?", driverID)

	if result.Error != nil {
		return nil, result.Error
	}

	return vehicle, nil;
}

func (r *vehiclerepoImpl)GetVehicle(ctx context.Context, vehicleID string)(*models.Vehicle, error){
	var vehicle *models.Vehicle
	result := r.db.WithContext(ctx).Table(vehicleTable).First(&vehicle, "vehicle_id = ?", vehicleID)

	if result.Error != nil {
		return nil, result.Error
	}

	return vehicle, nil;
}