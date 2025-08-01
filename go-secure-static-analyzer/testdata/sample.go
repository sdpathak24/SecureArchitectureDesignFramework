package main

func main() {
	password := "supersecret" // hardcoded secret -- should be flagged
	api_key := "123456"       // hardcoded API key -- should be flagged
	normalVar := 32
	// print values
	println(password, api_key, normalVar)
}
