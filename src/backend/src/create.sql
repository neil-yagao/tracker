-- --------------------------------------------------
--  Table Structure for `models.UserInfo`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `user_info` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `username` varchar(128) NOT NULL DEFAULT '' ,
    `useridentity` varchar(128) NOT NULL DEFAULT '' 
) ENGINE=InnoDB;

-- --------------------------------------------------
--  Table Structure for `models.Movement`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `movement` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `target_muscle` varchar(32) NOT NULL DEFAULT '' ,
    `secondary_muscle` varchar(32) NOT NULL DEFAULT '' ,
    `name` varchar(128) NOT NULL DEFAULT '' ,
    `description` longtext NOT NULL,
    `dividable` tinyint NOT NULL DEFAULT 0 ,
    `add_by_id` bigint NOT NULL
) ENGINE=InnoDB;

-- --------------------------------------------------
--  Table Structure for `models.AssignedPlan`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `assigned_plan` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `assign_to_id` bigint NOT NULL,
    `assigned_plan_id` bigint NOT NULL,
    `assign_time` date NOT NULL,
    `executing_status` varchar(16)
) ENGINE=InnoDB;

-- --------------------------------------------------
--  Table Structure for `models.UserSession`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `user_session` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `assign_to_id` bigint NOT NULL,
    `expecting_date` date NOT NULL,
    `execution_date` date,
    `status` varchar(16) NOT NULL DEFAULT '' 
) ENGINE=InnoDB;

-- --------------------------------------------------
--  Table Structure for `models.SessionWorkout`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `session_workout` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `belong_session_id` bigint NOT NULL,
    `mapped_movement_id` bigint NOT NULL
) ENGINE=InnoDB;

-- --------------------------------------------------
--  Table Structure for `models.Exercise`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `exercise` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `belong_workout_id` bigint NOT NULL,
    `weight` numeric(8, 2) NOT NULL DEFAULT 0 ,
    `reps` tinyint NOT NULL DEFAULT 0 ,
    `sequence` tinyint NOT NULL DEFAULT 0 
) ENGINE=InnoDB;

-- --------------------------------------------------
--  Table Structure for `models.SessionMovement`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `session_movement` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `movement_id` bigint NOT NULL,
    `sets` numeric(8, 2) NOT NULL DEFAULT 0 ,
    `reps` numeric(8, 2) NOT NULL DEFAULT 0 ,
    `sequence` bigint NOT NULL DEFAULT 0 ,
    `need_warmup` tinyint NOT NULL DEFAULT 0 ,
    `session_id` bigint NOT NULL
) ENGINE=InnoDB;

-- --------------------------------------------------
--  Table Structure for `models.Session`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `session` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(128) NOT NULL DEFAULT '' ,
    `target_muscle` varchar(128) NOT NULL DEFAULT '' ,
    `repeat` bigint NOT NULL DEFAULT 0 ,
    `weekly` varchar(32) NOT NULL DEFAULT '' ,
    `plan_id` bigint NOT NULL
) ENGINE=InnoDB;

-- --------------------------------------------------
--  Table Structure for `models.Plan`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `plan` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(128) NOT NULL DEFAULT '' ,
    `create_by_id` bigint NOT NULL
) ENGINE=InnoDB;
