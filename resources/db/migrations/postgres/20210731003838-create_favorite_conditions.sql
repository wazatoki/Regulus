
-- +migrate Up
create table favorite_conditions (
      query_conditions_id text not null REFERENCES query_conditions (id)
    , staffs_id text not null REFERENCES staffs (id)
    , primary key(query_conditions_id, staffs_id)
)
;

-- +migrate Down
drop table favorite_conditions;