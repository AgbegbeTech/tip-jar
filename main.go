package main

import (
	"github.com/gin-gonic/gin"
	"tipjar/contracts"
)

var tipJar *contracts.TipJarContract // Initialize a global variable for the TipJar contract

func main() {
	// Initialize your TipJar contract
	tipJar = contracts.NewTipJarContract("content_creator_public_key", 30, 0)

	// Create an HTTP server with Gin
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/totalTips", GetTotalTips)
		v1.POST("/deposit", Deposit)
		v1.POST("/distributeTips", DistributeTips)
		v1.POST("/setDuration", SetDuration)             // Endpoint for setting the duration
		v1.POST("/addContributor", AddContributor)       // Endpoint for adding a contributor
		v1.POST("/removeContributor", RemoveContributor) // Endpoint for removing a contributor
		v1.GET("/ipfs/:hash", GetIPFSContent)
	}

	r.Run(":8080")
}

// ...

func SetDuration(c *gin.Context) {
	newDuration := c.PostForm("newDuration")

	// Convert newDuration to an integer or handle validation as needed

	// Set the new duration in the TipJar contract
	// tipJar.SetDuration(newDuration)

	c.JSON(200, gin.H{"message": "Duration set successfully"})
}

func AddContributor(c *gin.Context) {
	contributorPublicKey := c.PostForm("contributorPublicKey")
	contributionAmount := c.PostForm("contributionAmount")

	// Convert contributionAmount to an integer or handle validation as needed

	// Add the contributor to the TipJar contract
	// tipJar.AddContributor(contributorPublicKey, contributionAmount)

	c.JSON(200, gin.H{"message": "Contributor added successfully"})
}

func RemoveContributor(c *gin.Context) {
	removeContributorKey := c.PostForm("removeContributorKey")

	// Remove the contributor from the TipJar contract
	// tipJar.RemoveContributor(removeContributorKey)

	c.JSON(200, gin.H{"message": "Contributor removed successfully"})
}
