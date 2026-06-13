drop database if exists aulago;
create database aulago;
use aulago;

create table users (
	id int not null auto_increment primary key,
    name_user varchar (50) not null,
    cpf varchar(14) not null unique,
    email varchar (100) not null unique,
    password_user varchar(255) not null
);

select * from users ;