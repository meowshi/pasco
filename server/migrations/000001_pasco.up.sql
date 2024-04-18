CREATE TABLE "public"."env" (
    "key" text NOT NULL,
    "value" text NOT NULL,
    "edited_at" timestamp NOT NULL,
    CONSTRAINT "env_key" PRIMARY KEY ("key")
) WITH (oids = false);


CREATE TABLE "public"."event" (
    "uuid" uuid NOT NULL,
    "name" text NOT NULL,
    "google_sheet_cell" text NOT NULL,
    "locker_event_id" integer NOT NULL,
    "created_at" timestamptz NOT NULL,
    "allowed_friends" boolean NOT NULL,
    CONSTRAINT "event_uuid" PRIMARY KEY ("uuid")
) WITH (oids = false);


CREATE SEQUENCE pick_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."pick" (
    "yandexoid_login" text NOT NULL,
    "event_uuid" uuid,
    "with_friends" boolean NOT NULL,
    "is_list_success" boolean NOT NULL,
    "is_gift_success" boolean NOT NULL,
    "is_bracelet_success" boolean NOT NULL,
    "picked_at" timestamptz NOT NULL,
    "id" bigint DEFAULT nextval('pick_id_seq') NOT NULL,
    CONSTRAINT "pick_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE SEQUENCE plus_one_events_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."plus_one_event" (
    "id" integer DEFAULT nextval('plus_one_events_id_seq') NOT NULL,
    "name" text NOT NULL,
    CONSTRAINT "plus_one_event_name" UNIQUE ("name"),
    CONSTRAINT "plus_one_events_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE TABLE "public"."registration" (
    "event_uuid" uuid NOT NULL,
    "yandexoid_login" text NOT NULL,
    "friends" integer NOT NULL,
    "status" smallint NOT NULL,
    "status_cell" text NOT NULL,
    CONSTRAINT "registration_event_uuid_yandexoid_login" PRIMARY KEY ("event_uuid", "yandexoid_login")
) WITH (oids = false);


CREATE TABLE "public"."yandexoid" (
    "login" text NOT NULL,
    "name" text NOT NULL,
    "surname" text NOT NULL,
    CONSTRAINT "yandexoid_login" PRIMARY KEY ("login")
) WITH (oids = false);


ALTER TABLE ONLY "public"."pick" ADD CONSTRAINT "pick_event_uuid_fkey" FOREIGN KEY (event_uuid) REFERENCES event(uuid) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;
ALTER TABLE ONLY "public"."pick" ADD CONSTRAINT "pick_yandexoid_login_fkey" FOREIGN KEY (yandexoid_login) REFERENCES yandexoid(login) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;

ALTER TABLE ONLY "public"."registration" ADD CONSTRAINT "registration_event_uuid_fkey" FOREIGN KEY (event_uuid) REFERENCES event(uuid) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;
ALTER TABLE ONLY "public"."registration" ADD CONSTRAINT "registration_yandexoid_login_fkey" FOREIGN KEY (yandexoid_login) REFERENCES yandexoid(login) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;
