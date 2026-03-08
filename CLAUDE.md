# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Development (runs client on :5173 and server on :8080 in parallel)
make dev

# Individual servers
make dev-client   # cd client && npm run dev
make dev-server   # cd server && go run ./cmd/server

# Tests
make test          # runs both
make test-client   # cd client && npm run test:run
make test-server   # cd server && go test ./...

# Single client test file
cd client && npx vitest run src/path/to/file.test.ts

# Linting
cd client && npx eslint .
```

## Architecture

Echelon is an AI agent orchestration system with a React frontend and Go backend.

**Communication:** Vite dev server proxies `/api/*` to `http://localhost:8080`. Server allows CORS from `http://localhost:5173`.

### Server (`server/`)

- `cmd/server/main.go` — entry point; loads `.env`, starts on `PORT` (default 8080)
- `internal/api/router.go` — chi router with logging, recovery, and CORS middleware
- `internal/api/handlers.go` — HTTP handlers (currently returning 501; stubs only)
- `internal/agents/` — agent logic (planner, spawner, budget) — all stubs awaiting implementation
- `internal/anthropic/client.go` — Anthropic API client stub
- `internal/types/types.go` — shared type definitions (empty)

API routes: `POST /api/intake`, `POST /api/plan`, `POST /api/spawn`, `GET /api/health`

### Client (`client/`)

- React 19 + TypeScript + Vite
- Styling: TailwindCSS 4 + shadcn/ui (radix-nova style, neutral base color)
- State: Zustand
- Animation: Framer Motion
- Graph/workflow visualization: `@xyflow/react` + `@dagrejs/dagre` (for agent DAG views)
- Path alias: `@/` → `src/`
- Components go in `src/components/ui/` (shadcn pattern); shared utilities in `src/lib/utils.ts` (`cn()` for class merging)
- Tests use Vitest + jsdom + @testing-library/react; setup file at `src/test/setup.ts`
