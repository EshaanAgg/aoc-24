pub fn part1(file_path: &str) -> u64 {
    let v = std::fs::read_to_string(file_path)
        .expect("Unable to read file")
        .lines()
        .map(|line| {
            line.split(" ")
                .map(|x| {
                    x.parse::<u64>()
                        .expect("Unable to parse the content of the line as an integer")
                })
                .collect::<Vec<u64>>()
        });
    0
}
