// naive intuition: translate every word in a hashset, then check if other words contains all the words
// in the first word.
//                  Time:  O(N^2)
//                  Space: O(N)

// intuition 2: copy and sort every word. use a map with word as key and index as values. 
// Create the result based on indecies.
//                  Time: O(N * WLogW)
//                  Space: O(N)

// Intuition 2 improved: counting sort at every word so that the sorting costs O(W) for each word.
// This is possible because each word only contain lower-case english letter and thus have a 
// fixed, stable maximum length.
// Explanation on counting sort: https://blog.csdn.net/wgiyq/article/details/54583399

func groupAnagrams(strs []string) [][]string {
        
    dict :=  make(map[string][]string)
    for _, str := range strs {
        key := countingSort(str)
        _, ok := dict[key]
        if !ok {
            dict[key] = make([]string, 0, 1)
        }
        dict[key] = append(dict[key], str)
    }
    
    result := make([][]string, 0, len(dict))
    for _, list := range dict {
        result = append(result, list)
    }
    return result
}

func countingSort(word string) string {
    // there are 26 lower case letters in English
    aux := make([]int, 26)
    
    for _, r := range word {
        aux[r-'a']++
    }
    
    // sum up all elements in aux
    c := 0
    for i, count := range aux {
        c += count
        aux[i] = c
    }
    
    sortedWord := make([]byte, len(word))
    for i := len(word) - 1; i >= 0; i-- {
        index := word[i] - 'a' 
        if aux[index] - 1 >= 0 {
            aux[index]--
            sortedWord[aux[index]] = word[i]
        }
    }
    return string(sortedWord)
}
