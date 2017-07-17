package genmodel

type Address struct {
	Street  string
	Zip     string
	City    string
	Country string
}

type Application struct {
	Profile     Profile
	JobPosition JobPosition
}
