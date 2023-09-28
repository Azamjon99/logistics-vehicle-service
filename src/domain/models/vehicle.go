package models
import "time"

type Vehicle struct {
    ID          string    `json:"id"`
    DriverID    string    `json:"driver_id"`
    Model       string    `json:"model"`
    Make        string    `json:"make"`
    PlateNumber string    `json:"plate_number"`
    ImageUrl    string    `json:"image_url"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}