// Exercise 1.11:
// Try fetchall with longer argument lists, such as samples from the top million web sites available at alexa.com.
// How does the program behave if a web site just doesn't respond?
// (Section 8.9 describes mechanisms for coping in such cases.)

/*
# Stop go-routine

	func fetch(url string, ch chan<- string) {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			ch <- fmt.Sprintf("%.3fs %v\n", time.Since(start).Seconds(), err) // send to channel ch
			return                                                            //  quit signal for current go-routine
		}
	...

# Run unit test (request-timeout set to 8 seconds)

		$ go test -v
		=== RUN   TestCreateOutDir
		--- PASS: TestCreateOutDir (0.00s)
		=== RUN   TestFetchAll
	    	ch1ex11_test.go:38:
		--- PASS: TestFetchAll (8.00s)
		PASS
		ok      ch1ex10/ch1ex11 8.011s
*/
package ch1ex11
