package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	xstr := r.FormValue("x")
	ystr := r.FormValue("y")

	x, err := strconv.Atoi(xstr)
	if err != nil {
		http.Error(w, "invalid x", http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(ystr)
	if err != nil {
		http.Error(w, "invalid y", http.StatusBadRequest)
		return
	}

	var result int = add(x, y)

	w.Header().Set("Content-Type", "text/html")
	w.Write(fmt.Appendf([]byte{}, "%d", result))
}

func add(x int, y int) int {
	return x + y
}

func HandleSub(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	xstr := r.FormValue("x")
	ystr := r.FormValue("y")

	x, err := strconv.Atoi(xstr)
	if err != nil {
		http.Error(w, "invalid x", http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(ystr)
	if err != nil {
		http.Error(w, "invalid y", http.StatusBadRequest)
		return
	}

	var result int = substract(x, y)

	w.Header().Set("Content-Type", "text/html")
	w.Write(fmt.Appendf([]byte{}, "%d", result))

}

func substract(x int, y int) int {
	return x - y
}

func HandleDiv(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	xstr := r.FormValue("x")
	ystr := r.FormValue("y")

	x, err := strconv.Atoi(xstr)
	if err != nil {
		http.Error(w, "invalid x", http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(ystr)
	if err != nil {
		http.Error(w, "invalid y", http.StatusBadRequest)
		return
	}

	result, err := divide(x, y)

	if err != nil {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "cannot divide by zero")
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.Write(fmt.Appendf([]byte{}, "%d", result))
	}

}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divide by zero")
	}
	return x / y, nil
}

func HandleMult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	xstr := r.FormValue("x")
	ystr := r.FormValue("y")

	x, err := strconv.Atoi(xstr)
	if err != nil {
		http.Error(w, "invalid x", http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(ystr)
	if err != nil {
		http.Error(w, "invalid y", http.StatusBadRequest)
		return
	}

	result := mult(x, y)

	w.Header().Set("Content-Type", "text/html")
	w.Write(fmt.Appendf([]byte{}, "%d", result))

}

func mult(x int, y int) int {
	return x * y
}
