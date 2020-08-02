
-- +migrate Up
create table query_display_items (
	id text primary key
	, del boolean not null default false
	, created_at timestamp
	, cre_staff_id text
	, updated_at timestamp
	, update_staff_id text
	, query_conditions_id text not null REFERENCES query_conditions (id)
	, display_field_id text not null
	, row_order integer not null
)
;

-- +migrate Down
drop table query_display_items;
