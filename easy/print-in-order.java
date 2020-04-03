import java.util.concurrent.*;
class Foo {
    Semaphore a,b;

    public Foo() {
        a = new Semaphore(0);
        b = new Semaphore(0);
    }
    

    public void first(Runnable printFirst) throws InterruptedException {
        
        // printFirst.run() outputs "first". Do not change or remove this line.
        printFirst.run();
        a.release();
    }

    public void second(Runnable printSecond) throws InterruptedException {
        a.acquire();
        // printSecond.run() outputs "second". Do not change or remove this line.
        printSecond.run();
        b.release();
    }

    public void third(Runnable printThird) throws InterruptedException {
        b.acquire();
        // printThird.run() outputs "third". Do not change or remove this line.
        printThird.run();
        
    }
}
