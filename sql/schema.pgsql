CREATE DATABASE fgusers;
 \c fgusers;


CREATE TABLE IF NOT EXISTS security_level (
  id serial PRIMARY KEY,
  level_name text NOT NULL,
  level integer NOT NULL,

  UNIQUE(level_name)
);

CREATE TABLE IF NOT EXISTS position (
  id serial PRIMARY KEY,
  name text NOT NULL,
);

CREATE TABLE IF NOT EXISTS user (
  id serial PRIMARY KEY,
  username text NOT NULL,
  full_name text NOT NULL,
  password_hash text NOT NULL,
  password_salt text NOT NULL,
  security_level_id integer NOT NULL REFERENCES security_level(id), 
  
  is_disabled boolean NOT NULL,
);

CREATE TABLE IF NOT EXISTS user_position (
  user_id integer NOT NULL REFERENCES user(id),
  position_id integer NOT NULL REFERENCES position(id),
  PRIMARY KEY (user_id, position_id)
);
