CREATE DATABASE if not exists apsmgm DEFAULT CHARACTER SET utf8mb4;

use apsmgm;

create table if not exists TBL_DICTIONARYITEM
(
    DIC_TYPE    VARCHAR(50)  not null,
    DIC_CODE    VARCHAR(50)  not null,
    DIC_NAME    VARCHAR(128) not null,
    DISP_ORDER  VARCHAR(50),
    UPDATE_TIME TIMESTAMP(6),
    MEMO        VARCHAR(500),
    primary key (DIC_TYPE, DIC_CODE)
);