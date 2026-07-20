 

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT,
    phone VARCHAR(20),
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL
);



 

CREATE TABLE hotels (
    id SERIAL PRIMARY KEY,
    hotel_name VARCHAR(100) NOT NULL,
    star INT,
    average_price NUMERIC(10,2)
);


 
CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    hotel_id INT NOT NULL REFERENCES hotels(id),
    room_type VARCHAR(30) NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    total_rooms INT NOT NULL,
    capacity INT NOT NULL
);


 

CREATE TABLE bookings (

    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    room_id INT NOT NULL REFERENCES rooms(id),
    room_count INT NOT NULL,
    check_in DATE NOT NULL,
    check_out DATE NOT NULL,
    guest_count INT NOT NULL,
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

 