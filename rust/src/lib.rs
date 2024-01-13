pub fn linear_search(haystack: Vec<i32>, needle: i32) -> bool {
    for i in haystack{
        if i == needle {
            return true
        }
    }
    return false
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_linear_search_found() {
        let input_haystack = vec![1, 2, 3];
        let result = linear_search(input_haystack, 2);
        assert_eq!(result, true);
    }
    #[test]
    fn test_linear_search_not_found() {
        let input_haystack = vec![1, 2, 3];
        let result = linear_search(input_haystack, 4);
        assert_eq!(result, false);
    }
}
