CREATE TYPE room_type AS ENUM ('single', 'double', 'suite', 'deluxe', 'family', 'presidential');
CREATE TYPE room_status AS ENUM ('available', 'occupied', 'maintenance', 'cleaning');

CREATE TABLE room (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hotel_id UUID NOT NULL REFERENCES hotel(id) ON DELETE CASCADE,
    room_number VARCHAR(10) NOT NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    type room_type NOT NULL,
    status room_status NOT NULL DEFAULT 'available',
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
CREATE INDEX rooms_status_idx ON room (status) WHERE status = 'available';
CREATE INDEX rooms_price_idx ON room (price);

CREATE TRIGGER update_rooms_updated_at
    BEFORE UPDATE ON room
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
