# Garage Management System

Welcome to the Garage Management System, a digital hub for organizing and managing cars in a garage. This system, built using GoLang, enables efficient tracking and management of car records, inspired by the bustling atmosphere of a vibrant automotive workshop.

.

### Digital Representation

The Garage Management System serves as the virtual counterpart to a real garage. Every car that finds its way into this haven is meticulously logged into this system. Each vehicle's make, model, license plate, owner's name, service history, and more are captured here.

### Functionality

The system boasts a suite of endpoints that mimic the activities within the garage:

- **Retrieve Cars**: Fetches the entire collection of cars stored in the garage.
- **Retrieve Car by ID**: Provides details of a specific car based on its unique identifier.
- **Add New Car**: Incorporates a new car into the collection, capturing its specifications and details.
- **Update Car Details**: Allows for modifications to existing car records, such as changing ownership, updating service history, or marking its status.
- **Delete Car Record**: Removes a car from the system when it's sold, retired, or no longer part of the collection.

## Installation

1. Clone this repository:

    bash
    git clone https://github.com/Mudita-Srivastava/API_GOLANG.git
    

2. Navigate to the project directory:

    bash
    cd API_GOLANG
    

3. Install dependencies:

    bash
    go mod tidy
    

## Usage

- Run the application:

    bash
    go run main.go
    

- Run tests:

    bash
    go test

### Database Schema-
	carsId       string 
	make         string 
	model        string 
	licenceplate string 
	ownername    string 
	date         string 
	status       string 


---


