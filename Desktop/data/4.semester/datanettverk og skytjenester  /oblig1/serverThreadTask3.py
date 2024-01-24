import socket
import _thread as thread

SERVER_HOST = '127.0.0.1'
SERVER_PORT = 12000

# Create socket
server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
try:
    server_socket.bind((SERVER_HOST, SERVER_PORT))
except : 
    print("Bind failed. Error : ")

server_socket.listen(99)
print('Listening on port %s ...' % SERVER_PORT)



def handle_client(client_connection):
    while True:
        # Get the client request
        request = client_connection.recv(1024).decode()
        print(request)

        # Parse HTTP headers
        Headers = request.split('\n')
        fileName = Headers[0].split()[1]

        # Get the content of the file
        if fileName == '/':
            fileName = '/index.html'
        
        try:
            Fin = open('oblig1' + fileName)
            content = Fin.read()
            Fin.close()

        # Send HTTP response
            response = 'HTTP/1.0 200 OK\n\n' + content

        except FileNotFoundError:

            response = 'HTTP/1.0 404 NOT FOUND\n\n404 File Not Found'

        client_connection.sendall(response.encode())
        client_connection.close()

while True:    
    # Wait for client connections
    connection, client_address = server_socket.accept()
    print('Server connected by ', client_address)
 
    # Start new thread to handle incoming client request
    thread.start_new_thread(handle_client, (connection,))

# Close server socket
server_socket.close()

