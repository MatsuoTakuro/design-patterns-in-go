package builder

type Worker struct {
	name, position string
}

type workerMod func(*Worker)
