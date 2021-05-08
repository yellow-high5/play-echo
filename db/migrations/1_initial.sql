-- +migrate Up
create table users(
  id int primary key,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  name text,
  avatar_path text,
  tel text,
  sex int,
  email text,
  password text
);

create table hosts(
  id int primary key,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  name text,
  representative text,
  avatar_path text,
  url_path text,
  tel text,
  address text,
  intro text,
  has_parking boolean,
  has_credit_payment boolean
);

create table plans(
  id int primary key,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  host_id int,
  name text,
  description text,
  price int,
  type text,
  target_fish text,
  capasity int,
  meeting_at datetime,
  departure_at datetime,
  return_at datetime,
  duration int
);
