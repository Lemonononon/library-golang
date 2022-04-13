package model_book

type Book struct {
	ID       int     `gorm:"column:id;primaryKey"`
	BookID   string  `gorm:"column:book_id"`
	Category string  `gorm:"column:category;size:20"`
	BookName string  `gorm:"column:book_name;size:40"`
	Press    string  `gorm:"column:press;size:40"`
	Year     int     `gorm:"column:year"`
	Author   string  `gorm:"column:author;size:20"`
	Price    float32 `gorm:"column:price;"` //TODO
	Total    int     `gorm:"column:total"`  //总藏书量
	Stock    int     `gorm:"column:stock"`  //目前库存
}

type QueryBookReq struct {
	BookName   string `json:"book_name"`
	Category   string `json:"category"`
	Press      string `json:"press"`
	LowerYear  int    `json:"lower_year"`
	UpperYear  int    `json:"upper_year"`
	Author     string `json:"author"`
	LowerPrice int    `json:"lower_price"`
	UpperPrice int    `json:"upper_price"`
}

type QueryBookResp struct {
	Books []Book `json:"books"`
}
