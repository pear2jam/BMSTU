use master
create database lab9;
use lab9;

--drop table users;

create table users
(
    Id         bigint primary key,
    name varchar(255) not null,
    nickname   varchar(128) not null,
    age        int,
    deleted    bit          not null default 0
)

--drop table balance;
create table balance
(
    Id bigint primary key identity (1, 1),
    user_id bigint not null foreign key references users(Id) on update cascade on delete cascade,
    num int,
)

create trigger trigger_delete
    on users
    instead of delete
    AS
begin
    update users set deleted = 1 where Id in (select Id from deleted)
end

create trigger trigger_insert
    on users
    instead of insert
    as
begin
    if exists(select * from inserted where age <= 0)
        begin
            throw 50001, 'Age cannot be less or equal to 0', 1;
            return;
        end

    insert into users(Id, name, nickname, age, deleted)
    select Id, name, nickname, age, deleted
    from inserted;
    print 'Data Inserted!';
end


create trigger trigger_update
    on users
    after update
    as
begin
    if update(Id)
        begin
            throw 50001, 'Id changing not allowed!', 1;
            return;
        end
    if exists(select * from inserted where age <= 0)
        begin
            throw 50001, 'Age cannot be less or equal to 0', 1;
            return
        end
end

create trigger trigger_after_update
    on balance
    after update
    as
begin
    if update(user_id)
        begin
            throw 50001, 'id changing not allowed', 1;
            return;
        end
end

insert into users (Id, name, nickname, age) values (10, 'Sergey', 'pearjam', 30);
insert into balance (user_id, num) values(10, 10000);

delete from users where nickname = 'pearjam';
-- Here s Error
update users set age = 0 where nickname = 'pearjam';
update balance set user_id = 3 where user_id = 4;


--drop view view_users;
create view view_users as
    select u.Id, u.name, u.nickname, u.age, w.num
    from users u
    inner join balance w on u.Id = w.user_id
    where deleted = 0;

create trigger trigger_delete_view
    on view_users
    instead of delete
    AS
begin
    delete from users where Id in (select Id from deleted)
end

create trigger trigger_instead_view
    on view_users
    instead of insert
    as
begin
    insert into users(Id, name, nickname, age)
    select Id, name, nickname, age
    from inserted;

    insert into balance(user_id, num)
    select Id, num
    from inserted;
end

create trigger trigger_view_after_update
    on view_users
    instead of update
    as
begin
    if update(user_id)
        begin
            throw 50001, 'id changing not allowed', 1;
            return;
        end
    update users
    set name = i.name, nickname = i.nickname, age = i.age
    from users u
    inner join inserted i on u.Id = i.Id;
    update balance
    set num = i.num
    from balance w
    inner join inserted i on w.user_id = i.Id;
end

update users set deleted = 0
             where nickname = 'pearjam';

--drop trigger trigger_view_after_update

select * from view_users;
update view_users set id = id+5
select * from view_users

delete from view_users where nickname = 'pearjam';

insert into view_users (Id, name, nickname, age, num) values (15, 'Alex', 'alex66', 50, 800);
update view_users set name = 'Alexander', age = 30 where nickname = 'alex66';
update view_users set name = 'Serg', age = 5 where nickname = 'pearjam';

