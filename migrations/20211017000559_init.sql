-- +goose Up

CREATE TABLE `course`
(
    `id`                int(11) NOT NULL AUTO_INCREMENT,
    `title`             varchar(50)  NOT NULL,
    `short_description` varchar(250) NOT NULL,
    `long_description`  text         NOT NULL,
    `price`             int(11) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

CREATE TABLE `lesson`
(
    `id`                    int(11) NOT NULL AUTO_INCREMENT,
    `id_module`             int(11) NOT NULL,
    `title`                 varchar(50)  NOT NULL,
    `short_description`     varchar(250) NOT NULL,
    `content`               text         NOT NULL,
    `is_available_in_trial` tinyint(4) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `module`
(
    `id`                     int(11) NOT NULL AUTO_INCREMENT,
    `id_course`              int(11) NOT NULL,
    `title`                  varchar(50)  NOT NULL,
    `short_description`      varchar(250) NOT NULL,
    `is_available_int_trial` tinyint(4) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `task`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `id_lesson` int(11) NOT NULL,
    `content`   text NOT NULL,
    `need_file` tinyint(4) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `user`
(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `user_completed_lesson`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `id_user`   int(11) NOT NULL,
    `id_lesson` int(11) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `user_course`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `id_user`   int(11) NOT NULL,
    `id_course` int(11) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `user_task`
(
    `id`      int(11) NOT NULL AUTO_INCREMENT,
    `id_user` int(11) NOT NULL,
    `id_task` int(11) NOT NULL,
    `answer`  text,
    `status`  tinyint(4) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `user_task_conversation`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `id_task`        int(11) NOT NULL,
    `id_user_task`   int(11) NOT NULL,
    `id_user_author` int(11) NOT NULL,
    `content`        text NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- +goose Down
DROP TABLE IF EXISTS `course`;
DROP TABLE IF EXISTS `lesson`;
DROP TABLE IF EXISTS `module`;
DROP TABLE IF EXISTS `task`;
DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `user_completed_lesson`;
DROP TABLE IF EXISTS `user_course`;
DROP TABLE IF EXISTS `user_task`;
DROP TABLE IF EXISTS `user_task_conversation`;