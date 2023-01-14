--drop database first15;
create database first15;
--drop database second15;
create database second15;

--drop table first15.dbo.people;

create table first15.dbo.people (
    id bigint primary key,
    name varchar(127) not null,
);

--drop table second15.dbo.information;
create table second15.dbo.information (
    id_user varchar(127) primary key,
    region varchar(127) not null,
);

use first15;

drop trigger dbo.trigger_update;
create trigger dbo.trigger_update
    on dbo.people
    after update
    as
    begin
        if update(id)
            begin
                throw 50001, 'Cant update Id', 1;
                return;
            end
    end

create trigger dbo.trigger_info_insert
    on dbo.information
    instead of insert
    as
    begin
        if exists(select u.id from inserted as i left join first15.dbo.people u on i.id_user = u.id where id is null)
            begin
                throw 50001, 'User does not exist!', 1;
                return;
            end

        insert into dbo.information select id_user, region from inserted;
    end

drop trigger first15.dbo.trigger_delete;
create trigger dbo.trigger_delete
    on dbo.people
    after delete
    as
    begin
        delete from second15.dbo.information where id_user in (select id from deleted);
    end

use second15;

drop trigger dbo.trigger_info_insert;

drop trigger dbo.trigger_userinfo;
create trigger dbo.trigger_userinfo
    on dbo.information
    after update
    as
    begin
        if update(id_user)
            begin
                throw 50001, 'Cant update id_user!', 1;
                return;
            end
    end

insert into first15.dbo.people (id, name) values (10, 'Nadya'), (40, 'Sergey');
insert into first15.dbo.people (id, name) values (20, 'Livia')
update first15.dbo.people set id = 300 where id = 20;
insert into second15.dbo.information (id_user, region) values (20, 'Brazil'), (40, 'Russia');
-- Error
insert into second15.dbo.information (id_user, region) values (100, 'Argentina');
update second15.dbo.information set id_user = 300 where id_user = 34;
