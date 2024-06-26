-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders
(
    id serial,
    created_at timestamp(0) default NULL::timestamp without time zone,
    patient_id INTEGER NOT NULL,
    service_id INTEGER[],
    is_active SMALLINT,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
