package never_use_this

//Read a file line by line
func readLine(path string) {
	inFile, ok := os.Open(path)
	if ok != nil {
		panic(fmt.Sprintf("open failed: %s", ok))
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	//scanner.Split(bufio.ScanLines) // This is not needed. ScanLine is the default split function

	for scanner.Scan() {
		fmt.Println(scanner.Text()) // The newline isn't part of the returned Text
	}
}

//split a string by whitespaces
func splitSting(in string) []fields {
	return strings.Fields(in)
}

//slice operations - https://github.com/golang/go/wiki/SliceTricks
func sliceOps() {
	//push at tail
	sl = append(sl, x)

	//pop at tail
	x, sl = sl[len(sl)-1], sl[:len(sl)-1]
}

//sort interface
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//map operations
func mapOps() {
	// simply add to a map. if new key, its created. if existing key, its updated.
	m[key] = value

	// check for a key and do some work if key exists (!ok does if key doesn't exist)
	// ok is a bool.
	if val, ok := m[key]; ok {
		//do something here
	}

	// delete a key
	delete(m, key)
}
