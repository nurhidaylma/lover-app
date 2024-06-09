# lover-app

This is a dating mobile app backend system designed to mimic the functionality of popular dating apps like Tinder and Bumble.

## Features

- User authentication (sign up & login)
- Swipe-based matching system
- Premium features:
  - Unlimited swipes
  - Verified profile label

## Technologies Used

- Go (Golang) for backend development
- MySQL for the database
- Redis for caching swipe data
- Gorilla Mux for HTTP routing
- JWT (JSON Web Tokens) for authentication

## Getting Started

To run the application locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/dating-app.git
   cd dating-app

2. Set up the PostgreSQL database:

    Create a database and configure the connection details in the database/database.go file.

3. Install dependencies:
   ```bash
   go mod tidy

4. Start the server:
   ```bash
   go run main.go


## API Endpoints

1. User Management
    POST /signup: Create a new user account.
    POST /login: Log in with existing user credentials.

2. Swiping
    POST /swipe: Record a swipe action (like or pass) on a dating profile.

3. Premium Features
    POST /premium: Upgrade to premium features (e.g., unlimited swipes, verified label).

4. Set Profile
    POST /profile: Set user's profile 
