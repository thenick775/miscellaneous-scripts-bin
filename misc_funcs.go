//adds spaces to front of string based on length to create proper fixed width string
func FixedLengthString(length int, str string) string {
	return fmt.Sprintf(fmt.Sprintf("%%%d.%ds", length, length), str)
}

//clamps integer number to provided range
func clamp(n, min, max int) int {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}

//dummy example illustrating quickly how to sort an arbitrary structure by date
type mytyp []myStructure //contains Date field

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func (a mytyp) Less(i, j int) bool {
	p, err := time.Parse("01/02/2006 15:04", a[i].Date)
	chkerr(err)
	q, err := time.Parse("01/02/2006 15:04", a[j].Date)
	chkerr(err)
	return p.Before(q)
}

func (a mytyp) Len() int { return len(a) }

func (a mytyp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//example of removing duplicates based on date
func removeDuplicateValues(d []mytyp) []lineData {
	keys := make(map[string]bool)
	list := []lineData{}

	//If the key(values of the slice) is not equal
	//to the already present value in new slice (list)
	//then we append it, else we move to the next element
	for _, entry := range d {
		if _, value := keys[entry.date]; !value {
			keys[entry.date] = true
			list = append(list, entry)
		} else {
			fmt.Println("found duplicate\n", entry)
		}
	}
	return list
}
