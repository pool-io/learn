package structs

type Brand string

type Model string

type Car interface {
	Brand() Brand
	Model() Model
	ChangeBrand()
}
