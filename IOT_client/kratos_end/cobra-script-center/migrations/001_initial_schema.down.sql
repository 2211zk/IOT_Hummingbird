-- Drop indexes
DROP INDEX IF EXISTS idx_scheduled_tasks_is_active;
DROP INDEX IF EXISTS idx_scheduled_tasks_script_id;
DROP INDEX IF EXISTS idx_executions_status;
DROP INDEX IF EXISTS idx_executions_user_id;
DROP INDEX IF EXISTS idx_executions_script_id;
DROP INDEX IF EXISTS idx_scripts_is_active;
DROP INDEX IF EXISTS idx_scripts_created_by;

-- Drop tables
DROP TABLE IF EXISTS scheduled_tasks;
DROP TABLE IF EXISTS executions;
DROP TABLE IF EXISTS scripts;
DROP TABLE IF EXISTS users;