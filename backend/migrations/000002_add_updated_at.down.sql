DROP TRIGGER IF EXISTS work_experiences_updated_at ON work_experiences;
DROP TRIGGER IF EXISTS projects_updated_at ON projects;
DROP TRIGGER IF EXISTS skills_updated_at ON skills;

DROP FUNCTION IF EXISTS set_updated_at();

ALTER TABLE work_experiences DROP COLUMN IF EXISTS updated_at;
ALTER TABLE projects DROP COLUMN IF EXISTS updated_at;
ALTER TABLE skills DROP COLUMN IF EXISTS updated_at;
