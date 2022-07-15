/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : PostgreSQL
 Source Server Version : 140004
 Source Host           : localhost:5432
 Source Catalog        : ojire
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140004
 File Encoding         : 65001

 Date: 15/07/2022 10:44:10
*/


-- ----------------------------
-- Table structure for produk
-- ----------------------------
DROP TABLE IF EXISTS "public"."produk";
CREATE TABLE "public"."produk" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "nama" varchar(255) COLLATE "pg_catalog"."default",
  "sku" varchar(255) COLLATE "pg_catalog"."default",
  "jumlah" int4
)
;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "nama" varchar(255) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "no_telp" varchar(255) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Primary Key structure for table produk
-- ----------------------------
ALTER TABLE "public"."produk" ADD CONSTRAINT "barang_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("email");
