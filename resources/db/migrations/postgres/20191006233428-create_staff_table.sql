
-- +migrate Up
create table staff (
	id text PRIMARY KEY,
	del boolean DEFAULT false,
	created_at timestamp,
	cre_staff_id text,
	updated_at timestamp,
	update_staff_id text,
	staff_id text NOT NULL,
	password text NOT NULL,
	name text NOT NULL
);

-- +migrate Down
drop table staff;
