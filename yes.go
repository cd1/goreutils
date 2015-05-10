package goreutils

// Yes creates a channel which outputs a string indefinitely.
func Yes(str string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			ch <- str
		}
	}()

	return ch
}
