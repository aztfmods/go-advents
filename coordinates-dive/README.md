```mermaid
graph TD
    A(Start) --> B[Input Command (cmd)]
    B -->|parts := strings.Fields(cmd)| C[Split Command]
    C -->|fmt.Sscanf(parts[1], '%d', &value)| D[Parse Number]
    D --> E{Command Type?}
    E -->|Forward| F[Increase Horizontal Position]
    E -->|Down| G[Increase Depth]
    E -->|Up| H[Decrease Depth]
    F --> I(End)
    G --> I
    H --> I
```
