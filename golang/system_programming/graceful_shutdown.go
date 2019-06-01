// ch12
func signalHandler() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	s := <-signals
	switch s {
	case syscall.SIGINT:
		fmt.Println("SIGINT")
	case syscall.SIGTERM:
		fmt.Println("SIGTERM")
	}
}

func gracefulShutdown() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Handler: http.HandleFunc(func (w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "server pid: %d %v", os.Getpid(), os.Environ())
		})
	}
	go server.Serve(listeners[0])

	<- signals
	server.Shutdown(context.Background())
}
