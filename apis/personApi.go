package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "goweb/models"
	. "goweb/utils"
	"net/http"
	"strconv"
)

// GET /
func IndexApi(c *gin.Context) {
	//value, _ := c.Get("example")
	//fmt.Println(value)
	c.String(http.StatusOK, "It works")
}

// POST /person
func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name") //url写first_name=a&last_name=b
	lastName := c.Request.FormValue("last_name")

	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	CheckErr(err)
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// GET /persons
func GetPersonsApi(c *gin.Context) {
	var p Person
	persons, err := p.GetPersons()
	CheckErr(err)

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})

}

// GET /person/:id
func GetPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	CheckErr(err)
	p := Person{Id: id}
	person, err := p.GetPerson()
	CheckErr(err)

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})

}

// PUT /person/:id
func ModPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	CheckErr(err)
	p := Person{Id: id}
	err = c.Bind(&p)
	CheckErr(err)
	ra, err := p.ModPerson()
	CheckErr(err)
	msg := fmt.Sprintf("Update person %d successful %d", p.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// DELETE /person/:id
func DelPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	CheckErr(err)
	p := Person{Id: id}
	ra, err := p.DelPerson()
	CheckErr(err)
	msg := fmt.Sprintf("Delete person %d successful %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
