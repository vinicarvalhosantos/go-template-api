package addressRoutes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/vinicius.csantos/go-template-api/internal/handler/address"
	constantUtils "gitlab.com/vinicius.csantos/go-template-api/internal/util/constant"
	"gitlab.com/vinicius.csantos/go-template-api/internal/util/jwt"
)

func SetupAddressRoutes(router fiber.Router) {
	addressRoute := router.Group("/address")

	//Get All Addresses
	addressRoute.Get("/", jwt.Protected(), address.GetAddresses)
	//Get Address By ID
	addressRoute.Get(constantUtils.PathAddressIdParam, jwt.Protected(), address.GetAddressById)
	//Get All User Addresses
	addressRoute.Get(constantUtils.PathUserIdParam+"/user", jwt.Protected(), address.GetUserAddressesById)
	//Create a New Address
	addressRoute.Post("/", jwt.Protected(), address.RegisterAddress)
	//Update an Address
	addressRoute.Put(constantUtils.PathAddressIdParam, jwt.Protected(), address.UpdateAddress)
	//Update the User Main Address
	addressRoute.Patch(constantUtils.PathUpdateUserMainAddress, jwt.Protected(), address.UpdateUserMainAddress)
	//Delete an Address
	addressRoute.Delete(constantUtils.PathAddressIdParam, jwt.Protected(), address.DeleteAddress)
}
