```mermaid
graph TD
    A(Start) -->|Input Command (cmd)| B
    B[Split Command] -->|parts := strings.Fields(cmd)| C
    C[Parse Number] -->|var value int; fmt.Sscanf(parts[1], '%d', &value)| D
    D{Command Type?} -->|Forward| E[Increase Horizontal Position]
    D -->|Down| F[Increase Depth]
    D -->|Up| G[Decrease Depth]
    E --> H(End)
    F --> H
    G --> H
```
