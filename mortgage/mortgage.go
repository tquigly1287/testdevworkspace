package mortgage

import (
	"fmt"
	"net/http"
)

// GetBalance prints mortgage info
func GetBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Account balance: 96000 GBP. Status Healthy.")
}
