-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE card (
    id varchar(30)
);

-- +migrate StatementEnd