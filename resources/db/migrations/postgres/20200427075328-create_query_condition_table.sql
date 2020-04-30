
-- +migrate Up
create table query_condition (
	id text PRIMARY KEY,
	del boolean DEFAULT false,
	created_at timestamp,
	cre_staff_id text,
	updated_at timestamp,
	update_staff_id text,
	pattern_name text NOT NULL,
	category_name text NOT NULL,
	is_disclose boolean NOT NULL,
	owner_id text NOT NULL
);
-- +migrate Down
drop table query_condition;
