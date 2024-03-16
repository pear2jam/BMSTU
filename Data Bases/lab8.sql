use master;
go

create table developers(
	id int identity(1,1),
	name nvarchar(50),
	level varchar(10));
go
insert into developers(name, level) values ('sergey', 'junior'), ('anna', 'middle');
go

create procedure dbo.first_cursor
@first_cursor cursor varying output
as
	set @first_cursor = cursor
	forward_only static for select id, name, level from developers;
	open @first_cursor;
go

declare @newcursor cursor;
exec dbo.first_cursor @first_cursor = @newcursor
output;

fetch next from @newcursor;

while (@@fetch_status = 0)
begin;
fetch next from @newcursor;
end;

close @newcursor;
deallocate @newcursor;
go

create view dbo.vrand(value) as select rand();
go

create function 
    dbo.getrandom(@min int, @max int) returns int
as
begin
	return (select floor(@min + @max + value) from dbo.vrand)
end
go

create function newfunc()
returns table
as return
select id, name, level, dbo.getrandom(0,10) as random
from developers

go
create function newfunc2() returns varchar(50)
as
begin
return(select name from developers)
end
go

alter procedure dbo.first_cursor
@first_cursor cursor varying output
as
set @first_cursor = cursor

forward_only static for
select * from dbo.newfunc();
open @first_cursor;
go
declare @newcursor cursor;
exec dbo.first_cursor @first_cursor = @newcursor
output;

fetch next from @newcursor;
while (@@fetch_status = 0)
begin;
fetch next from @newcursor;
end;

close @newcursor;
deallocate @newcursor;
go

create function is_n()
returns nvarchar(1)
as
begin
	return(select level
	from developers)
	end
go
create procedure dbo.scrolling as
	declare @id int
	declare @name nvarchar(50)
	declare @level nvarchar(10)
	declare @scrolling_cursor cursor
	exec dbo.first_cursor @first_cursor = @scrolling_cursor output

	fetch next from @scrolling_cursor
	into @id, @name, @level

	while(@@fetch_status = 0)
	begin
		if(dbo.is_n() = 'n')
			print @level
		fetch next from @scrolling_cursor into @id, @name, @level
	end

	close @scrolling_cursor
	deallocate @scrolling_cursor
go

create procedure dbo.first_cursor_2
@first_cursor_2 cursor varying output
as
set @first_cursor_2 = cursor
forward_only static for
select dbo.newfunc2()
from developers;
open @first_cursor_2;
go

declare @newcursor cursor;
exec dbo.first_cursor @first_cursor = @newcursor output;
fetch next from @newcursor;

while (@@fetch_status = 0)
begin;
fetch next from @newcursor;
end;

close @newcursor;
deallocate @newcursor;
