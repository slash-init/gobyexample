# Go Run vs Go Build

## Key Differences

| Aspect            | `go run`                           | `go build`                           |
| ----------------- | ---------------------------------- | ------------------------------------ |
| **What it does**  | Compiles and runs code in one step | Only compiles, creates an executable |
| **Binary saved?** | No (temporary)                     | Yes (persistent file)                |
| **Execution**     | Direct output                      | Run with `./hello-world`             |
| **Speed**         | Recompiles every time              | One-time compilation                 |

## When to Use Each

### `go run` — Development/Testing

- Use while actively coding
- Quick iterations without manual recompilation
- Speed is not critical

### `go build` — Production/Distribution

- Creates a reusable binary that:
  - Runs without Go installed
  - Shares easily (single file, no source code needed)
  - Avoids recompilation overhead
  - Can be committed to CI/CD pipelines

## Why Build is "Better"

In production, if you run the program multiple times:

- `go run hello-world.go` recompiles each time (slower)
- `./hello-world` runs the pre-built binary (faster)

For a simple program the difference is negligible, but for larger apps with many dependencies, compilation takes seconds—and you'd waste time rebuilding if you run it 10 times.

## Summary

- **Development**: Use `go run` for convenience
- **Production**: Use `go build` for performance, distribution, and reliability
