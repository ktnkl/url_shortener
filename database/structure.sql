CREATE TABLE IF NOT EXISTS shortened_urls (
  id serial primary key,
  url text not null,
  visited int,
  created_at timestamp not null,
  updated_at timestamp not null
);
