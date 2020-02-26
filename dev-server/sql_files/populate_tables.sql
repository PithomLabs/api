-- Insert a test user
-- avatar_url isn't specified because there is a default value
INSERT INTO users(
    username, 
    password, 
    email, 
    fullname,
    created_at
) VALUES(
    'Komfy',
    'komfy_rocks',
    'komfy@test.test',
    'Komfy Social',
    1000
);

-- Insert two posts
-- Select the id from the user table and give the content manually
INSERT INTO entities(
    user_id,
    description,
    created_at
) SELECT id, 'This is the test post of Komfy account.', 1000 FROM users WHERE username='Komfy';

INSERT INTO entities(
    user_id,
    description,
    created_at
) SELECT id, 'This is the second test post of Komfy account.', 1000 FROM users WHERE username='Komfy';

INSERT INTO settings(
    user_id
) SELECT id FROM users WHERE username='Komfy';
