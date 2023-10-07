package testRepo

import (
	"ecommercestore/internal/database"
	"ecommercestore/internal/models"
	"ecommercestore/internal/repo"
	"github.com/gin-gonic/gin"
)

func testSetup() {
	gin.SetMode(gin.TestMode)
	database.ResetTestDatabase()
}

// user
func createMockUser() (*models.User, error) {
	user := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Phone:     "1234567890",
		Password:  "password",
	}
	err := repo.CreateUser(user)
	return user, err
}

func createMockUser2() (*models.User, error) {
	user := &models.User{
		FirstName: "Alice",
		LastName:  "Johnson",
		Email:     "alicejohnson@example.com",
		Phone:     "5555555555",
		Password:  "password",
	}
	err := repo.CreateUser(user)
	return user, err
}

// user address
func createMockUserAddress() (*models.UserAddress, error) {
	user, _ := createMockUser()
	userId := user.Id
	userAddress := &models.UserAddress{
		UserId:     userId,
		Address:    "123 Main Street",
		City:       "New York",
		Country:    "USA",
		PostalCode: "10001",
		Phone:      "1234567890",
	}

	err := repo.CreateUserAddress(userAddress)
	return userAddress, err
}

// discount
func createMockDiscount() (*models.Discount, error) {
	active := new(bool)
	*active = true
	discount := &models.Discount{
		Name:    "Summer Sale",
		Desc:    "Get 20% off on all summer products",
		Percent: 20.0,
		Active:  active,
	}

	err := repo.CreateDiscount(discount)
	return discount, err
}

func createMockDiscount2() (*models.Discount, error) {
	active := new(bool)
	*active = false
	discount := &models.Discount{
		Name:    "Back-to-School Sale",
		Desc:    "Save 15% on back-to-school essentials",
		Percent: 15.0,
		Active:  active,
	}

	err := repo.CreateDiscount(discount)
	return discount, err
}

// category
func createMockCategory() (*models.Category, error) {
	category := &models.Category{
		Name: "Men's Clothing",
		Desc: "Category for men's clothing.",
	}
	err := repo.CreateCategory(category)
	return category, err
}

func createMockCategory2() (*models.Category, error) {
	category := &models.Category{
		Name: "Men's Clothing2",
		Desc: "Category for men's clothing2.",
	}
	err := repo.CreateCategory(category)
	return category, err
}

// inventory
func createMockInventory() (*models.Inventory, error) {
	inventory := &models.Inventory{
		Quantity: 200,
	}
	err := repo.CreateInventory(inventory)
	return inventory, err
}

func createMockInventory2() (*models.Inventory, error) {
	inventory := &models.Inventory{
		Quantity: 100,
	}
	err := repo.CreateInventory(inventory)
	return inventory, err
}

// product
func createMockProduct() (*models.Product, error) {
	discount, _ := createMockDiscount()
	category, _ := createMockCategory()
	inventory, _ := createMockInventory()
	product := &models.Product{
		Name:  "Product 1",
		Code:  "P0001",
		Color: "Red",
		Size:  10,
		Desc:  "This is the description of product 1.",
		Price: 19.99,
	}
	product.DiscountId = discount.Id
	product.InventoryId = inventory.Id
	product.CategoryId = category.Id

	err := repo.CreateProduct(product)
	return product, err
}

func createMockProduct2() (*models.Product, error) {
	discount, _ := createMockDiscount2()
	category, _ := createMockCategory2()
	inventory, _ := createMockInventory2()
	product := &models.Product{
		Name:  "Product 2",
		Code:  "P0002",
		Color: "Blue",
		Size:  8,
		Desc:  "Description of product 2 goes here.",
		Price: 24.99,
	}
	product.DiscountId = discount.Id
	product.InventoryId = inventory.Id
	product.CategoryId = category.Id

	err := repo.CreateProduct(product)
	return product, err
}

// cart
func createMockCart() (*models.Cart, error) {
	user, _ := createMockUser()
	product, _ := createMockProduct()
	cart := &models.Cart{
		Quantity: 2,
	}
	cart.UserId = user.Id
	cart.ProductId = product.Id

	err := repo.AddItemToCart(cart)
	return cart, err
}

