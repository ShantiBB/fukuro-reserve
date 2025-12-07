CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE hotel (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    owner_id BIGINT NOT NULL,
    description TEXT,
    address TEXT NOT NULL,
    location GEOGRAPHY(Point, 4326) NOT NULL,
    rating NUMERIC(3,2) CHECK (rating >= 0 AND rating <= 5),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX hotels_location_idx ON hotel USING GIST (location);
CREATE INDEX hotels_owner_idx ON hotel (owner_id);
CREATE INDEX hotels_name_idx ON hotel (name);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_hotels_updated_at
    BEFORE UPDATE ON hotel
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();