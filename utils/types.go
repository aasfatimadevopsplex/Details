package utils

type book struct {
	ISBN   string  `json:"ISBN" bson:"ISBN"`
	Title  string  `json:"Title" bson:"Title"`
	Author string  `json:"Author" bson:"Author"`
	Price  float64 `json:"Price" bson:"Price"`
	Rating  float64 `json:"Rating" bson:"Rating"`

}
type books []book

type Response struct {
	Books []book `json:"book"`
}

