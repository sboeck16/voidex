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
			ressourcesNeeded: map[int]float64{
				fabric: 10,
			},
			unlocksBntID: []int{
				fabric + collectorBuilding,
				fabric + collectorBuilding + collectorAdd,
				fabric + collectorBuilding + collectorSub,
				rareMatter,
			},
		},
		{
			ressourcesNeeded: map[int]float64{
				rareMatter: 10,
			},
			unlocksBntID: []int{
				rareMatter + collectorBuilding,
				rareMatter + collectorBuilding + collectorAdd,
				rareMatter + collectorBuilding + collectorSub,
				components,
			},
		},
		{
			ressourcesNeeded: map[int]float64{
				components: 10,
			},
			unlocksBntID: []int{
				components + collectorBuilding,
				components + collectorBuilding + collectorAdd,
				components + collectorBuilding + collectorSub,
				moduls,
				// TODO FTR
			},
		},
		{
			ressourcesNeeded: map[int]float64{
				moduls: 10,
			},
			unlocksBntID: []int{
				moduls + collectorBuilding,
				moduls + collectorBuilding + collectorAdd,
				moduls + collectorBuilding + collectorSub,
				exoticMatter,
			},
		},
		{
			ressourcesNeeded: map[int]float64{
				exoticMatter: 10,
			},
			unlocksBntID: []int{
				exoticMatter + collectorBuilding,
				exoticMatter + collectorBuilding + collectorAdd,
				exoticMatter + collectorBuilding + collectorSub,
				structures,
			},
		},
		{
			ressourcesNeeded: map[int]float64{
				structures: 10,
			},
			unlocksBntID: []int{
				structures + collectorBuilding,
				structures + collectorBuilding + collectorAdd,
				structures + collectorBuilding + collectorSub,
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
	}
)
