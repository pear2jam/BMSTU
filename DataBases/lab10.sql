use lab_10;

-- Dirty read
set transaction isolation level read uncommitted;
-- solve
-- set transaction isolation level read committed;

-- Fantom read
set transaction isolation level repeatable read;
-- solve
-- set transaction isolation level serializable;
begin transaction;
    select count(*) as count from people;
    waitfor delay '00:00:10';
    select count(*) as count from people;
commit;

-- Repeatable read
set transaction isolation level read committed;
-- solve
-- set transaction isolation level repeatable read;

begin transaction;
    select * from people where id = 1;
    waitfor delay '00:00:10';
    select * from people where id = 1;
commit;
