CREATE TABLE IF NOT EXISTS profiles (
    profile_code BIGSERIAL PRIMARY KEY,
    wanted_job_title VARCHAR(255),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(20),
    country VARCHAR(100),
    city VARCHAR(100),
    address VARCHAR(255),
    postal_code VARCHAR(10),
    driving_license VARCHAR(50),
    nationality VARCHAR(100),
    place_of_birth VARCHAR(100),
    date_of_birth VARCHAR(100) NOT NULL,
    photo_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