func createMockListItemsOfCart() (int, error) {
	user, _ := createMockUser()
	product, _ := createMockProduct()
	product2, _ := createMockProduct2()
	cart1 := &models.Cart{
		Quantity: 2,
	}
	cart1.UserId = user.Id
	cart1.ProductId = product.Id
	err := repo.AddItemToCart(cart1)
	if err != nil {
		return 0, err
	}

	cart2 := &models.Cart{
		Quantity: 1,
	}
	cart2.UserId = user.Id
	cart2.ProductId = product2.Id
	err = repo.AddItemToCart(cart2)
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

// product like
func createMockProductLike() (*models.ProductLike, error) {
	user, _ := createMockUser()
	product, _ := createMockProduct()
	proLike := &models.ProductLike{}
	proLike.ProductId = product.Id
	proLike.UserId = user.Id
	err := repo.CreateActionLike(proLike)
	return proLike, err
}

func createMockListUserLikeProduct() (int, error) {
	user1, _ := createMockUser()
	user2, _ := createMockUser2()
	product, _ := createMockProduct()

	proLike1 := &models.ProductLike{}
	proLike1.ProductId = product.Id
	proLike1.UserId = user1.Id
	err := repo.CreateActionLike(proLike1)
	if err != nil {
		return 0, err
	}

	proLike2 := &models.ProductLike{}
	proLike2.ProductId = product.Id
	proLike2.UserId = user2.Id
	err = repo.CreateActionLike(proLike2)
	if err != nil {
		return 0, err
	}

	return product.Id, nil
}

// order
func createMockOrder() (*models.Order, error) {
	user, err := createMockUser()
	order := &models.Order{
		TotalPrice: 0,
	}
	order.UserId = user.Id
	_, err = repo.CreateOrder(order)

	return order, err
}

func createMockListOrdersByUser() (int, error) {
	user, err := createMockUser()
	order1 := &models.Order{
		TotalPrice: 0,
	}
	order1.UserId = user.Id
	_, err = repo.CreateOrder(order1)
	if err != nil {
		return 0, err
	}

	order2 := &models.Order{
		TotalPrice: 0,
	}
	order2.UserId = user.Id
	_, err = repo.CreateOrder(order2)
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func createMockListOrders() error {
	user1, err := createMockUser()
	order1 := &models.Order{
		TotalPrice: 0,
	}
	order1.UserId = user1.Id
	_, err = repo.CreateOrder(order1)
	if err != nil {
		return err
	}

	user2, err := createMockUser()
	order2 := &models.Order{
		TotalPrice: 0,
	}
	order2.UserId = user2.Id
	_, err = repo.CreateOrder(order2)
	if err != nil {
		return err
	}

	return nil
}

// order item
func createMockOrderItem() (*models.OrderItem, *models.Order, error) {
	order, _ := createMockOrder()
	product, _ := createMockProduct()

	orderItem := &models.OrderItem{
		Price:    10,
		Quantity: 1,
	}
	orderItem.OrderId = order.Id
	orderItem.ProductId = product.Id
	order, err := repo.AddItemToOrder(order.UserId, orderItem)

	return orderItem, order, err
}

func createMockListOrderItems() (int, error) {
	order, _ := createMockOrder()

	product1, _ := createMockProduct()

	orderItem1 := &models.OrderItem{
		Price:    10,
		Quantity: 1,
	}
	orderItem1.OrderId = order.Id
	orderItem1.ProductId = product1.Id
	_, err := repo.AddItemToOrder(order.UserId, orderItem1)
	if err != nil {
		return 0, err
	}

	product2, _ := createMockProduct2()

	orderItem2 := &models.OrderItem{
		Price:    20,
		Quantity: 2,
	}
	orderItem2.OrderId = order.Id
	orderItem2.ProductId = product2.Id
	_, err = repo.AddItemToOrder(order.UserId, orderItem2)
	if err != nil {
		return 0, err
	}

	return orderItem1.OrderId, nil
}
