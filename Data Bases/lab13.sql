--drop database first;
create database first;
--drop database second;
create database second;

--split id into (0, 100), (101, 200)

create table first.dbo.users (
    id bigint primary key check (id between 1 and 100),
    login varchar(127) not null,
    name varchar(127) not null,
)

create table second.dbo.users (
    id bigint primary key check (id between 101 and 200),
    login varchar(127) not null,
    name varchar(127) not null,
)

use first;
create view users as
    select * from second.dbo.users
    union all
    select * from first.dbo.users;

use second;
create view users as
    select * from second.dbo.users
    union all
    select * from first.dbo.users;


select * from users;
insert into users (id, name, login) values(11, 'Alex', 'alex66');
update users set name = 'Sergey' where id = 11;
