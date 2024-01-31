create table public.Users
(
	user_id SERIAL PRIMARY KEY,
	username varchar(32) UNIQUE NOT NULL,
	user_email varchar(255) UNIQUE NOT NULL,
	hashed_password text NOT NULL,
	create_time TIMESTAMP NOT NULL
);