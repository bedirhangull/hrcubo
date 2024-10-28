ALTER TABLE "saved_applicants" DROP CONSTRAINT IF EXISTS "saved_applicants_applier_id_foreign";
ALTER TABLE "saved_applicants" DROP CONSTRAINT IF EXISTS "saved_applicants_user_id_foreign";
ALTER TABLE "applications" DROP CONSTRAINT IF EXISTS "applications_job_id_foreign";
ALTER TABLE "applications" DROP CONSTRAINT IF EXISTS "applications_applicant_id_foreign";
ALTER TABLE "applicants" DROP CONSTRAINT IF EXISTS "applicants_resume_id_foreign";
ALTER TABLE "jobs" DROP CONSTRAINT IF EXISTS "jobs_created_from_foreign";
ALTER TABLE "jobs" DROP CONSTRAINT IF EXISTS "jobs_team_id_foreign";
ALTER TABLE "invites" DROP CONSTRAINT IF EXISTS "invites_team_id_foreign";
ALTER TABLE "invites" DROP CONSTRAINT IF EXISTS "invites_to_id_foreign";
ALTER TABLE "invites" DROP CONSTRAINT IF EXISTS "invites_from_id_foreign";
ALTER TABLE "user_teams" DROP CONSTRAINT IF EXISTS "user_teams_team_id_foreign";
ALTER TABLE "user_teams" DROP CONSTRAINT IF EXISTS "user_teams_user_id_foreign";
ALTER TABLE "team" DROP CONSTRAINT IF EXISTS "team_team_image_id_foreign";
ALTER TABLE "assets" DROP CONSTRAINT IF EXISTS "assets_user_id_foreign";
ALTER TABLE "users" DROP CONSTRAINT IF EXISTS "users_image_id_foreign";