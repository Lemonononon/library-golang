package model_book

type Book struct {
	BookID   int     `gorm:"column:book_id;primaryKey"`
	Category string  `gorm:"column:category;size:20"`
	BookName string  `gorm:"column:book_name;size:40"`
	Press    string  `gorm:"column:press;size:40"`
	Year     int     `gorm:"column:year"`
	Author   string  `gorm:"column:author;size:20"`
	Price    float32 `gorm:"column:price;"` //TODO
	Total    int     `gorm:"column:total"`  //总藏书量
	Stock    int     `gorm:"column:stock"`  //目前库存
}
