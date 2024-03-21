create table public.post
(
	post_id SERIAL PRIMARY KEY,
	post_image character varying(200),
	post_description character varying(350),
	create_time timestamp(6) without time zone,
	user_id integer
);
