CREATE TABLE skills (
    id SERIAL PRIMARY KEY,
    profile_code BIGINT NOT NULL,
    skill VARCHAR(255) NOT NULL,
    level VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);