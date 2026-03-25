<div align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/master/icons/folder-aws-open.svg" width="100" height="100" alt="Go-Storage Logo">
  
  # Go Cloud Storage ☁️
  
  ### A lightweight Go storage abstraction focused on S3-compatible providers (AWS S3, MinIO) and multitenant scenarios. It provides a minimal, interface-driven API so application code can remain provider-agnostic.
  
  [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
  [![Go Report Card](https://goreportcard.com/badge/github.com/D-Kale/go-storage)](https://goreportcard.com/report/github.com/D-Kale/go-storage)
  [![Go Reference](https://pkg.go.dev/badge/github.com/D-Kale/go-storage.svg)](https://pkg.go.dev/github.com/D-Kale/go-storage)

  **S3 Compatible** • **MinIO Ready** • **Multitenant focused** • **Lightweight**
  
  ---
  
  ### 🌐 Choose your language / Selecciona tu idioma
  
  [**English Documentation**](#-english-version) | [**Documentación en Español**](#-versión-en-español)

  ---
</div>

<br/>

## 🇺🇸 English Version

### 🚀 Why this project is useful
- **Swap storage providers** without changing business logic (interface-based design).
- **Built for multitenancy**: configure per-tenant buckets and credentials at runtime.
- **Includes an S3-compatible driver** (`s3/`) that works with AWS S3 and MinIO.

### 🏗️ Quick overview
Project layout (top-level files):
- `storage.go`: core Storage interface (Upload, Delete, Exists, GetURL).
- `s3/`: S3/MinIO driver implementation (connector, buckets, driver).
- `go.mod`, `go.sum`: module and dependency info.
- `LICENSE.md`: project license.

### 🛠️ Installation
```bash
go get [github.com/D-Kale/go-storage@latest](https://github.com/D-Kale/go-storage@latest)
```

### Import the s3 driver where needed:

```go
import (

    "context"

    "github.com/D-Kale/go-storage/s3"

)
```

### 💻 Getting started — example

```
import (
    "context"
    "fmt"
    "strings"
    "[github.com/D-Kale/go-storage/s3](https://github.com/D-Kale/go-storage/s3)"
)

func main() {
    cfg := s3.Config{
        AccessKey:    "minioadmin",
        SecretKey:    "minioadmin",
        Bucket:       "tenant-1",
        Endpoint:     "http://localhost:9000",
        Region:       "us-east-1",
        UsePathStyle: true,
    }

    driver, err := s3.New(cfg)
    if err != nil { /* handle error */ }

    ctx := context.Background()
    url, err := driver.Upload(ctx, "path/to/object.txt", strings.NewReader("hello"))
    if err != nil { /* handle error */ }
    
    fmt.Println("Uploaded to:", driver.GetURL(url))
}
```

See s3 package for configuration options and helper methods (EnsureBucketExists, bucket management).

### 🔌 API surface (high level)
- The Storage interface (storage.go) exposes:

- Upload(ctx, path string, file io.Reader) (string, error)

- Delete(ctx, path string) error

- Exists(ctx, path string) (bool, error)

- GetURL(path string) string

### 💡 Usage patterns & tips
- Use per-tenant driver instances to isolate buckets/credentials.

- For local development prefer MinIO with UsePathStyle=true and a local endpoint.

- Rely on driver.EnsureBucketExists (s3 driver) during onboarding to create tenant buckets.

## 🇪🇸 Versión en Español

### 🚀 Por qué es útil este proyecto
- **Intercambia proveedores de almacenamiento** sin cambiar la lógica de negocio (diseño basado en interfaces).

- **Diseñado para multitenancy:** configura buckets y credenciales por inquilino en tiempo de ejecución.

- **Incluye un driver compatible con S3** (s3/) que funciona con AWS S3 y MinIO.

### 🏗️ Resumen rápido
Estructura del proyecto:

- `storage.go`: Interfaz principal de almacenamiento (Upload, Delete, Exists, GetURL).

- `s3/`: Implementación del driver S3/MinIO (conector, buckets, driver).

- `go.mod`, go.sum: Información del módulo y dependencias.

- `LICENSE.md`: Licencia del proyecto.

### 🛠️ Instalación
```
go get [github.com/D-Kale/go-storage@latest](https://github.com/D-Kale/go-storage@latest)
```

### Importar la libreria de S3 donde corresponda:

```go
import (

    "context"

    "github.com/D-Kale/go-storage/s3"

)
```

### 💻 Primeros pasos — ejemplo
```
import (
    "context"
    "fmt"
    "strings"
    "[github.com/D-Kale/go-storage/s3](https://github.com/D-Kale/go-storage/s3)"
)

func main() {
    config := s3.Config{
        AccessKey:    "minioadmin",
        SecretKey:    "minioadmin",
        Bucket:       "tenant-1",
        Endpoint:     "http://localhost:9000",
        Region:       "us-east-1",
        UsePathStyle: true,
    }

    driver, err := s3.New(config)
    if err != nil { /* manejar error */ }

    ctx := context.Background()
    url, err := driver.Upload(ctx, "ruta/al/objeto.txt", strings.NewReader("hola"))
    if err != nil { /* manejar error */ }
    
    fmt.Println("Subido a:", driver.GetURL(url))
}
```

Revisa el paquete de s3 para revisar la configuración, opciones y helpers (EnsureBucketExists, bucket management).

### 🔌 Interfaz de la API (Nivel alto)
La interfaz Storage expone:

- Upload(ctx, path string, file io.Reader) (string, error)

- Delete(ctx, path string) error

- Exists(ctx, path string) (bool, error)

- GetURL(path string) string

### 💡 Patrones de uso y consejos
- Usa instancias de driver por inquilino para aislar buckets/credenciales.

- Para desarrollo local, prefiere MinIO con UsePathStyle=true y un endpoint local.

- Confía en driver.EnsureBucketExists (driver s3) durante el onboarding para crear los buckets de los inquilinos.

### 🤝 Contributing & support
Contributions welcome — please open issues or pull requests.

Report bugs or request features via the repository issues.

### 📄 License
This project is licensed under the MIT License — see LICENSE.md.

### 👤 Maintainers
Maintained by the repository owner. For contact or maintainer list, check the repository metadata or open an issue.

<div align="center">
<sub>Built with ❤️ by <b>D-Kale</b></sub>
</div>