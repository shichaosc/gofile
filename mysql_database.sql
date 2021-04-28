-- drop table book_reptile_source;
-- drop table book_content;
-- drop table book_title;
-- drop table book_book;
-- drop table book_type;
create database book DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

create table user_auther(
    id int(11) primary key AUTO_INCREMENT not null comment 'id',
    username varchar(64) not null comment '作者用户名',
	realname varchar(32) comment '真实姓名',
    introduction  text comment '简介',
	avatar varchar(256) comment '头像',
	gender int(11) comment '性别',
	on_delete int(11) comment '是否删除',
	created_on TIMESTAMP default CURRENT_TIMESTAMP,
	updated_on TIMESTAMP default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) comment '作者表';

create table book_type(
    id int(11)  PRIMARY kEY AUTO_INCREMENT NOT NULL,
    type_name VARCHAR(32) NOT NULL comment '书籍类型',
    sort int(11)  not null,
	created_on TIMESTAMP default CURRENT_TIMESTAMP,
	updated_on TIMESTAMP default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)comment '书籍类型表';


create table book_book(
	id bigint PRIMARY key AUTO_INCREMENT not null,
	book_name varchar(64) not null comment '书籍名称',
	book_introduction text not null comment '书籍简介',
	cover_img varchar(256) not null,
	type_id int(11),
	auther_id int(11) not null,
	on_delete int(11) default 0 comment '是否删除',
	created_on TIMESTAMP default CURRENT_TIMESTAMP,
	updated_on TIMESTAMP default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	constraint book_book_book_type_type_id foreign key (type_id) references book_type(id) on delete set null,
	constraint book_book_user_auther_auther_id foreign key (auther_id) references user_auther(id)
);

create table book_title(
	id bigint primary key AUTO_INCREMENT not null,
	title_name varchar(64) not null comment '书籍章节名称',
	sort int not null comment '章节排序',
	book_id bigint not null,
	created_on TIMESTAMP default CURRENT_TIMESTAMP,
	updated_on TIMESTAMP default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	constraint book_book_title_book_book_id foreign key (book_id) references book_book(id)
);

create table book_content(
	id BIGINT primary key AUTO_INCREMENT not null,
	content text not null comment '书籍每章内容',
	title_id bigint not null,
	created_on TIMESTAMP default CURRENT_TIMESTAMP,
	updated_on TIMESTAMP default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)comment '书籍内容';


create table book_reptile_source(
    id bigint primary key AUTO_INCREMENT not null,
    content_id bigint,
    source varchar(128) comment '章节来源',
    source_url varchar(256) comment '章节来源url',
    created_on TIMESTAMP default CURRENT_TIMESTAMP,
	updated_on TIMESTAMP default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)comment '书籍章节爬取来源';



