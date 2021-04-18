package entity

type Person struct {
	Age       uint8  `json:"age" validate:"gte=1,lte=150"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type Video struct {
	Author      Person `json:"author" validate:"required"`
	Description string `json:"description" validate:"max=140"`
	Title       string `json:"title" validate:"min=2,max=10,containsProfanity"`
	Url         string `json:"url" validate:"required,url,contains=watch?v="`
}
