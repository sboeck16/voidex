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
				fighters + collectorBuilding,
				fighters + collectorBuilding + collectorAdd,
				fighters + collectorBuilding + collectorSub,
				fighters + shipBuildingMax,
				fighters + shipBuildingMax + collectorAdd,
				fighters + shipBuildingMax + collectorSub,
			},
		},
		{
			ressourcesNeeded: map[int]float64{
				matter: 5,
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
				fabric: 5,
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
				rareMatter: 5,
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
				components: 5,
			},
			unlocksBntID: []int{
				components + collectorBuilding,
				components + collectorBuilding + collectorAdd,
				components + collectorBuilding + collectorSub,
				moduls,
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
