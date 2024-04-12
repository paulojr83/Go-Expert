create table orders(
   id          varchar(255) not null primary key,
   price       float        not null,
   tax         float        not null,
   final_price float        not null
);
