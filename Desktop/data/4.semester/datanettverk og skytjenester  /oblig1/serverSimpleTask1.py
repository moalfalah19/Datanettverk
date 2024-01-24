import socket

SERVER_HOST = '127.0.0.1'
SERVER_PORT = 12000

# Create socket
server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
try:
    server_socket.bind((SERVER_HOST, SERVER_PORT))
except : 
    print("Bind failed. Error : ")

# Listen for incoming connections
server_socket.listen(1)

# Print message to confirm server is running and listening
print('Listening on port %s ...' % SERVER_PORT)
 
 #This function handles incoming client requests by receiving the request, parsing the request's HTTP headers, and
 #returning the appropriate response to the client.
def handle_client(client_connection):

    # Receive client request and decode the data
    request = client_connection.recv(1024).decode()
    print(request)

    # Parse HTTP headers
    Headers = request.split('\n')
    fileName = Headers[0].split()[1]

    # Get the content of the file
    if fileName == '/':
        fileName = '/index.html'

    try:
         # Open file and read its contents
        fin = open('oblig1' + fileName)
        content = fin.read()
        fin.close()

        # Send HTTP response
        response = 'HTTP/1.0 200 OK\n\n' + content

    except FileNotFoundError:
        # If file not found, construct 404 error response
        response = 'HTTP/1.0 404 NOT FOUND\n\n404 File Not Found'

    # Send response to client and close connection
    client_connection.sendall(response.encode())
    client_connection.close()

while True:    
    # Wait for client connections
    connection, client_address = server_socket.accept()
    print('Server connected by ', client_address)

    # Handle the client request
    handle_client(connection)
    
server_socket.close()
