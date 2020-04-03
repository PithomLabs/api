-- Users table 
-- PK: id
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY NOT NULL,
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
-- FK: user_id 1---* users.user_id
--     entities.answer_of 1---1 entities.entity_id
CREATE TABLE IF NOT EXISTS entities (
    entity_id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,

    type ENTITY_TYPE NOT NULL DEFAULT 'post',
    likes INT DEFAULT 0,
    created_at INT NOT NULL,
    edited_at INT DEFAULT NULL,

    content_type CONTENT_TYPE NOT NULL DEFAULT 'text',
    text TEXT NOT NULL,
    source TEXT DEFAULT NULL,
    NSFW BOOLEAN NOT NULL DEFAULT FALSE,
    -- answer_of is not a primary key because it needs to be null
    answer_of INT DEFAULT NULL REFERENCES entities(entity_id) ON DELETE CASCADE
);

-- Likes Table, allow us to know if somebody likes
-- someone post/comment, or if we liked a given post/comment
-- PK: like_id
-- FK: user_id 1---1 users.id
--     entity_id 1---1 entities.id
CREATE TABLE IF NOT EXISTS likes (
    like_id SERIAL NOT NULL,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    entity_id INT NOT NULL REFERENCES entities(entity_id) ON DELETE CASCADE,
    -- Disable like's duplication
    PRIMARY KEY(user_id, entity_id)
);

-- Settings of every user account, only one by account
-- PK: setting_id
-- FK: user_id 1---1 users.id
CREATE TABLE IF NOT EXISTS settings (
    setting_id SERIAL NOT NULL,
    user_id INT NOT NULL UNIQUE REFERENCES users(user_id) ON DELETE CASCADE,
    show_likes BOOLEAN NOT NULL DEFAULT TRUE,
    show_nsfw BOOLEAN NOT NULL DEFAULT FALSE,
    nsfw_page BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(setting_id, user_id)
);

-- Custom enum based on https://cloudinary.com/documentation/image_upload_api_reference#optional_parameters
CREATE TYPE RESOURCE_TYPE AS ENUM (
    'image',
    'raw',
    'video',
    'auto'
);

-- Assets of all Komfy's media
-- PK: asset_id
-- FK: entity_id 1---0/* entities.entity_id
CREATE TABLE IF NOT EXISTS assets (
    asset_id SERIAL PRIMARY KEY NOT NULL,
    entity_id INT REFERENCES entities(entity_id) ON DELETE CASCADE,
    width INT NOT NULL,
    height INT NOT NULL,
    resource_type RESOURCE_TYPE NOT NULL DEFAULT 'image',
    url text NOT NULL,
    secure_url text NOT NULL,
    created_at INT NOT NULL
);