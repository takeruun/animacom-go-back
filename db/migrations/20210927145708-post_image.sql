
-- +migrate Up
CREATE TABLE IF NOT EXISTS post_images (
  id int(15) AUTO_INCREMENT NOT NULL,
  post_id bigint(20) NOT NULL,
  image VARCHAR(255),
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  KEY `index_post_images_on_post_id` (`post_id`)
);

-- +migrate Down
DROP TABLE IF EXISTS post_images;
