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
), (
    'Twitter',
    'mastodon_rocks',
    'twitter@test.test',
    'Twitter Social',
    1584971400
);

INSERT INTO settings(
    user_id
) SELECT users.user_id FROM users;

-- Insert two posts
-- Select the id from the user table and give the content manually
INSERT INTO entities(
    user_id,
    text,
    created_at,
    content_type
) VALUES (
    1, -- User is certain to be 1 because we just initialized the users' table
    'This is the test post of Komfy account.', 
    1584974743,
    'image'
), (
    1,
    'This is the second test post of Komfy account.',
    1584974743,
    'image'
), (
    1,
    'This is a post without sources.',
    1584974999,
    'text'
), (
    1,
    'This is a post with multiple sources.',
    1594975999,
    'image'
);

INSERT INTO entities (
    user_id,
    answer_of,
    text,
    created_at,
    content_type,
    type
) VALUES (
    2,
    1,
    'This is a comment.',
    1594975999,
    'text',
    'comment'
), (
    2,
    5,
    'This is a comment with an image',
    1594975999,
    'image',
    'comment'
);

INSERT INTO assets (
    entity_id,
    width,
    height,
    resource_type,
    url,
    secure_url,
    created_at
) VALUES(
    1,
    1157,
    1280,
    'image',
    'http://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/dubai_rabg8t.jpg',
    'https://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/dubai_rabg8t.jpg',
    1584974743
), (
    2,
    687,
    761,
    'image',
    'http://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/palms_sdw5gf.jpg',
    'https://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/palms_sdw5gf.jpg',
    1584974743
), (
    4,
    1157,
    1280,
    'image',
    'http://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/dubai_rabg8t.jpg',
    'https://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/dubai_rabg8t.jpg',
    1584974999
), (
    4,
    687,
    761,
    'image',
    'http://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/palms_sdw5gf.jpg',
    'https://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/palms_sdw5gf.jpg',
    1584975010
), (
    6,
    687,
    761,
    'image',
    'http://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/palms_sdw5gf.jpg',
    'https://res.cloudinary.com/dlcfinrwj/image/upload/v1584973868/palms_sdw5gf.jpg',
    1584975010
);