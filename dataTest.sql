create database dataTest;
use dataTest;

create table users
(
    id       varchar(200)  not null primary key,
    password varchar(200)  not null,
    max_todo int default 5 not null
);

create table tasks
(
    id           varchar(200) not null primary key,
    content      varchar(200) not null,
    user_id      varchar(200) not null,
    created_date varchar(200) not null,
    constraint user_id
	foreign key (user_id) references users (id)
);

create index user_id_idx
    on tasks (user_id);

insert into users value ('firstUser','example',5)