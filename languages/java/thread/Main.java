package thread;


/* 
 * In Java, multiple programs aren't actually running in parallel
 * JVM just switches between threads very quickly 
 */

public class Main {
	/* 
	 * Runnable is a job, and thread is a worker that will do the job 
	 * call .start() to start the thread 
	 */
	public static void main(String[] args) {
		Runnable threadJob = new MyThread();
		Thread alpha = new Thread(threadJob);
		Thread beta = new Thread(threadJob);
		
		alpha.setName("Alpha");
		beta.setName("Beta");
		
		alpha.start();
		beta.start();
		System.out.println("This is in main");
	}
}
