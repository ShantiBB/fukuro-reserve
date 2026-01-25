-- +goose Up
-- +goose StatementBegin
CREATE TYPE room_type AS ENUM ('ROOM_TYPE_UNSPECIFIED',
                               'ROOM_TYPE_SINGLE',
                               'ROOM_TYPE_DOUBLE',
                               'ROOM_TYPE_SUITE',
                               'ROOM_TYPE_DELUXE',
                               'ROOM_TYPE_FAMILY',
                               'ROOM_TYPE_PRESIDENTIAL');

CREATE TYPE room_status AS ENUM ('ROOM_STATUS_AVAILABLE',
                                 'ROOM_STATUS_OCCUPIED',
                                 'ROOM_STATUS_MAINTENANCE',
                                 'ROOM_STATUS_CLEANING');

CREATE TABLE IF NOT EXISTS room (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hotel_id UUID NOT NULL REFERENCES hotel(id) ON DELETE CASCADE,
    room_number VARCHAR(10) NOT NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    type room_type NOT NULL,
    status room_status NOT NULL DEFAULT 'ROOM_STATUS_AVAILABLE',
    price NUMERIC(10,2) NOT NULL CHECK (price > 0),
    capacity INT NOT NULL CHECK (capacity > 0 AND capacity <= 10),
    area_sqm NUMERIC(6,2) NOT NULL,
    floor INT NOT NULL CHECK (floor >= 0),
    amenities TEXT[],
    images TEXT[],
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(hotel_id, room_number)
);

CREATE INDEX rooms_hotel_id_idx ON room (hotel_id);
CREATE INDEX rooms_type_idx ON room (type);
CREATE INDEX rooms_status_idx ON room (status) WHERE status = 'ROOM_STATUS_AVAILABLE';
CREATE INDEX rooms_price_idx ON room (price);

CREATE TRIGGER update_rooms_updated_at
    BEFORE UPDATE ON room
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_rooms_updated_at ON room;

DROP INDEX IF EXISTS rooms_hotel_id_idx;
DROP INDEX IF EXISTS rooms_type_idx;
DROP INDEX IF EXISTS rooms_status_idx;
DROP INDEX IF EXISTS rooms_price_idx;

DROP TABLE IF EXISTS room;

DROP TYPE IF EXISTS room_type;
DROP TYPE IF EXISTS room_status;
-- +goose StatementEnd
