// This function allows to pass any struct/map, and creates json from it, then writes to http.ResponseWriter
func WriteJson(w http.ResponseWriter, i interface{}) error {

	defer Recover("WriteJson")

	var j []byte

	j, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return err
	}

	// set Headers
	w.Header().Set("Content-Type", "application/json")
	// write to writer
	w.Write(j)
	return nil
}
