-- Users table 
-- PK: id
-- FK: X
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    fullname VARCHAR(100) NOT NULL,
    bio TEXT,
    avatar_url TEXT default 'default_komfy_profile_url' NOT NULL,
    created_at INT NOT NULL,
    checked BOOLEAN NOT NULL DEFAULT FALSE
);

-- Custom enum for content_type of the posts
CREATE TYPE CONTENT_TYPE AS ENUM (
    'text',
    'audio',
    'video',
    'image'
);

-- Custom enum to identify post from comments
CREATE TYPE ENTITY_TYPE AS ENUM (
    'post',
    'comment'
);

-- Entities Table (Posts and Comments)
-- PK: id, user_id
-- FK: user_id 1---* users.id
--     entities.answer_of 1---1 entities.id
CREATE TABLE IF NOT EXISTS entities (
    id SERIAL NOT NULL UNIQUE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    type ENTITY_TYPE NOT NULL DEFAULT 'post',
    likes INT DEFAULT 0,
    created_at INT NOT NULL,
    edited_at INT DEFAULT NULL,

    content_type CONTENT_TYPE NOT NULL DEFAULT 'text',
    description TEXT NOT NULL,
    source TEXT DEFAULT NULL,
    NSFW BOOLEAN NOT NULL DEFAULT FALSE,
    -- answer_of is not a primary key because it needs to be null
    answer_of INT DEFAULT NULL REFERENCES entities(id) ON DELETE CASCADE,

    PRIMARY KEY(id, user_id)
);

-- Likes Table, allow us to know if somebody likes
-- someone post/comment, or if we liked a given post/comment
-- PK: id, user_id, entity_id
-- FK: user_id 1---1 users.id
--     entity_id 1---1 entities.id
CREATE TABLE IF NOT EXISTS likes (
    id SERIAL NOT NULL,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    entity_id INT NOT NULL REFERENCES entities(id) ON DELETE CASCADE,

    PRIMARY KEY(id, user_id, entity_id)
);

-- Settings of every user account, only one by account
-- PK: id, user_id
-- FK: user_id 1---1 users.id
CREATE TABLE IF NOT EXISTS settings (
    id SERIAL NOT NULL,
    user_id INT NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    show_likes BOOLEAN NOT NULL DEFAULT TRUE,
    show_nsfw BOOLEAN NOT NULL DEFAULT FALSE,
    nsfw_page BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id, user_id)
);
