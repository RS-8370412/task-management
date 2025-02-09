# Task Management System

A full-stack task management application built with Go, MySQL, and modern web technologies. Manage tasks efficiently with features like email notifications, currency conversion, and task exports.

## Features
- ✨ CRUD operations for tasks
- 📅 Task scheduling
- 💰 Cost tracking with currency conversion
- 📧 Email notifications
- 🔍 Search functionality
- 📊 CSV export
- 🎯 Status management
- 💾 Bulk operations

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/RS-8370412/task-management.git
cd task-management
```

2. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configurations
```

3. Install dependencies:
```bash
go mod download
```

4. Run the application:
```bash
go run main.go
```

Visit `http://localhost:3090` to access the application.

## Prerequisites
- Go 1.21+
- MySQL 5.7+
- SendGrid account
- ExchangeRate API account

## Documentation
For detailed documentation, see [DOCUMENTATION.md](DOCUMENTATION.md)

## API Endpoints

### Task Management
- `POST /api/tasks` - Create task
- `GET /api/tasks` - List tasks
- `GET /api/tasks/{id}` - Get task
- `PUT /api/tasks/{id}` - Update task
- `DELETE /api/tasks/{id}` - Delete task
- `PUT /api/tasks/{id}/status` - Update status

### Additional Features
- `GET /api/tasks/search` - Search tasks
- `GET /api/tasks/export` - Export to CSV
- `DELETE /api/tasks/bulk` - Bulk delete

## Tech Stack
- Backend: Go (Golang)
- Database: MySQL
- Frontend: HTML, JavaScript, Tailwind CSS
- Services: SendGrid, ExchangeRate API

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
