CREATE TABLE "team" (
    "id" UUID PRIMARY KEY,
    "team_name" VARCHAR(255) NOT NULL,
    "team_image_id" UUID NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "user_teams" (
    "id" UUID PRIMARY KEY,
    "user_id" UUID NOT NULL,
    "team_id" UUID NOT NULL,
    "joined_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "invites" (
    "id" UUID PRIMARY KEY,
    "from_id" UUID NOT NULL,
    "to_id" UUID NOT NULL,
    "team_id" UUID NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "team_name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX "team_team_name_index" ON "team"("team_name");
CREATE INDEX "invites_from_id_index" ON "invites"("from_id");
CREATE INDEX "invites_to_id_index" ON "invites"("to_id");