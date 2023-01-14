--drop database first14;
create database first14;
--drop database second14;
create database second14;

--split columns into 2 tables

create table first14.dbo.users (
    id bigint primary key,
    name varchar(127) not null,
);
create table second14.dbo.users (
    id bigint primary key,
    login varchar(127) not null,
);

use first14;

create view users_view as
    select u1.id, u1.name, u2.login from first14.dbo.users u1
    inner join second14.dbo.users as u2 on u1.id = u2.id;

create trigger trigger_view_upd
    on users_view
    instead of update
as
    if update(id)
        begin
            throw 50001, 'Id changing not allowed!', 1;
            return;
        end
    update first14.dbo.users set name = i.name
    from inserted i
    where first14.dbo.users.id = i.id;

    update second14.dbo.users set id = i.id, login = i.login
    from inserted i
    where second14.dbo.users.id = i.id;
    

create trigger trigger_view_insert
    on users_view
    instead of insert
as
    insert into first14.dbo.users (id, name) select id, name from inserted;
    insert into second14.dbo.users (id, login) select id, login from inserted;

--drop trigger trigger_view_upd;

create trigger trigger_view_delete
    on users_view
    instead of delete
as
    delete from first14.dbo.users where id in (select id from deleted)
    delete from second14.dbo.users where id in (select id from deleted)

use second14;
create view users_view as
    select u1.id, u1.name, u2.login from first14.dbo.users u1
    inner join second14.dbo.users as u2 on u1.id = u2.id;

create trigger trigger_view_insert
    on users_view
    instead of insert
as
    insert into first14.dbo.users (id, name) select id, name from inserted;
    insert into second14.dbo.users (id, login) select id, login from inserted;

drop trigger trigger_view_upd;
create trigger trigger_view_upd
    on users_view
    instead of update
as
    if update(id)
        begin
            throw 50001, 'Id changing not allowed', 1;
            return;
        end

    update first14.dbo.users set id = i.id, name = i.name
    from inserted i
    where first14.dbo.users.id = i.id;

    update second14.dbo.users set id = i.id, login = i.login
    from inserted i
    where second14.dbo.users.id = i.id;

create trigger trigger_view_delete
    on users_view
    instead of delete
as
    delete from first14.dbo.users where id in (select id from deleted)
    delete from second14.dbo.users where id in (select id from deleted)

select * from users_view;

insert into users_view (id, name, login) values(11, 'Alex', 'alex66');
update users_view set id = 50, name = 'Sergey' where id = 11;

select * from users_view;
