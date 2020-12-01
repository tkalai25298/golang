package handlers

import (
	"fmt"
	"net/http"
)


//ListRoles to check server health
func ListRoles(rw http.ResponseWriter,req *http.Request) {
	fmt.Fprintf(rw," client \n admin \n orderer \n peer \n")
}