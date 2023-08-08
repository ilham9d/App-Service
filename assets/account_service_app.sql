-- DDL
-- create database
create database db_app_service;

-- use database
use db_app_service;

-- create table
create table accounts(
id varchar(20) not null primary key,
name varchar(100),
email varchar(100) not null,
phone_number varchar(20) not null,
password varchar(100) not null,
balance int,
create_at datetime default current_timestamp,
update_at datetime,
delete_at datetime
);

create table transaction(
id varchar(20) not null primary key,
status varchar(10),
account_id_pengirim varchar(20),
account_id_penerima varchar(20),
balance int,
date_time_transaction datetime,
constraint fk_account_pengirim foreign key (account_id_pengirim) references accounts(id),
constraint fk_account_penerima foreign key (account_id_penerima) references accounts(id)
);

alter table transaction
drop column status;

alter table transaction
add column status varchar(10) after account_id_penerima;

select * from transaction;

select * from accounts;

insert into accounts(id, name, email, phone_number, password, balance)
values("USR-00001", "ilham", "ilham@mail.com", "085804", "qwerty", 0);

insert into accounts(id, name, email, phone_number, password, balance)
values("USR-00002", "mawar", "mawar@mail.com", "085640", "qwerty", 0);