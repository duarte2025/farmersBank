CREATE TABLE IF NOT EXISTS accounts (
  id uuid not null primary key,
  name TEXT not null,
  balance integer not null,
  created_at TIMESTAMP default now()
);