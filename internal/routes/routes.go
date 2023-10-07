package routes

import (
	"ecommercestore/internal/controllers"
	"ecommercestore/internal/middleware"
	"ecommercestore/internal/models"
	"github.com/gin-gonic/gin"
)

func SetRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.RedirectTrailingSlash = true

	api := router.Group("/api")

	api.Use(middleware.CustomErrors, middleware.Recover)

	api.POST("/signup", gin.Bind(models.User{}), controllers.SignUp())       // user
	api.POST("/signin", gin.Bind(models.UserSignIn{}), controllers.SignIn()) // user

	// user
	user := api.Group("/user", middleware.Authorization())
	{
		user.GET("/:id", controllers.GetUserById())                                              // admin
		user.GET("/", controllers.GetListUsers())                                                // admin
		user.GET("/profile", controllers.GetProfileUser())                                       // user
		user.PUT("/info", gin.Bind(models.InfoUser{}), controllers.UpdateInfoUser())             // user
		user.PUT("/password", gin.Bind(models.UserPassword{}), controllers.UpdatePasswordUser()) // user
	}

	// discount
	discount := api.Group("/discount", middleware.Authorization())
	{
		discount.POST("/", gin.Bind(models.Discount{}), controllers.CreateDiscount())         // admin
		discount.PUT("/:id", gin.Bind(models.DiscountUpdate{}), controllers.UpdateDiscount()) // admin
		discount.GET("/:id", controllers.GetDiscount())                                       // admin
		discount.GET("/", controllers.GetListDiscounts())                                     // admin
		discount.DELETE("/:id", controllers.DeleteDiscount())                                 // admin
	}

	// category
	category := api.Group("/category")
	{
		category.GET("/:id", controllers.GetCategory())                                                                   // admin or user
		category.GET("/", controllers.GetListCategories())                                                                // admin or user
		category.POST("/", middleware.Authorization(), gin.Bind(models.Category{}), controllers.CreateCategory())         // admin
		category.PUT("/:id", middleware.Authorization(), gin.Bind(models.CategoryUpdate{}), controllers.UpdateCategory()) // admin
		category.DELETE("/:id", middleware.Authorization(), controllers.DeleteCategory())                                 // admin
	}

	// inventory
	inventory := api.Group("/inventory", middleware.Authorization())
	{
		inventory.POST("/", gin.Bind(models.Inventory{}), controllers.CreateInventory())     // admin
		inventory.GET("/:id", controllers.GetInventory())                                    // admin
		inventory.GET("/", controllers.GetListInventories())                                 // admin
		inventory.PATCH("/:id", gin.Bind(models.Inventory{}), controllers.UpdateInventory()) // admin
		inventory.DELETE("/:id", controllers.DeleteInventory())                              //admin
	}

	// product
	product := api.Group("/product")
	{
		product.POST("/", middleware.Authorization(), gin.Bind(models.Product{}), controllers.CreateProduct())           // admin
		product.PATCH("/:id", middleware.Authorization(), gin.Bind(models.ProductUpdate{}), controllers.UpdateProduct()) // admin
		product.GET("/:id", controllers.GetProduct())                                                                    // user
		product.GET("/", controllers.GetListProducts())                                                                  // user
		product.DELETE("/:id", middleware.Authorization(), controllers.DeleteProduct())                                  // admin

		product.POST("/like/:id", middleware.Authorization(), controllers.CreateActionLike())         // user
		product.DELETE("/dislike/:id", middleware.Authorization(), controllers.CreateActionDislike()) // user
		product.GET("/liked-users/:id", middleware.Authorization(), controllers.GetListUsersLike())   // admin or user
	}

	// user address
	address := api.Group("/address", middleware.Authorization())
	{
		address.POST("/", gin.Bind(models.UserAddress{}), controllers.CreateUserAddress())      // user
		address.GET("/", controllers.GetUserAddress())                                          // user
		address.PUT("/", gin.Bind(models.UserAddressUpdate{}), controllers.UpdateUserAddress()) // user
	}

	// cart
	cart := api.Group("/cart", middleware.Authorization())
	{
		cart.POST("/", controllers.AddItemToCart())           // user
		cart.DELETE("/:id", controllers.RemoveItemFromCart()) // user
		cart.GET("/", controllers.GetListItemsFromCart())     // user
	}

	// order
	order := api.Group("/order", middleware.Authorization())
	{
		order.POST("/", controllers.CreateOrder())                                    // user
		order.GET("/:orderId", controllers.GetOrder())                                // user
		order.GET("/user", controllers.GetListOrdersByUser())                         // user
		order.GET("/total", controllers.GetTotalOrders())                             // admin
		order.PATCH("/:orderId", gin.Bind(models.Order{}), controllers.UpdateOrder()) // user
		order.DELETE("/:orderId", controllers.DeleteOrder())                          // user

		// order item
		order.POST("/:orderId/item", gin.Bind(models.OrderItem{}), controllers.AddItemToOrder())                 // user
		order.GET("/:orderId/item/:itemId", controllers.GetOrderItem())                                          // user
		order.GET("/:orderId/items", controllers.GetListOrderItems())                                            // user
		order.PATCH("/:orderId/item/:itemId", gin.Bind(models.OrderItemUpdate{}), controllers.UpdateOrderItem()) // user
		order.DELETE("/:orderId/item/:itemId", controllers.RemoveItemFromOrder())                                // user
	}

	return router
}
