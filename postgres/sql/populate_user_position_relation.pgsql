\c authdb

INSERT INTO user_position(fg_user_id, position_id) VALUES 
(
  (SELECT id FROM fg_user WHERE username='carosa'),
  (SELECT id FROM position WHERE name='Webutvikler')
),
(
  (SELECT id FROM fg_user WHERE username='carosa'),
  (SELECT id FROM position WHERE name='Koordineringssjef')
),
(
  (SELECT id FROM fg_user WHERE username='pernilak'),
  (SELECT id FROM position WHERE name='Webutvikler')
);
