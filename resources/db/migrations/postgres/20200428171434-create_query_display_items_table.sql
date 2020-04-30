
-- +migrate Up
create table query_display_items (
	id text PRIMARY KEY,
	del boolean DEFAULT false,
	created_at timestamp,
	cre_staff_id text,
	updated_at timestamp,
	update_staff_id text,
	query_condition_id text,
	display_field_id text NOT NULL,
	row_order integer NOT NULL
);
-- +migrate Down
drop table query_display_items;
