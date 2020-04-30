
-- +migrate Up
create table staff_groups (
	id text primary key
	, del boolean default false
	, created_at timestamp
	, cre_staff_id text
	, updated_at timestamp
	, update_staff_id text
	, staff_id text not null
	,name text not null
)
;

-- +migrate Down
drop table staff_group;
