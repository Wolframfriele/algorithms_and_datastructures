use std::cmp::Ordering;

pub fn linear_search(haystack: Vec<i32>, needle: i32) -> Option<usize> {
    for (i, value) in haystack.iter().enumerate(){
        if *value == needle {
            return Some(i)
        }
    }
    None
}

pub fn binary_search(haystack:Vec<i32>, needle: i32) -> Option<usize> {
    let mut l: usize = 0;
    let mut h = haystack.len();
    while l < h{
        let m = l + (h-l)/2;
        match haystack[m].cmp(&needle) {
            Ordering::Equal => return Some(m),
            Ordering::Less => l = m + 1,
            Ordering::Greater => h = m,
        }
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_linear_search_found() {
        let input_haystack = vec![1, 2, 3];
        let result = linear_search(input_haystack, 2);
        assert_eq!(result, Some(1));
    }

    #[test]
    fn test_linear_search_not_found() {
        let input_haystack = vec![1, 2, 3];
        let result = linear_search(input_haystack, 4);
        assert_eq!(result, None);
    }

    #[test]
    fn test_binary_search_found() {
        let input_haystack = vec![1, 2, 3];
        let result = binary_search(input_haystack, 2);
        assert_eq!(result, Some(1));
    }

    #[test]
    fn test_binary_search_found_2() {
        let input_haystack = vec![1, 2, 3, 4, 5, 6, 7, 8, 9];
        let result = binary_search(input_haystack, 8);
        assert_eq!(result, Some(7));
    }

    #[test]
    fn test_binary_search_not_found() {
        let input_haystack = vec![1, 2, 3];
        let result = binary_search(input_haystack, 4);
        assert_eq!(result, None);
    }
}
