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
    1584971203
);

INSERT INTO settings(
    user_id
) SELECT users.user_id FROM users WHERE username='Komfy';

INSERT INTO assets (
    width,
    height,
    resource_type,
    url,
    secure_url,
    created_at
) VALUES(
    1157,
    1280,
    'image',
    'http://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/dubai_rabg8t.jpg',
    'https://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/dubai_rabg8t.jpg',
    1584974743
), (
    687,
    761,
    'image',
    'http://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/palms_sdw5gf.jpg',
    'https://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/palms_sdw5gf.jpg',
    1584974743
);

-- Insert two posts
-- Select the id from the user table and give the content manually
INSERT INTO entities(
    user_id,
    text,
    created_at
) VALUES (
    1, -- User is certain to be 1 because we just initialized the users' table
    'This is the test post of Komfy account.', 
    1584974743
), (
    1,
    'This is the second test post of Komfy account.',
    1584974743
);

INSERT INTO contain (
    entity_id,
    asset_id
) VALUES (
    1, -- This correspond to the first entity of the two we just created
    1 -- This correspond to the first asset of the two we just created
), (
    2,
    2
);