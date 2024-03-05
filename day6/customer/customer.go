package customer

type MedBooks struct {
	Annotation bool
	Diagnose   string
}

type Organs struct {
	Leg    int
	Hands  int
	Livers int
	Heart  bool
}

type Patients struct {
	Name    string
	Age     int
	Order   int
	Organs  Organs
	MedBook MedBooks
}
