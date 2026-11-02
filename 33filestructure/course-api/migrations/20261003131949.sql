-- Create "courses" table
CREATE TABLE "courses" (
  "course_id" uuid NOT NULL,
  "course_name" text NOT NULL,
  "course_price" bigint NOT NULL,
  "author_name" text NULL,
  "author_website" text NULL,
  PRIMARY KEY ("course_id")
);
