CREATE TABLE IF NOT EXISTS educations (
    id SERIAL PRIMARY KEY,
    profile_code BIGINT NOT NULL,
    school VARCHAR(255) NOT NULL,
    degree VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    city VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (profile_code) REFERENCES profiles (profile_code)
);

CREATE INDEX IF NOT EXISTS idx_profile_code ON educations (profile_code);
CREATE INDEX IF NOT EXISTS idx_deleted_at ON educations (deleted_at);
