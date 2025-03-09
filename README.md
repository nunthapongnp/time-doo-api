# Time-Doo API

Time-Doo API is a powerful and efficient task management system designed to help you organize and manage your tasks and subtasks with ease. This API provides a robust backend service for creating, updating, retrieving, and deleting tasks and subtasks, ensuring that your productivity is maximized.

## Features

- **Task Management**: Create, update, retrieve, and delete tasks.
- **Subtask Management**: Manage subtasks under each task for better organization.
- **User Authentication**: Secure authentication using Firebase Authentication.
- **Data Storage**: Persistent data storage using Google Firestore.
- **Caching**: Enhanced performance with Redis caching.
- **API Security**: Middleware for secure API access.

## Technology Stack

- **Go**: The primary programming language used for developing the API.
- **Gin**: A high-performance HTTP web framework for Go.
- **Firebase**: Used for authentication and Firestore for data storage.
- **Redis**: Used for caching to improve performance.
- **Docker**: Containerization for easy deployment and scalability.

## Getting Started

### Prerequisites

- Go 1.24.1 or later
- Docker (optional, for containerization)
- Firebase project with Firestore and Authentication enabled
- Redis instance

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/nunthapongnp/time-doo-api.git
   cd time-doo-api
   ```

2. Create a `.env` file in the root directory and add your environment variables:

   ```env
   PORT=3000
   FIRESTORE_PROJECT_ID=your-firestore-project-id
   FIREBASE_API_KEY=your-firebase-api-key
   GOOGLE_APPLICATION_CREDENTIALS=path-to-your-firebase-credentials.json
   REDIS_ADDRESS=your-redis-address
   REDIS_PASSWORD=your-redis-password
   ```

3. Install dependencies:

   ```sh
   go mod tidy
   ```

4. Run the application:
   ```sh
   go run main.go
   ```

### API Endpoints

- **Authentication**

  - `POST /api/v1/auth/get-id-token`: Get Firebase ID token using email and password.

- **Tasks**

  - `POST /api/v1/tasks`: Create a new task.
  - `GET /api/v1/tasks/:taskId`: Retrieve a task by ID.
  - `PUT /api/v1/tasks/:taskId`: Update a task by ID.
  - `DELETE /api/v1/tasks/:taskId`: Delete a task by ID.

- **Subtasks**
  - `POST /api/v1/tasks/:taskId/subtasks`: Create a new subtask under a task.
  - `GET /api/v1/tasks/:taskId/subtasks/:subtaskId`: Retrieve a subtask by ID.
  - `PUT /api/v1/tasks/:taskId/subtasks/:subtaskId`: Update a subtask by ID.
  - `DELETE /api/v1/tasks/:taskId/subtasks/:subtaskId`: Delete a subtask by ID.
