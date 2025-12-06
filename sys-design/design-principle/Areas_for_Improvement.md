# System Design Journey

## Architectural Principle: Program to an Interface, Not an Implementation

[design-principles/program-to-interface](https://officialcto.com/interview-section/design-principles/program-to-interface)

A core principle for building robust, maintainable, and testable software is to depend on abstractions (interfaces) rather than concrete implementations. This means the business logic of our application should not have direct knowledge of *how* an external service or data store works, only on *what* operations it can perform.

This principle should be applied to all modules in the `internal/` directory that communicate with an external, swappable system.

### 1. Database Abstraction (Repository Pattern)

-   **Problem:** The application's core logic is directly tied to PostgreSQL. Migrating to another database (e.g., BigQuery for analytics, or even a mock for testing) would require widespread code changes.
-   **Solution:** Define a `Repository` interface in `internal/database/` that specifies the data operations required (e.g., `CreateGroundTruth`, `GetLatestRecord`).
    -   The concrete implementation (`PostgresRepository`) will live alongside it and use the `sqlc`-generated code.
    -   Business logic in `workflows/` will only ever depend on the `Repository` interface.
-   **Benefit:** Decouples the application from PostgreSQL, improves testability by allowing for mock repositories, and localizes all database-specific code to a single place.

### 2. AI Model Abstraction (Strategy Pattern)

-   **Problem:** The `gemini` client is tightly coupled to the Google Gemini SDK. Supporting other models from providers like OpenAI, Alchemy, or others is impossible without significant refactoring. The existing `GenAIClient` interface still uses Gemini-specific types, preventing true abstraction.
-   **Solution:** Define a generic `ImageGenerator` interface (e.g., in `internal/imaging/`) that uses simple, application-defined request and response structs, not vendor-specific ones.
    -   Each provider (`gemini`, `openai`, etc.) will have its own client that implements this common interface.
    -   Workflows will depend only on the `ImageGenerator` interface.
-   **Benefit:** The application becomes "model-agnostic." We can easily add support for new image generation APIs or switch between them via configuration, enabling true A/B testing or cost optimization.

### 3. Other External Services

This same pattern should be applied to other external services as they are developed:

-   **`storage/`**: An interface would allow switching between Google Cloud Storage (GCS), AWS S3, or a local filesystem for testing.
-   **`pubsub/`**: An interface would allow the messaging backend to be swapped between Google Pub/Sub, Kafka, or RabbitMQ without changing the core logic.

By consistently applying this principle, we isolate our core business logic from the ever-changing world of external vendors and services, resulting in a more flexible and long-lasting application.

## Area of Improvements

- *   **Structured Logging:** The current logging is great for interactive development. For a tool that might run as part of an automated CI/CD pipeline, switching to a structured logging library (like Go's built-in `slog`) would be a major improvement. This would allow logs to be easily parsed, filtered, and ingested by monitoring tools.

- *   **Testability of `main`:** The `main` function itself is hard to test directly. A common pattern is to move the core application logic from `main` into a `run(cfg Config)` function that returns an `error`. This makes your main logic testable and keeps the `main` function as a thin wrapper for setup and exit codes.

- *   **Error Handling and Retries:** The current tool logs an error and moves on. For a more robust system, you could implement a simple retry mechanism (e.g., using a library like `go-retryablehttp`) for transient network errors when calling the Gemini API.
