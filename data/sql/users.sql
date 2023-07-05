CREATE TABLE `users` (
    `id` BIGINT unsigned AUTO_INCREMENT NOT NULL,
    `name` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `auth_tokens` (
    `id` BIGINT unsigned AUTO_INCREMENT NOT NULL,
    `user_id` BIGINT unsigned NOT NULL,
    `token` varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`)
        ON DELETE CASCADE
);

-- Indexes
CREATE UNIQUE INDEX auth_tokens_user_id_unique_index
ON `auth_tokens` (`user_id`);

CREATE INDEX auth_tokens_token_index
ON `auth_tokens` (`token`);

INSERT INTO `users` (`name`, `password`) VALUES ("Test", "123123");