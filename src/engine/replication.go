package engine

type Replicatable interface {
	GetId() string
	ResetToChanged()
	IsChanged() bool
	BuildChangeSet() map[string]any
	BuildSnapshot() map[string]any
	ApplyDataSet(dataSet map[string]any)
}

func BuildChangeSet(replicatable Replicatable) map[string]any {
	return replicatable.BuildChangeSet()
}

func BuildSnapshot(replicatable Replicatable) map[string]any {
	return replicatable.BuildSnapshot()
}

func ResetToChanged(replicatable Replicatable) {
	replicatable.ResetToChanged()
}

func ApplyDataSet(replicatable Replicatable, dataSet map[string]any) {
	replicatable.ApplyDataSet(dataSet)
}
