CREATE TABLE people (
  id UUID PRIMARY KEY,
  nick VARCHAR(32) NOT NULL,
  name VARCHAR(100) NOT NULL,
  birth_date DATE NOT NULL,
  stack VARCHAR(32)[],
  CONSTRAINT unique_nick UNIQUE (nick)
)
