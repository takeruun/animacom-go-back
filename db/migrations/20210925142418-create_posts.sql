
-- +migrate Up
CREATE TABLE IF NOT EXISTS posts (
  id int(15) AUTO_INCREMENT NOT NULL,
  user_id bigint(20),
  category_id bigint(20),
  title VARCHAR(255),
  sub_title VARCHAR(255),
  body text,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  KEY `index_posts_on_user_id` (`user_id`),
  KEY `index_posts_on_category_id` (`category_id`),
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS posts;
