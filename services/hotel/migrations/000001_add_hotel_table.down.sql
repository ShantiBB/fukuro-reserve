DROP TRIGGER IF EXISTS update_hotels_updated_at ON hotel;

DROP FUNCTION IF EXISTS update_updated_at_column();

DROP INDEX IF EXISTS hotels_location_idx;
DROP INDEX IF EXISTS hotels_owner_idx;
DROP INDEX IF EXISTS hotels_name_idx;

DROP TABLE IF EXISTS hotel;

DROP EXTENSION IF EXISTS postgis;
