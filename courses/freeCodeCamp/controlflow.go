/*
Agenda - Control Flows
- Defer
- Panic
- Recover
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	// Defer
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	str := "start"
	defer fmt.Println(str)
	str = "end"

	// Panic
	// fmt.Println("start")
	// panic("something bad happened")
	// fmt.Println("end")

	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err = http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// Order: Defer --> Panic
	// fmt.Println("start")
	// defer fmt.Println("this was defered")
	// panic("something bad happened")
	// fmt.Println("end")

	// fmt.Println("start")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Error: ", err)
	// 	}
	// }()
	// panic("something bad happened")
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

/* Summary
- Defer
  - Used to delay execution of a statement until function exists
  - Useful to group "open" and "close" functions together
	- Be careful in loops
  - Run in LIFO(last-in, first-out) order
  - Arguments evaluatd at time defer is executed, not at time of called function execution
- Panic
  - Occur when program cannot continue at all
	- Don't use when file can't be opened, unless it is critical
	- Use for unrecoverable events - cannot obtain TCP port for web server
  - Function will stop executing
    - Deferred functions will still fire
  - If nothing handles panic, program will exit
- Recover
  - Used to recover from panics
  - Only useful in deferred functions
  - Current function will not attempt to continue, but higher functions in call stack will
*/
