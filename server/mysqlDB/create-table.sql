DROP TABLE IF EXISTS user;
CREATE TABLE `user` (
`id` int AUTO_INCREMENT NOT NULL,
`firstname` varchar(255) NOT NULL,
`lastname` varchar(255),
`email` varchar(255) UNIQUE NOT NULL,
`password` varchar(255) NOT NULL,
`gender` varchar(1),
`phone` int NOT NULL,
`created_at` timestamp DEFAULT (now()),
`updated_at` timestamp DEFAULT (now()),
PRIMARY KEY (`id`)
);
CREATE INDEX `user_index_0` ON `user` (`email`);
CREATE INDEX `user_index_1` ON `user` (`phone`);

INSERT INTO user 
(firstname, lastname, email, password, gender, phone)
VALUES
("Test", "User", "test.user@gmail.com", "Text@123", 'F', 123457890);

