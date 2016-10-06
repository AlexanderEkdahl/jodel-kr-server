CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    user_id BIT(256),
    message TEXT,
    location GEOGRAPHY(Point, 4326),
    created_at TIMESTAMP DEFAULT current_timestamp
);
