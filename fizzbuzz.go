package main

func fizzBuzz() {
	for i := 1; i<100; i++ {
		// if i is divisible by 15 then it is divisible by 3 and 5
		if i % 15 == 0 {
			println("FizzBuzz")
		} else if i % 3 == 0 {
			println("Fizz")
		} else if i % 5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}