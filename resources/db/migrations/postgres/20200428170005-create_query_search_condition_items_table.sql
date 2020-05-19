
-- +migrate Up
create table query_search_condition_items (
	id text primary key
	, del boolean not null default false
	, created_at timestamp
	, cre_staff_id text
	, updated_at timestamp
	, update_staff_id text
    , query_conditions_id text not null REFERENCES query_conditions (id)
	, search_field_id text not null
    , condition_value text not null
    , match_type text not null
    , operator text not null
    , row_order integer not null
)
;

-- +migrate Down
drop table query_search_condition_items;
