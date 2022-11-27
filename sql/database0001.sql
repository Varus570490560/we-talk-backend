create table we_talk.avatar
(
    id  bigint auto_increment,
    url varchar(256) not null,
    constraint avatar_pk
        unique (id)
);

create table we_talk.user
(
    id        bigint auto_increment
        primary key,
    nick_name varchar(64) not null,
    avatar    bigint      not null,
    constraint user_avatar_id_fk
        foreign key (avatar) references we_talk.avatar (id)
);

