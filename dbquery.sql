CREATE TABLE users (
                       ID int PRIMARY KEY AUTO_INCREMENT NOT NULL,
                       Name varchar(255),
                       Email varchar(255) UNIQUE,
                       Password varchar(255),
                       Address varchar(255),
                       Created_At datetime,
                       Updated_At datetime,
                       Deleted_At datetime
);

CREATE TABLE admins (
                        ID int PRIMARY KEY AUTO_INCREMENT NOT NULL,
                        Username varchar(255),
                        Password varchar(255),
                        Created_At datetime,
                        Updated_At datetime,
                        Deleted_At datetime
);