package main

/*
holds science definition and what it does
*/
type scienceDef struct {
	name string
	typ  int // uses int from ressources to specify in which science field
	pay  map[int]float64

	// building->resource->typeOfBoni->amount
	// can be used to reduce cost of dependent ressource
	numBoni map[int]map[int]map[int]float64

	// ship->typeOfBoni->amount
	limitBoni map[int]map[int]float64
}

// #############################################################################
// #							DATA
// #############################################################################

var (
	scienceData = []*scienceDef{
		{
			name: "Better Matter Collector",
			typ:  science,
			pay:  map[int]float64{science: 100.0},
			numBoni: map[int]map[int]map[int]float64{
				matter + collectorBuilding: {
					matter: {
						increaseBasePerc: 0.5,
					},
				},
			},
		},
	}
)
