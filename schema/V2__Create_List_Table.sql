CREATE TABLE list(
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  owner VARCHAR(100),  -- Keep the owner's email
  created DATE,
  active BOOLEAN
);