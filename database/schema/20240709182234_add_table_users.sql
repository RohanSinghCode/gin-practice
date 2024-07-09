-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users(
    Id UUID PRIMARY KEY,
    FirstName TEXT NOT NULL,
    LastName TEXT NOT NULL,
    Password TEXT NOT NULL,
    EmailAddresss TEXT NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users
-- +goose StatementEnd
