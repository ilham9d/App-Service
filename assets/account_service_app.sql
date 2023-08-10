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

select * from transaction order by date_time_transaction desc;

select * from accounts;

insert into accounts(id, name, email, phone_number, password, balance)
values("USR-00001", "ilham", "ilham@mail.com", "085804", "qwerty", 0);

insert into accounts(id, name, email, phone_number, password, balance)
values("USR-00002", "mawar", "mawar@mail.com", "085640", "qwerty", 0);

alter table accounts
modify column balance int(11) not null;

delete from accounts where id = "1234";

update accounts set delete_at = null where id = "USR-00003";

select status, t.balance, t.date_time_transaction from transaction as t
inner join accounts as a on  a.id=t.account_id_penerima
where a.id = "USR-00003";

select id, account_id_pengirim, account_id_penerima, status, balance, date_time_transaction from transaction
where status = "Top-Up" and account_id_penerima = "6vcpf1M9nEUMevkDh8PH" and account_id_pengirim = "6vcpf1M9nEUMevkDh8PH";

select pengirim.name as nama_pengirim, penerima.name as nama_penerima, pengirim.phone_number, penerima.phone_number, t.balance, t.date_time_transaction from transaction as t
inner join accounts as pengirim on t.account_id_pengirim = pengirim.id
inner join accounts as penerima on t.account_id_penerima = penerima.id
where (pengirim.id = "6vcpf1M9nEUMevkDh8PH" or penerima.id="6vcpf1M9nEUMevkDh8PH") and status = "Transfer";

insert into transaction(id, account_id_pengirim,account_id_penerima, status, balance)
values("transfer-00001", "6vcpf1M9nEUMevkDh8PH", "USR-00002", "Transfer", 10000),
("transfer-00002", "USR-00003", "6vcpf1M9nEUMevkDh8PH", "Transfer", 20000);

insert into transaction(id, account_id_pengirim,account_id_penerima, status, balance)
values("transfer-00003", "USR-00003", "USR-00002", "Transfer", 30000);