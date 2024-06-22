insert into articles (title, contents, username, nice, created_at) values ('Hello', 'Hello, world!', 'Alice', 2, now());
insert into articles (title, contents, username, nice) values ('Goodbye', 'Goodbye, world!', 'Bob', 3);

insert into comments (article_id, message, created_at) values (1, 'Nice to meet you!', now());
insert into comments (article_id, message) values (1, 'Nice to meet you, too!');