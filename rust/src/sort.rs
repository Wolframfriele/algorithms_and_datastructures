pub fn bubble_sort(arr: &mut Vec<i32>) {
   for i in 0..arr.len(){
        for j in 0..arr.len() -1 -i{
            if arr[j] > arr[j+1] {
                arr.swap(j, j+1);
            }
        }
    } 
}

#[cfg(test)]
mod tests{
    use super::*;

    #[test]
    fn test_bubble_sort() {
        let mut sortable = vec![2, 1, 4, 3, 10, 9, 5];
        bubble_sort(&mut sortable);
        assert_eq!(sortable, vec![1, 2, 3, 4, 5, 9, 10]);
    }
} 
    
    
