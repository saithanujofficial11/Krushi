# Krushi

## Overview of the Service
Krushi is a service dedicated to empowering farmers by providing them with the tools and information needed to optimize their agricultural practices, improve crop yields, and enhance overall sustainability.

## Technical Implementation
The system currently includes a backend management service built with **Go**.

### Backend Server
- **Language**: Go
- **Architecture**: Layered architecture (Handler -> Repository -> Model)
- **Storage**: Persistent JSON storage (`data/farmers.json`)
- **API Endpoints**:
    - `POST /farmers`: Register a new farmer with their contact information.
    - `GET /farmers/{farmerId}`: Retrieve specific farmer details using their unique ID.
