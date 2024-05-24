CREATE DATABASE mindsculpt;

CREATE TABLE IF NOT EXISTS image_generation (
  id UUID PRIMARY KEY,
  url TEXT NOT NULL,
  censored BOOLEAN NOT NULL,
  create_time BIGINT DEFAULT 0
);