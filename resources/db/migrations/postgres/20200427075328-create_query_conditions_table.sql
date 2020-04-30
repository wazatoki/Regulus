
-- +migrate Up
create table query_conditions (
	id text primary key
	, del boolean default false
	, created_at timestamp
	, cre_staff_id text REFERENCES staffs (id)
	, updated_at timestamp
	, update_staff_id text REFERENCES staffs (id)
	, pattern_name text not null
	, category_name text not null
	, is_disclose boolean not null
	, owner_id text not null REFERENCES staffs (id)
)
;

-- +migrate Down
drop table query_condition;
