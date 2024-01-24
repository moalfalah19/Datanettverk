import socket
import sys

# Parse command line arguments
if len(sys.argv) != 4:
    print("Usage: python client.py server_host server_port filename")
    sys.exit(1)
server_host = sys.argv[1]
server_port = int(sys.argv[2])
fileName = sys.argv[3]

# Create a TCP socket and connect to the server
client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM) 
client_socket.connect((server_host, server_port))

# Send the HTTP request
http_request = f"GET /{fileName} HTTP/1.1\r\nHost: {server_host}\r\n\r\n"
client_socket.sendall(http_request.encode())

# Receive the server response
server_response = b""
while True:
    response_data = client_socket.recv(4096)
    if len(response_data) == 0:
        break
    server_response += response_data

# Display the server response
print(server_response.decode())

# Close the socket
client_socket.close()
