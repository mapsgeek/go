package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/mapsgeek/tolling/types"
)

func GenLocation(bbox *types.BoundingBox) (float64, float64) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	lat := rand.Float64()*(bbox.MaxLat-bbox.MinLat) + bbox.MinLat
	lng := rand.Float64()*(bbox.MaxLng-bbox.MinLng) + bbox.MinLng
	return lat, lng
}

func GenerateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}
