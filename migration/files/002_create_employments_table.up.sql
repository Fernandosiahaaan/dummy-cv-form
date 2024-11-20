CREATE TABLE IF NOT EXISTS employments (
    id SERIAL PRIMARY KEY NOT NULL,
    profile_code BIGINT NOT NULL,
    job_title VARCHAR(255) NOT NULL,
    employer VARCHAR(255) NOT NULL,
    start_date VARCHAR(255) NOT NULL,
    end_date VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (profile_code) REFERENCES profiles (profile_code)
);

CREATE INDEX IF NOT EXISTS idx_profile_code ON employments (profile_code);
CREATE INDEX IF NOT EXISTS idx_deleted_at ON employments (deleted_at);
