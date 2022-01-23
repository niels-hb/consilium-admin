package handlers

func startGeneration(userCount int, transactionsMin int, transactionsMax int, schedulesMin int, schedulesMax int, dryRun bool) {
	println("Running generate with:")
	println("count:", userCount)
	println("transactions-min:", transactionsMin)
	println("transactions-max:", transactionsMax)
	println("schedules-min:", schedulesMin)
	println("schedules-max:", schedulesMax)
	println("dry-run:", dryRun)
}
