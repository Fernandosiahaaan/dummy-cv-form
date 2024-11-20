CREATE TABLE IF NOT EXISTS skills (
    id SERIAL PRIMARY KEY,
    profile_code BIGINT NOT NULL,
    skill VARCHAR(255) NOT NULL,
    level VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (profile_code) REFERENCES profiles (profile_code)
);

CREATE INDEX IF NOT EXISTS idx_profile_code ON skills (profile_code);
CREATE INDEX IF NOT EXISTS idx_deleted_at ON skills (deleted_at);
