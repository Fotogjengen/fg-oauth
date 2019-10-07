\c authdb

INSERT INTO fg_user
(
  username,
  full_name,
  password_hash,
  security_level_id,
  is_disabled
)
VALUES
(
  'carosa',
  'Caroline Sandsbråten',
  'hash',
  (SELECT id FROM security_level WHERE level='FG'),
  DEFAULT
),
(
  'sjsivert',
  'Sindre Sivertsen',
  'hash',
  (SELECT id FROM security_level WHERE level='FG'),
  DEFAULT
),
(
  'pernilak',
  'Pernille Klevstuen',
  'hash',
  (SELECT id FROM security_level WHERE level='FG'),
  DEFAULT
),
(
  'huskfolk',
  'Husfolk Folkesen',
  'hash',
  (SELECT id FROM security_level WHERE level='HUSFOLK'),
  DEFAULT
);
