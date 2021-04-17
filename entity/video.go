package entity

type Person struct {
	Age       uint8  `json:"age" binding:"gte=1,lte=150"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type Video struct {
	Author      Person `json:"author" binding:"required"`
	Description string `json:"description" binding:"max=140"`
	Title       string `json:"title" binding:"min=2,max=10"`
	Url         string `json:"url" binding:"required,url"`
}
