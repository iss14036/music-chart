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

INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('All i need', 'Jacob Collier', '00:03:00', 'Djesse Vol. 3', '2020');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Isn''t she', 'Stevie Wonder', '00:03:00', 'Songs in the Key of Life', '1976');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Stay With Me', 'Miki Matsubara', '00:03:00', 'Pocket Park', '1980');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Love theory', 'Kirk Franklin', '00:03:00', 'Long, Live, Love', '2019');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Smile', 'Kirk Franklin', '00:03:00', 'Hello Fear', '2011');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Mungkin Nanti', 'Ariel', '00:03:00', 'Bintang di Surga', '2004');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('OverJoyed', 'Stevie Wonder', '00:03:00', 'In Square Circle', '1985');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('I just called to say i love you', 'Stevie Wonder', '00:03:00', 'The Woman in Red', '1984');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Januari', 'Glenn Fredly', '00:03:00', 'Selamat Pagi, Dunia!', '2002');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Akhir cerita cinta', 'Glenn Fredly', '00:03:00', 'Selamat Pagi, Dunia!', '2002');
INSERT INTO t_music (title, singer, duration, album, release_year) VALUES('Terserah', 'Glenn Fredly', '00:03:00', 'Private Collection', '2008');