CREATE TABLE `users` (
    `id` int unsigned AUTO_INCREMENT NOT NULL,
    `name` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`name`, `password`) VALUES ("Test", "123123");