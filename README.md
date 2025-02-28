# Go Socket.IO Server Example Tutorial

Socket.io server written in Go that supports godot and hyperpad clients.
This tutorial will guide you through setting up a Go Socket.IO server using the example provided in the [Dangerbites/Go-Socket.io-Server-Example](https://github.com/Dangerbites/Go-Socket.io-Server-Example) repository.

---

## Step 1: Initialize a Go Project

1. **Create a new folder** for your Go project.
2. Open your terminal and navigate to the folder.
3. Run the following command to initialize a Go module:

   ```bash
   go mod init <your-module-name>
   ```

   Replace `<your-module-name>` with a name of your choice, e.g., `go-socketio-server`.

---

## Step 2: Install the Go Socket.IO Library

To use Socket.IO in your Go project, you need to install the `socket.io` library. Run the following command in your terminal:

```bash
go get github.com/zishang520/socket.io/v2/socket
```

This will download and install the necessary dependencies.

---

## Step 3: Add the Code

1. **Create a new file** named `main.go` in your project folder.
2. Copy the code from the [main.go file in the repository](https://github.com/Dangerbites/Go-Socket.io-Server-Example/blob/main/main.go) and paste it into your `main.go` file.

   Alternatively, you can download the `main.go` file directly from the repository and place it in your project folder.

3. **Add the `index.html` file**:
   - Create a new file named `index.html` in your project folder.
   - Copy the contents of the [index.html file from the repository](https://github.com/Dangerbites/Go-Socket.io-Server-Example/blob/main/index.html) and paste it into your `index.html` file.

   The `index.html` file is a simple client-side interface to interact with the Socket.IO server.

---

## Step 4: Run the Server

Once the code is in place, you can run the server using the following command:

```bash
go run main.go
```

This will start the Socket.IO server on the default port (usually `localhost:3000`).

To test the server, open your browser and navigate to `http://localhost:3000`. You should see the `index.html` page, which allows you to send and receive messages via the Socket.IO server.

---

## Important Notes

### Godot Client Compatibility
- If you're using a **Godot client**, ensure it is on version **4.0.x**.
- Use this plugin, [msabaeian/godot-socketio](https://github.com/msabaeian/godot-socketio).

### HyperPad Client Compatibility
- If you're using a **HyperPad client**, ensure it is on version **3.0.x**.
- If the connection doesn't work, try turning off the **reconnect** feature in the client.

