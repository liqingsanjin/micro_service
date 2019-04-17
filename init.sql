CREATE DATABASE if not exists apsmgm DEFAULT CHARACTER SET utf8mb4;

use apsmgm;

create table if not exists TBL_AUTH_ASSIGNMENT
(
  ITEM_NAME  VARCHAR(64) not null,
  USER_ID    INT(10) not null,
  CREATED_AT INT(22),
  primary key (USER_ID, ITEM_NAME)
);

create table if not exists TBL_AUTH_ITEM
(
  NAME        VARCHAR(64) not null primary key,
  TYPE        INT       not null,
  DESCRIPTION VARCHAR(200),
  RULE_NAME   VARCHAR(64),
  ITEM_DATA   VARCHAR(500),
  CREATED_AT  INT(10),
  UPDATED_AT  INT(10),
  ITEM_CODE   VARCHAR(10),
  PARENT_ITEM VARCHAR(64),
  CREATE_USER INT
);

create table if not exists TBL_AUTH_ITEM_CHILD
(
  PARENT VARCHAR(64) not null,
  CHILD  VARCHAR(64) not null,
  primary key (PARENT, CHILD)
);

create table if not exists TBL_MENU
(
  ID         int(10) auto_increment not null primary key,
  NAME       VARCHAR(128),
  PARENT     int(10),
  MENU_ROUTE VARCHAR(256),
  MENU_ORDER int(10),
  MENU_DATA  VARCHAR(500)
);

create table if not exists TBL_USER
(
  USER_ID              int(10)  auto_increment   not null primary key,
  LEAGUER_NO           VARCHAR(20)  not null,
  USER_NAME            VARCHAR(32)  not null unique,
  AUTH_KEY             VARCHAR(32)  not null,
  PASSWORD_HASH        VARCHAR(256) not null,
  PASSWORD_RESET_TOKEN VARCHAR(256),
  EMAIL                VARCHAR(256),
  USER_TYPE            VARCHAR(10)  not null,
  USER_INFO            VARCHAR(255),
  USER_STATUS          int not null,
  USER_NOTICE          VARCHAR(1024),
  REC_CRT_TS           DATETIME,
  REC_UPD_TS           DATETIME,
  PARENT_USER_NAME     VARCHAR(32)
) collate utf8_bin;

create table if not exists TBL_INS_INF
(
  INS_ID_CD          VARCHAR(11) not null primary key,
  INS_COMPANY_CD     VARCHAR(11),
  INS_TYPE           VARCHAR(1)  not null,
  INS_NAME           VARCHAR(100),
  INS_PROV_CD        VARCHAR(6),
  INS_CITY_CD        VARCHAR(6),
  INS_REGION_CD      VARCHAR(10),
  INS_STA            VARCHAR(2)  not null,
  INS_STLM_TP        VARCHAR(1),
  INS_ALO_STLM_CYCLE VARCHAR(1),
  INS_ALO_STLM_MD    VARCHAR(1),
  INS_STLM_C_NM      VARCHAR(80),
  INS_STLM_C_ACCT    VARCHAR(40),
  INS_STLM_C_BK_NO   VARCHAR(12),
  INS_STLM_C_BK_NM   VARCHAR(200),
  INS_STLM_D_NM      VARCHAR(80),
  INS_STLM_D_ACCT    VARCHAR(40),
  INS_STLM_D_BK_NO   VARCHAR(12),
  INS_STLM_D_BK_NM   VARCHAR(200),
  MSG_RESV_FLD1      VARCHAR(255),
  MSG_RESV_FLD2      VARCHAR(255),
  MSG_RESV_FLD3      VARCHAR(255),
  MSG_RESV_FLD4      VARCHAR(255),
  MSG_RESV_FLD5      VARCHAR(255),
  MSG_RESV_FLD6      VARCHAR(255),
  MSG_RESV_FLD7      VARCHAR(255),
  MSG_RESV_FLD8      VARCHAR(255),
  MSG_RESV_FLD9      VARCHAR(255),
  MSG_RESV_FLD10     VARCHAR(255),
  REC_OPR_ID         VARCHAR(1),
  REC_UPD_OPR        VARCHAR(32),
  REC_CRT_TS         DATETIME,
  REC_UPD_TS         DATETIME,
  index(INS_NAME)
);