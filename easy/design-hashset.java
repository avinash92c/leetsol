import java.util.HashMap;
class MyHashSet {
    private HashMap map;

    /** Initialize your data structure here. */
    public MyHashSet() {
        map = new HashMap();
    }
    
    public void add(int key) {
        map.put(key,new Object());
    }
    
    public void remove(int key) {
        map.remove(key);
    }
    
    /** Returns true if this set contains the specified element */
    public boolean contains(int key) {
        return map.containsKey(key);
    }
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * MyHashSet obj = new MyHashSet();
 * obj.add(key);
 * obj.remove(key);
 * boolean param_3 = obj.contains(key);
 */
