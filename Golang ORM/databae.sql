create table samples
(
    id   varchar(100) not null primary key,
    name varchar(100) not null
) engine = InnoDB;

select *
from samples;

create table users
(
    id         varchar(100) not null primary key,
    password   varchar(100) not null,
    name       varchar(100) not null,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp on update current_timestamp
) engine = InnoDB;

select *
from users;

alter table users rename column name to first_name;

alter table users
    add column middle_name varchar(100) after first_name;
alter table users
    add column last_name varchar(100) after middle_name;

select *
from users;
SELECT *
FROM users
WHERE first_name like '%Ahmad%'
  AND password = 'rahasia';
select *
from users
where id = '0';

create table user_logs
(
    id         integer auto_increment primary key,
    user_id    varchar(100) not null,
    action     varchar(100) not null,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp on update current_timestamp
) engine = InnoDB;

select *
from user_logs;

desc user_logs;

delete
from user_logs;

alter table user_logs
    modify created_at bigint not null;
alter table user_logs
    modify updated_at bigint not null;

select *
from users
where id = '999';

create table todos
(
    id          integer auto_increment primary key,
    user_id     varchar(100) not null,
    title       varchar(100) not null,
    description text,
    created_at  timestamp    not null default current_timestamp,
    updated_at  timestamp    not null default current_timestamp on update current_timestamp,
    deleted_at  timestamp
) engine = InnoDB;

select *
from todos;

create table wallets
(
    id         varchar(100) primary key,
    user_id    varchar(100) not null,
    balance    bigint       not null default 0,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp on update current_timestamp,
    foreign key (user_id) references users (id)
) engine InnoDB;

select *
from wallets;

select *
from wallets w
         join users u on w.user_id = u.id;

create table addresses
(
    id         bigint primary key auto_increment,
    user_id    varchar(100) not null,
    address    varchar(100) not null,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp on update current_timestamp,
    foreign key (user_id) references users (id)
) engine = InnoDB;

select *
from addresses;

create table products
(
    id         varchar(100) primary key,
    name       varchar(100) not null,
    price      bigint       not null,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp on update current_timestamp
) engine = InnoDB;

select *
from products;

create table user_like_product
(
    user_id    varchar(100) not null,
    product_id varchar(100) not null,
    primary key (user_id, product_id),
    foreign key (user_id) references users (id),
    foreign key (product_id) references products (id)
) engine = InnoDB;

desc user_like_product;

select *
from user_like_product;

select *
from users;

SELECT users.id,
       users.password,
       users.first_name,
       users.middle_name,
       users.last_name,
       users.created_at,
       users.updated_at
FROM users
         JOIN user_like_product
              ON user_like_product.user_id = users.id AND user_like_product.product_id = 'P001'
WHERE first_name like 'Solikhin%';

SELECT *
FROM `addresses`
WHERE `addresses`.`user_id` = '2';

desc guest_books;