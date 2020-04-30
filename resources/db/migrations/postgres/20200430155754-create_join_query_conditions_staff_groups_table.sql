
-- +migrate Up
create table join_query_conditions_staff_groups (
    query_conditions_id text not null REFERENCES query_conditions (id)
    , staff_groups_id text not null REFERENCES staff_groups (id)
    , primary key(query_conditions_id, staff_groups_id)
)
;

-- +migrate Down
drop table join_query_conditions_staff_groups;
