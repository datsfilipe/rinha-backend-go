CREATE EXTENSION pg_trgm;

CREATE OR REPLACE FUNCTION array_to_string_immutable(anyarray, text)
  RETURNS text LANGUAGE SQL IMMUTABLE STRICT AS $$
    SELECT array_to_string($1, $2)
  $$;

CREATE TABLE people (
  id UUID PRIMARY KEY,
  nick VARCHAR(32) NOT NULL,
  name VARCHAR(100) NOT NULL,
  birth_date DATE NOT NULL,
  stack VARCHAR(32)[],
  search TEXT GENERATED ALWAYS AS (
    name || ' ' || nick || ' ' || COALESCE(ARRAY_TO_STRING_IMMUTABLE(stack, ' '), '')
  ) STORED,
  CONSTRAINT unique_nick UNIQUE (nick)
);

CREATE INDEX people_search_idx ON people USING GIST (search gist_trgm_ops);
