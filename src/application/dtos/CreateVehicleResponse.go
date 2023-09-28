package dtos

import (
	"encoding/json"
	pb "github.com/Azamjon99/logistic-vehicle-service/src/application/protos/logistics_vehicle"

)

func ConvertToCreateVehicleResponse(result interface{}) (*pb.CreateVehicleResponse, error) {
    resultJSON, err := json.Marshal(result)
    if err != nil {
        return nil, err
    }

    var resultMap map[string]interface{}
    if err := json.Unmarshal(resultJSON, &resultMap); err != nil {
        return nil, err
    }


    Vehicle := &pb.Vehicle{
    }

    response := &pb.CreateVehicleResponse{
        Vehicle: Vehicle,
    }

    return response, nil
}