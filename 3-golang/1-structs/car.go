package structs

type Brand string

const (
	BRAND_BMW    Brand = "BMW"
	BRAND_TOYOTA Brand = "TOYOTA"
)

type Model string

type Car interface {
	Brand() Brand
	Model() Model
	ChangeBrand(Brand)
	ChangeModel(Model)
}

type Owner struct {
	Name string
}
