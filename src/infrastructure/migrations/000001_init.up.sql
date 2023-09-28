-- support schema
CREATE SCHEMA IF NOT EXISTS support;

-- vehicle
CREATE TABLE IF NOT EXISTS support.vehicle (
	id varchar(36) PRIMARY KEY,
    driver_id VARCHAR(255),
    model VARCHAR(255),
    make VARCHAR(255),
    PlateNumber VARCHAR(255),
    ImageUrl VARCHAR(255),
	created_at timestamp,
	updated_at timestamp
);

-- idx_driver_vehicle
CREATE INDEX IF NOT EXISTS idx_driver_vehicle ON support.vehicle(driver_id);




