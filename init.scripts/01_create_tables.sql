CREATE TABLE IF NOT EXISTS shortened_urls (
  id serial primary key,
  url text not null,
  visited int default 0,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);
