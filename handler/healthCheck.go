package handler

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	res := fmt.Sprintf("Health Check, Status: %d\n", status)
	w.Write([]byte(res))
}
