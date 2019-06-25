SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS Contacts;
create table Contacts
(
  id        int auto_increment
    primary key,
  Firstname varchar(50) null,
  lastname  varchar(50) null
);

insert into Contacts (`id`,`firstname`, `lastname`) values
    (1,'ivan','ivanov'),
    (2,'petr','petrov');




DROP TABLE IF EXISTS Phonenumber;
create table Phonenumber
(
  id          int         not null,
  phonenumber varchar(10) null
);

insert into Phonenumber (contact_id, Phonenumber) VALUES
('1', '952000001'),
('1', '952000002'),
('2', '952000004');