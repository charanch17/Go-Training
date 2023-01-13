package main

import "fmt"

var diseaseMapper = make(map[int]string)

type Patient interface {
	Diagnoisedwith() string
}
type HumanBeing struct {
	name          string
	claimid       int
	diagnosiscode int
}

func (h HumanBeing) Diagnoisedwith() string {
	if h.claimid != 0 && h.diagnosiscode != 0 {
		return h.name + " is diagnosied with " + diseaseMapper[h.diagnosiscode]
	}
	return h.name + " is healthy"
}

func main() {
	diseaseMapper[1] = "cough"
	diseaseMapper[2] = "cold"
	diseaseMapper[3] = "covid"
	HumanBeing1 := HumanBeing{name: "charan", claimid: 2, diagnosiscode: 1}
	var patient1 Patient
	// intialily dynamic type will be nil  and value will be nil
	fmt.Printf("before %T\n", patient1)
	fmt.Printf("before %v\n", patient1)
	patient1 = HumanBeing1
	fmt.Printf("after %T\n", patient1)
	fmt.Printf("after %v\n", patient1)

}
