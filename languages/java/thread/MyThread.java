package thread;

// Only need to implement run() in Runnable interface 
public class MyThread implements Runnable {
	public void run() {
		for (int i = 0; i < 20; i++) {
			String threadName = Thread.currentThread().getName();
			System.out.println(threadName + " is running");				
		}
		
		this.withdraw();
		this.deposit();
	}
	
	// synchronized only allows one thread to access the method at a given time (lock)
	private synchronized void withdraw() {
		for (int i = 0; i < 30; i++) {
			String threadName = Thread.currentThread().getName();
			System.out.println(threadName + " takes money");				
		}		
	}
	
	/* 
	 * Locks are not per 'method', but per 'object' 
	 * 
	 * Once alpha acquires a lock for deposit(), 
	 * beta cannot acquire a lock even if it's trying to run withdraw() 
	 * 
	 * In other words a synchronized function must fully execute before 
	 * another thread can access any of the synchronized functions 
	 */
	private synchronized void deposit() {
		for (int i = 0; i < 30; i++) {
			String threadName = Thread.currentThread().getName();
			System.out.println(threadName + " puts money");				
		}			
	}
}
