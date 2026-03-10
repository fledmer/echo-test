-- Modify "request_logs" table
ALTER TABLE "request_logs" ADD COLUMN "jopa_time" timestamptz NOT NULL;
