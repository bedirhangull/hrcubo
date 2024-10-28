CREATE TABLE "jobs" (
    "id" UUID PRIMARY KEY,
    "team_id" UUID NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "created_from" UUID NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "applicants" (
    "id" UUID PRIMARY KEY,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "stage" SMALLINT NOT NULL,
    "applied_position" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "phone" VARCHAR(15) NOT NULL UNIQUE,
    "resume_id" UUID NOT NULL,
    "rate" SMALLINT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "applications" (
    "id" UUID PRIMARY KEY,
    "applicant_id" UUID NOT NULL,
    "job_id" UUID NOT NULL,
    "applied_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "saved_applicants" (
    "id" UUID PRIMARY KEY,
    "user_id" UUID NOT NULL,
    "applier_id" UUID NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX "jobs_title_index" ON "jobs"("title");
CREATE INDEX "applicants_first_name_index" ON "applicants"("first_name");
CREATE INDEX "applicants_last_name_index" ON "applicants"("last_name");
CREATE INDEX "applicants_resume_id_index" ON "applicants"("resume_id");
CREATE INDEX "saved_applicants_user_id_index" ON "saved_applicants"("user_id");