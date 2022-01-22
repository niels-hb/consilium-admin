package handlers

import "github.com/thatisuday/commando"

func Export(_ map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	uid, _ := flags["uid"].GetString()
	target, _ := flags["target"].GetString()

	startExport(uid, target)
}

func Import(_ map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	uid, _ := flags["uid"].GetString()
	source, _ := flags["source"].GetString()
	dryRun, _ := flags["dry-run"].GetBool()

	startImport(uid, source, dryRun)
}

func Migrate(_ map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	from, _ := flags["from"].GetString()
	to, _ := flags["to"].GetString()
	dryRun, _ := flags["dry-run"].GetBool()

	startMigration(from, to, dryRun)
}
