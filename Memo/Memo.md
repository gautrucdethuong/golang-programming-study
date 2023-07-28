1. %T\n: Get type of variable 
2. Declare variable
a. var <name> <type>
b. var <name> = <value>
c. <name> := <value>
3. GET INPUT FROM USERS:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter your answer with numer 0 - 2: ")
	scanner.Scan()
	input := scanner.Text()
	answer, err := strconv.Atoi(input)
    	if err != nil {
		log.Panic(err)
	}
    ----------
    fmt.Print("Enter an Integer Number: ")
	fmt.Scan(&number)
	fmt.Print("Enter the range or end value: ")
	fmt.Scan(&max_range)
4. 