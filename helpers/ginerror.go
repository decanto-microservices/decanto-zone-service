package helpers

import "github.com/gin-gonic/gin"

func CheckForError(c *gin.Context, err error) {
	// Personalizable by using the gin Context to take some more actions other than panic
	if err != nil {
		panic(err)
	}
}
