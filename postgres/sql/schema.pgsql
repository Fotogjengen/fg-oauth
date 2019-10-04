 \c authdb;

CREATE TABLE IF NOT EXISTS security_level (
  id serial PRIMARY KEY,
  level text NOT NULL,

  UNIQUE(level)
);

CREATE TABLE IF NOT EXISTS position (
  id serial PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE IF NOT EXISTS fg_user (
  id serial PRIMARY KEY,
  username text NOT NULL,
  full_name text NOT NULL,
  password_hash text NOT NULL,
  password_salt text NOT NULL,
  security_level_id integer NOT NULL REFERENCES security_level(id),  
  is_disabled boolean NOT NULL
);

CREATE TABLE IF NOT EXISTS user_position (
  fg_user_id integer NOT NULL REFERENCES fg_user(id),
  position_id integer NOT NULL REFERENCES position(id),
  PRIMARY KEY (fg_user_id, position_id)
);

