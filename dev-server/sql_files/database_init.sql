-- Users table 
-- PK: id
-- FK: X
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
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
    entity_id SERIAL PRIMARY KEY,
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
    like_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    entity_id INT NOT NULL REFERENCES entities(entity_id) ON DELETE CASCADE
);

-- Settings of every user account, only one by account
-- PK: setting_id
-- FK: user_id 1---1 users.id
CREATE TABLE IF NOT EXISTS settings (
    setting_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE REFERENCES users(user_id) ON DELETE CASCADE,
    show_likes BOOLEAN NOT NULL DEFAULT TRUE,
    show_nsfw BOOLEAN NOT NULL DEFAULT FALSE,
    nsfw_page BOOLEAN NOT NULL DEFAULT FALSE
);

-- Custom enum based on https://cloudinary.com/documentation/image_upload_api_reference#optional_parameters
CREATE TYPE RESOURCE_TYPE AS ENUM (
    'text',
    'audio',
    'video',
    'image'
);

-- Assets of all Komfy's media
-- PK: asset_id
-- FK: asset_id n---1 contain.contain_id
CREATE TABLE IF NOT EXISTS assets (
    asset_id SERIAL PRIMARY KEY,
    width INT NOT NULL,
    height INT NOT NULL,
    resource_type RESOURCE_TYPE NOT NULL DEFAULT 'image',
    url text NOT NULL,
    secure_url text NOT NULL,
    created_at INT NOT NULL
);

-- Contains is the bridgge between sources and entities
-- PK: contain_id
-- FK: contain_id 1---n assets.asset_id 
--     contain_id 1---n entities.entity_id
CREATE TABLE IF NOT EXISTS contain (
    contain_id SERIAL PRIMARY KEY,
    entity_id INT NOT NULL REFERENCES entities(entity_id) ON DELETE CASCADE,
    asset_id INT NOT NULL REFERENCES assets(asset_id) ON DELETE CASCADE
);