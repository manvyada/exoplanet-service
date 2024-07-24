# exoplanet-service
The Exoplanet Service is a simple RESTful microservice designed to manage and interact with data related to exoplanets. This service allows you to add, list, retrieve, update, and delete exoplanet data. Additionally, it can estimate the fuel cost required to travel to a specific exoplanet based on crew capacity.

Features

	•	Add Exoplanet: POST request to add a new exoplanet.
	•	List Exoplanets: GET request to retrieve a list of all exoplanets.
	•	Get Exoplanet by ID: GET request to retrieve a specific exoplanet by its ID.
	•	Update Exoplanet: PUT request to update an existing exoplanet by its ID.
	•	Delete Exoplanet: DELETE request to remove an exoplanet by its ID.
	•	Estimate Fuel: GET request to estimate the fuel cost for travel based on crew capacity.

Installation

Prerequisites

	•	Go 1.18 or higher
	•	Docker (optional, for containerization)

Clone the Repository
git clone https://github.com/manvyada/exoplanet-service.git
cd exoplanet-service

Build and Run the Service

Running Locally

	1.	Build the application:
      go build -o exoplanet-service
  2.  Run the application:
      ./exoplanet-service
         or
       cd to the project directory
       go run -v main.go

Docker

	1.	Build the Docker image:
      docker build -t exoplanet-service .
  2.  Run the Docker container:
      docker run -p 8080:8080 exoplanet-service

      The service will be accessible at http://localhost:8080/exoplanets

  





