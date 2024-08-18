-- +goose Up
-- +goose StatementBegin
create table post (
	id text primary key,
	user_id text not null,
	title text not null,
	content text not null,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null
);

create table post_settings (
	post_id text primary key,
	settings jsonb not null
);

create table comment (
	id text primary key,
	parent_id text not null,
	post_id text not null,
	user_id text not null,
	content text not null,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null
);

create index comment_post_id_idx on comment using btree (post_id);
create index comment_post_id_parent_id_idx on comment using btree (parent_id);
-- +goose StatementEnd
