package dtos

import (
	"encoding/json"
	pb "github.com/Azamjon99/logistic-vehicle-service/src/application/protos/logistics_vehicle"

)

func ConvertToGetVehicleResponse(result interface{}) (*pb.GetVehicleResponse, error) {
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

    response := &pb.GetVehicleResponse{
        Vehicle: Vehicle,
    }

    return response, nil
}