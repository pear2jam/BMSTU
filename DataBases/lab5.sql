use master;
go
create database UsersList
go
use UsersList
go
create schema OldSchema;
go

create table OldSchema.user_ (
	userid int not null,
	userlogin nvarchar(100) not null,
	name nvarchar(100) not null,
	lastname nvarchar(100) not null,
	city varchar(100) not null,
	region varchar(100) not null,
	);
alter database UsersList
	add filegroup MyFiles
go

alter database UsersList
	add file
	(name = 'Log',
	filename = '/users/naumovserge/MyFiles.ndf',
	size = 1000mb,
	maxsize = 2500,
	filegrowth = 1)
to filegroup MyFiles;

alter database UsersList
modify filegroup MyFiles default;
go
create table Chatzone(
	chat int not null,
	admin nvarchar(100) not null,
	city varchar(100) not null,
	region varchar(100) not null,
	)
on MyFiles;
go
alter database UsersList
modify filegroup [primary] default;
alter database UsersList
remove file Log;
go
create clustered index chat_admin_city
on Chatzone (chat, admin, city, region)
on [primary];
go
alter database UsersList
remove filegroup MyFiles;
go
create schema NewSchema;
go
alter schema NewSchema transfer OldSchema.user_;
go
drop table NewSchema.user_;
go
drop schema NewSchema
go


