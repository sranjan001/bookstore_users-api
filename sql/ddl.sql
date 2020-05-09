CREATE TABLE users (
    id BIGINT(20)  not null auto_increment,
    first_name VARCHAR(45),
    last_name VARCHAR(45),
    email VARCHAR(45) not null,
    date_created DATETIME,
    primary key (id),
    unique index `email_UNIQUE` (`email` ASC)
);

alter table users add status VARCHAR(45) not null;
alter table users add password VARCHAR(32) not null;