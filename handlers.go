package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createEmployeeHandler(dbAdapter *DatabaseAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newEmp Employee
		if err := c.BindJSON(&newEmp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := dbAdapter.CreateEmployee(newEmp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Employee added"})
	}
}

func getEmployeeHandler(dbAdapter *DatabaseAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
			return
		}
		emp, err := dbAdapter.GetEmployee(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		fmt.Println("Employee:", emp)
		c.JSON(http.StatusOK, emp)
	}
}

func updateEmployeeHandler(dbAdapter *DatabaseAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp Employee
		if err := c.BindJSON(&emp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := dbAdapter.UpdateEmployee(emp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Employee updated"})
	}
}

func deleteEmployeeHandler(dbAdapter *DatabaseAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
			return
		}
		if err := dbAdapter.DeleteEmployee(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
	}
}
