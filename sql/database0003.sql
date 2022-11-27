alter table we_talk.user
    add password int not null;

alter table we_talk.user
    modify password varchar(32) not null;
