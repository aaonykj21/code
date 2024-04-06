package models

type Bread struct {
	Bread_id      int    `json:"bread_id"`
	Bread_nameTH  string `json:"bread_nameTH"`
	Bread_nameENG string `json:"bread_nameENG"`
	Bread_price   int    `json:"bread_price"`
	Bread_stock   int    `json:"bread_stock"`
}

type Veg struct {
	Veg_id      int    `json:"veg_id"`
	Veg_nameTH  string `json:"veg_nameTH"`
	Veg_nameENG string `json:"veg_nameENG"`
	Veg_price   int    `json:"veg_price"`
	Veg_stock   int    `json:"veg_stock"`
}

type Meat struct {
	Meat_id      int    `json:"meat_id"`
	Meat_nameTH  string `json:"meat_nameTH"`
	Meat_nameENG string `json:"meat_nameENG"`
	Meat_price   int    `json:"meat_price"`
	Meat_stock   int    `json:"meat_stock"`
}

type Sauce struct {
	Sauce_id      int    `json:"sauce_id"`
	Sauce_nameTH  string `json:"sauce_nameTH"`
	Sauce_nameENG string `json:"sauce_nameENG"`
	Sauce_price   int    `json:"sauce_price"`
	Sauce_stock   int    `json:"sauce_stock"`
}

type Topping struct {
	Topping_id      int    `json:"topping_id"`
	Topping_nameTH  string `json:"topping_nameTH"`
	Topping_nameENG string `json:"topping_nameENG"`
	Topping_price   int    `json:"topping_price"`
	Topping_stock   int    `json:"topping_stock"`
}

type Order_detailENG struct {
	Order_id        int      `json:"order_id"`
	Bread_nameENG   string   `json:"bread_nameENG"`
	Meat_nameENG    []string `json:"meat_nameENG"`
	Veg_nameENG     []string `json:"veg_nameENG"`
	Sauce_nameENG   []string `json:"sauce_nameEng"`
	Topping_nameENG []string `json:"topping_nameEng"`
	Sum_price       int      `json:"price"`
}

type Order_detailTH struct {
	Order_id       int      `json:"order_id"`
	Bread_nameTH   string   `json:"bread_nameTH"`
	Meat_nameTH    []string `json:"meat_nameTH"`
	Veg_nameTH     []string `json:"veg_nameTH"`
	Sauce_nameTH   []string `json:"sauce_nameTH"`
	Topping_nameTH []string   `json:"topping_nameTH"`
	Sum_price      int      `json:"price"`
}
