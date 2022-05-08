package builder

type WorkerBuilder struct {
	actions []workerMod
}

func (wb *WorkerBuilder) Build() *Worker {
	w := Worker{}
	for _, a := range wb.actions {
		a(&w)
	}
	return &w
}

func (wb *WorkerBuilder) Called(name string) *WorkerBuilder {
	wb.actions = append(wb.actions, func(w *Worker) {
		w.name = name
	})
	return wb
}

func (wb *WorkerBuilder) WorksAsA(position string) *WorkerBuilder {
	wb.actions = append(wb.actions, func(w *Worker) {
		w.position = position
	})
	return wb
}
