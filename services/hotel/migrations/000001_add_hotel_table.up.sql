CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE hotel (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    country_code CHAR(2) NOT NULL CHECK (country_code ~ '^[a-z]{2}$'),
    city_slug VARCHAR(100) NOT NULL CHECK (city_slug ~ '^[a-z0-9]+(-[a-z0-9]+)*$'),
    title VARCHAR(100) NOT NULL,
    slug VARCHAR(100) NOT NULL CHECK (slug ~ '^[a-z0-9]+(-[a-z0-9]+)*$'),
    owner_id BIGINT NOT NULL,
    description TEXT,
    address TEXT NOT NULL,
    location GEOGRAPHY(Point, 4326) NOT NULL,
    rating NUMERIC(3,2) CHECK (rating >= 0 AND rating <= 5),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(country_code, city_slug, slug)
);

CREATE INDEX hotels_location_idx ON hotel USING GIST (location);
CREATE INDEX hotels_owner_idx ON hotel (owner_id);

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