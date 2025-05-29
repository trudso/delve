package engine

type ReplicatableFactory func(id string, ds map[string]any) Replicatable

type ReplicationCollection struct {
	id            string
	replicatables []Replicatable
	fixedFields   map[string]string
	factory       ReplicatableFactory
}

func (r ReplicationCollection) GetId() string {
	return r.id
}

func (r ReplicationCollection) IsChanged() bool {
	for _, replicatable := range r.replicatables {
		if replicatable.IsChanged() {
			return true
		}
	}
	return false
}

func (r ReplicationCollection) ResetToChanged() {
	for _, replicatable := range r.replicatables {
		replicatable.ResetToChanged()
	}
}

func (r ReplicationCollection) BuildChangeSet() map[string]any {
	if !r.IsChanged() {
		return nil
	}
	data := map[string]any{}
	for _, replicatable := range r.replicatables {
		var ds = BuildChangeSet(replicatable)
		for k, v := range ds {
			data[k] = v
		}
	}

	if len(data) != 0 {
		for key, v := range r.fixedFields {
			data[key] = v
		}
	}

	return map[string]any{r.id: data}
}

func (r ReplicationCollection) BuildSnapshot() map[string]any {
	data := map[string]any{}
	for _, replicatable := range r.replicatables {
		var ds = BuildSnapshot(replicatable)
		for k, v := range ds {
			data[k] = v
		}
	}
	if len(data) != 0 {
		for key, v := range r.fixedFields {
			data[key] = v
		}
	}
	return map[string]any{r.id: data}
}

func (r ReplicationCollection) ApplyDataSet(dataSet map[string]any) {
	if value, found := dataSet[r.id]; found {
		ds := value.(map[string]any)

		for dsKey, dsValue := range ds {
			replicatable, hasElement := r.getReplicatable(dsKey)
			if !hasElement && r.factory != nil && !r.isFixedField(dsKey) {
				replicatable = r.factory(dsKey, dsValue.(map[string]any))
				ApplyDataSet(replicatable, map[string]any{dsKey: dsValue})
				r.AddElement(replicatable) // issue as r is not a reference
				continue
			}

			if replicatable != nil {
				//fmt.Println("Applying", dsKey, "to", replicatable)
				ApplyDataSet(replicatable, map[string]any{dsKey: dsValue})
			} else {
				//fmt.Fprintln(os.Stderr, "No replicatable found for:", dsKey)
			}
		}
	}
}

func (r ReplicationCollection) getReplicatable(key string) (Replicatable, bool) {
	for _, replicatable := range r.replicatables {
		if replicatable.GetId() == key {
			return replicatable, true
		}
	}
	return nil, false
}

func (r ReplicationCollection) isFixedField(key string) bool {
	_, found := r.fixedFields[key]
	return found
}

func (r ReplicationCollection) AddCollection(other ReplicationCollection) {
	r.replicatables = append(r.replicatables, other.replicatables...)
}

func (r *ReplicationCollection) AddElement(replicatable Replicatable) {
	r.replicatables = append(r.replicatables, replicatable)
}

func NewReplicationCollection(id string, replicatables []Replicatable, fixedFields map[string]string, factory ReplicatableFactory) ReplicationCollection {
	return ReplicationCollection{
		id:            id,
		replicatables: replicatables,
		fixedFields:   fixedFields,
		factory:       factory,
	}
}
