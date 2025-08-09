# MCP Addition Server

A simple Go project demonstrating a Model Context Protocol (MCP) server that provides a tool for adding two numbers.

## How it Works

This project consists of two main components:
-   **MCP Server (`main.go`)**: A Go application that starts an MCP server and exposes an `add` tool. It communicates over standard input/output (stdio).
-   **MCP Client (`client/main.go`)**: A Go application that acts as a client to the server. It starts the server as a subprocess and communicates with it over stdio to call the `add` tool.

The server's `add` tool takes two string parameters, `a` and `b`, converts them to integers, calculates their sum, and returns the result as a string.

## How to Run and Test

### Prerequisites
- Go (version 1.24.3 or later)
- Visual Studio Code (optional, for manual testing)

### Running the Test Suite
To run the automated unit tests for the server's `Add` function, execute the following command in your terminal from the root of the project:
```bash
go test -v
```

### Running the Client (and Server)
The client application is designed to automatically start the server. To run the client, which will in turn run the server and call the `add` tool, execute the following command:
```bash
go run client/main.go
```
You should see the output `2025/xx/xx xx:xx:xx 30`, which is the sum of 10 and 20 as hardcoded in the client.

### Manual Testing from VS Code

You can also test the server manually by interacting with it directly. This is useful for understanding the MCP protocol.

1.  **Start the Server**:
    Open a terminal in VS Code and run the server directly:
    ```bash
    go run main.go
    ```
    The server is now running and waiting for JSON-RPC messages on its standard input.

2.  **Send a Request**:
    The server expects JSON-RPC 2.0 messages. To call the `add` tool, you need to send a `callTool` request. You can type or paste the following JSON message into the same terminal where the server is running and then press Enter.

    *Make sure to send the `Content-Length` header followed by two newlines (`\r\n\r\n`) before the JSON payload.*

    ```
    Content-Length: 133

    {
        "jsonrpc": "2.0",
        "method": "mcp/callTool",
        "params": {
            "name": "add",
            "arguments": {
                "a": "5",
                "b": "7"
            }
        },
        "id": 1
    }
    ```
    *(Note: You might need to be careful with newlines when pasting into the terminal. It might be easier to write this to a file and pipe it to the server, e.g., `cat request.json | go run main.go`)*

3.  **Receive the Response**:
    The server will process the request and print the JSON-RPC response to its standard output. You should see a response like this:

    ```
    Content-Length: 135

    {
        "jsonrpc": "2.0",
        "id": 1,
        "result": {
            "isError": false,
            "content": [
                {
                    "type": "text",
                    "text": "12"
                }
            ]
        }
    }
    ```

## How an AI Assistant Client Would Use This Tool

An AI assistant or a similar "claude code like client" would interact with this server by following the Model Context Protocol. When the assistant determines that it needs to perform an addition, it would formulate a `callTool` request.

For example, if a user asks the assistant "What is 123 + 456?", the assistant would:

1.  Identify that the `add` tool provided by this server is suitable for the task.
2.  Construct a JSON-RPC request to call the `add` tool with the appropriate arguments.
3.  Send the request to the server.

The raw request sent by the assistant's MCP client would look like this:

```json
{
    "jsonrpc": "2.0",
    "method": "mcp/callTool",
    "params": {
        "name": "add",
        "arguments": {
            "a": "123",
            "b": "456"
        }
    },
    "id": "request-123"
}
```

4.  Receive the response from the server:

```json
{
    "jsonrpc": "2.0",
    "id": "request-123",
    "result": {
        "isError": false,
        "content": [
            {
                "type": "text",
                "text": "579"
            }
        ]
    }
}
```
5.  Parse the response, extract the result "579", and present it to the user as the answer.
