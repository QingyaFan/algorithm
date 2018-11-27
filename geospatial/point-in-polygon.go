package geospatial

type point struct {
	x float64
	y float64
}

func PointInPolygon(nvert int, vertx []float64, verty []float64, testx float64, testy float64) bool {
	var i int = 0
	var j int
	var c bool = false

	for j = nvert - 1; i < nvert; i++ {
		j = i
		if ((verty[i] > testy) != (verty[j] > testy)) && (testx < (vertx[j]-vertx[i])*(testy-verty[i])/(verty[j]-verty[i])+vertx[i]) {
			c = !c
		}
	}

	return c
}
