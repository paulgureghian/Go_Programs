package inventory



type Inventory struct {
	Name      string
	Date      string
	Cake      int
	Chocolate int
	Chips     int
}

f := excelize.Newfile()