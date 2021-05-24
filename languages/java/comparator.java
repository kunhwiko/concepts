import java.util.*;

/* 
 * 'Comparable' is an interface that allows objects to become comparable
 *  All we have to do is implement the compareTo() method  
 *
 *  You can only ever have one version of compareTo()
 *  What if you want multiple methods that compare objects in different ways
 *  This is what 'comparators' are for  
 */

public class Main {
    private ArrayList<Song> arrList = new ArrayList<Song>();
    
    // Method 1: implementing comparable interface 
    class Song implements Comparable<Song> {
        String title, musician; 
        
        Song(String title, String musician) {
            this.title = title;
            this.musician = musician;
        }
        
        String getTitle() { return this.title; }
        
        String getMusician() { return this.musician; }
        
        // criteria for sorting 
        public int compareTo(Song o) {
            return getTitle().compareTo(o.getTitle());
        }
        
        // override equals to customize what makes an object equal 
        public boolean equals(Song o) {
            return getMusician().equals(o.getMusician());
        }
    }
    
    // Method 2: implementing comparator class 
    class ExampleComparator implements Comparator<Song> {
        /* 
         * if the compare function returns a negative int, o1 goes left 
         * this means elements will be sorted as o1 o2 ....
         * 
         * if the compare function returns a positive int, o1 goes right 
         * this means elements will be sorted as o2 o1 ...
         */
        public int compare(Song o1, Song o2) {
            return o1.getTitle().compareTo(o2.getTitle());
        }
    }
    
    private void runSort() {
        arrList.add(new Song("The Scientist", "Coldplay")); arrList.add(new Song("Kill This Love", "Blackpink"));
        arrList.add(new Song("Paradise", "Coldplay")); arrList.add(new Song("Blueming", "IU"));   
        Collections.sort(arrList);

        // Our criteria for equals is based on musician name 
        Song obj1 = arrList.get(2); Song obj2 = arrList.get(3);
        System.out.println(obj1.equals(obj2));
        
        // Method 3: anonymous classes 
        Collections.sort(arrList, new Comparator<Song>(){
            public int compare(Song s1, Song s2) {
                return s1.getMusician().compareTo(s2.getMusician());
            }           
        });
    }
    
    public static void main(String[] args) {
        // instance of class Main not assigned 
        new Main().runSort();
    }
}

