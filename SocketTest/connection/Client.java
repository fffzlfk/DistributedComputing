import java.io.*;
import java.net.*;

public class Client {
	public static void main(String[] args) throws IOException {
		Socket s = new Socket("localhost", 4999);
		PrintWriter pw = new PrintWriter(s.getOutputStream());
		pw.println("Is it working?");
		pw.flush();

		InputStreamReader in = new InputStreamReader(s.getInputStream());
		BufferedReader br = new BufferedReader(in);
		System.out.println("Echo: " + br.readLine());
		s.close();
	}
}
