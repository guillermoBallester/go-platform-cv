-- +goose Up
-- +goose StatementBegin

-- Core tables
CREATE TABLE experiences (
    id SERIAL PRIMARY KEY,
    company_name TEXT NOT NULL,
    job_title TEXT NOT NULL,
    location TEXT,
    start_date DATE NOT NULL,
    end_date DATE,                    -- NULL = current position
    description TEXT NOT NULL,        -- Rich text for RAG
    highlights TEXT,                  -- Key achievements/responsibilities
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,        -- Rich text for RAG
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE achievements (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,        -- Rich text for RAG
    date DATE,
    experience_id INT REFERENCES experiences(id) ON DELETE SET NULL,
    project_id INT REFERENCES projects(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Junction tables
CREATE TABLE experience_skills (
    experience_id INT NOT NULL REFERENCES experiences(id) ON DELETE CASCADE,
    skill_id INT NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    PRIMARY KEY (experience_id, skill_id)
);

CREATE TABLE project_skills (
    project_id INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    skill_id INT NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    PRIMARY KEY (project_id, skill_id)
);

CREATE TABLE achievement_skills (
    achievement_id INT NOT NULL REFERENCES achievements(id) ON DELETE CASCADE,
    skill_id INT NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    PRIMARY KEY (achievement_id, skill_id)
);

CREATE TABLE experience_projects (
    experience_id INT NOT NULL REFERENCES experiences(id) ON DELETE CASCADE,
    project_id INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    PRIMARY KEY (experience_id, project_id)
);

-- Indexes for common queries
CREATE INDEX idx_experiences_dates ON experiences(start_date, end_date);
CREATE INDEX idx_projects_dates ON projects(start_date, end_date);
CREATE INDEX idx_achievements_experience ON achievements(experience_id);
CREATE INDEX idx_achievements_project ON achievements(project_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS experience_projects;
DROP TABLE IF EXISTS achievement_skills;
DROP TABLE IF EXISTS project_skills;
DROP TABLE IF EXISTS experience_skills;
DROP TABLE IF EXISTS achievements;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS experiences;

-- +goose StatementEnd
