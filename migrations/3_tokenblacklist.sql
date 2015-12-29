-- +migrate Up
CREATE  TABLE `gochat`.`token` (
  `id` INT NOT NULL AUTO_INCREMENT ,
  `token` TEXT CHARACTER SET 'utf8' NOT NULL ,
  PRIMARY KEY (`id`) );


-- +migrate Down
DROP TABLE `gochat`.`token`;