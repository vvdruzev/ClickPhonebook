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

insert into Contacts (`firstname`, `lastname`) values
    ('ivan','ivanov'),
    ('petr','petrov');




DROP TABLE IF EXISTS Phonenumber;
create table Phonenumber
(
  contact_id  int         not null,
  phonenumber varchar(10) null,
  constraint Phonenumber_Contacts_id_fk
  foreign key (contact_id) references Contacts (id)
    on update cascade
    on delete cascade
);


insert into Phonenumber (contact_id, Phonenumber) VALUES
('1', '952000001'),
('1', '952000002'),
('2', '952000004');