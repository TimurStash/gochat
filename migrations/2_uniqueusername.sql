-- +migrate Up
ALTER TABLE `gochat`.`user`
ADD UNIQUE INDEX `username_UNIQUE` (`username` ASC) ;


-- +migrate Down
ALTER TABLE `gochat`.`user`
DROP INDEX `username_UNIQUE` ;