# Go LDAP REST API

## 🌟 Project Description

This project is a REST API developed to perform LDAP (Lightweight Directory Access Protocol) operations. Using Go, it enables interaction with LDAP servers through a flexible and secure API.

## ✨ Features

- Connect to LDAP servers
- User and group querying
- Simple and secure REST API endpoints
- Swagger documentation
- Docker and docker-compose support

## 🚀 Requirements

- Go 1.20+
- Docker (optional)
- LDAP server

## 🔧 Installation

### Local Setup

1. Clone the repository:

```bash
git clone https://github.com/cihantaylan/go-ldap-rest.git
cd go-ldap-rest
```

2. Install dependencies:

```bash
go mod download
```

3. Set up environment variables:

```bash
cp .env.example .env
# Edit the .env file with your configuration
```

4. Run the application:

```bash
go run main.go
```

### Docker Setup

1. Ensure Docker and Docker Compose are installed

2. Start the application:

```bash
docker-compose up --build
```

## 📖 API Documentation

Swagger documentation is available at `http://localhost:8082/swagger/index.html`

## 🔒 Configuration

You can configure the following environment variables in the `.env` file:

- `LDAP_HOST`: LDAP server address
- `LDAP_PORT`: LDAP server port
- `LDAP_BIND_DN`: Binding username
- `LDAP_BIND_PASSWORD`: Binding password

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 💬 Support & Community

- 📫 Author: Cihan TAYLAN
  - Website: [cihantaylan.com](https://cihantaylan.com)
  - GitHub: [@cihantaylan](https://github.com/cihantaylan)
  - LinkedIn: [cihantaylan](https://www.linkedin.com/in/cihantaylan/)

### Issue Reporting

Found a bug or have a feature request? Please open an issue on [GitHub](https://github.com/cihanTAYLAN/grpc-boilerplate-realtime/issues) or reach out on [X](https://x.com/cihantaylan24).

---

<div align="center">
  <sub>Built with ❤️ by Cihan TAYLAN</sub>
  <br>
  ⭐ Don't forget to star this project if you found it helpful!
</div>
