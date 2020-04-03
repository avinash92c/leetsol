import java.util.concurrent.*;
class FooBar {
    private int n;
    Semaphore a,b;

    public FooBar(int n) {
        a = new Semaphore(1);
        b = new Semaphore(0);
        this.n = n;
    }

    public void foo(Runnable printFoo) throws InterruptedException {
        
        for (int i = 0; i < n; i++) {
            a.acquire();
        	// printFoo.run() outputs "foo". Do not change or remove this line.
        	printFoo.run();
            b.release();
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {
        
        for (int i = 0; i < n; i++) {
            b.acquire();
            // printBar.run() outputs "bar". Do not change or remove this line.
        	printBar.run();
            a.release();
        }
    }
}
