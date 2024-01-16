package main

type unlockGoal struct {
	ressourcesNeeded map[int]float64
	unlocksBntID     []int
}

var (
	goalsForButtons = []*unlockGoal{
		// not much here, but enables first ressource click
		{
			ressourcesNeeded: map[int]float64{},
			unlocksBntID: []int{
				matter,
			},
		},
		{
			ressourcesNeeded: map[int]float64{
				matter: 10,
			},
			unlocksBntID: []int{
				matter + collectorBuilding,
				matter + collectorBuilding + collectorAdd,
				matter + collectorBuilding + collectorSub,
				fabric,
			},
		},
		{
			ressourcesNeeded: map[int]float64{},
			unlocksBntID:     []int{},
		},
		{
			ressourcesNeeded: map[int]float64{},
			unlocksBntID:     []int{},
		},
		{
			ressourcesNeeded: map[int]float64{},
			unlocksBntID:     []int{},
		},
		{
			ressourcesNeeded: map[int]float64{},
			unlocksBntID:     []int{},
		},
		{
			ressourcesNeeded: map[int]float64{},
			unlocksBntID:     []int{},
		},
	}
)
