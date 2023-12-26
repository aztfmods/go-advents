```mermaid
graph TD
    A(Start) --> B[Input Command (cmd)]
    B --> C[Split Command<br>parts := strings.Fields(cmd)]
    C --> D[Parse Number<br>var value int; fmt.Sscanf(parts[1], "%d", &value)]
    D --> E{Command Type?<br>switch parts[0]}
    E -->|Forward| F[Forward<br>Increase Horizontal Position]
    E -->|Down| G[Down<br>Increase Depth]
    E -->|Up| H[Up<br>Decrease Depth]
    F --> I(End)
    G --> I
    H --> I
```
