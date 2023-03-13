DROP TABLE IF EXISTS posts, authors;

CREATE TABLE public.authors (
        id SERIAL PRIMARY KEY,
       "name" text NOT NULL
);

CREATE TABLE public.posts (
       id SERIAL PRIMARY KEY,
       author_id int4 NOT NULL,
       title text NOT NULL,
       "content" text NOT NULL,
       created_at int4 NOT NULL,
       "authorname" text NOT NULL,
       "publishedat" int4 NOT NULL,
       CONSTRAINT posts_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.authors(id)
);

INSERT INTO authors (id, name) VALUES (0, 'Михаил');
INSERT INTO posts (id, author_id, title, content, created_at, authorname, publishedat) VALUES (0, 0, 'Статья', 'Содержание статьи', 0, 'Михаил', 0);