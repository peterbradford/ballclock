package ballclock;

public class Main {

    public static void main(String... args) {
        int balls = 123;
        Queues bClock = new Queues(balls);
        double start = System.currentTimeMillis();
        bClock.pushMinute();
        while (!bClock.IsOriginalPosition()) {
            bClock.pushMinute();
        }
        double totalTime = System.currentTimeMillis()-start;
        System.out.printf("%s balls cycle after %s days. \nCompleted in %s milliseconds (%1.2f seconds)", balls, bClock.getDays()/2, totalTime, (double)totalTime/1000);
    }
}
