
-- +migrate Up
create table maker (
	id text PRIMARY KEY,
	del boolean DEFAULT false,
	created_at timestamp,
	cre_staff_id text,
	updated_at timestamp,
	update_staff_id text,
	name text NOT NULL
);

-- +migrate Down
drop table maker;