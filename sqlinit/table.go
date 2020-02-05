package sqlinit

const SQL_CREATETABLE = `

create table if not exists user (
	id bigint not null auto_increment comment '自增id',
	num varchar(255) not null default '' comment '用户编号',
	name varchar(30) not null default '' comment '用户名',
	pass varchar(255) not null  default '' comment '用户密码',
	phone varchar(20) not null default '' comment '用户手机号',
	email varchar(125) not null default '' comment '用户邮箱',
	ctime datetime default null comment '添加时间',
	status tinyint(4) DEFAULT 1 comment '用户状态',
	primary key (id),
	unique key user_num_index (num),
	key user_name_index (name)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4/;


create table if not exists class (
        id bigint not null auto_increment comment '自增id',
        name varchar(255) not null default '' comment '分类名',
		description varchar(255) not null default '' comment '分类描述',
        primary key (id),
        unique key class_name_index (name)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4/;


create table if not exists article (
	id bigint not null auto_increment comment '自增id',
	cid int default null comment 'class分类id',
	uid int default null comment '用户id',
	title varchar(100) default null comment '文章标题',
	origin varchar(255) default null comment '文章出处',
	author varchar(30) default null  comment '文章作者',
	content text default null comment '文章内容',
	hits int default 0 comment '阅读量',
	ctime datetime default null comment '文章发表时间',
	utime datetime default null comment '文章更新时间',
	primary key (id),
	key article_cid_index (cid),
	key article_title_index (title),
	key article_utime_index (utime)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4/;
`
