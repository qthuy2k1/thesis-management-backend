# Thesis Management Backend

## Overview

An API provides the features to manage the thesis graduation for IT students and lecturers

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Configuration](#configuration)

## Installation

### Requirements

- Docker (required)
- Kubernetes (optional, for advanced usage)
- Minikube (optional, for local development with Kubernetes)


### Steps
1. Clone this repository
    ```bash
    git clone https://github.com/qthuy2k1/thesis-management-backend.git
    cd thesis-management-backend
    ```

2. If you want to use Docker, run the following command
    ```bash
    docker compose up
    ```
3. If you want to use Kubernetes and Minikube, run the following command
    ```bash
    # Run the Minikube
    minikube start --namespace thesis-management-backend

    # Create the Kubernetes namespace
    kubectl create namespace thesis-management-backend

    # Apply all the kubernetes deployments and services
    kubectl apply -f kubernetes/

    # Expose the service port to the local port
    minikube tunnel
    ```

4. You can access to the Kubernetes dashboard by running the following command
    ```bash
    minikube dashboard
    ```
5. After all the pods are running, you can use the API right now


## Architecture
1. Server architecture
![server_architecture](https://github.com/user-attachments/assets/9b0d65b9-7297-4b35-9d68-f6e856fcc3b6)

2. Database diagram
1.1 Classroom service
![classroom_service](https://github.com/user-attachments/assets/d148593f-ea87-4743-ad82-80d3b1907267)

1.2 User service
![user_service](https://github.com/user-attachments/assets/43da94cd-edd6-4d13-af22-e5f5779b86aa)

1.3 Schedule service
![schedule_service](https://github.com/user-attachments/assets/fb161388-02bd-4e6d-af1a-b3586f3fd70e)

## Usage

For the examples, please refer to the [api-docs](https://github.com/qthuy2k1/thesis-management-backend/tree/master/api-docs) folder


## Configuration
The API requires the following configuration:
1. For authentication, you have to download the firebase admin sdk json file from Firebase Console in the Project Settings. Here is an example of configuration:
    ```bash
    {
        "type": "service_account",
        "project_id": "YOUR_PROJECT_ID",
        "private_key_id": "YOUR_PRIVATE_KEY_ID",
        "private_key": "YOUR_PRIVATE_KEY",
        "client_email": "YOUR_CLIENT_EMAIL",
        "client_id": "YOUR_CLIENT_ID",
        "auth_uri": "https://accounts.google.com/o/oauth2/auth",
        "token_uri": "https://oauth2.googleapis.com/token",
        "auth_provider_x509_cert_url": "YOUR_AUTH_PROVIDER_X509_CERT_URL",
        "client_x509_cert_url": "YOUR_CLIENT_X509_CERT_URL",
        "universe_domain": "googleapis.com"
    }
    ```
    

2. For uploading file, you have to download the credentials file from Google Cloud Console. Here is an example of configuration:

    ```bash
        {
            "installed": 
            {
                "client_id":"YOUR_CLIENT_ID",
                "project_id":"YOUR_PROJECT_ID",
                "auth_uri":"https://accounts.google.com/o/oauth2/auth",
                "token_uri":"https://oauth2.googleapis.com/token",
                "auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs",
                "client_secret":"YOUR_CLIENT_SECRET",
                "redirect_uris":["http://localhost"]
            }
        }
    ```
