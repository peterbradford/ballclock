package ballclock;

import java.util.*;

class Queues {
    private Deque<Integer> mainQueue = new ArrayDeque<>();
    private Deque<Integer> minQueue = new ArrayDeque<>();
    private Deque<Integer> fiveMinQueue = new ArrayDeque<>();
    private Deque<Integer> hourQueue = new ArrayDeque<>();

    private int days = 0;

    int getDays() {
        return days;
    }

    Queues(int balls) {
        for (int i = 1; i <= balls; i++) {
            mainQueue.add(i);
        }
    }

    boolean IsOriginalPosition()  {
        Iterator<Integer> mainQIterator = mainQueue.iterator();
        int i = 1;
        while (mainQIterator.hasNext()) {
            if (mainQIterator.next() != i) {
                return false;
            }
            i++;
        }
        return true;
    }

    void pushMinute() {
        int x = (int)mainQueue.poll();
        if (minQueue.size() == 4) {
            while(!minQueue.isEmpty()) {
                mainQueue.addLast(minQueue.removeLast());
            }
            pushFiveMinute(x);
        } else {
            minQueue.add(x);
        }
    }

    private void pushFiveMinute(int x) {
        if (fiveMinQueue.size() == 11) {
            while(!fiveMinQueue.isEmpty()) {
                mainQueue.addLast(fiveMinQueue.removeLast());
            }
            pushHour(x);
        } else {
            fiveMinQueue.add(x);
        }
    }

    private void pushHour(int x) {
        if (hourQueue.size() == 11) {
            while(!hourQueue.isEmpty()) {
                mainQueue.addLast(hourQueue.removeLast());
            }
//            System.out.println("adding day: " + days);
            days++;
            mainQueue.add(x);
        } else {
            hourQueue.add(x);
        }
    }
}