-- +goose Up
-- +goose StatementBegin
CREATE TABLE skills (
                        id SERIAL PRIMARY KEY,
                        name TEXT NOT NULL,
                        category TEXT NOT NULL, -- "Backend", "Cloud", etc.
                        proficiency INT DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE skills;
-- +goose StatementEnd
