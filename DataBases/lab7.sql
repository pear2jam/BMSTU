use master;
go

create table people (
	id int identity(1,1) not null primary key,
	name varchar(255) not null,
	sex varchar(1),
	age int,
	height int,
	constraint peoplefilter check (sex = 'male' and age > 16)
);

insert into people (name, age, sex) values ('Nikita', 20, 'male'), ('Sergey', 19, 'male'), ('Artyom', 31,'male');

create table phone(
	phoneid int identity(1,1) not null primary key,
	rowguid uniqueidentifier rowguidcol not null
	constraint con_phone default (newid()),
	numberviewber varchar(255) not null,
	region varchar(255),
	peopleid int foreign key references people(id)
	on delete cascade
	);

	insert into phone (numberviewber, region, peopleid) values ('88005553535', 'rus', 2),
	                                                         ('554124214242', 'br', 1),
	                                                         ('352342405844', 'ger', 1);

go

create view people_index with schemabinding as
	select id, name, age
	from dbo.people
	go

create unique clustered index i_ina on people_index(id);
go
create view numberview as
select phoneid, phone_number, country
from phone
go
create view peoplephone as
select pe.name as nameofperson,
       pe.age as ageofperson,
	   ph.phone_number

from person pe inner join phone ph
	on pe.id = ph.phoneid with check option

go
create index ind_person
on person (id, name, age) include (sex)
go
select * from person
go
