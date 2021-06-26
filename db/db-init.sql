CREATE TABLE t_favorite (
                            user_id integer NOT NULL,
                            music_id integer NOT NULL,
                            PRIMARY KEY (user_id,music_id)
);

CREATE TABLE t_music (
                         id serial NOT NULL,
                         title varchar(100) NOT NULL,
                         singer varchar(100) NOT NULL,
                         duration varchar(100) DEFAULT NULL,
                         album varchar(100) DEFAULT NULL,
                         release_year varchar(4) DEFAULT NULL,
                         PRIMARY KEY (id)
);

CREATE TABLE t_user (
                        id serial NOT NULL,
                        username varchar(64) NOT NULL,
                        password varchar(255) NOT NULL,
                        full_name varchar(255) NOT NULL,
                        hobby varchar(64) DEFAULT NULL,
                        gender varchar(64) DEFAULT NULL,
                        address varchar(255) DEFAULT NULL,
                        PRIMARY KEY (id),
                        UNIQUE(username)
);

create index users_username_IDX on t_user (username);

INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Happier', 'Olivia Rodrigo', '00:02:57', 'Sour', '2021');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('De javu', 'Olivia Rodrigo', '00:03:52', 'Sour', '2021');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Monster', 'Katy', '00:03:00', 'After Hours', '2014');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Begin', 'Maneskin', '00:03:00', 'Chosen', '2017');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('No roots', 'Alice Merton', '00:03:00', 'No roots', '2016');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Here''s Your Perfect', 'Jammi miller', '00:03:00', 'Here''s Your Perfect', '2021');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Take Me To Church E.P.', 'Hozier', '00:03:00', 'Take Me To Church E.P.', '2013');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('build a bi * ch', 'Bella Poarch', '00:03:00', 'build a bi * ch', '2021');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Talking to the moon', 'Bruno Mars', '00:03:00', 'Doo-Wops & Hooligans', '2010');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('At My Worst', 'Pink Sweat', '00:03:00', 'Pink Planet', '2021');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Save your tears', 'The Weeknd', '00:03:00', 'After Hours', '2020');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Driver License', 'Olivia Rodrigo', '00:03:00', 'Sour', '2021');