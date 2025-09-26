# Library Management System

A comprehensive library management system built with Go (Golang) and Gin framework, featuring role-based access control and complete library operations management.

## 📚 Features

- **Multi-Role System**
  - Owner (Super Admin)
  - Library Admin
  - User (Reader)

- **Authentication & Authorization**
  - JWT-based authentication
  - Role-based access control
  - Secure password handling

- **Library Operations**
  - Book management (Add, Update, Delete, Search)
  - Issue/Return management
  - User management
  - Multiple libraries support

## 🏗 System Architecture

### Database Schema
- **Books**: Manages book inventory
- **Libraries**: Stores library information
- **Users**: Handles user data with role-based access
- **IssueRegistry**: Tracks book lending
- **RequestEvents**: Manages book requests

### API Endpoints

#### Authentication
- `POST /signup`: User registration
- `POST /login`: User login
- `GET /logout`: User logout

#### Owner Routes (/owner)
- Create new libraries
- Manage library admins

#### Admin Routes (/admin)
- Book management
- Issue approval
- Inventory management

#### User Routes (/user)
- Browse books
- Request books
- Search functionality

## 🚀 Setup and Installation

### Prerequisites
1. Go (version 1.16 or higher)
2. SQLite
3. Git

### Installation Steps

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd library
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up the database**
   - The system uses SQLite by default
   - Database file will be created automatically as `Library.db`

4. **Run the application**
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:8080`

## 🔧 Configuration

The system uses the following default configuration:
- Port: 8080
- Database: SQLite (Library.db)
- JWT Secret: Configured in auth middleware

## 📝 API Documentation

### Authentication APIs

#### Sign Up
```http
POST /signup
Content-Type: application/json

{
    "name": "User Name",
    "email": "user@example.com",
    "contact_number": 1234567890,
    "password": "secure_password",
    "role": "user"
}
```

#### Login
```http
POST /login
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "secure_password"
}
```

### Book Management APIs

#### Add New Book (Admin)
```http
POST /admin/addBook
Content-Type: application/json
Authorization: Bearer <token>

{
    "isbn": 123456789,
    "title": "Book Title",
    "author": "Author Name",
    "publisher": "Publisher Name",
    "version": "1.0",
    "totalCopies": 5
}
```

#### Search Books (User)
```http
GET /user/searchBook?query=<search_term>
Authorization: Bearer <token>
```

## 🔐 Security Features

1. JWT-based authentication
2. Password hashing
3. Role-based access control
4. Input validation
5. CORS protection

## 📊 Database Structure

![Database Schema](https://i.imgur.com/example.png)

### Key Tables:
- Books
- Users
- Libraries
- IssueRegistry
- RequestEvents

## 🛠 Development

### Project Structure
```
library/
├── admin/          # Admin-specific operations
├── controllers/    # Route handlers
├── database/       # Database models and connection
├── DB/            # Database initialization
├── Middlewares/   # Authentication middleware
├── otp/           # OTP functionality
├── owner/         # Owner-specific operations
├── services/      # Business logic
├── user/          # User-specific operations
└── utils/         # Utility functions
```

## 🧪 Testing

To run tests:
```bash
go test ./...
```

## 📈 Future Improvements

1. Email notifications
2. Fine management system
3. Book reservation system
4. Analytics dashboard
5. PDF generation for reports

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For support, email [support@example.com](mailto:support@example.com) or create an issue in the repository.
