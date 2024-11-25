CREATE TABLE "saved_applicants"(
    "id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "applier_id" UUID NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "saved_applicants" ADD PRIMARY KEY("id");
CREATE INDEX "saved_applicants_user_id_index" ON
    "saved_applicants"("user_id");
CREATE TABLE "assets"(
    "id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "file_name" VARCHAR(255) NOT NULL,
    "file_type" VARCHAR(255) NOT NULL,
    "file_path" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "assets" ADD PRIMARY KEY("id");
CREATE INDEX "assets_user_id_index" ON
    "assets"("user_id");
ALTER TABLE
    "assets" ADD CONSTRAINT "assets_file_path_unique" UNIQUE("file_path");
CREATE TABLE "invites"(
    "id" UUID NOT NULL,
    "from_id" UUID NOT NULL,
    "to_id" UUID NOT NULL,
    "team_id" UUID NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "team_name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "invites" ADD PRIMARY KEY("id");
CREATE INDEX "invites_from_id_index" ON
    "invites"("from_id");
CREATE INDEX "invites_to_id_index" ON
    "invites"("to_id");
CREATE TABLE "users"(
    "id" UUID NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "role" VARCHAR(255) NULL,
    "image_id" UUID NULL,
    "is_user_premium" BOOLEAN NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "last_login_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_email_unique" UNIQUE("email");
ALTER TABLE
    "users" ADD CONSTRAINT "users_phone_unique" UNIQUE("phone");
CREATE TABLE "jobs"(
    "id" UUID NOT NULL,
    "team_id" UUID NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "created_from" UUID NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "jobs" ADD PRIMARY KEY("id");
CREATE INDEX "jobs_title_index" ON
    "jobs"("title");
CREATE TABLE "applications"(
    "id" UUID NOT NULL,
    "applicant_id" UUID NOT NULL,
    "job_id" UUID NOT NULL,
    "applied_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "applications" ADD PRIMARY KEY("id");
CREATE TABLE "team"(
    "id" UUID NOT NULL,
    "team_name" VARCHAR(255) NOT NULL,
    "team_image_id" UUID NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "team" ADD PRIMARY KEY("id");
CREATE INDEX "team_team_name_index" ON
    "team"("team_name");
CREATE TABLE "applicants"(
    "id" UUID NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "stage" SMALLINT NOT NULL,
    "applied_position" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(15) NOT NULL,
    "resume_id" UUID NOT NULL,
    "rate" SMALLINT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "applicants" ADD PRIMARY KEY("id");
CREATE INDEX "applicants_first_name_index" ON
    "applicants"("first_name");
CREATE INDEX "applicants_last_name_index" ON
    "applicants"("last_name");
ALTER TABLE
    "applicants" ADD CONSTRAINT "applicants_email_unique" UNIQUE("email");
ALTER TABLE
    "applicants" ADD CONSTRAINT "applicants_phone_unique" UNIQUE("phone");
CREATE INDEX "applicants_resume_id_index" ON
    "applicants"("resume_id");
CREATE TABLE "user_teams"(
    "id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "team_id" UUID NOT NULL,
    "joined_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "user_teams" ADD PRIMARY KEY("id");
ALTER TABLE
    "saved_applicants" ADD CONSTRAINT "saved_applicants_applier_id_foreign" FOREIGN KEY("applier_id") REFERENCES "applicants"("id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_image_id_foreign" FOREIGN KEY("image_id") REFERENCES "assets"("id");
ALTER TABLE
    "jobs" ADD CONSTRAINT "jobs_team_id_foreign" FOREIGN KEY("team_id") REFERENCES "team"("id");
ALTER TABLE
    "applications" ADD CONSTRAINT "applications_applicant_id_foreign" FOREIGN KEY("applicant_id") REFERENCES "applicants"("id");
ALTER TABLE
    "applicants" ADD CONSTRAINT "applicants_resume_id_foreign" FOREIGN KEY("resume_id") REFERENCES "assets"("id");
ALTER TABLE
    "saved_applicants" ADD CONSTRAINT "saved_applicants_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "applications" ADD CONSTRAINT "applications_job_id_foreign" FOREIGN KEY("job_id") REFERENCES "jobs"("id");
ALTER TABLE
    "team" ADD CONSTRAINT "team_team_image_id_foreign" FOREIGN KEY("team_image_id") REFERENCES "assets"("id");
ALTER TABLE
    "user_teams" ADD CONSTRAINT "user_teams_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "jobs" ADD CONSTRAINT "jobs_created_from_foreign" FOREIGN KEY("created_from") REFERENCES "users"("id");
ALTER TABLE
    "invites" ADD CONSTRAINT "invites_to_id_foreign" FOREIGN KEY("to_id") REFERENCES "users"("id");
ALTER TABLE
    "assets" ADD CONSTRAINT "assets_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "invites" ADD CONSTRAINT "invites_from_id_foreign" FOREIGN KEY("from_id") REFERENCES "users"("id");
ALTER TABLE
    "invites" ADD CONSTRAINT "invites_team_id_foreign" FOREIGN KEY("team_id") REFERENCES "team"("id");
ALTER TABLE
    "user_teams" ADD CONSTRAINT "user_teams_team_id_foreign" FOREIGN KEY("team_id") REFERENCES "team"("id");