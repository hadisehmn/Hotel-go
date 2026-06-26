CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT,
    phone VARCHAR(20),
    password_hash VARCHAR(255) NOT NULL
);



CREATE TABLE hotels (
    id SERIAL PRIMARY KEY,
    hotel_name VARCHAR(100) NOT NULL,
    star INT
);

CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    hotel_id INT REFERENCES hotels(id),
    room_name VARCHAR(100),
    room_type VARCHAR(50),
    price NUMERIC(10,2)
);


CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    room_id INT REFERENCES rooms(id),
    check_in TIMESTAMP,
    check_out TIMESTAMP,
    guest_count INT,
    total_price NUMERIC(10,2),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE booking_guests (
    id SERIAL PRIMARY KEY,
    booking_id INT REFERENCES bookings(id),
    guest_type VARCHAR(20)
);

CREATE TABLE pricing_rules (
    id SERIAL PRIMARY KEY,
    room_id INT REFERENCES rooms(id),
    guest_type VARCHAR(20),
    price NUMERIC(10,2)
);

 