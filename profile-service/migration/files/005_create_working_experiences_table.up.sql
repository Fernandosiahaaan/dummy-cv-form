CREATE TABLE IF NOT EXISTS working_experiences (
    id SERIAL PRIMARY KEY,
    profile_code BIGINT NOT NULL,
    experience TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (profile_code) REFERENCES profiles (profile_code)
);

CREATE INDEX IF NOT EXISTS idx_profile_code ON working_experiences (profile_code);
CREATE INDEX IF NOT EXISTS idx_deleted_at ON working_experiences (deleted_at);