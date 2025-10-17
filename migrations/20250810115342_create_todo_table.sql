-- +goose Up
create table deliveries (
    order_uid text primary key,
    name text not null,
    phone text not null,
    zip text not null,
    city text not null,
    address text not null,
    region text not null,
    email text not null
);
create table orders(
  order_uid text primary key,
  track_number text not null,
  entry text not null,
  locale text not null,
  internal_signature text not null,
  customer_id text not null,
  delivery_service text not null,
  shardkey text not null,
  sm_id int not null,
  date_created timestamp with time zone not null,
  oof_shard text not null
);

create table payments(
  transaction text unique primary key,
  request_id text not null,
  currency text not null,
  provider text not null,
  amount bigint not null,
  payment_dt int not null,
  bank text not null,
  delivery_cost bigint not null,
  goods_total bigint not null,
  custom_fee bigint not null,
  order_uid text not null references orders(order_uid)
);

create table items(
  chrt_id int not null,
  track_number text not null,
  price int not null,
  rid text not null,
  name text not null,
  sale int not null,
  size int not null,
  total_price int not null,
  nm_id int not null,
  brand text not null,
  status int not null,
  order_uid text not null references orders(order_uid),
  primary key (order_uid,chrt_id)
);

-- +goose Down

drop table payments,items,orders,deliveries;
