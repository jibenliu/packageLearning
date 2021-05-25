package schema

type Person struct {
	Id   string `bson:"_id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}


