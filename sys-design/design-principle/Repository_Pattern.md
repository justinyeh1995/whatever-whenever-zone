# The Repository Pattern

The Repository Pattern is an architectural design pattern that abstracts the data layer of an application, providing a clean separation between the business logic and the persistence mechanism. It defines a contract (an interface) for data operations, allowing the application's core logic to remain independent of the specific database technology or ORM (Object-Relational Mapper) being used.

## Core Concepts

1.  **Interface (Contract):** An interface is defined that declares all the necessary data access methods for a specific entity or aggregate. This interface represents the "contract" that any data access layer must fulfill. It expresses *what* data operations are available, not *how* they are performed.

    *   **Example in Go:**
        ```go
        package database

        import (
            "context"
            "github.com/google/uuid"
        )

        // GroundTruthRepository defines the operations for managing ground truth data,
        // abstracting the underlying storage mechanism.
        type GroundTruthRepository interface {
            InsertFromRecords(ctx context.Context, modelVersion string) error
            GetGroundTruthByID(ctx context.Context, id uuid.UUID) (*GroundTruth, error)
            // ... other methods like Update, Delete, List, etc.
        }
        ```

2.  **Concrete Implementation:** A concrete class or struct implements the defined interface, providing the actual logic for interacting with a specific data store (e.g., PostgreSQL, BigQuery, an in-memory database, a file system, or a remote API). This is the only part of the application that "knows" about the database details.

    *   **Example in Go (PostgreSQL implementation using sqlc-generated code):**
        ```go
        package database

        import (
            "context"
            "database/sql"
            "your-project/internal/database/sqlc" // sqlc-generated code
            "github.com/google/uuid"
        )

        // PostgresRepository is the PostgreSQL implementation of the GroundTruthRepository.
        type PostgresRepository struct {
            queries *sqlc.Queries
        }

        func NewPostgresRepository(dbConnection *sql.DB) *PostgresRepository {
            return &PostgresRepository{
                queries: sqlc.New(dbConnection),
            }
        }

        func (r *PostgresRepository) InsertFromRecords(ctx context.Context, modelVersion string) error {
            // Calls the sqlc-generated method for a specific query
            return r.queries.InsertGroundTruthFromRecords(ctx, modelVersion)
        }

        func (r *PostgresRepository) GetGroundTruthByID(ctx context.Context, id uuid.UUID) (*GroundTruth, error) {
            // ... calls another sqlc method ...
            return nil, nil // Placeholder
        }
        ```

3.  **Dependency Injection:** The business logic (e.g., service layers, workflow orchestrators) depends on the *interface* of the repository, not its concrete implementation. The specific implementation is "injected" at application startup.

    *   **Example in Go:**
        ```go
        package workflows

        import "your-project/internal/database" // Import the package with the interface

        // EvalWorkflow orchestrates the evaluation process.
        type EvalWorkflow struct {
            repo database.GroundTruthRepository // Depends on the interface!
        }

        func NewEvalWorkflow(repo database.GroundTruthRepository) *EvalWorkflow {
            return &EvalWorkflow{repo: repo}
        }

        func (w *EvalWorkflow) PopulateGroundTruth(ctx context.Context) error {
            // This code doesn't know or care if it's talking to Postgres or BigQuery.
            return w.repo.InsertFromRecords(ctx, "gemini-2.5-flash-image")
        }
        ```

## Benefits

*   **Database Agnosticism:** The most significant advantage. You can swap the underlying database technology (e.g., PostgreSQL to BigQuery, or even to a NoSQL database) by simply writing a new implementation of the repository interface. The rest of your application code remains unchanged.
*   **Enhanced Testability:** For unit tests, you can easily create mock implementations of the repository interface. This allows you to test your business logic in isolation, without requiring a real database connection, making tests faster and more reliable.
*   **Clear Separation of Concerns:** Database-specific logic, SQL queries, and ORM interactions are encapsulated within the repository implementation. Business rules and application logic remain clean and free of persistence details.
*   **Maintainability and Scalability:** Changes to the database schema or migration to a new database technology are localized to the repository implementation, reducing the risk of introducing bugs elsewhere in the application.

## Repository Pattern in Frameworks (e.g., Java Spring Boot)

Frameworks like Java's Spring Boot (especially with Spring Data JPA/JDBC) provide built-in support for the Repository Pattern, often automating much of the implementation.

In Spring Data, you typically define a simple interface that extends `JpaRepository` (or similar base interfaces). Spring then, at runtime, automatically generates a concrete implementation of this interface based on naming conventions and the entity definition. Your application code then just `@Autowired`s (injects) this interface, without ever needing to see or write the concrete implementation.

While the implementation details differ (manual in Go with `sqlc` vs. highly automated in Spring Boot), the underlying principle of abstracting data access through an interface remains the same, providing the identical benefits of flexibility, testability, and maintainability.
