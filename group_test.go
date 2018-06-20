package httprouter_test

/*
	// 使用方法
func main() {
	gs := httprouter.NewGroup("/v1",
		httprouter.NSGroup("/api",
			httprouter.NSRouter("/index", "GET",
				func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
					fmt.Fprint(w, "Welcome!\n")
				}),
			httprouter.NSRouter("/hello", "POST",
				func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
					fmt.Fprint(w, "hi,boy!\n")
				}),
		),
	)
	r := httprouter.AddGroups(gs)
	log.Fatal(http.ListenAndServe(":8080", r))
}
*/
