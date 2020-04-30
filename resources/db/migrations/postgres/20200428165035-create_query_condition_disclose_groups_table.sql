
-- +migrate Up
create table query_condition_disclose_groups (
	id text PRIMARY KEY,
	del boolean DEFAULT false,
	created_at timestamp,
	cre_staff_id text,
	updated_at timestamp,
	update_staff_id text,
	query_condition_id text,
    group_id text
);
-- +migrate Down
drop table query_condition;
