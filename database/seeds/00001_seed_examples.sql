-- +goose Up
INSERT INTO examples (name, lucky_number, is_premium) VALUES
('Example One', 7.0, false),
('Example Two', 42.0, true),
('Example Three', 3.14, false);

-- +goose Down
DELETE FROM examples WHERE name IN ('Example One', 'Example Two', 'Example Three');
