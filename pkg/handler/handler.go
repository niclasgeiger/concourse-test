package handler

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
)

type Operation string

const (
	ADD  Operation = "addition"
	SUB  Operation = "subtraction"
	PROD Operation = "multiplication"
)

var (
	unknownOperationErr = errors.New("unknown operation")
)

var calc Calculator = new(calculator)

func HandleRandomCalculation(w http.ResponseWriter, req *http.Request) {
	operation := randomOperation()
	a, b := rand.Intn(1000), rand.Intn(1000)
	result, err := calc.Do(a, b, operation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	output := struct {
		A         int       `json:"a"`
		B         int       `json:"b"`
		Operation Operation `json:"operation"`
		Result    int       `json:"result"`
	}{
		A:         a,
		B:         b,
		Operation: operation,
		Result:    result,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

type Calculator interface {
	Do(a, b int, op Operation) (int, error)
	Sum(a, b int) int
	Sub(a, b int) int
	Prod(a, b int) int
}

type calculator struct {
}

func (c calculator) Do(a, b int, op Operation) (int, error) {
	switch op {
	case ADD:
		return c.Sum(a, b), nil
	case SUB:
		return c.Sub(a, b), nil
	case PROD:
		return c.Prod(a, b), nil
	}
	return 0, unknownOperationErr
}

func (c calculator) Sum(a, b int) int {
	return a + b
}

func (c calculator) Sub(a, b int) int {
	return a - b
}

func (c calculator) Prod(a, b int) int {
	return a * b
}

func randomOperation() Operation {
	return []Operation{ADD, SUB, PROD}[rand.Intn(3)]
}
