DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS posts CASCADE;

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) not null,
                       lastName VARCHAR(100) not null,
                       email VARCHAR(100) UNIQUE not null,
                       password VARCHAR(100) not null
);

CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                       title VARCHAR(200) not null,
                       content TEXT not null,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);

-- COMANDOS
-- sudo mysql -u admin -p curso_go < database/schema.sql