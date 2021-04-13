package main

/*func PostNaviHandler(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Codificar")
		fmt.Println("Codificar")
		Satelites := []satellite{
			satellite{name: "kenobi", distance: "100.0", message: "mensaje"},
		}
		fmt.Println(Satelites)
		json.NewEncoder(w).Encode(Satelites)
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}*/
