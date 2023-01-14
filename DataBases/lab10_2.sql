create database lab_10;
use lab_10;

create table people (
    id bigint identity(1, 1) primary key,
    age int,
);
insert into people (age) values (28);


-- Dirty read
update people set age = 200 where id = 1;
begin transaction;
    update people set age = 100 where id = 1;
    waitfor delay '00:00:10';
rollback;

-- Fantom read
delete from people where id != 1;
insert into people (age) values (18);

-- Repeatable read
update people set age = 100 where id = 1;
update people set age = 200 where id = 1;
