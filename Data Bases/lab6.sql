use master;
go
create database peoples
go
use peoples
go
create table people (
	id int identity(1,1) not null primary key,
	name varchar(255) not null,
	sex varchar(1),
	age int,
	height int,
	constraint peoplefilter check (sex = 'male' and age > 16)
);

insert into people (name, age, sex) values ('nikita',20,'male'), ('sergey',19, 'male'), ('artyom',31,'male');
select * from people;

go

create table phone(
	phoneid int identity(1,1) not null primary key,
	rowguid uniqueidentifier rowguidcol not null
	constraint con_phone default (newid()),
	region varchar(255),
	phonenumber varchar(255) not null ,
	peopleid int foreign key references people(id)
	on delete cascade
	);

	insert into phone (phonenumber, region, peopleid) values ('88005553535', 'rus', 2),
	                                                         ('554124214242', 'br', 1),
	                                                         ('352342405844', 'ger', 1);
	select * from phone;


delete from people where people.id = 1
select * from phone
go

IF OBJECT_ID (N'People_FK') is not null
	begin
		alter table People
			drop constraint People_FK
	end
go

create table houses(
	propery_number int primary key,
	type varchar(30),
	floors int,
	);

create sequence houses_seq
	start with 0
	increment by 1
	minvalue 0;
insert into houses
	values(next value for houses_seq, 'flat', 1),
	      (next value for houses_seq, 'house', 2),
		  (next value for houses_seq, 'house', 3);
	select * from houses;

