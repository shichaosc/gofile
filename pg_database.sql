-- drop table book_reptile_source;
-- drop table book_content;
-- drop table book_title;
-- drop table book_book;
-- drop table book_type;
create database book DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

create table user_auther(
    id int primary key not null,
    username varchar(64) not null,
		realname varchar(32),
    introduction  text,
		avatar varchar(256),
		gender int
);

comment on table user_auther is '作者表';
comment on column user_auther.username is '作者用户名';
comment on column user_auther.realname is '真实姓名';
comment on column user_auther.introduction is '简介';
comment on column user_auther.avatar is '作者头像';
comment on column user_auther.gender is '性别';


create table book_type(
    id int  PRIMARY kEY NOT NULL,
    type_name VARCHAR(32) NOT NULL,
    sort int  not null,
		created_on TIMESTAMP,
		updated_on TIMESTAMP
);

comment on column book_type.type_name is '书籍类型';

create table book_book(
	id bigint PRIMARY key not null,
	book_name varchar(64) not null,
	book_introduction text not null,
	cover_img varchar(256) not null,
	type_id int not null REFERENCES book_type(id),
	auther_id int not null references user_auther(id),
	created_on TIMESTAMP,
	updated_on TIMESTAMP
);

comment on column book_book.book_name is '书籍名称';
comment on column book_book.type_id is '书籍类型id';
comment on column book_book.auther_id is '作者id';

create table book_title(
	id bigint primary key not null,
	title_name varchar(64) not null,
	sort int not null,
	book_id bigint references book_book(id),
	created_on TIMESTAMP,
	updated_on TIMESTAMP
);

comment on column book_title.title_name is '书籍章节名称';
comment on column book_title.book_id is '书籍id';

create table book_content(
	id BIGINT primary key not null,
	content text not null,
	title_id bigint references book_title(id),
	created_on TIMESTAMP,
	updated_on TIMESTAMP
);

comment on column book_content.content is '书籍每章内容';
comment on column book_content.title_id is '书籍章节id';

create table book_reptile_source(
    id bigint primary key not null,
    content_id bigint references book_content(id),
    source varchar(128),
    source_url varchar(256)
);

comment on table book_reptile_source is '书籍章节爬取来源';
comment on column book_reptile_source.source is '章节来源';
comment on column book_reptile_source.source_url is '章节来源url';



