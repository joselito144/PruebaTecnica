-- DATABASE

CREATE DATABASE Test;


-- TABLES

	CREATE TABLE Persons (
	  Id int NOT NULL PRIMARY KEY,
	  Name varchar(150) NOT NULL,
	  LastName varchar(150) NOT NULL,
    CreationDate timestamp NOT NULL,
	  Active tinyint(1) NOT NULL
	);