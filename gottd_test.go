package gotdd

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Example() {
	// Contains the requirements as discussed with the PO...

	// I would like to have a program where I can register lunchtalks and review them,
	// As a user I would like to

	// Specify some lunch talks

	lt1 := LunchTalk{Title: "Tdd with Go", Speaker: "Ken Lomax"}
	lt2 := LunchTalk{Title: "Kyma with drobes", Speaker: "Mr Engelke"}
	fmt.Printf("lt1 %+v\n", lt1)
	fmt.Printf("lt2 %+v\n", lt2)

	// List them
	r := Register{}
	r.AddLunchTalk(lt1)
	r.AddLunchTalk(lt2)
	lts3 := r.GetLunchTalks()
	fmt.Printf("lt3 %+v\n", lts3)

	// Add reviews to them
	rev4 := Review{Comment: "Pile of poo"}
	r.AddReview(0, rev4)
	lts4 := r.GetLunchTalks()
	fmt.Printf("lt4 %+v\n", lts4)

	// Adjust the reviews
	rev5 := Review{Comment: "Amazing!!!!"}
	r.AdjustReview(0, 0, rev5)
	lts5 := r.GetLunchTalks()
	fmt.Printf("lt5 %+v\n", lts5)

	// They'll be some more requirements later,,,,

	// Please have some performance KPI,
	// You can run go test -bench=.

	// Error Handling!!
	e6 := r.AddLunchTalk(LunchTalk{Title: "", Speaker: ""})
	fmt.Printf("e6 %+v\n", e6)

	e7 := r.AddReview(0, Review{})
	fmt.Printf("e7 %+v\n", e7)

	e8 := r.AddReview(99, Review{Comment: "Bum"})
	fmt.Printf("e8 %+v\n", e8)

	e9 := r.AdjustReview(0, 99, Review{Comment: "Bum"})
	fmt.Printf("e9 %+v\n", e9)

	e10 := r.AdjustReview(0, 0, Review{Comment: ""})
	fmt.Printf("e10 %+v\n", e10)
	// More to come in our next meerting

	// Can multiple users use this at once!?

	// Output:
	// lt1 {Title:Tdd with Go Speaker:Ken Lomax Reviews:[]}
	// lt2 {Title:Kyma with drobes Speaker:Mr Engelke Reviews:[]}
	// lt3 [{Title:Tdd with Go Speaker:Ken Lomax Reviews:[]} {Title:Kyma with drobes Speaker:Mr Engelke Reviews:[]}]
	// lt4 [{Title:Tdd with Go Speaker:Ken Lomax Reviews:[{Comment:Pile of poo}]} {Title:Kyma with drobes Speaker:Mr Engelke Reviews:[]}]
	// lt5 [{Title:Tdd with Go Speaker:Ken Lomax Reviews:[{Comment:Amazing!!!!}]} {Title:Kyma with drobes Speaker:Mr Engelke Reviews:[]}]
	// e6 Missing Data
	// e7 Missing Data
	// e8 Out of bounds
	// e9 Out of bounds
	// e10 Missing Data
}

func Benchmark(b *testing.B) {
	r := Register{}
	for i := 0; i < b.N; i++ {
		r.AddLunchTalk(LunchTalk{Title: "Tdd with Go", Speaker: "Ken Lomax"})
		r.AddReview(rand.Intn(100), Review{Comment: "Pile of poo"})
		r.AdjustReview(rand.Intn(100), rand.Intn(100), Review{Comment: "Pile of poo"})
	}
}

func TestMultipleCalls(t *testing.T) {
	r := Register{}
	n := 100
	for i := 0; i < n; i++ {
		go r.AddLunchTalk(LunchTalk{Title: "Tdd with Go", Speaker: "Ken Lomax"})
	}
	for i := 0; i < n; i++ {
		go r.AddReview(rand.Intn(100), Review{Comment: "Pile of poo"})
	}
	for i := 0; i < n; i++ {
		go r.AdjustReview(rand.Intn(100), rand.Intn(100), Review{Comment: "Pile of poo"})
	}
	time.Sleep(time.Second * 2)

}
