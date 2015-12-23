-- +migrate Up
CREATE  TABLE `gochat`.`user` (
  `id` INT NOT NULL AUTO_INCREMENT ,
  `username` VARCHAR(45) CHARACTER SET 'utf8' NOT NULL ,
  `password` VARCHAR(45) CHARACTER SET 'utf8' NOT NULL ,
  PRIMARY KEY (`id`) );


-- +migrate Down
DROP TABLE `gochat`.`user`;