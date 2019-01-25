
-- name: table-profile-create
create table profiles(
    profile_id int,
    first_name varchar(20),
    last_name varchar(20),
    primary key (profile_id)
);

-- name: table-profile-drop
drop table if exists profiles;
