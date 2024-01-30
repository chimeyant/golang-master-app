CREATE TABLE users (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  uuid varchar(36) NOT NULL UNIQUE,
  name varchar(255) NOT NULL,
  email varchar(100) NOT NULL,
  password varchar(255) NOT NULL,
  authent varchar(100) NOT NULL,
  avatar varchar(255),
  google_id varchar(26),
  remember_token varchar(255),
  status int NOT NULL DEFAULT 1,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  deleted_at datetime(3),
  PRIMARY KEY (id),
  KEY idx_users_created_at (created_at),
  KEY idx_users_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
