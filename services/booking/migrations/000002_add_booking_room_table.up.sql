CREATE TABLE IF NOT EXISTS booking_room (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,

    room_id UUID NOT NULL,

    adults INT8 NOT NULL DEFAULT 1 CHECK (adults >= 1 AND adults <= 10),
    children INT8 NOT NULL DEFAULT 0 CHECK (children >= 0 AND children <= 10),

    price_per_night NUMERIC(12,2) NOT NULL CHECK (price_per_night >= 0),

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_booking_room_booking ON booking_room(booking_id);
CREATE INDEX IF NOT EXISTS idx_booking_room_room ON booking_room(room_id);
