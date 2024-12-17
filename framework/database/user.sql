drop table if exists db_user.user;
create table db_user.user(
     id bigint(20)  unsigned not null auto_increment primary key comment '用户ID',
     user_name varchar(256) not null  comment '用户名',
     password varchar(100) not null comment '密码',
     id_card  varchar(18) not null comment '身份证号',
     phone  varchar(11)  not null comment '手机号',
     email  varchar(32)  null comment '邮箱',
     sex  tinyint unsigned  null comment '性别',
     create_time datetime default current_timestamp not null comment '创建时间',
     update_time datetime default  current_timestamp on update current_timestamp not null comment '更新时间',
     unique key `uk_phone`(phone),
     key `idx_create_time`(create_time),
     key `idx_update_time`(update_time)
) comment '用户信息表' charset=utf8mb4