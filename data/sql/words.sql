START TRANSACTION;

-- Tables
CREATE TABLE `languages` (
    `id` tinyint unsigned AUTO_INCREMENT NOT NULL,
    `iso` varchar(2) NOT NULL,
    `created_at` timestamp NOT NULL default(current_timestamp()),
    PRIMARY KEY (`id`)  
);

CREATE TABLE `words` (
    `id` bigint unsigned AUTO_INCREMENT NOT NULL,
    `text` varchar(255) NOT NULL,
    `language_id` tinyint unsigned NOT NULL,
    `type` tinyint(1) unsigned NOT NULL DEFAULT(0),
    `created_at` timestamp NOT NULL default(current_timestamp()),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`language_id`) 
        REFERENCES `languages`(`id`)
        ON DELETE CASCADE
);

CREATE TABLE `translations` (
    `id` BIGINT unsigned AUTO_INCREMENT NOT NULL,
    `word_id` BIGINT unsigned NOT NULL,
    `translation_id` bigint unsigned NOT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY (`word_id`)
        REFERENCES `words`(`id`)
        ON DELETE CASCADE,
    FOREIGN KEY (`translation_id`)
        REFERENCES `words`(`id`)
        ON DELETE CASCADE
);

-- Idexes
CREATE UNIQUE INDEX words_text_language_id_unique_index
ON `words` (`text`, `language_id`);
CREATE UNIQUE INDEX translations_word_id_translation_id_unique_index
ON `translations` (`word_id`, `translation_id`);

CREATE INDEX words_language_index
ON `words`(`language_id`, `type`);
CREATE INDEX translations_word_translation_index
ON `translations`(`word_id`, `translation_id`);

-- Data
INSERT INTO `languages` (`iso`) VALUES ("en"), ("ua");

COMMIT;