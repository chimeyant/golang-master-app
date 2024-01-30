CREATE TABLE app_infos (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  uuid varchar(36) NOT NULL  UNIQUE ,
  app_name varchar(255) NOT NULL,
  app_ver varchar(10) NOT NULL ,
  app_desc varchar(255),
  app_logo varchar(100),
  app_theme varchar(100),
  app_color varchar(100),
  app_bg varchar(100),
  app_bg_nav varchar(100),
  app_url varchar(100),
  app_company varchar(100),
  app_slogan varchar(100),
  app_address varchar(100),
  app_wa varchar(100),
  app_fb varchar(100),
  app_tw varchar(100),
  app_ig varchar(100),
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  deleted_at datetime(3),
  PRIMARY KEY (id),
  KEY idx_app_infos_created_at (created_at),
  KEY idx_app_infos_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;