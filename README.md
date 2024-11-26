# Pubnode

A free, open-source hub for Software Artisans to share, collaborate, and grow together.

## Features

- 🚀 Project showcase and discovery
- 💬 Developer-focused social feed
- 📨 Inbox for direct communication
- 🔍 Advanced search capabilities
- 🎨 Clean, minimal UI 

## Tech Stack

- Backend: Go with Fiber framework
- Database: SQLite with GORM
- Frontend: Server-side rendered HTML with HTMX
- Styling: TailwindCSS
- Authentication: Multi-provider support (including GitHub)

## Getting Started

### Prerequisites

- Go 1.23.2 or higher
- Node.js and pnpm
- Air (for hot reload during development)

### Development

1. Clone the repository:
```bash
git clone https://github.com/d3cie/pubnode.git
cd pubnode
```

2. Install dependencies:
```bash
pnpm install
go mod download
```

3. Start the development server:
```bash
make dev
```

4. In a separate terminal, start the Tailwind watcher:
```bash
make dev/tw
```

The application will be available at `http://localhost:3000`

## Contributing

Pubnode is open to contributions! Whether it's bug fixes, feature additions, or documentation improvements, we welcome your input.

## License

MIT License - feel free to use and modify as you see fit.
