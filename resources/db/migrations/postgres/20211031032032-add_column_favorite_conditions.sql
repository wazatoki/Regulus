
-- +migrate Up
alter table favorite_conditions add column row_order integer;

-- +migrate Down
alter table favorite_conditions drop column row_order row_order;