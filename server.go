import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// This ensures that log messages will include file and line number.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	router := gin.Default()
	router.GET("/patients", getPatients)
	router.GET("/patients/:id", getPatientByID)
	router.POST("/patients", postPatients)
	router.Run(":8080")
}

// Example of adding logging to the getPatients function
func getPatients(c *gin.Context) {
	log.Println("Fetching all patients")
	c.IndentedJSON(http.StatusOK, patients)
	log.Println("Fetched all patients successfully")
}

// Example of adding detailed logging to the postPatients function
func postPatients(c *gin.Context) {
	var newPatient patient

	if err := c.BindJSON(&newPatient); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patients = append(patients, newPatient)
	log.Printf("Added new patient: %v\n", newPatient)
	c.IndentedJSON(http.StatusCreated, newPatient)
}

// Similarly, you can add logging to the getPatientByID function to track specific requests
func getPatientByID(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Fetching patient with ID: %s\n", id)

	for _, patient := range patients {
		if patient.ID == id {
			log.Printf("Found patient: %v\n", patient)
			c.IndentedJSON(http.StatusOK, patient)
			return
		}
	}

	log.Printf("Patient with ID %s not found\n", id)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "patient not found"})
}
