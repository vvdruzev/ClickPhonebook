	DROP TABLE IF EXISTS PhoneBook;
	CREATE TABLE PhoneBook (
		CountryCode varchar(10) NOT NULL,
		CountryName varchar(45) DEFAULT NULL,
		PhoneCode varchar(45) DEFAULT NULL,
		PRIMARY KEY (CountryCode));
