-- drop table book_content;
-- drop table book_title;
-- drop table book_book;
-- drop table book_type;

create table book_type(
    id int  PRIMARY kEY NOT NULL,
    type_name VARCHAR(32) NOT NULL,
    sort int  not null
);

comment on column book_type.type_name is '书籍类型';

create table book_book(
	id bigint PRIMARY key not null,
	book_name varchar(64) not null,
	type_id int not null REFERENCES book_type(id)
);

comment on column book_book.book_name is '书籍名称';
comment on column book_book.type_id is '书籍类型id';

create table book_title(
	id bigint primary key not null,
	title_name varchar(64) not null,
	sort int not null,
	book_id bigint references book_book(id),
	create
);

comment on column book_title.title_name is '书籍章节名称';
comment on column book_title.book_id is '书籍id';

create table book_content(
	id BIGINT primary key not null,
	content text not null,
	title_id bigint references book_title(id)
);

comment on column book_content.content is '书籍每章内容';
comment on column book_content.title_id is '书籍章节id';

