-- +goose Up
create table orders(
  order_uid text primary key,
  track_number text not null,
  entry text not null,
  locale text not null,
  internal_signature text not null,
  customer_id text not null,
  delivery_service text not null,
  shardkey text not null,
  sm_id bigint not null,
  date_created timestamp with time zone not null,
  oof_shard text not null
);

create table deliveries (
    order_uid text primary key references orders(order_uid) on delete cascade,
    name text not null,
    phone text not null,
    zip text not null,
    city text not null,
    address text not null,
    region text not null,
    email text not null
);

create table payments(
  order_uid text primary key references orders(order_uid) on delete cascade,
  transaction text unique,
  request_id text not null,
  currency text not null,
  provider text not null,
  amount bigint not null,
  payment_dt bigint not null,
  bank text not null,
  delivery_cost bigint not null,
  goods_total bigint not null,
  custom_fee bigint not null
  );

create table items(
  order_uid text primary key references orders(order_uid) on delete cascade,
  chrt_id bigint not null,
  track_number text not null,
  price bigint not null,
  rid text not null,
  name text not null,
  sale bigint not null,
  size text not null,
  total_price bigint not null,
  nm_id bigint not null,
  brand text not null,
  status bigint not null
);

-- +goose Down

drop table payments,items,deliveries,orders;
