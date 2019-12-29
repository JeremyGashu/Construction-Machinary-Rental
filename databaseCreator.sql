create table users(
    username varchar(256) primary key,
    firstname varchar(256) not null,
    lastname varchar(256) not null,
    email varchar(256) not null,
    phone varchar(20) not null,
    address varchar(256) not null,
    postnum varchar(20) not null,
    password varchar(256) not null,
    imagepath text default 'defaultuser.jpg',
    account numeric default 200000
);

create table companies (
    id int primary key,
    name varchar(256) not null,
    email varchar(256) not null,  
    phone varchar(256) not null,
    address varchar(256) not null,
    rating int default 0;
    imagepath text default 'defaultcompany.jpg',
    password varchar(256) not null
);
--material type

create table materials(
    id int primary key,
    name varchar(256) not null,
    type int references materialtype(id)
    owner int references companies(id) on delete cascade,
    priceperday numeric default 0,
    ondiscount boolean default false,
    discount numeric default 0,
    onsale boolean default false,
    imagepath text default 'defaultmaterial.jpg'
);


