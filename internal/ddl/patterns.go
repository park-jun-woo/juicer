package ddl

import "regexp"

var (
	reCreateTable    = regexp.MustCompile(`(?i)^\s*CREATE\s+TABLE\s+(?:IF\s+NOT\s+EXISTS\s+)?(\w+)\s*\(`)
	reDropTable      = regexp.MustCompile(`(?i)^\s*DROP\s+TABLE\s+(?:IF\s+EXISTS\s+)?(\w+)`)
	reAlterTable     = regexp.MustCompile(`(?is)^\s*ALTER\s+TABLE\s+(\w+)\s+(.+)`)
	reRenameTable    = regexp.MustCompile(`(?i)^\s*RENAME\s+TO\s+(\w+)`)
	reAddColumn      = regexp.MustCompile(`(?i)^\s*ADD\s+COLUMN\s+(?:IF\s+NOT\s+EXISTS\s+)?(.+)`)
	reDropColumn     = regexp.MustCompile(`(?i)^\s*DROP\s+COLUMN\s+(?:IF\s+EXISTS\s+)?(\w+)`)
	reAlterColumn    = regexp.MustCompile(`(?i)^\s*ALTER\s+COLUMN\s+(\w+)\s+(.+)`)
	reAddConstraint  = regexp.MustCompile(`(?i)^\s*ADD\s+CONSTRAINT\s+(.+)`)
	reDropConstraint = regexp.MustCompile(`(?i)^\s*DROP\s+CONSTRAINT\s+(?:IF\s+EXISTS\s+)?(\w+)`)
	reCreateIndex    = regexp.MustCompile(`(?i)^\s*CREATE\s+(?:UNIQUE\s+)?INDEX\s+(?:IF\s+NOT\s+EXISTS\s+)?(\w+)\s+ON\s+(\w+)\s*\(`)
	reDropIndex      = regexp.MustCompile(`(?i)^\s*DROP\s+INDEX\s+(?:IF\s+EXISTS\s+)?(\w+)`)
	reConstraintLine = regexp.MustCompile(`(?i)^\s*(?:FOREIGN\s+KEY|PRIMARY\s+KEY|UNIQUE|CHECK|CONSTRAINT)\b`)
)
