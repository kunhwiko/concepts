// When using comparators, note the difference from the Comparable interface 

public static void exampleOne() {
    Comparator<String> cmp = new Comparator<>() {
        public int compare(String o1, String o2) {
            // if the compare function returns a negative int, o1 goes left 
            // this means elements will be sorted as o1 o2 ....

            // if the compare function returns a positive int, o1 goes right 
            // this means elements will be sorted as o2 o1 ...

            // in the example below, if o2 is lexicographically larger, 
            // the function will return a positive value (o1 goes right)
            // therefore, the comparator will sort in descending order 
            return o2.compareTo(o1);   
        }
    };

    TreeSet<String> res = new TreeSet<>(cmp);
    res.add("Cat"); res.add("Batch"); res.add("Resident");
}

public static void exampleTwo() {
    // Assume there is an array [(5,10),(3,14)]
    // we ultimately want [(3,14),(5,10)] 
    Arrays.sort(arr, new Comparator<int[]>(){
        public int compare(int[] i1, int[] i2) {
            // compare based on the first element
            // if i1[0] is smaller, i1[0] comes first 
            // therefore, this will sort in ascending order based on index 0
            return i1[0] - i2[0];
        }
    });
}

