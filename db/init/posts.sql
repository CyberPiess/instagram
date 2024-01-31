CREATE TABLE IF NOT EXISTS public."Posts"
(
    "postId" integer NOT NULL DEFAULT nextval('"Posts_postId_seq"'::regclass),
    "postImage" character varying(200) COLLATE pg_catalog."default" NOT NULL,
    "postDescription" character varying(350) COLLATE pg_catalog."default",
    "createTime" timestamp(6) without time zone NOT NULL,
    "userId" integer NOT NULL
)

TABLESPACE pg_default;