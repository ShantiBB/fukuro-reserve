DROP TRIGGER IF EXISTS update_rooms_updated_at ON room;

DROP INDEX IF EXISTS rooms_hotel_id_idx;
DROP INDEX IF EXISTS rooms_type_idx;
DROP INDEX IF EXISTS rooms_status_idx;
DROP INDEX IF EXISTS rooms_price_idx;

DROP TABLE IF EXISTS room;

DROP TYPE IF EXISTS room_type;
DROP TYPE IF EXISTS room_status;
