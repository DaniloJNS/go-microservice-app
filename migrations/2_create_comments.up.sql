CREATE TABLE IF NOT EXISTS comments (
  id SERIAL PRIMARY KEY,
  comment VARCHAR NOT NULL,
  comment_date DATE DEFAULT CURRENT_DATE,
  user_id BIGINT REFERENCES users(id)
)