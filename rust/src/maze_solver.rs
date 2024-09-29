#[derive(Debug, PartialEq, Clone)]
pub struct Point {
    x: usize,
    y: usize,
}

impl Point {
    fn north(&self) -> Option<Point> {
        if self.y > 0 {
            Some(Point {
                x: self.x,
                y: self.y.saturating_sub(1),
            })
        } else {
            None
        }
    }

    fn east(&self) -> Option<Point> {
        Some(Point{
            x: self.x.saturating_add(1),
            y: self.y,
        })
    }

    fn south(&self) -> Option<Point> {
        Some(Point{
            x: self.x,
            y: self.y.saturating_add(1),
        })
    }

    fn west(&self) -> Option<Point> {
        if self.x > 0 {
            Some(Point {
                x: self.x.saturating_sub(1),
                y: self.y,
            })
        } else {
            None
        }
    }

    fn directions(&self) -> [Option<Point>; 4] {
        [
            self.north(),
            self.east(),
            self.south(),
            self.west(),
        ]
    }
}

pub struct Maze {
    start: Point,
    end: Point,
    width: usize,
    data: Vec<char>,
}

impl Maze {
    pub fn new(start: Point, end: Point, width: usize, data: Vec<char>) -> Self {
        Maze{
            start,
            end,
            width,
            data,
        }
    }

    fn index_from_coordinate(&self, coordinate: &Point) -> usize {
        coordinate.y * self.width + coordinate.x
    }

    fn coordinate_out_bound(&self, coordinate: &Point) -> bool {
        coordinate.x > self.width || coordinate.y > self.width
    }

    fn coordinate_is_destination(&self, coordinate: &Point) -> bool {
        *coordinate == self.end
    }

    fn coordinate_is_wall(&self, coordinate: &Point) -> bool {
        self.data[self.index_from_coordinate(coordinate)] == '#'
    }
}

struct SolvingProgress {
    seen: Vec<bool>,
    path: Vec<Point>,
}

impl SolvingProgress {
   fn new(maze: &Maze) -> Self {
        SolvingProgress{
            seen: vec![false ;maze.data.len()],
            path: Vec::new(),
        }
    }

    fn index_is_seen(&self, index: usize) -> bool {
        self.seen[index]
    }

    fn set_seen(&mut self, index: usize) {
        self.seen[index] = true;
    }

    fn add_to_path(&mut self, coordinate: Point) {
        self.path.push(coordinate);
    }

    fn pop_from_path(&mut self) {
        self.path.pop();
    }

    fn get_route(&mut self) -> Vec<Point> {
        self.path.clone()
    }
}

fn walk(maze: &Maze, state: &mut SolvingProgress, current_position: Point) -> bool {
    if maze.coordinate_out_bound(&current_position) {
        return false;
    }
    if maze.coordinate_is_wall(&current_position) {
        return false;
    }
    if state.index_is_seen(maze.index_from_coordinate(&current_position)) {
        return false;
    }
    if maze.coordinate_is_destination(&current_position) {
        state.add_to_path(current_position);
        return true;
    }

    //pre
    state.add_to_path(current_position.clone()); 
    state.set_seen(maze.index_from_coordinate(&current_position));
    //recurse
    for new_position in current_position.directions() {
        if new_position.is_some() && walk(maze, state, new_position.unwrap().clone()) {
            return true;
        }
    }

    //post
    state.pop_from_path();
    false
}

pub fn solve_maze(maze: Maze) -> Vec<Point>{
    let mut state = SolvingProgress::new(&maze);
    walk(&maze, &mut state, maze.start.clone());
    state.get_route() 
}



#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_maze_solver_line() {
        let maze = Maze::new(Point{x: 0, y: 1}, Point{x: 2, y: 1}, 3, vec![
                '#', '#', '#',
                ' ', ' ', ' ',
                '#', '#', '#'
            ]);
        let result = solve_maze(maze);
        assert_eq!(result, vec![Point{x: 0, y: 1}, Point{x: 1, y:1}, Point{x: 2, y: 1}]);
    }

    #[test]
    fn test_complicated_maze() {
        let maze = Maze::new(Point{x: 1, y: 9}, Point{x: 6, y: 0}, 9, vec![
                '#', '#', '#', '#', '#', '#', ' ', '#', '#',
                '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
                '#', '#', '#', '#', '#', ' ', '#', ' ', '#',
                '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', '#',
                '#', ' ', '#', '#', '#', ' ', '#', ' ', '#',
                '#', ' ', '#', ' ', ' ', ' ', '#', ' ', '#',
                '#', '#', '#', '#', ' ', '#', '#', ' ', '#',
                '#', '#', ' ', '#', ' ', '#', ' ', ' ', '#',
                '#', ' ', ' ', ' ', ' ', '#', ' ', '#', '#',
                '#', ' ', '#', '#', '#', '#', '#', '#', '#',
            ]);
        let result = solve_maze(maze);
        assert_eq!(result, vec![
            Point{x: 1, y: 9}, 
            Point{x: 1, y: 8}, 
            Point{x: 2, y: 8}, 
            Point{x: 3, y: 8},
            Point{x: 4, y: 8},
            Point{x: 4, y: 7},
            Point{x: 4, y: 6},
            Point{x: 4, y: 5},
            Point{x: 5, y: 5},
            Point{x: 5, y: 4},
            Point{x: 5, y: 3},
            Point{x: 5, y: 2},
            Point{x: 5, y: 1},
            Point{x: 6, y: 1},
            Point{x: 6, y: 0},
        ]);
    }       

    #[test]
    fn test_index_from_coordinate(){
        let maze = Maze::new(Point{x: 0, y: 1}, Point{x: 2, y: 1}, 3, vec![
            '#', '#', '#',
            ' ', ' ', ' ',
            '#', '#', '#'
        ]);
        let result = maze.index_from_coordinate(&Point{x: 0, y: 1});
        
        assert_eq!(result, 3);
    }

    #[test]
    fn test_coordinate_out_of_bounds() {
        let maze = Maze::new(Point{x: 0, y: 1}, Point{x: 2, y: 1}, 2, vec![
            '#', '#',
            '#', '#',
        ]);
        let with_out_of_bounds_coordinanate = Point{x: 3, y:0};
        let with_in_bounds_coordinate = Point{x:0, y:0};
        assert!(maze.coordinate_out_bound(&with_out_of_bounds_coordinanate));
        assert!(!maze.coordinate_out_bound(&with_in_bounds_coordinate));
    }
}
