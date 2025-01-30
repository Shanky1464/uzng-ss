package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

// AuthToken is the token used for authorization.
const AuthToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkxYejBQYWVFRERuS3dfcjdkU2s2UyJ9.eyJodHRwczovL2FwaS51bmlwaG9yZS5jb20vdGVuYW50IjoiZmU1NzIxNTYtMjIxMi00NThlLTg3M2EtNmFjNjBiZWJmNjQyIiwiaHR0cHM6Ly9hcGkudW5pcGhvcmUuY29tL29yZ19pZCI6Im9yZ19TcjJkRHJXSTlvaFY5OFJxIiwiaXNzIjoiaHR0cHM6Ly9kZXYtZzJqNW11MnoudXMuYXV0aDAuY29tLyIsInN1YiI6ImtNbXAzcXNKTWN2NjNOZ0k1RXJyS0hkdlJCM0YxNjdmQGNsaWVudHMiLCJhdWQiOiJhcGkudW5pcGhvcmUuY29tIiwiaWF0IjoxNzM4MDAxOTQzLCJleHAiOjE3MzgwODgzNDMsInNjb3BlIjoicmVhZDp1c2VycyByZWFkOnV6bmctcHJvY2Vzc2luZy1tZXRhZGF0YSB3cml0ZTp1em5nLXByb2Nlc3NpbmctbWV0YWRhdGEiLCJndHkiOiJjbGllbnQtY3JlZGVudGlhbHMiLCJhenAiOiJrTW1wM3FzSk1jdjYzTmdJNUVycktIZHZSQjNGMTY3ZiIsInBlcm1pc3Npb25zIjpbInJlYWQ6dXNlcnMiLCJyZWFkOnV6bmctcHJvY2Vzc2luZy1tZXRhZGF0YSIsIndyaXRlOnV6bmctcHJvY2Vzc2luZy1tZXRhZGF0YSJdfQ.ZaQovC7KD9__Wa2Lw2O9G7nUCwP-1M7Zm10AK4DnKQdDG8ZjM3gX6-gPXUhCGQ3kpzAds6ryLfvhl1XcSoq5PP3_byNIM14sgdP4mH89UT0ExMlstRPxHzkV30Sf1xkwmTU55qMfTIE4qdBvt_nd-YYEalvXztq2PAWJAUxWTHqwL8veKPyM6FuCWzH0ZaPPmS7Pkvo994dMcTIT0nc1gFmMnlWGxtfMdnBSJlvuqHyJhCr7_0O3GkAOOPIN8N8xg2k3YD-Jj2PW5kFmsn28vWWxifPZJfCCu1OAvjiD_KbtnnZQ5BnSslug2WPZCmoQXsDOkXYTIGykx9mhBq4YYQ" // Replace with your actual token

// Response struct to return status code as JSON
type Response struct {
	StatusCode11  int
	StatusCode12  int
	StatusCode13  int
	StatusCode21  int
	StatusCode22  int
	StatusCode23  int
	StatusCode31  int
	StatusCode32  int
	StatusCode33  int
	StatusCode41  int
	StatusCode42  int
	StatusCode43  int
	StatusCode51  int
	StatusCode52  int
	StatusCode53  int
	StatusCode71  int
	StatusCode72  int
	StatusCode73  int
	StatusCode81  int
	StatusCode82  int
	StatusCode83  int
	StatusCode91  int
	StatusCode92  int
	StatusCode93  int
	StatusCode101 int
	StatusCode102 int
	StatusCode103 int
	StatusCode111 int
	StatusCode112 int
	StatusCode113 int
	StatusCode121 int
	StatusCode122 int
	StatusCode123 int
	Message       string
}


