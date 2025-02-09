# Task Management System

A robust task management system built with Go and a modern web interface. The system allows users to create, manage, and track tasks with features like email notifications, currency conversion, and task exports.

## Table of Contents
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Environment Variables](#environment-variables)
- [API Endpoints](#api-endpoints)
- [Frontend Interface](#frontend-interface)

## Features
- ✨ Create, read, update, and delete tasks
- 📅 Task scheduling with due dates
- 💰 Cost tracking with currency conversion
- 📧 Email notifications using SendGrid
- 🔍 Search functionality
- 📊 Task export to CSV
- 🎯 Task status management
- 💾 Bulk operations support

## Tech Stack
- Backend: Go (Golang)
- Database: MySQL
- Frontend: HTML, JavaScript, Tailwind CSS
- External Services:
  - SendGrid for email notifications
  - ExchangeRate API for currency conversion

## Project Structure
```
task-management/
├── config/
│   └── config.go           # Configuration management
├── database/
│   └── database.go         # Database initialization
├── handlers/
│   └── task_handler.go     # HTTP request handlers
├── models/
│   └── task.go            # Data models
├── services/
│   ├── email_service.go    # Email notification service
│   ├── exchange_service.go # Currency exchange service
│   └── task_service.go     # Task business logic
├── static/
│   └── index.html         # Frontend interface
├── .env                   # Environment variables
├── go.mod                # Go dependencies
├── go.sum                # Go dependencies checksums
└── main.go               # Application entry point
```

## Prerequisites
- Go 1.21 or higher
- MySQL 5.7 or higher
- SendGrid account
- ExchangeRate API account

## Installation

1. Clone the repository:
```bash
git clone https://github.com/RS-8370412/task-management.git
cd task-management
```

2. Install dependencies:
```bash
go mod download
```

3. Set up the database:
   - Create a MySQL database
   - The tables will be automatically created when the application starts

4. Configure environment variables:
   - Copy `.env.example` to `.env`
   - Update the values with your configurations

5. Run the application:
```bash
go run main.go
```

The server will start on port 3090 (http://localhost:3090)

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
SENDGRID_API_KEY=your_sendgrid_api_key
EXCHANGERATE_API_KEY=your_exchange_rate_api_key
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=taskdb
```

## API Endpoints

### Task Management
- `POST /api/tasks` - Create a new task
- `GET /api/tasks` - Get all tasks
- `GET /api/tasks/{id}` - Get a specific task
- `PUT /api/tasks/{id}` - Update a task
- `DELETE /api/tasks/{id}` - Delete a task
- `PUT /api/tasks/{id}/status` - Update task status

### Additional Features
- `GET /api/tasks/search` - Search tasks
- `GET /api/tasks/export` - Export tasks to CSV
- `DELETE /api/tasks/bulk` - Bulk delete tasks

### Query Parameters
- Status filtering: `?status=pending|completed|in-progress`
- Currency conversion: `?currency=USD|EUR|GBP|JPY|INR`
- Search: `?q=search_term`

## Frontend Interface

The web interface provides a user-friendly way to interact with the task management system:

### Features
- Clean, responsive design using Tailwind CSS
- Real-time task creation and updates
- Task filtering by status
- Interactive task editing
- Delete confirmation
- Success/error notifications
- Automatic task list refresh

### Usage
1. Access the web interface at `http://localhost:3090`
2. Use the navigation bar to switch between task creation and viewing
3. Create new tasks using the form
4. View and manage existing tasks in the task list view
5. Edit tasks by clicking the edit icon
6. Delete tasks using the delete icon
7. Filter tasks by status using the dropdown

## Error Handling
The system includes comprehensive error handling:
- Database connection errors
- Invalid request validation
- Email service errors
- Currency conversion failures
- Resource not found errors

## Security Considerations
- SQL injection protection through parameterized queries
- CORS configuration for API security
- Environment variables for sensitive data
- Input validation on both frontend and backend

## Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License
This project is licensed under the MIT License - see the LICENSE file for details
