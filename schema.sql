CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    message TEXT,
    location GEOGRAPHY(Point, 4326),
    user_id TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    content TEXT,
    user_id TEXT,
    message_id SERIAL REFERENCES messages(id),
    created_at TIMESTAMP DEFAULT current_timestamp
);