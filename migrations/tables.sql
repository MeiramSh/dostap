CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    password VARCHAR(255),
    avatar_link VARCHAR(255),
    gender BOOLEAN,
    age INTEGER,
    phone_number VARCHAR(255) UNIQUE,
    city_of_residence VARCHAR(255),
    description VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    is_email_verified BOOLEAN,
    username VARCHAR(255) UNIQUE,
    is_private BOOLEAN,
    receive_push BOOLEAN,
    role_id int,
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS blocks (
    who_blocked_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    blocked_id INTEGER REFERENCES users(id) ON DELETE CASCADE
);



CREATE TABLE IF NOT EXISTS friend_requests (
    id SERIAL PRIMARY KEY,
    requestor_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    recipient_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(255),
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS friendships (
    user_id1 INTEGER REFERENCES users(id) ON DELETE CASCADE,
    user_id2 INTEGER REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP,
    PRIMARY KEY (user_id1, user_id2)
);

CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    title VARCHAR (255),
    creator_id INTEGER REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS event_participants (
    event_id INTEGER REFERENCES events(id) ON DELETE CASCADE,
    participant_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (event_id, participant_id)
);

CREATE TABLE IF NOT EXISTS event_organisers (
    event_id INTEGER REFERENCES events(id) ON DELETE CASCADE,
    organiser_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (event_id, organiser_id)
);

CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS interests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    category INTEGER REFERENCES category(id) ON DELETE CASCADE
);