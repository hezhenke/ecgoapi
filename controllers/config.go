package controllers

import (
	"ecshopGoApi/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllConfig(c *gin.Context) {
	data := "MJWe1D7HNo8d7bQEkCbsHwMD03yD4J6999\\/+gKmNfzPkP3d+cPs6qFoFLhIfU\\/pBJc\\/+ZGtq580248ZDdaEuBpmRAno4CuI9kHmYITPVNQ6\\/0DZDHegK66eKZOrwdEI53LJ058+SCwTdPHRa8mo846S5LH4sSnwBcRs\\/9yQIH2SpRVT\\/9PVznW1t7tsNiIv2L1zZ40iFfkmgy57diCNZoCUGEPfAFB6Y0C8xDsbZN21TLk29oKSyCpB\\/m4yXK2l\\/DUHz7iI11+zCWOKL\\/4akfAQ+woXEgQy\\/SDSD\\/oun+3PhJ+Uatz6ISuP+KEAmQsWtx\\/31G6gBpFcrbbvomqfgGAGwX\\/YTKwpIUy8zwkOIAXUlR+d9CM2FfI2+lI7Vh7I686rUZyvXphhyCp7haSxCNHOpQ\\/Zzz1pnTIrIVSJOf8Eiovd4ZwrqjJXzt\\/o3jxfgUy7KaUh1IYe8R5mbNvLHy5JP0gJmvtJWGgyDDyebK3TTA7SzyDYNQoaHpSxajtl\\/CPN1yyKw+JTTvCLb80qU\\/SDVLU6xEjba5Q6FDaZjZfsf\\/WeXPdxDK6a2poO2Zs2EN6soC53Vrl54yhCiVWA\\/byuwcEAeM+U29RFGhyfLJHCFkJh4mev0ndAjhPiM2dVbVXoyghSLT9+xeCW7dx9SmmZxMReWz9asQ7RcnAnz3DlRMC6iZE6DA4vOZqRQ922JweAV8b7PT7TK11xci2hcXSWl3zmiF91iRkv\\/Jy86T+G0WVQXivXRApWGOhiqKjHjtUobGmdxVgVbwgLTfd6sSrT0suf3TDZ9krfQRQEvpOZsltsUTyyUUWqR1CAvIV8gQ\\/Qbe2i6C68mYNBHZ2bTd8uIQ\\/BNLjs7Mfo2k2K\\/2hQzT68Axa2uBA2TqIt+uTq3kpg+JWSspYd8NN6a3fx1sgeDKrmyvqwxttkv5MeRLTrL77GXDfKBj68OlDIrlAtJRma8SasfWTjrErysgHiGc13q2T2GEh\\/8q6RN1dOjZ7wVjR+BYc3tJ8emu0ZNG554OkD\\/XN2SJpC7l831DO+VIr\\/Uaci5DnW1hkqJmy4U9iEU00p0YKK5lxX9rtTvFCGK1eOVFC8+7qKuNwHzCT4oP+G6gJecHwhKWQAkdULvfSezi0CVD+2bTIfKC7PV7tGbR91vFI5JRniJ5xioF\\/8Ut63iq5sVFwDeqWWYUvovEoG5JLkQckPr+uc8hTI="
	c.JSON(http.StatusOK,dtos.FormatBody(gin.H{"data":data}))
}