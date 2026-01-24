# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Your Role in This Repository

**This is a learning repository.** The user is learning web development through hands-on implementation.

**What you SHOULD do:**

- Explain concepts, patterns, and best practices
- Answer questions about the code, architecture, or technologies
- Help document learnings in the appropriate .md files (see below)
- Point out issues or suggest improvements, but DO NOT implement them
- Provide guidance on how to implement features when asked

**What you SHOULD NOT do:**

- DO NOT directly modify code unless explicitly requested by the user
- DO NOT proactively implement features - the user wants to code themselves
- The user learns by doing, so let them write the code

## Learning Documentation

When the user learns something new or asks you to document learnings, create or update the appropriate file:

- **`docs/learn-React.md`** - React-specific concepts (hooks, components, state, effects, etc.)
- **`docs/learn-ReactRouter.md`** - React Router concepts (loaders, actions, navigation, routing, etc.)
- **`docs/learn-web.md`** - Web fundamentals (JavaScript, HTML, CSS, browser APIs, HTTP, etc.)

Always categorize learnings appropriately. If something involves multiple technologies, choose the most specific category.

**Goal:** These documents should become comprehensive, well-structured references that enable future development with React and React Router. Build them incrementally as a knowledge base that the user can refer to later.

### Documentation Principles

When writing or updating learning documentation, follow these principles:

**1. Development Flow Structure**

- Organize content following the natural development progression
- Use "Problem → Solution" format to show why each concept is needed
- Example flow: Static pages → Display data → User interaction → Dynamic values → External data

**2. Show Necessity Before Implementation**

- Always explain **why** a concept is needed before showing **how** to use it
- Start each section with a problem statement
- Make it clear what issue the concept solves

**3. General and Reusable Examples**

- Use "使用例" (usage examples) not "このプロジェクトでの例" (this project's example)
- Write examples that can be applied to any project, not just the current one
- Keep examples focused and minimal while still being realistic

**4. Incremental Complexity**

- Start with the simplest version of a concept
- Gradually introduce more complex patterns
- Show common pitfalls (❌ bad example) alongside correct approaches (✅ good example)

**5. Comprehensive Coverage**

- Include all essential information needed for future development
- Add tables for quick reference (comparison tables, decision matrices)
- Cover edge cases and common scenarios
- Include best practices and common patterns

**6. Code Examples Best Practices**

- Always include type annotations (TypeScript)
- Show both basic and advanced usage
- Use realistic variable names and scenarios
- Add comments only where the logic isn't self-evident

**7. Consistency in Formatting**

- Use consistent heading levels (## for main topics, ### for subtopics)
- Keep code blocks properly formatted with language tags
- Use tables for comparisons and structured information
- Maintain consistent terminology throughout

## Project Overview

A simple memo application built as a learning project for RESTful API design and SPA architecture. The backend uses Go with Gin framework, and the frontend uses React with TypeScript and React Router v7.

**Tech Stack:**

- Backend: Go + Gin Web Framework + MySQL 8.0
- Frontend: React 19 + TypeScript + React Router v7 + Vite + Tailwind CSS v4

## Database

**Database Name:** `memo_app`

**Table:** `memos`

- `id` (BIGINT, PK, AUTO_INCREMENT)
- `title` (VARCHAR(255), NOT NULL)
- `content` (TEXT)
- `created_at` (DATETIME, DEFAULT CURRENT_TIMESTAMP)

**Environment Variables Required:**

- `DBUSER`: MySQL username
- `DBPASS`: MySQL password

Database connection is configured in backend/main.go:266-272 to connect to `127.0.0.1:3306`.

## Development Commands

### Backend (Go/Gin)

Navigate to `backend/` directory first:

```bash
# Run the API server
go run main.go

# Regenerate Swagger documentation (after updating godoc comments)
swag init

# Install dependencies
go mod download
```

The backend server runs on `http://localhost:8080`.

**API Endpoints:**

- `GET /memos` - Get all memos
- `GET /memos/:id` - Get memo by ID
- `POST /memos` - Create new memo
- `PUT /memos/:id` - Update memo
- `DELETE /memos/:id` - Delete memo
- `GET /swagger/*any` - Swagger UI documentation

### Frontend (React/TypeScript)

Navigate to `frontend/` directory first:

```bash
# Development server (with hot reload)
npm run dev

# Production build
npm run build

# Start production server
npm run start

# Type checking
npm run typecheck
```

The frontend dev server typically runs on `http://localhost:5173`.

## Architecture

### Backend Architecture (backend/)

**Single-file architecture:** All backend logic is in `main.go` (~300 lines).

- **Data layer functions:** `getAllMemos()`, `getMemoByID()`, `addMemo()`, `editMemo()`, `deleteMemoByID()` - direct SQL operations using database/sql
- **Handler functions:** `getAllMemosHandler()`, `getMemoByIDHandler()`, `addMemoHandler()`, `editMemoHandler()`, `deleteMemoHandler()` - Gin HTTP handlers
- **Server struct:** Holds database connection, methods are HTTP handlers
- **Swagger annotations:** Inline godoc comments for API documentation

Database connection uses `database/sql` with `go-sql-driver/mysql`. The `Server` struct pattern allows handlers to access the database connection.

### Frontend Architecture (frontend/app/)

**React Router v7 file-based routing:**

Routes are defined in `app/routes.ts`:

- `/` → `routes/index.tsx` (memo list page)
- `/memos/:id` → `routes/memo.tsx` (memo detail page)
- `/memos/:id/edit` → `routes/edit_memo.tsx` (memo edit page)

**Data fetching pattern:**

- API client functions in `app/api/memos.ts` handle all backend communication
- Route components use `loader()` functions for data fetching (React Router v7 pattern)
- Route components use `action()` functions for mutations (React Router v7 pattern)
- All API calls target `http://localhost:8080`

**Component pattern:**

- Each route file exports: `loader()` (for data fetching), `action()` (for mutations), and default component
- Type safety with `Route.ComponentProps`, `Route.LoaderArgs`, `Route.ActionArgs`
- Uses React Router's `Form` component for mutations with automatic revalidation

**Styling:** Tailwind CSS v4 with utility classes, no separate CSS files.

## Current Implementation Status

**Implemented:**

- Memo list view with cards (index page)
- Memo detail view with back button
- Memo edit functionality with form submission
- Navigation between list → detail → edit pages
- Responsive layout with Tailwind

**Not Yet Implemented:**

- Delete memo functionality (buttons exist but have no onClick handler)
- Create new memo functionality
- Error handling in frontend
- Form validation
- Loading states

## Important Notes

- The API client in `frontend/app/api/memos.ts` is incomplete - missing create and delete functions
- Delete buttons in the UI are non-functional (frontend/app/routes/index.tsx:44 and routes/memo.tsx:41)
- Backend uses `c.IndentedJSON()` which adds pretty-printing overhead - consider `c.JSON()` for production
- No CORS configuration in backend - add if deploying frontend separately
- MySQL connection requires `ParseTime: true` to correctly handle `created_at` as `time.Time`
