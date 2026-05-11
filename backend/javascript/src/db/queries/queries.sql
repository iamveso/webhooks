/* @name RegisterUser */
INSERT INTO usernames (username) VALUES (@username) ON CONFLICT DO NOTHING RETURNING *;
