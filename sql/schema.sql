CREATE SCHEMA `powerlift` ;


/*user name powerlift password p@ssw0rd*/

CREATE TABLE `powerlift`.`workout` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL DEFAULT '',
  `target` VARCHAR(128) NULL DEFAULT '',
  `description` VARCHAR(512) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC));
ALTER TABLE `powerlift`.`workout` 
ADD COLUMN `perform_date` VARCHAR(32) NOT NULL DEFAULT '2016-11-11'


CREATE TABLE `powerlift`.`movement` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL DEFAULT '',
  `target_muscle` VARCHAR(128) NOT NULL DEFAULT '',
  `secondary_muscle` VARCHAR(256) NULL DEFAULT '',
  `description` VARCHAR(1024) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC));

CREATE TABLE `powerlift`.`working_set` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `workout` INT NOT NULL,
  `movement` INT NOT NULL,
  `target_weight` DOUBLE NOT NULL DEFAULT 0.0,
  `target_number` INT NOT NULL DEFAULT 10,
  `acheive_number` INT NOT NULL,
  `sequence` SMALLINT(4) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC));

