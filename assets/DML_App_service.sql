select * from transaction order by date_time_transaction desc;

select * from accounts;

insert into accounts(id, name, email, phone_number, password, balance)
values("USR-00001", "ilham", "ilham@mail.com", "085804", "qwerty", 0);

insert into accounts(id, name, email, phone_number, password, balance)
values("USR-00002", "mawar", "mawar@mail.com", "085640", "qwerty", 0);

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