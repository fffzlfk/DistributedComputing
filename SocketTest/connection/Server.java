import java.net.*;
import java.io.*;

public class Server {

	public static void main(String[] args) throws IOException {
		try {
			ServerSocket serverSocket = new ServerSocket(4999);
			Socket socket = null;
			int cnt = 0;
			System.out.println("Start listen...");

			while (true) {
				socket = serverSocket.accept();
				ServerThread serverThread = new ServerThread(socket);
				serverThread.start();
				cnt++;
				System.out.println("Current the number of Client: " + cnt);
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
	}
}
