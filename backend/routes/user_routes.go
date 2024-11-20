package routes

import (
	"fiber-api/controllers"

	"github.com/gofiber/fiber/v2"
)

// UserRoute sets up routes related to user operations.
func UserRoute(app *fiber.App) {

	// Sample testing route
	app.Get("/sample", controllers.Sample)

	// Register a new user
	app.Post("/register", controllers.RegisterUser)

	// Fetch all users
	app.Get("/getall", controllers.GetAllUsers)

	// Fetch the pizza menu
	app.Get("/getmenu", controllers.GetMenu)

	// Add an item to the cart
	app.Post("/addtocart", controllers.InsertCart)

	// Fetch cart data
	app.Get("/retrivetocart", controllers.RetriveToCart)

	// User login
	app.Post("/login", controllers.Login)

	// User logout
	app.Get("/logout", controllers.LogoutUser)

	// Get ingredients
	app.Get("/getingredients", controllers.GetIngredients)

	// Delete from cart
	app.Post("/deletefromcart", controllers.DeleteFromCart)

	// Build a custom pizza
	app.Post("/build", controllers.BuildPizza)
}
