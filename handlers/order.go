package handlers

import (

	// "github.com/SoNim-LSCM/maxbot_oms/models"

	"github.com/gofiber/fiber/v2"
)

const AW2_RESPONSE string = `{
	"messageCode": "LOCATION_UPDATE",
	"userId": 1,
	"dutyLocationId" : 1,
	"dutyLocationName" : "5/F DSC"
}`

// @Summary		Test AW2 websocket response.
// @Description	Get the response of AW2 (Server notify the user which location selected).
// @Tags			Test
// @Param parameters body string false "AW2 response"
// @Produce		plain
// @Success		200	"OK"
// @Router			/testAW2 [post]
func HandleAddOrder(c *fiber.Ctx) error {

	// mqtt.PublishMqtt("direct/publish", []byte("packet scheduled message"))
	// var response wsTest.ReportDutyLocationUpdateResponse
	// if err := c.BodyParser(&response); errorHandler.CheckError(err, "Invalid/missing input: ") {
	// 	err := json.Unmarshal([]byte(AW2_RESPONSE), &response)
	// 	errorHandler.CheckError(err, "translate string to json in wsTest")
	// }
	// if err := websocket.SendBoardcastMessage(response); err != nil {
	// 	return c.SendString(err.Error())
	// }
	return c.SendString("OK")
}
