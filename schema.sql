CREATE TABLE "public"."course_users"
(
    "id"          bigserial    NOT NULL,
    "status"      smallint     NOT NULL,
    "uid"         bigint       NOT NULL,
    "name"        varchar(256) NOT NULL,
    "email"       varchar(256) NOT NULL,
    "inserted_at" timestamptz  NOT NULL,
    "updated_at"  timestamptz  NOT NULL,
    PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "unique_course_users_uid" ON "public"."course_users" USING btree ("uid");

CREATE TABLE "public"."course_courses"
(
    "id"          bigserial    NOT NULL,
    "status"      smallint     NOT NULL,
    "name"        varchar(256) NOT NULL,
    "desc"        text         NOT NULL,
    "quota"       bigint       NOT NULL,
    "started_at"  timestamptz  NOT NULL,
    "ended_at"    timestamptz  NOT NULL,
    "inserted_at" timestamptz  NOT NULL,
    "updated_at"  timestamptz  NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."course_user_courses"
(
    "id"          bigserial   NOT NULL,
    "status"      smallint    NOT NULL,
    "user_id"     bigint      NOT NULL,
    "course_id"   bigint      NOT NULL,
    "inserted_at" timestamptz NOT NULL,
    "updated_at"  timestamptz NOT NULL,
    PRIMARY KEY ("id")
);
CREATE INDEX "index_course_user_courses_status_user_id" ON "public"."course_user_courses" USING btree ("status", "user_id");
CREATE INDEX "index_course_user_courses_status_course_id" ON "public"."course_user_courses" USING btree ("status", "course_id");
