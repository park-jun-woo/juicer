package ddl

import "regexp"

// tblName matches table names including schema-qualified and/or double-quoted identifiers:
//   users, public.profiles, "UserEvents", public."UserEvents"
const tblName = `((?:"[^"]+"|\w+)(?:\.(?:"[^"]+"|\w+))?)`

var (
	reCreateTable    = regexp.MustCompile(`(?i)^\s*CREATE\s+TABLE\s+(?:IF\s+NOT\s+EXISTS\s+)?` + tblName + `\s*\(`)
	reDropTable      = regexp.MustCompile(`(?i)^\s*DROP\s+TABLE\s+(?:IF\s+EXISTS\s+)?` + tblName)
	reAlterTable     = regexp.MustCompile(`(?is)^\s*ALTER\s+TABLE\s+(?:ONLY\s+)?` + tblName + `\s+(.+)`)
	reRenameTable    = regexp.MustCompile(`(?i)^\s*RENAME\s+TO\s+(\w+)`)
	reAddColumn      = regexp.MustCompile(`(?i)^\s*ADD\s+COLUMN\s+(?:IF\s+NOT\s+EXISTS\s+)?(.+)`)
	reDropColumn     = regexp.MustCompile(`(?i)^\s*DROP\s+COLUMN\s+(?:IF\s+EXISTS\s+)?(\w+)`)
	reAlterColumn    = regexp.MustCompile(`(?i)^\s*ALTER\s+COLUMN\s+(\w+)\s+(.+)`)
	reAddConstraint  = regexp.MustCompile(`(?is)^\s*ADD\s+CONSTRAINT\s+(.+)`)
	reDropConstraint = regexp.MustCompile(`(?i)^\s*DROP\s+CONSTRAINT\s+(?:IF\s+EXISTS\s+)?(\w+)`)
	reCreateIndex    = regexp.MustCompile(`(?i)^\s*CREATE\s+(?:UNIQUE\s+)?INDEX\s+(?:(?:IF\s+NOT\s+EXISTS\s+)?(?:"[^"]+"|\w+)\s+)?ON\s+` + tblName + `\s*\(`)
	reDropIndex      = regexp.MustCompile(`(?i)^\s*DROP\s+INDEX\s+(?:IF\s+EXISTS\s+)?` + tblName)
	reConstraintLine = regexp.MustCompile(`(?i)^\s*(?:FOREIGN\s+KEY|PRIMARY\s+KEY|UNIQUE|CHECK|CONSTRAINT)\b`)
)