func callExternalAPI(apiURL string) (int, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Return the status code
	return resp.StatusCode, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Replace with your actual API URLs
	// Gateway APIs
	api11 := "https://api.us.cloud.uniphorestaging.com/uzng-gateway/health/startup"
	api12 := "https://api.us.cloud.uniphorestaging.com/uzng-gateway/health/readiness"
	api13 := "https://api.us.cloud.uniphorestaging.com/uzng-gateway/health/liveness"

	// Flow Manager APIs
	api21 := "https://api.us.cloud.uniphorestaging.com/uzng-flow-manager/health/startup"
	api22 := "https://api.us.cloud.uniphorestaging.com/uzng-flow-manager/health/readiness"
	api23 := "https://api.us.cloud.uniphorestaging.com/uzng-flow-manager/health/liveness"

	// Score Card Studio Backend APIs
	api31 := "https://api.us.cloud.uniphorestaging.com/score-card-studio-backend/health/startup"
	api32 := "https://api.us.cloud.uniphorestaging.com/score-card-studio-backend/health/readiness"
	api33 := "https://api.us.cloud.uniphorestaging.com/score-card-studio-backend/health/liveness"

	// Score Card Studio Runtime APIs
	api41 := "https://api.us.cloud.uniphorestaging.com/score-card-studio-runtime/health/startup"
	api42 := "https://api.us.cloud.uniphorestaging.com/score-card-studio-runtime/health/readiness"
	api43 := "https://api.us.cloud.uniphorestaging.com/score-card-studio-runtime/health/liveness"

	// Audio Processor APIs
	api51 := "https://api.us.cloud.uniphorestaging.com/uzng-audio-processor/health/startup"
	api52 := "https://api.us.cloud.uniphorestaging.com/uzng-audio-processor/health/readiness"
	api53 := "https://api.us.cloud.uniphorestaging.com/uzng-audio-processor/health/liveness"



	// Conversation Facts Backend APIs
	api71 := "https://api.us.cloud.uniphorestaging.com/uzng-conversation-facts-backend/health/startup"
	api72 := "https://api.us.cloud.uniphorestaging.com/uzng-conversation-facts-backend/health/readiness"
	api73 := "https://api.us.cloud.uniphorestaging.com/uzng-conversation-facts-backend/health/liveness"

	// Data Replicator APIs
	api81 := "https://api.us.cloud.uniphorestaging.com/uzng-data-replicator/health/startup"
	api82 := "https://api.us.cloud.uniphorestaging.com/uzng-data-replicator/health/readiness"
	api83 := "https://api.us.cloud.uniphorestaging.com/uzng-data-replicator/health/liveness"

	// Facts Resolver APIs
	api91 := "https://api.us.cloud.uniphorestaging.com/uzng-facts-resolver/health/startup"
	api92 := "https://api.us.cloud.uniphorestaging.com/uzng-facts-resolver/health/readiness"
	api93 := "https://api.us.cloud.uniphorestaging.com/uzng-facts-resolver/health/liveness"

	// Orchestrator APIs
	api101 := "https://api.us.cloud.uniphorestaging.com/uzng-orchestrator/health/startup"
	api102 := "https://api.us.cloud.uniphorestaging.com/uzng-orchestrator/health/readiness"
	api103 := "https://api.us.cloud.uniphorestaging.com/uzng-orchestrator/health/liveness"

	// Orchestrator Sweeper APIs
	api111 := "https://api.us.cloud.uniphorestaging.com/uzng-orchestrator-sweeper/health/startup"
	api112 := "https://api.us.cloud.uniphorestaging.com/uzng-orchestrator-sweeper/health/readiness"
	api113 := "https://api.us.cloud.uniphorestaging.com/uzng-orchestrator-sweeper/health/liveness"

	// Sentiments APIs
	api121 := "https://api.us.cloud.uniphorestaging.com/uzng-sentiments/health/startup"
	api122 := "https://api.us.cloud.uniphorestaging.com/uzng-sentiments/health/readiness"
	api123 := "https://api.us.cloud.uniphorestaging.com/uzng-sentiments/health/liveness"

	statusCode11, _ := callExternalAPI(api11)
	statusCode12, _ := callExternalAPI(api12)
	statusCode13, _ := callExternalAPI(api13)
	
	statusCode21, _ := callExternalAPI(api21)
	statusCode22, _ := callExternalAPI(api22)
	statusCode23, _ := callExternalAPI(api23)
	
	statusCode31, _ := callExternalAPI(api31)
	statusCode32, _ := callExternalAPI(api32)
	statusCode33, _ := callExternalAPI(api33)
	
	statusCode41, _ := callExternalAPI(api41)
	statusCode42, _ := callExternalAPI(api42)
	statusCode43, _ := callExternalAPI(api43)
	
	statusCode51, _ := callExternalAPI(api51)
	statusCode52, _ := callExternalAPI(api52)
	statusCode53, _ := callExternalAPI(api53)

	
	statusCode71, _ := callExternalAPI(api71)
	statusCode72, _ := callExternalAPI(api72)
	statusCode73, _ := callExternalAPI(api73)
	
	statusCode81, _ := callExternalAPI(api81)
	statusCode82, _ := callExternalAPI(api82)
	statusCode83, _ := callExternalAPI(api83)
	
	statusCode91, _ := callExternalAPI(api91)
	statusCode92, _ := callExternalAPI(api92)
	statusCode93, _ := callExternalAPI(api93)
	
	statusCode101, _ := callExternalAPI(api101)
	statusCode102, _ := callExternalAPI(api102)
	statusCode103, _ := callExternalAPI(api103)
	
	statusCode111, _ := callExternalAPI(api111)
	statusCode112, _ := callExternalAPI(api112)
	statusCode113, _ := callExternalAPI(api113)
	
	statusCode121, _ := callExternalAPI(api121)
	statusCode122, _ := callExternalAPI(api122)
	statusCode123, _ := callExternalAPI(api123)
	

	// If any error occurred during API calls
	// if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil || err8 != nil || err9 != nil || err10 != nil || err11 != nil || err12 != nil {
	// 	http.Error(w, fmt.Sprintf("Error: %v", err1), http.StatusInternalServerError)
	// 	return
	// }

	// Send the status code as a JSON response
	response := Response{
		StatusCode11:  statusCode11,
		StatusCode12:  statusCode12,
		StatusCode13:  statusCode13,
		StatusCode21:  statusCode21,
		StatusCode22:  statusCode22,
		StatusCode23:  statusCode23,
		StatusCode31:  statusCode31,
		StatusCode32:  statusCode32,
		StatusCode33:  statusCode33,
		StatusCode41:  statusCode41,
		StatusCode42:  statusCode42,
		StatusCode43:  statusCode43,
		StatusCode51:  statusCode51,
		StatusCode52:  statusCode52,
		StatusCode53:  statusCode53,
		StatusCode71:  statusCode71,
		StatusCode72:  statusCode72,
		StatusCode73:  statusCode73,
		StatusCode81:  statusCode81,
		StatusCode82:  statusCode82,
		StatusCode83:  statusCode83,
		StatusCode91:  statusCode91,
		StatusCode92:  statusCode92,
		StatusCode93:  statusCode93,
		StatusCode101: statusCode101,
		StatusCode102: statusCode102,
		StatusCode103: statusCode103,
		StatusCode111: statusCode111,
		StatusCode112: statusCode112,
		StatusCode113: statusCode113,
		StatusCode121: statusCode121,
		StatusCode122: statusCode122,
		StatusCode123: statusCode123,
		Message:       "API call completed successfully",
	}
	

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Create an HTTP server and enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"}, // Modify this with your frontend URL if needed
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})

	// Routes
	http.HandleFunc("/check-api", handler)

	// Start server with CORS handling
	fmt.Println("Server running on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler.Handler(http.DefaultServeMux)))
}
